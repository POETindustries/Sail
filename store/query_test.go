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
	Add().In("test").Exec()
}

func TestQueryGet(t *testing.T) {
	Get().In("test").Exec()
	Get().In("test").All().Exec()
}

func TestQueryUpdate(t *testing.T) {
	Update().In("test").Exec()
}

func TestQueryDelete(t *testing.T) {
	Delete().In("test").Exec()
}
