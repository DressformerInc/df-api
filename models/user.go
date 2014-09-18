package models

import (
	. "df/api/utils"
	r "github.com/dancannon/gorethink"
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
	*Base
	Dummy  *Dummy
	Object *UserScheme
}

func (*User) Construct(args ...interface{}) interface{} {
	user := &User{
		&Base{r.Db("dressformer").Table("users")},
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

	if user.Dummy != nil {
		user.Object.Dummy = user.Dummy.Default()
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
		u, _ := this.Create(&UserScheme{Token: i.(Token).Get()})
		return u.(*UserScheme)

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
	this.Put(this.Object.Id, this.Object)
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
