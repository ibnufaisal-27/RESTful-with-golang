package config

//Config Struct
type Config struct {
	DB *DBConfig
}

//DBConfig Struct
type DBConfig struct {
	Dialect  string
	Username string
	Password string
	Name     string
	Charset  string
}

//GetConfig DB
func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "mysql",
			Username: "root",
			Password: "",
			Name:     "employee",
			Charset:  "utf8",
		},
	}
}
