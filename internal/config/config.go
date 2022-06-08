package config

type HOGConfig struct {
	MySQL MySQL `json:"mysql"`
}

type MySQL struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Address  string `json:"address"`
	DB       string `json:"db"`
}
