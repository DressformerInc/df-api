package utils

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
