package composites

import (
	"github.com/IvSen/shareThings/pkg/client/postgresql"
	"github.com/IvSen/shareThings/pkg/config"
	"gorm.io/gorm"
)

type PgClientComposite struct {
	Db *gorm.DB
}

func NewPgClientComposite(config *config.Config) (PgClientComposite, error) {
	pgConfig := postgresql.NewPgConfig(
		config.PostgreSQL.Username, config.PostgreSQL.Password,
		config.PostgreSQL.Host, config.PostgreSQL.Port, config.PostgreSQL.Database,
	)
	pgClient, err := postgresql.NewClient(pgConfig)
	return PgClientComposite{Db: pgClient}, err
}
