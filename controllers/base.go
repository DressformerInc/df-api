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

func (this *BaseController) Find(u *models.User, r render.Render, params martini.Params) (int, []byte) {
	r.JSON(http.StatusOK, this.model.Find(params["id"]))
}

*/
