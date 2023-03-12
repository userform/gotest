package conf_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/userform/gotest/conf"
)

func TestLoadConfigFromToml(t *testing.T) {
	should := assert.New(t)
	err := conf.LoadConfigFromToml("../etc/demo.toml")
	if should.NoError(err) {
		should.Equal("demo1",conf.C().App.Name)
		// fmt.Println(conf.C().App.Name)
	}
}

func TestLoadConfigFromEnv(t *testing.T) {
	should := assert.New(t)

	os.Setenv("MYSQL_DATABASE","unit_test")
	err := conf.LoadConfigFromEnv()
	if should.NoError(err) {
		should.Equal("unit_test",conf.C().MySQL.Database)
		// fmt.Println(conf.C().MySQL.Database)
	}
}
