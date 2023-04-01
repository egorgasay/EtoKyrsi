package repository

import (
	"checkwork/internal/entity"
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type Config struct {
	DriverName     string
	DataSourceCred string
}

type IStorage interface {
	Disconnect() error
	DeleteAccount() error
	CreateUser(username, password string) error
	UpdateUserScore(student string) error
	GetUsers() ([]entity.User, error)

	CheckPassword(username, password string) (bool, error)

	//ChangeNick(username, nick string) error
	ChangePassword(username, oldPassword, newPassword string) error

	SetPending(username string, status int) error
	CheckIsPending(username string) (bool, error)

	GetTaskIDAndMsg(username string) (int, sql.NullString, error)
	GetWorks() ([]Work, error)

	SetVerdict(student, verdict string) error
	DeletePullRequest(student string) error
	AddPullRequest(link, student string) error

	UpdateTask(num int, title string) error
	DeleteTask(num int) error
	GetTasks() ([]entity.Task, error)
	GetTitle(number int) (string, error)
}

type Storage struct {
	DB *sql.DB
}

func Init(cfg *Config) (IStorage, error) {
	if cfg == nil {
		panic("конфигурация задана некорректно")
	}

	db, err := sql.Open(cfg.DriverName, cfg.DataSourceCred)
	if err != nil {
		return nil, err
	}

	return New(db, "file://internal/repository/migrations"), nil
}

func New(db *sql.DB, pathToMigrations string) IStorage {
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		log.Fatal(err)
		return nil
	}

	m, err := migrate.NewWithDatabaseInstance(
		pathToMigrations,
		"gomarket", driver)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	err = m.Up()
	if err != nil {
		if err.Error() != "no change" {
			log.Fatal(err)
		}
	}

	return Storage{db}
}

func (s Storage) Disconnect() error {
	return nil
}

func (s Storage) DeleteAccount() error {
	return nil
}
