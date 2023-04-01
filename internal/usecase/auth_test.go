package usecase

import (
	"checkwork/config"
	"checkwork/internal/repository"
	"log"
	"testing"
)

func Test_CreateUser(t *testing.T) {
	tests := []struct {
		name     string
		username string
		password string
		wantErr  bool
	}{
		{
			name:     "OK",
			username: "admin",
			password: "admin",
			wantErr:  false,
		},
		{
			name:     "Duplicate",
			username: "admin",
			password: "admin",
			wantErr:  true,
		},
	}

	cfg := config.New()
	cfg.DBConfig.DataSourceCred = "test"
	storage, err := repository.Init(cfg.DBConfig)
	if err != nil {
		log.Fatalf("Failed to initialize: %s", err.Error())
	}

	logic := New(storage)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err = logic.CreateUser(test.username, test.password)
			if test.wantErr && err == nil {
				t.Fail()
			}
		})
	}
}
