package main

import (
	"log"

	"github.com/hamza72x/blog-in-your-email/data"
	"github.com/hamza72x/blog-in-your-email/feed"
	"github.com/hamza72x/blog-in-your-email/helper"
	"github.com/hamza72x/blog-in-your-email/mail"
)

func main() {

	checkConfigs()

	blogs := data.GetBlogDataFromCSV()
	isFirstRun := helper.IsFirstRun()

	for i := range blogs {
		feed.CheckBlogFeed(blogs[i])
	}

	if isFirstRun {
		mail.SendWelcomeEmail()
	}
}

func checkConfigs() {

	ini := helper.GetIni()
	iniFile := helper.GetIniFile()

	if len(ini.RECEIVER_EMAIL) == 0 {
		log.Printf("Please set RECEIVER_EMAIL in %s\n", iniFile)
		panic("Please set RECEIVER_EMAIL in " + iniFile)
	}

	if len(ini.SENDER_EMAIL) == 0 {
		log.Printf("Please set SENDER_EMAIL in %s\n", iniFile)
		panic("Please set SENDER_EMAIL in " + iniFile)
	}

	if len(ini.SENDER_PASSWORD) == 0 {
		log.Printf("Please set SENDER_PASSWORD in %s\n", iniFile)
		panic("Please set SENDER_PASSWORD in " + iniFile)
	}

	// check data file exists or not
	if !helper.FileExists(helper.GetDataFile()) {
		log.Printf("Data file %s not found\n", helper.GetDataFile())
		panic("Data file " + helper.GetDataFile() + " not found")
	}
}
