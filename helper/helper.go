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
	return path.Join(GetHomeDir(), ".config", "blog-in-your-email", constant.DATA_FILE)
}

func GetIniFile() string {
	return path.Join(GetHomeDir(), ".config", "blog-in-your-email", constant.INI_FILE)
}

func GetDBFile() string {
	return path.Join(GetHomeDir(), ".config", "blog-in-your-email", constant.DB_FILE)
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

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func IsFirstRun() bool {
	return !FileExists(GetDBFile())
}
