package conf

import "flag"

type Config struct {
	Name string
	Host string
	Port int

	DbHost string
	DbPort int
	DbUser string
	DbPwd  string
	DbName string

	JwtKey        string
	JwtTimeout    int
	JwtMaxRefresh int
}

func ParseConfig() *Config {
	cfg := &Config{}

	flag.StringVar(&cfg.Name, "name", "base-go", "server name")
	flag.StringVar(&cfg.Host, "host", "127.0.0.1", "server host")
	flag.IntVar(&cfg.Port, "port", 5000, "server port")
	flag.StringVar(&cfg.DbHost, "db-host", "127.0.0.1", "database host")
	flag.IntVar(&cfg.DbPort, "db-port", 3306, "database port")
	flag.StringVar(&cfg.DbUser, "db-user", "root", "database user")
	flag.StringVar(&cfg.DbPwd, "db-pwd", "root", "database password")
	flag.StringVar(&cfg.DbName, "db-name", "base", "database name")
	flag.StringVar(&cfg.JwtKey, "jwt-key", "a-secret-key", "jwt key")
	flag.IntVar(&cfg.JwtTimeout, "jwt-timeout", 7*24, "jwt timeout hours")
	flag.IntVar(&cfg.JwtMaxRefresh, "jwt-max-refresh", 24, "jwt max refresh hours")

	flag.Parse()

	return cfg
}
