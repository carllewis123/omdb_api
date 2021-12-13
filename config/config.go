package config

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type config struct {
	Database          string `yaml:"database"`
	Connection        string `yaml:"db_connection"`
	MaxConnection     int    `yaml:"max_connection"`
	MaxIdleConnection int    `yaml:"max_idle_connection"`
	Username          string `yaml:"db_username"`
	Password          string `yaml:"db_password"`
	DBName            string `yaml:"db_name"`
	Host              string `yaml:"db_host"`
	Port              string `yaml:"db_port"`
	ClientId          string `yaml:"client_id"`
	ClientSecret      string `yaml:"client_secret"`

	FilePubKey  string `yaml:"public_key"`
	FilePrivKey string `yaml:"private_key"`
}

type url_config struct {
	UrlSearchByParam     string `yaml:"url_search_by_param"`
	UrlGetDetailFilm     string `yaml:"url_get_detail_film"`
	UrlOmdbSearchByParam string `yaml:"url_omdb_search_by_param"`

	IpAddr          string `yaml:"url_ip_address"`
	ServicePortHttp string `yaml:"http_service_port"`
	ServicePortGrpc string `yaml:"grpc_service_port"`
}

func ConfigYAML() (conf *config, err error) {
	data, err := ioutil.ReadFile("./config/config.yaml")
	if err != nil {
		return
	}
	conf = &config{}
	err = yaml.Unmarshal(data, conf)
	if err != nil {
		panic(err)
	}
	return
}

func ConfigUrl() (conf *url_config, err error) {
	data, err := ioutil.ReadFile("./config/url_config.yaml")
	if err != nil {
		return
	}
	conf = &url_config{}
	err = yaml.Unmarshal(data, conf)
	return
}
