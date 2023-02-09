package global

type Config struct {
	Serverini serverini `mapstructure:"serverini" json:"serverini" toml:"serverini" yaml:"serverini"`
	App       App       `mapstructure:"app" json:"app" toml:"app" yaml:"app"`
	Log       Log       `mapstructure:"log" json:"log" toml:"log" yaml:"log"`
}
