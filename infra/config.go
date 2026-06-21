package infra

import (
	"github.com/Rian-rgb/ewallet-common-lib/config"
)

func InitConfig() {
	config.LoadEnv()
}
