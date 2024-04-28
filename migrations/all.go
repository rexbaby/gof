package migrations

import (
	migrations1714271290 "gof/migrations/1714271290"

	"gofr.dev/pkg/gofr/migration"
)

var migrateAll = map[int64]migration.Migrate{
	1714271290: migrations1714271290.Migrate(),
}

func All() map[int64]migration.Migrate {
	return migrateAll
}
