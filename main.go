package main

import (
	"github.com/imfulee/punch/hr_system"
)

func main() {
	LoadConfig()
	config := GetConfig()

	nueip := hr_system.NUEIP{
		Company:  config.company,
		Username: config.username,
		Password: config.password,
		URL:      config.url,
	}

	nueip.Punch(hr_system.PunchIn)
}
