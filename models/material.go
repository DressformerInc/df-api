package models

import (
	r "github.com/dancannon/gorethink"
)

type MaterialScheme struct {
	Id        string  `gorethink:"id,omitempty"     json:"id"                  mtl:"-"`
	Name      string  `gorethink:"name,omitempty"   json:"name,omitempty"      mtl:"newmtl"`
	Ka        string  `gorethink:"ka,omitempty"     json:"ka,omitempty"        mtl:"Ka"`
	Kd        string  `gorethink:"kd,omitempty"     json:"kd,omitempty"        mtl:"Kd"`
	Ks        string  `gorethink:"ks,omitempty"     json:"ks,omitempty"        mtl:"Ks"`
	Ts        string  `gorethink:"ts,omitempty"     json:"ts,omitempty"        mtl:"Ts"`
	Tf        string  `gorethink:"tf,omitempty"     json:"tf,omitempty"        mtl:"Tf"`
	Tr        string  `gorethink:"tr,omitempty"     json:"tr,omitempty"        mtl:"Tr"`
	Illum     int     `gorethink:"illum,omitempty"  json:"illum,omitempty"     mtl:"illum"`
	D         string  `gorethink:"d,omitempty"      json:"d,omitempty"         mtl:"d"`
	Ns        string  `gorethink:"ns,omitempty"     json:"ns,omitempty"        mtl:"Ns"`
	Sharpness int     `gorethink:"sharpness,omitempty" json:"sharpness,omitempty" mtl:"sharpness"`
	Ni        string  `gorethink:"ni,omitempty"     json:"ni,omitempty"        mtl:"Ni"`
	Map_Ka    *Source `gorethink:"map_ka,omitempty" json:"map_ka,omitempty"    mtl:"map_Ka"`
	Map_Kd    *Source `gorethink:"map_kd,omitempty" json:"map_kd,omitempty"    mtl:"map_Kd"`
	Map_Ks    *Source `gorethink:"map_ks,omitempty" json:"map_ks,omitempty"    mtl:"map_Ks"`
	Map_Ns    *Source `gorethink:"map_ns,omitempty" json:"map_ns,omitempty"    mtl:"map_Ns"`
	Map_D     *Source `gorethink:"map_d,omitempty"  json:"map_d,omitempty"     mtl:"map_d"`
	Disp      *Source `gorethink:"disp,omitempty"   json:"disp,omitempty"      mtl:"disp"`
	Decal     *Source `gorethink:"decal,omitempty"  json:"decal,omitempty"     mtl:"decal"`
	Bump      *Source `gorethink:"bump,omitempty"   json:"bump,omitempty"      mtl:"bump"`
	Refl      *Source `gorethink:"refl,omitempty"   json:"refl,omitempty"      mtl:"refl"`
}

type Material struct {
	*Base
}

func (*Material) Construct(args ...interface{}) interface{} {
	return &Material{
		&Base{r.Db("dressformer").Table("materials")},
	}
}

func (this *Material) Find(id string) *MaterialScheme {
	i, err := this.Base.Find(id, &MaterialScheme{})
	if err != nil {
		return nil
	}

	return i.(*MaterialScheme)
}

func (this *Material) FindAll(ids []string, opts URLOptionsScheme) []MaterialScheme {
	i, err := this.Base.FindAll(ids, opts, &[]MaterialScheme{})
	if err != nil {
		return nil
	}

	return *(i.(*[]MaterialScheme))
}
