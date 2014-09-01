package models

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

	user.Dummy = "//localhost:6500/geometry/e12c24c9-c8b7-45ac-bbd6-f57ee8c362e9"
	return user
}
