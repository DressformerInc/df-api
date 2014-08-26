package models

type UserScheme struct {
	Dummy string `gorethink:"dummy" json:"dummy"`

	Body struct {
		Height    float32 `gorethink:"height"    json:"height"`
		Chest     float32 `gorethink:"chest"     json:"chest"`
		Underbust float32 `gorethink:"underbust" json:"underbust"`
		Waist     float32 `gorethink:"waist"     json:"waist"`
		Hips      float32 `gorethink:"hips"      json:"hips"`
	} `gorethink:"body,omitempty" json:"body,omitempty"`
}

type User struct {
	Object *UserScheme
}

func (*User) Construct(args ...interface{}) interface{} {
	this := &User{
		Object: guest(),
	}

	return this
}

func guest() *UserScheme {
	user := &UserScheme{}

	user.Dummy = "//assets.dressformer.com/model/"
	return user
}
