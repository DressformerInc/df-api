package models

import (
	. "df/api/utils"
	"errors"
	r "github.com/dancannon/gorethink"
	enc "github.com/dancannon/gorethink/encoding"
	"log"
)

const H_LEN = 10

type UserScheme struct {
	Id      string           `gorethink:"id,omitempty"      json:"-"`
	Token   string           `gorethink:"token,omitempty"   json:"-"`
	Dummy   *DummyScheme     `json:"dummy,omitempty"`
	Name    string           `gorethink:"name,omitempty"    json:"name,omitempty"`
	History []*GarmentScheme `gorethink:"history,omitempty" json:"history,omitempty"`
}

type User struct {
	r.Term
	dummy  *Dummy
	Object *UserScheme
}

func (*User) Construct(args ...interface{}) interface{} {
	user := &User{
		r.Db("dressformer").Table("users"),
		(*Dummy).Construct(nil).(*Dummy),
		&UserScheme{},
	}

	log.Println("args:", args)

	if len(args) > 0 {
		if user.Object = user.constructFrom(args[0]); user.Object == nil {
			log.Println("Unexpected error, unable to proceed. Error: user.Object is nil")
			return user
		}
	}

	if user.dummy != nil {
		user.Object.Dummy = user.dummy.Default()
	}

	return user
}

func (this *User) constructFrom(args ...interface{}) *UserScheme {
	var i interface{}

	if len(args) > 0 {
		i = args[0]
	}

	switch t := i.(type) {
	case Token:
		// if the token has been restored from cookie, we've already have this user, so find it
		if i.(Token).IsRestored {
			return this.Find(i.(Token))
		}

		// if not, return newly created one
		u, _ := this.Create(UserScheme{Token: i.(Token).Get()})
		return u

	default:
		log.Println("Unexpected type:", t)
	}

	return nil
}

func (this *User) Find(args ...interface{}) *UserScheme {
	var i interface{}

	user := &UserScheme{}

	if len(args) > 0 {
		i = args[0]
	}

	switch i.(type) {
	case Token:
		token := i.(Token).Get()
		log.Println("find by token, ", token)

		rows, err := this.GetAllByIndex("token", token).Run(session())
		if err != nil {
			log.Println("Unable to fetch cursor for index token:", token, "Error:", err)
			return nil
		}

		if err = rows.One(&user); err != nil {
			log.Println("Error getting data. Error:", err)
			return nil
		}
	}

	log.Println("result:", user)

	return user
}

func (this *User) Put(payload interface{}) (*UserScheme, error) {
	id := this.Object.Id

	log.Println("Updating:", id, "with:", payload)

	result, err := this.Get(id).Update(payload, r.UpdateOpts{ReturnChanges: true, Durability: "soft"}).Run(session())
	if err != nil {
		log.Println("Error updating:", id, "with data:", payload, "error:", err)
		return nil, errors.New("Wrong data")
	}

	response := &r.WriteResponse{}

	if err = result.One(response); err != nil {
		log.Println("Unable to iterate cursor:", err)
		return nil, errors.New("Internal server error")
	}

	if len(response.Changes) != 1 {
		log.Println("Unexpected length of Changes:", len(response.Changes))
		return nil, errors.New("Internal server error")
	}

	newval := &UserScheme{}

	if err = enc.Decode(newval, response.Changes[0].NewValue); err != nil {
		log.Println("Decode error:", err)
		return nil, errors.New("Internal server error")
	}

	return newval, nil
}

func (this *User) Create(payload UserScheme) (*UserScheme, error) {
	result, err := this.Insert(payload, r.InsertOpts{ReturnChanges: true, Durability: "soft"}).Run(session())
	if err != nil {
		log.Println("Error inserting data:", err)
		return nil, errors.New("Internal server error")
	}

	response := &r.WriteResponse{}

	if err = result.One(response); err != nil {
		log.Println("Unable to iterate cursor:", err)
		return nil, errors.New("Internal server error")
	}

	log.Println("inserted :", response.Inserted)

	if len(response.Changes) != 1 {
		log.Println("Unexpected length of Changes:", len(response.Changes))
		return nil, errors.New("Internal server error")
	}

	newval := &UserScheme{}

	if err = enc.Decode(newval, response.Changes[0].NewValue); err != nil {
		log.Println("Decode error:", err)
		return nil, errors.New("Internal server error")
	}

	return newval, nil
}

func (this *User) UpdateHistory(g *GarmentScheme) {
	history := []*GarmentScheme{}

	history = append(history, g)
	for i, _ := range this.Object.History {
		if i == H_LEN-1 {
			break
		}

		if this.Object.History[i].Id == g.Id {
			continue
		}

		history = append(history, this.Object.History[i])
	}

	this.Object.History = history
	this.Put(this.Object)
}

// Just an example of usage subqueries in RethinkDB
//
//	rows, err := r.Db("dressformer").Table("tokens").Get(token).Merge(map[string]interface{}{
//		"user": r.Db("dressformer").Table("users").Get(r.Row.Field("user_id")),
//	}).Run(session())
//	if err != nil {
//		log.Println("Error finding user object by token. Error:", err)
//		return user
//	}
//
//	result := &struct {
//		User *UserScheme `gorethink:"user,omitempty"`
//	}{&UserScheme{}}
