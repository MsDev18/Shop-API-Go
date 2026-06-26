package migrator

import (
	"fmt"
	"strings"

	"github.com/golang-migrate/migrate/v4"
)

func (m Migrator) Step(step int) {
	err := m.migrate.Steps(step)
	if err != nil {
		switch {
		case err == migrate.ErrNoChange:
			fmt.Println("no change in migration")
			return
		case strings.Contains(err.Error(), "limit") && strings.Contains(err.Error(), "short"):
			fmt.Println("limit reached")
			return
		case strings.Contains(err.Error(), "file does not exist"):
			fmt.Println("version not exist")
			return
		}
		fmt.Println("error in step migration : ", err)
	}
}
