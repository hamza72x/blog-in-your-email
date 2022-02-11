package helper

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/hamza72x/blog-in-your-email/constant"
	"github.com/hamza72x/blog-in-your-email/model"
	"gopkg.in/ini.v1"
)

// panics if fails
func GetHomeDir() string {
	dir, err := os.UserHomeDir()
	if err != nil {
		log.Println("Error getting user home dir")
		panic(err)
	}
	return dir
}

func GetDataFile() string {
	return path.Join(GetHomeDir(), ".config", constant.DATA_FILE)
}

func GetIniFile() string {
	return path.Join(GetHomeDir(), ".config", constant.INI_FILE)
}

func GetDBFile() string {
	return path.Join(GetHomeDir(), ".config", constant.DB_FILE)
}

func GetIni() model.INI {

	cfg, err := ini.Load(GetIniFile())

	if err != nil {
		fmt.Println("Error loading ini file")
		panic(err)
	}

	ini := model.INI{}

	err = cfg.MapTo(&ini)

	if err != nil {
		fmt.Println("Error mapping ini file")
		panic(err)
	}

	return ini
}