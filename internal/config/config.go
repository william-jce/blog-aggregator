package config

const configFileName = ".gatorconfig.json"

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	filePath, err := getConfigFilePath()
}

func (cfg *Config) SetUser(currentUserName string) error {

}

func getConfigFilePath() (string, error) {

}

func write(cfg Config) error {

}
