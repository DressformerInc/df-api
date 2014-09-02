package models

import (
	. "df/api/utils"
	r "github.com/dancannon/gorethink"
	"log"
)

type UserScheme struct {
	Dummy string `gorethink:"dummy" json:"dummy"`

	Body struct {
		Height    float32 `gorethink:"height"    json:"height,omitempty"`
		Chest     float32 `gorethink:"chest"     json:"chest,omitempty"`
		Underbust float32 `gorethink:"underbust" json:"underbust,omitempty"`
		Waist     float32 `gorethink:"waist"     json:"waist,omitempty"`
		Hips      float32 `gorethink:"hips"      json:"hips,omitempty"`
	} `gorethink:"body,omitempty" json:"body,omitempty"`
}

type User struct {
	r.Term
}

func (*User) Construct(args ...interface{}) interface{} {
	return &User{
		r.Db("dressformer").Table("users"),
	}
}

func (this *User) Find(args ...interface{}) *UserScheme {
	result := &UserScheme{
		Dummy: AppConfig.AssetsUrl() + "/geometry/" + getDefaultDummy(),
	}

	return result
}

func getDefaultDummy() string {
	result := map[string]interface{}{"id": ""}

	rows, err := r.Db("dressformer").Table("geometry").GetAllByIndex("default_dummy", true).Run(session())
	if err != nil {
		log.Println("Unable to fetch cursor for GetAllByIndex(default_dummy:true). Error:", err)
		return result["id"].(string)
	}

	if err := rows.One(&result); err != nil {
		log.Println("Unable to get data, err:", err)
		return result["id"].(string)
	}

	return result["id"].(string)
}
