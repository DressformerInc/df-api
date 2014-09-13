package models

import (
	. "df/api/utils"
	r "github.com/dancannon/gorethink"
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
