package migration

import (
	"context"
	"database/sql"
	"fmt"
	"godeploy/pkg/config"

	migrate "github.com/rubenv/sql-migrate"
	"github.com/sirupsen/logrus"
	_ "modernc.org/sqlite"
)

var db *sql.DB

type Migration struct {
	conf       *config.Config
	folder     string
	migrations *migrate.FileMigrationSource
}

func New(conf *config.Config, folder string) *Migration {
	return &Migration{
		conf:   conf,
		folder: folder,
	}
}

func (m *Migration) Up(ctx context.Context) error {
	migrate.SetTable("migrations")
	m.migrations = &migrate.FileMigrationSource{
		Dir: m.folder,
	}
	logrus.WithField("folder", m.folder).Info("migrate database")

	db, err := sql.Open("sqlite", fmt.Sprintf("file:%s", m.conf.DatabasePath))
	if err != nil {
		return err
	}

	_, err = migrate.ExecMax(db, "SQLITE", m.migrations, migrate.Up, -1)

	db.Close()
	return err
}
