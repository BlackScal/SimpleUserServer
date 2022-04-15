package configs

type UserServiceConf struct {
	AppID string `yaml:"appid"`
	Addr  struct {
		IP   string `yaml:"ip"`
		Port int    `yaml:"port"`
	}
	Log struct {
		Level  string `yaml:"level"`
		Output string `yaml:"output"`
		Format string `yaml:"format"`
	}
}
