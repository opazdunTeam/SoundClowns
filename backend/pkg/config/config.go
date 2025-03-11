package config

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func GetDBConfig() DBConfig {
	return DBConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "music_user",
		Password: "music_password",
		DBName:   "music_db",
	}
}
