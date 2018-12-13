package setting

import (
	"fmt"
	"gopkg.in/ini.v1"
	"gopkg.in/macaron.v1"
	"os"
)

var (
	// Global settings.
	Cfg *ini.File
)

func init() {
	iniPath := "conf/app.ini"
	envExist := PathExist("conf/app.ini.env")
	if envExist {
		iniPath = "conf/app.ini.env"
	}

	sources := []interface{}{iniPath}

	var err error
	Cfg, err = macaron.SetConfig(sources[0], sources[1:]...)

	if err != nil {
		//log.FatalD(4, "Fail to set configuration: %v", err)
		fmt.Println("fail", err)
	}

}

func PathExist(_path string) bool {
	_, err := os.Stat(_path)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}
