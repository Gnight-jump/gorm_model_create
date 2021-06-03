/**
配置文件解析
*/
package conf

type Conf struct {
	Dbase       Database `yaml:"Database""`
	TargetModel GModel   `yaml:"GModel"`
}

// 数据库配置
type Database struct {
	Host     string `yaml:"Host"`
	Port     int    `yaml:"Port"`
	User     string `yaml:"User"`
	Password string `yaml:"Password"`
	DBName   string `yaml:"DBName"`
}

// 目标表单数据
type GModel struct {
	StorePath   string   `yaml:"StorePath"`
	PackageName string   `yaml:"PackageName"`
	ModelCover  bool     `yaml:"ModelCover"`
	TableName   []string `yaml:"TableName"`
}
