package config

import "common/utils"

func ConsulInit() {
	go utils.RegisterService("member", "localhost", 50051)
}
