package sql

type ISql interface {
	Init(setting any)
}

type Database struct {
	MySQL *MySQLSetting `yaml:"MySQL,omitempty"`
}

func New(db Database) ISql {
	var sqlInstance ISql
	if db.MySQL != nil {
		sqlInstance = &MySQL{}
		sqlInstance.Init(db.MySQL)
	} else {
		sqlInstance = nil
	}

	return sqlInstance
}
