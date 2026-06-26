package migrator

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
)

func (m Migrator) Down() {
	err := m.migrate.Down()
	if err != nil {
		if err == migrate.ErrNoChange {
			fmt.Println("No migrations to roll back")
			return
		}
		panic(fmt.Errorf("error in down migration : %v", err))
	}
	fmt.Println("Migration down completed successfully")
}