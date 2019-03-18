package lmdb

import (
	"go-study/src/log"
	"go.uber.org/zap"
)

func DemoLmdb() {
	log.LmdbLogger.Info("------lmdb---------")
	size := int64(102400)
	path := "F:\\lmdb"
	name := "db"
	client := NewClient(size, path, name)
	flag := client.Set("sex", "nan")
	log.LmdbLogger.Info("Set", zap.Bool("flag", flag))
	sex := client.Get("sex")
	log.LmdbLogger.Info(sex)
	log.LmdbLogger.Info(client.Stat())
}
