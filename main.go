package main

import (
	"fmt"

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

	punchErrs := nueip.Punch(hr_system.PunchOut)
	if punchErrs != nil {
		fmt.Println(punchErrs)
	}
}
