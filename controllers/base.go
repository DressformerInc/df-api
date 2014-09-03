package controllers

type Controller interface {
	Construct(arg ...interface{}) interface{}
}

/*
	@todo

type BaseController struct{}

func (*BaseController) Construct(args ...interface{}) interface{} {
	return &BaseController{

	}
}

func (this *BaseController) Find(u *models.User, enc encoder.Encoder, params martini.Params) (int, []byte) {
	return http.StatusOK, encoder.Must(enc.Encode(this.model.Find(params["id"])))
}
*/
