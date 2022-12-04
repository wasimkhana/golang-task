package config

import "github.com/spf13/viper"

// keys for database configuration.
const (
	DbName = "db.name"
	DbPath = "db.path"
)

func init() {
	// env var for db
	_ = viper.BindEnv(DbName, "DB_NAME")
	_ = viper.BindEnv(DbPath, "DB_PATH")

	// defaults
	viper.SetDefault(DbName, "sqlite3")
	viper.SetDefault(DbPath, "../database.sqlite")
}
