package models

import (
	. "df/api/utils"
)

type Source struct {
	Id     string  `gorethink:"id,omitempty"          json:"id,omitempty"`
	Weight float64 `gorethink:"weight,omitempty"      json:"weight,omitempty"`
	Name   string  `gorethink:"origin_name,omitempty" json:"origin_name,omitempty"`
	Url    string  `gorethink:"-"                     json:"url,omitempty" binding:"-"`
}

func (this *Source) LinkTo(name string) string {
	return AppConfig.AssetsUrl() + "/" + name + "/" + this.Id
}

func url(s *Source, name string) {
	if s == nil || s.Id == "" {
		return
	}

	s.Url = s.LinkTo(name)
}