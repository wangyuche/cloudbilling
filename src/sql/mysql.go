package sql

type MySQLSetting struct {
	Account  string `yaml:"Account,omitempty"`
	Password string `yaml:"Password,omitempty"`
	Database string `yaml:"Database,omitempty"`
}

type MySQL struct {
}

func (this *MySQL) Init() {

}
