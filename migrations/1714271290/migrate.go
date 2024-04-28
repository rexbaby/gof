package migrations1714271290

import (
	_ "embed"

	"gofr.dev/pkg/gofr/migration"
)

//go:embed create_table_departments.sql
var createDepartmentsSQL string

//go:embed create_table_districts.sql
var createDistrictsSQL string

//go:embed create_table_ranks.sql
var createRanksSQL string

//go:embed create_table_staffs.sql
var createStaffsSQL string

func Migrate() migration.Migrate {
	return migration.Migrate{
		UP: func(d migration.Datasource) error {

			_, err := d.SQL.Exec(createDepartmentsSQL)
			if err != nil {
				return err
			}

			_, err = d.SQL.Exec(createDistrictsSQL)
			if err != nil {
				return err
			}

			_, err = d.SQL.Exec(createRanksSQL)
			if err != nil {
				return err
			}

			_, err = d.SQL.Exec(createStaffsSQL)
			if err != nil {
				return err
			}

			return nil
		},
	}
}
