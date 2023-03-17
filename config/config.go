package config

import (
	"checkwork/internal/globals"
	"checkwork/internal/repository"
	"flag"
	"os"
)

const (
	defaultHost = ":80"
)

type Flag struct {
	host *string
	dsn  *string
	key  *string
}

var f Flag

func init() {
	f.host = flag.String("a", defaultHost, "-a=host")
	f.dsn = flag.String("d", "etokyrsi", "-d=connection_string")
	f.key = flag.String("k", "change_me", "-k=key")
}

type Config struct {
	DBConfig  *repository.Config
	MentorKey string
	Host      string
}

func New() *Config {
	if addr, ok := os.LookupEnv("RUN_ADDRESS"); ok {
		f.host = &addr
	}

	if dsn, ok := os.LookupEnv("DATABASE_URI"); ok {
		f.dsn = &dsn
	}

	if key, ok := os.LookupEnv("KEY"); ok {
		f.key = &key
	}

	globals.Secret = []byte(*f.key)
	return &Config{
		DBConfig: &repository.Config{
			DriverName:     "sqlite3",
			DataSourceCred: *f.dsn,
		},
		MentorKey: *f.key,
		Host:      *f.host,
	}
}
