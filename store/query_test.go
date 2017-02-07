package store

import (
	"sail/conf"
	"testing"
)

func TestQuery(t *testing.T) {
	conf.Instance().DBName = "sl_test"

	for _, d := range []string{"sqlite3", "mysql", "postgres"} {
		conf.Instance().DBDriver = d
		if ok := DB().init(); !ok {
			t.Error(d + " init failed")
			continue
		}
		TestQueryAdd(t)
		TestQueryGet(t)
		TestQueryUpdate(t)
		TestQueryDelete(t)
	}
}

func TestQueryAdd(t *testing.T) {
	Add().In("sl_test").Exec()
}

func TestQueryGet(t *testing.T) {
	Get().In("sl_test").Exec()
	Get().In("sl_test").All().Exec()
}

func TestQueryUpdate(t *testing.T) {
	Update().In("sl_test").Exec()
}

func TestQueryDelete(t *testing.T) {
	Delete().In("sl_test").Exec()
}
