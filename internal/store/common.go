package store

import (
	"github.com/fennay/tgo/internal/utils/config"
)

func Store() *config.DB {

	return &config.GetConfig().DB
}
