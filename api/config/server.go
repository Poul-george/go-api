package config

type Server struct {
	StartAddress string `mapstructure:"start_address"`
	IdleTimeout  int64  `mapstructure:"idle_timeout"`
}

func GetServerConfig() Server {
	c := getConfig().Server

	return Server{
		StartAddress: c.StartAddress,
		IdleTimeout:  c.IdleTimeout,
	}
}
