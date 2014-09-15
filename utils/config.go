package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var AppConfig *ConfigScheme

func init() {
	AppConfig = &ConfigScheme{}
}

type ConfigScheme struct {
	App struct {
		ListenOn string `json:"listen_on"`
		HttpsOn  string `json:"https_on"`
		SSLCert  string `json:"ssl_cert"`
		SSLKey   string `json:"ssl_key"`
		HashKey  string `json:"hash_key"`
		BlockKey string `json:"block_key"`
	} `json:"application"`

	Endpoints struct {
		Api    string `json:"api"`
		Assets string `json:"assets"`
	} `json:"endpoints"`

	Connections struct {
		Rethink struct {
			Spec   string `json:"spec"`
			DbName string `json:"db_name"`
		} `json:"rethink"`
	} `json:"connections"`
}

func InitConfigFrom(file string) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Println("Unable to read", file, "Error:", err)
		return
	}

	err = json.Unmarshal(data, AppConfig)
	if err != nil {
		log.Println("Unable to read config.", err)
	}
}

func (this *ConfigScheme) ListenOn() string {
	return this.App.ListenOn
}

func (this *ConfigScheme) HttpsOn() string {
	return this.App.HttpsOn
}

func (this *ConfigScheme) SSLCert() string {
	return this.App.SSLCert
}

func (this *ConfigScheme) SSLKey() string {
	return this.App.SSLKey
}

func (this *ConfigScheme) ApiUrl() string {
	return this.Endpoints.Api
}

func (this *ConfigScheme) AssetsUrl() string {
	return this.Endpoints.Assets
}

func (this *ConfigScheme) HashKey() []byte {
	if this.App.HashKey == "" {
		return nil
	}

	return []byte(this.App.HashKey)
}

func (this *ConfigScheme) BlockKey() []byte {
	if this.App.BlockKey == "" {
		return nil
	}

	return []byte(this.App.BlockKey)
}

func (this *ConfigScheme) RethinkAddress() string {
	return this.Connections.Rethink.Spec
}

func (this *ConfigScheme) RethinkDbName() string {
	return this.Connections.Rethink.DbName
}
