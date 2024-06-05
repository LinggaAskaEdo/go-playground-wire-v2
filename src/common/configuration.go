package common

type Configuration struct {
	Token    TokenConfiguration
	Logger   LoggerConfiguration
	Database DatabaseConfiguration
	Cache    CacheConfiguration
	Server   ServerConfiguration
	Handler  HandlerConfiguration
	Business BusinessConfiguration
}

type TokenConfiguration struct {
	Auth string
}

type LoggerConfiguration struct {
}

type DatabaseConfiguration struct {
	MySQL MySQLConfiguration
}

type MySQLConfiguration struct {
	MySQLUser            string
	MySQLPassword        string
	MySQLHost            string
	MySQLPort            string
	MySQLName            string
	MySQLMaxIdleConns    int
	MySQLMaxOpenConns    int
	MySQLConnMaxLifetime int
	MySQLConnMaxIdleTime int
}

type CacheConfiguration struct {
	Redis RedisConfiguration
}

type RedisConfiguration struct {
	RedisHost     string
	RedisPort     string
	RedisPassword string
}

type ServerConfiguration struct {
	ServerHost string
	ServerPort string
}

type HandlerConfiguration struct {
	Scheduler SchedulerConfiguration
}

type SchedulerConfiguration struct {
	SchedulerNewsEnable    bool
	SchedulerNewsTimeInSec int
}

type BusinessConfiguration struct {
	Usecase UsecaseConfiguration
}

type UsecaseConfiguration struct {
	News NewsConfiguration
}

type NewsConfiguration struct {
	RssURL string
}
