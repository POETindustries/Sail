package store

import "fmt"

type ErrNoDriver struct {
	driver string
}

func (e ErrNoDriver) Error() string {
	return fmt.Sprintf("%s is not a supported database driver", e.driver)
}
