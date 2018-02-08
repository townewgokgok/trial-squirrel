package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v2"
	//sq "github.com/Masterminds/squirrel"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
)

type settings struct {
	MySQL struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
		DB   string `yaml:"db"`
		User string `yaml:"user"`
		Pass string `yaml:"pass"`
	} `yaml:"mysql"`
}

func mustLoadSettings(s *settings) {
	data, err := ioutil.ReadFile("settings.yml")
	if err != nil {
		panic(errors.Wrap(err, "Failed to load settings"))
	}
	err = yaml.Unmarshal(data, s)
	if err != nil {
		panic(errors.Wrap(err, "Failed to unmarshall settings"))
	}
}

func main() {
	var s settings
	mustLoadSettings(&s)

	datasrc := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", s.MySQL.User, s.MySQL.Pass, s.MySQL.Host, s.MySQL.Port, s.MySQL.DB)
	//fmt.Printf("%s\n", datasrc)

	db, err := sql.Open("mysql", datasrc)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT archiveid, name FROM archive")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		id, name := 0, ""
		rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%5d: %s\n", id, name)
	}
}
