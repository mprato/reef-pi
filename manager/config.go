package manager

import (
	"github.com/reef-pi/reef-pi/controller/storage"
	"github.com/reef-pi/reef-pi/controller/utils"
	"log"
)

const Bucket = "config"

type Config struct {
	Creds        utils.Credentials `json:"creds"`
	Address      string            `json:"address"`
	HTTPS        bool              `json:"https"`
	Notification bool              `json:"notification"`
}

var DefaultConfig = Config{
	Creds: utils.Credentials{
		User:     "reef-pi",
		Password: "reef-pi",
	},
	Address: ":8088",
}

type Error struct {
	Message string `json:"message"`
	ID      string `json:"id"`
	Time    string `json:"time"`
}

const ConfigKey = "config"
const TelemetryKey = "telemetry"

func loadConfiguration(store storage.Store) (Config, error) {
	var c Config
	return c, store.Get(Bucket, ConfigKey, &c)
}

func initializeConfiguration(store storage.Store) (Config, error) {
	if err := store.CreateBucket(Bucket); err != nil {
		log.Println("ERROR:Failed to create bucket:", Bucket, ". Error:", err)
		return DefaultConfig, err
	}
	if err := store.Update(Bucket, "credentials", DefaultConfig.Creds); err != nil {
		return DefaultConfig, err
	}
	return DefaultConfig, store.Update(Bucket, ConfigKey, DefaultConfig)
}
