package conf

type Config struct {
	// common
	App   AppConfig
	Log   LogConfig
	MySQL MySQLConfig
	Redis RedisConfig
	Cache CacheConfig
}

// AppConfig
type AppConfig struct {
	Name      string
	RunMode   string `mapstructure:"run_mode"`
	Addr      string
	Url       string
	JwtSecret string
}

// LogConfig
type LogConfig struct {
	Writers          string `mapstructure:"writers"`
	LoggerLevel      string `mapstructure:"logger_level"`
	LoggerFile       string `mapstructure:"logger_file"`
	LoggerWarnFile   string `mapstructure:"logger_warn_file"`
	LoggerErrorFile  string `mapstructure:"logger_error_file"`
	LogFormatText    bool   `mapstructure:"log_format_text"`
	LogRollingPolicy string `mapstructure:"log_rolling_policy"`
	LogRotateDate    int    `mapstructure:"log_rotate_date"`
	LogRotateSize    int    `mapstructure:"log_rotate_size"`
	LogBackupCount   int    `mapstructure:"log_backup_count"`
}

// MySQLConfig
type MySQLConfig struct {
	Name            string
	Addr            string
	UserName        string
	Password        string
	ShowLog         bool `mapstructure:"show_log"`
	MaxIdleConn     int  `mapstructure:"max_idle_conn"`
	MaxOpenConn     int  `mapstructure:"max_open_conn"`
	ConnMaxLifeTime int  `mapstructure:"conn_max_life_time"`
}

// RedisConfig
type RedisConfig struct {
	Addr         string
	Password     string
	Db           int
	DialTimeout  int `mapstructure:"dial_timeout"`
	ReadTimeout  int `mapstructure:"read_timeout"`
	WriteTimeout int `mapstructure:"write_timeout"`
	PoolTimeOut  int `mapstructure:"pool_timeout"`
	PoolSize     int `mapstructure:"pool_size"`
}

// CacheConfig
type CacheConfig struct {
	Driver string
	Prefix string
}
