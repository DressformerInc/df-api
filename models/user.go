package models

import (
	. "df/api/utils"
	r "github.com/dancannon/gorethink"
	"log"
)

type UserScheme struct {
	Dummy *DummyScheme `json:"dummy,omitempty"`
	Name  string       `gorethink:"name,omitempty"  json:"name,omitempty"`
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
		user.Object = user.Find(args[0])
	}

	if user.dummy != nil {
		user.Object.Dummy = user.dummy.Default()
	}

	return user
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

		rows, err := r.Db("dressformer").Table("tokens").Get(token).Merge(map[string]interface{}{
			"user": r.Db("dressformer").Table("users").Get(r.Row.Field("user_id")),
		}).Run(session())
		if err != nil {
			log.Println("Error finding user object by token. Error:", err)
			return user
		}

		result := &struct {
			User *UserScheme `gorethink:"user,omitempty"`
		}{&UserScheme{}}

		if err = rows.One(&result); err != nil {
			log.Println("Error getting data. Error:", err)
			return user
		}
		log.Println("result:", result.User)
		user = result.User
	}

	return user
}
