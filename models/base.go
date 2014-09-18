package models

import (
	. "df/api/utils"
	"errors"
	r "github.com/dancannon/gorethink"
	enc "github.com/dancannon/gorethink/encoding"
	"log"
	"time"
)

type Model interface {
	Construct(arg ...interface{}) interface{}
}

var rs *r.Session

func session() *r.Session {
	if rs != nil {
		return rs
	}

	rs, err := r.Connect(r.ConnectOpts{
		Address:     AppConfig.RethinkAddress(),
		Database:    AppConfig.RethinkDbName(),
		MaxIdle:     600,
		IdleTimeout: time.Second * 10,
	})

	if err != nil {
		log.Fatalln(err.Error())
	}

	return rs
}

type Base struct {
	r.Term
}

func (this *Base) Create(payload interface{}) (interface{}, error) {
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

	log.Println("response:", response)
	log.Println("inserted :", response.Inserted)

	if len(response.Changes) != 1 {
		log.Println("Unexpected length of Changes:", len(response.Changes))
		return nil, errors.New("Internal server error")
	}

	newval := T(payload)

	if err = enc.Decode(newval, response.Changes[0].NewValue); err != nil {
		log.Println("Decode error:", err)
		return nil, errors.New("Internal server error")
	}

	return newval, nil
}

func (this *Base) Put(id string, payload interface{}) (interface{}, error) {
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

	newval := T(payload)

	if err = enc.Decode(newval, response.Changes[0].NewValue); err != nil {
		log.Println("Decode error:", err)
		return nil, errors.New("Internal server error")
	}

	return newval, nil
}

func (this *Base) Remove(id string) error {
	_, err := this.Get(id).Delete().Run(session())
	if err != nil {
		log.Println("Error deleting:", id, "error:", err)
		return errors.New("Internal server error")
	}

	return nil
}
