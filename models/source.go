package models

import (
	. "df/api/utils"
)

type Source struct {
	Id      string  `gorethink:"id,omitempty"        json:"id,omitempty"`
	Weight  float64 `gorethink:"weight,omitempty"    json:"weight,omitempty"`
	Name    string  `gorethink:"orig_name,omitempty" json:"orig_name,omitempty"`
	Options string  `gorethink:"options,omitempty"   json:"options,omitempty"`
}

func (this *Source) LinkTo(name string) string {
	if name == "" {
		return AppConfig.AssetsUrl() + "/" + this.Id
	}

	return AppConfig.AssetsUrl() + "/" + name + "/" + this.Id
}
