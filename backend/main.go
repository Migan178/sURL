package main

import (
	"fmt"

	"github.com/Migan178/surl/configs"
)

func main() {
	if err := r.Run(fmt.Sprintf(":%d", configs.GetConfigs().Port)); err != nil {
		panic(err)
	}
}
