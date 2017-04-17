package main

import (
	"database/sql"
	"gopkg.in/yaml.v2"
	_ "github.com/go-sql-driver/mysql"
	//sq "github.com/Masterminds/squirrel"
	"io/ioutil"
	"github.com/pkg/errors"
	"fmt"
)

type settings struct {
	mysql struct {
		host string
		db string
		user string
		pass string
	}
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

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", s.mysql.user, s.mysql.pass, s.mysql.host))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//rows, err := db.Query("SELECT * FROM users")
	//if err != nil {
	//	panic(err.Error())
	//}
}
