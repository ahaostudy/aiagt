package conf

import (
	"github.com/aiagt/aiagt/common/confutil"
	"path/filepath"

	ktconf "github.com/aiagt/kitextool/conf"
)

var conf = new(ServerConf)

func init() {
	confutil.LoadConf(conf,
		filepath.Join("conf"),
		filepath.Join("apps", "knowledge", "conf"),
	)
}

func Conf() *ServerConf {
	return conf
}

type ServerConf struct {
	ktconf.ServerConf

	Metrics   Metrics   `yaml:"metrics"`
	Tracing   Tracing   `yaml:"tracing"`
	Milvus    Milvus    `yaml:"milvus"`
	Embedding Embedding `yaml:"embedding"`
	Cos       Cos       `yaml:"cos"`
}

type Metrics struct {
	Addr string `yaml:"addr"`
}

type Tracing struct {
	ExportAddr string `yaml:"export_addr"`
}

type Milvus struct {
	Address    string `yaml:"address"`
	Username   string `yaml:"username"`
	Password   string `yaml:"password"`
	Collection string `yaml:"collection"`
}

type Embedding struct {
	APIKey  string `yaml:"api_key"`
	Model   string `yaml:"model"`
	BaseURL string `yaml:"base_url"`
}

type Cos struct {
	URL       string `yaml:"url"`
	SecretID  string `yaml:"secret_id"`
	SecretKey string `yaml:"secret_key"`
}
