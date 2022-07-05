package sql

import (
	"database/sql"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLSetting struct {
	Account    string `yaml:"Account,omitempty"`
	Password   string `yaml:"Password,omitempty"`
	Database   string `yaml:"Database,omitempty"`
	URL        string `yaml:"URL,omitempty"`
	Connection int    `yaml:"Connection,omitempty"`
}

type MySQL struct {
	DB   *sql.DB
	Once sync.Once
}

func (this *MySQL) Init(setting any) {
	s := setting.(*MySQLSetting)
	CreateDatabase(s)
	this.Once.Do(func() {
		db, err := sql.Open("mysql", s.Account+":"+s.Password+"@tcp("+s.URL+")/"+s.Database+"?parseTime=true")
		if err != nil {
			panic(err.Error())
		}
		err = db.Ping()
		if err != nil {
			panic(err.Error())
		}
		db.SetMaxOpenConns(s.Connection)
		db.SetMaxIdleConns(s.Connection)
		db.SetConnMaxLifetime(0)
		this.DB = db
	})
	CreateGroupsTable(this.DB)
	CreateProjectsTable(this.DB)
	CreateBillingInfoTable(this.DB)
}

func CreateDatabase(s *MySQLSetting) {
	db, err := sql.Open("mysql", s.Account+":"+s.Password+"@tcp("+s.URL+")/")
	if err != nil {
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)
	db.SetConnMaxLifetime(0)
	sql := `create database if not exists ` + s.Database + ` default character set utf8mb4 collate utf8mb4_general_ci;`
	tx, err := db.Begin()
	if err != nil {
		panic(err.Error())
	}
	_, err = tx.Exec(sql)
	if err != nil {
		panic(err.Error())
	}
	tx.Commit()
}

func CreateGroupsTable(db *sql.DB) {
	sql := `create table if not exists groups (
		Id INT NOT NULL AUTO_INCREMENT,
		Description VARCHAR(100) NULL,
		PRIMARY KEY (Id));`
	tx, err := db.Begin()
	if err != nil {
		panic(err.Error())
	}
	defer tx.Commit()
	_, err = tx.Exec(sql)
	if err != nil {
		panic(err.Error())
	}
}

func CreateProjectsTable(db *sql.DB) {
	sql := `create table if not exists projects (
			Id INT NOT NULL AUTO_INCREMENT,
			ProjectName VARCHAR(100) NULL,
			ProjectID VARCHAR(100) NULL,
			Groups JSON NULL,
			PRIMARY KEY (Id));`
	tx, err := db.Begin()
	if err != nil {
		panic(err.Error())
	}
	defer tx.Commit()
	_, err = tx.Exec(sql)
	if err != nil {
		panic(err.Error())
	}
}

func CreateBillingInfoTable(db *sql.DB) {
	sql := `create table if not exists billinginfo (
		Id INT NOT NULL AUTO_INCREMENT,
		Idprojects INT NULL,
		Time TIMESTAMP NULL,
		Info JSON NULL,
		Price INT NULL,
		PRIMARY KEY (Id));`
	tx, err := db.Begin()
	if err != nil {
		panic(err.Error())
	}
	defer tx.Commit()
	_, err = tx.Exec(sql)
	if err != nil {
		panic(err.Error())
	}
}
