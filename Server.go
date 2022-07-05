package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/wangyuche/cloudbilling/src/sql"
	yaml "gopkg.in/yaml.v3"
)

func main() {
	s := ParserSetting(os.Getenv("SettingPath"))
	sql.New(s.Data.Database)
}

type Setting struct {
	Data struct {
		Database sql.Database `yaml:"Database,omitempty"`
	} `yaml:"Data,omitempty"`
}

func ParserSetting(filepath string) Setting {
	config, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	setting := Setting{}
	err = yaml.Unmarshal(config, &setting)
	if err != nil {
		log.Fatal(err)
	}
	return setting
}
