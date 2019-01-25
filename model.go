package gotuil

import "time"

// SettingInfo 配置信息
type SettingInfo struct {
	App       *AppSetting     `json:"app" yaml:"app"`
	Server    *ServerSetting  `json:"server" yaml:"server"`
	RedisList []*RedisSetting `json:"redis_list" yaml:"redis_list"`
	DBList    []*DBSetting    `json:"db_list" yaml:"db_list"`
}

// AppSetting 应用配置
type AppSetting struct {
	JwtSecret   string `json:"jwt_secret" yaml:"jwt_secret"`
	JwtTokenExp int64  `json:"jwt_token_exp" yaml:"jwt_token_exp"`
	PageSize    int    `json:"page_size" yaml:"page_size"`
	PrefixURL   string `json:"prefix_url" yaml:"prdfix_url"`
	TimeFormat  string `json:"time_format" yaml:"time_format"`
	LogPath     string `json:"log_path" yaml:"log_path"`
}

// ServerSetting 服务配置信息
type ServerSetting struct {
	RunMode      string        `json:"run_mode" yaml:"run_mode"`
	HTTPPort     int           `json:"http_port" yaml:"http_port"`
	ReadTimeout  time.Duration `json:"read_timeout" yaml:"read_timeout"`
	WriteTimeout time.Duration `json:"write_timeout" yaml:"write_timeout"`
}

// DBSetting 数据库配置
type DBSetting struct {
	Type        string `json:"db_type" yaml:"db_type"`
	Name        string `json:"name" yaml:"name"`
	TablePrefix string `json:"table_prefix" yaml:"table_prefix"`
	DataSource  string `json:"data_source" yaml:"data_source"`
}

// RedisSetting Redis配置信息
type RedisSetting struct {
	Host        string        `json:"host" yaml:"host"`
	Port        int           `json:"port" yaml:"port"`
	Password    string        `json:"password" yaml:"password"`
	MaxIdle     int           `json:"max_idle" yaml:"max_idle"`
	MaxActive   int           `json:"max_active" yaml:"max_active"`
	IdleTimeout time.Duration `json:"idle_timeout" yaml:"idle_timeout"`
	Name        string        `json:"name" yaml:"name"`
}
