package config

// Etcd etcd配置
type Etcd struct {
	Enabled   bool     `mapstructure:"enabled" json:"enabled" yaml:"enabled"`                         // 是否启用 etcd
	Endpoints []string `mapstructure:"endpoints" json:"endpoints" yaml:"endpoints"`                 // etcd 集群地址
	Prefix    string   `mapstructure:"prefix" json:"prefix" yaml:"prefix"`                             // key 前缀
	Username  string   `mapstructure:"username" json:"username" yaml:"username"`                       // 用户名
	Password  string   `mapstructure:"password" json:"password" yaml:"password"`                       // 密码
	Timeout   int      `mapstructure:"timeout" json:"timeout" yaml:"timeout"`                           // 连接超时(秒)
}
