package config

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var db *sql.DB
var once sync.Once

func DbConn() *sql.DB {
	once.Do(func() {
		var err error
		env := GetEnv()
		if db, err = sql.Open("postgres", env.DATABASE_URL); err != nil {
			log.Panic(err)
		}
		db.SetMaxOpenConns(100)
		db.SetMaxIdleConns(20)
	})
	return db
}

func DbClose() {
	db.Close()
}
