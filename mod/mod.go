package mod

var Conf Config

type Config struct {
	*App
	*Logger
	*Mysql
	*Snowflake
	*Rides
}

type App struct {
	Name    string `mapstructure:"name"`
	Version string `mapstructure:"version"`
	Port    int    `mapstructure:"port"`
	Mode    string `mapstructure:"mode"`
}

type Logger struct {
	Filename   string `mapstructure:"Filename"`
	MaxSize    int    `mapstructure:"MaxSize"`
	MaxAge     int    `mapstructure:"MaxAge"`
	MaxBackups int    `mapstructure:"MaxBackups"`
	Compress   bool   `mapstructure:"Compress"`
	Level      string `mapstructure:"Level"`
}

type Mysql struct {
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
	Dbname   string `mapstructure:"dbname"`
	Username string `mapstructure:"username"`
}

type Snowflake struct {
	StartTime string `mapstructure:"startTime"`
	MachineId int64  `mapstructure:"machineId"`
}

type Rides struct {
	Addr     string `mapstructure:"addr"`
	Db       int    `mapstructure:"db"`
	Password string `mapstructure:"password"`
	PoolSize int    `mapstructure:"poolsize"`
}
