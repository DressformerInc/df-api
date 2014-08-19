package models

import (
	r "github.com/dancannon/gorethink"
	"log"
)

type GarmentScheme struct {
	Id   string `gorethink:"id,omitempty"`
	Name string `gorethink:"name"`
}

type Garment struct {
	Object *UserScheme
}

func (*Garment) Construct(args ...interface{}) interface{} {
	return &Garment{}
}

func (this *Garment) FindAll(ids interface{}) []GarmentScheme {
	rows, err := r.Table("users").GetAll(r.Args(ids)).Run(session())
	if err != nil {
		log.Println("Unable to fetch cursor for args:", ids, "Error:", err)
		return nil
	}

	var result []GarmentScheme

	if err = rows.All(&result); err != nil {
		log.Println("Unable to get data, err:", err)
	}

	return result
}
