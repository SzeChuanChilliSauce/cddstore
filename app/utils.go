package app

import (
	"flag"
	"github.com/tendermint/tendermint/config"
	tmdb "github.com/tendermint/tm-db"
	"path/filepath"
)
var (
	home = flag.String("home", ".", "Dir in ")
)
func InitConfig() *config.Config {
	return &config.Config{
		BaseConfig: config.BaseConfig{RootDir: *home},
	}
}

func CreateDB(cfg *config.Config) tmdb.DB {
	return tmdb.NewDB("store_app", tmdb.GoLevelDBBackend, filepath.Join(cfg.RootDir, "data"))
}
