package models

type UserScheme struct {
	Avatar struct {
		Model string `gorethink:"model" json:"model"`
	} `gorethink:"avatar,omitempty" json:"avatar, omitempty"`

	Body struct {
		Height    float32 `gorethink:"height"    json:"height"`
		Chest     float32 `gorethink:"chest"     json:"chest"`
		Underbust float32 `gorethink:"underbust" json:"underbust"`
		Waist     float32 `gorethink:"waist"     json:"waist"`
		Hips      float32 `gorethink:"hips"      json:"hips"`
	} `gorethink:"body,omitempty" json:"body, omitempty"`
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

	user.Avatar.Model = "//assets.dressformer.com/model/53d11d10fcb05d8ed2000042"
	user.Body.Chest = 90
	user.Body.Height = 170
	user.Body.Hips = 90
	user.Body.Underbust = 70
	user.Body.Waist = 60

	return user
}
