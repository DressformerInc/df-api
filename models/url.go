package models

type URLOptionsScheme struct {
	Ids   string `form:"ids"`
	Start int    `form:"start"`
	Limit int    `form:"limit"`
}
