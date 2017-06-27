package config

import (
	"asana/utils"

	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v1"
)

type Conf struct {
	Personal_access_token string
	Workspace             int
}

func ConfigFile() string {
	os.Mkdir(utils.Home()+"/.asana", 0700)
	return utils.Home() + "/.asana/config.yml"
}

func Load() Conf {
	var dat []byte
	var err error
	dat, err = ioutil.ReadFile(ConfigFile())
	if err != nil {
		fmt.Println("Config file isn't set.\n  ==> $ asana config")
		os.Exit(1)
	}
	conf := Conf{}
	err = yaml.Unmarshal(dat, &conf)
	utils.Check(err)
	return conf
}
