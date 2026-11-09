package config

type Server struct {
	JWT       JWT     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap       Zap     `mapstructure:"zap" json:"zap" yaml:"zap"`
	Redis     Redis   `mapstructure:"redis" json:"redis" yaml:"redis"`
	RedisList []Redis `mapstructure:"redis-list" json:"redis-list" yaml:"redis-list"`
	System    System  `mapstructure:"system" json:"system" yaml:"system"`
	Captcha   Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	// gorm
	Pgsql  Pgsql           `mapstructure:"pgsql" json:"pgsql" yaml:"pgsql"`
	Sqlite Sqlite          `mapstructure:"sqlite" json:"sqlite" yaml:"sqlite"`
	DBList []SpecializedDB `mapstructure:"db-list" json:"db-list" yaml:"db-list"`
	// oss
	Local        Local        `mapstructure:"local" json:"local" yaml:"local"`
	Minio        Minio        `mapstructure:"minio" json:"minio" yaml:"minio"`

	DiskList []DiskList `mapstructure:"disk-list" json:"disk-list" yaml:"disk-list"`

	// 跨域配置
	Cors CORS `mapstructure:"cors" json:"cors" yaml:"cors"`
}
