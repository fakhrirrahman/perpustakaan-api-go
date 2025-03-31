package connection

import (
	"database/sql"
	"fmt"
	"go-web-native/internal/config"
	"log"
)

func GetDatabase(conf config.Database)* sql.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=%s",
		conf.Host,
		conf.Port,
		conf.User,
		conf.Pass,
		conf.Name,
		conf.Tz,
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Error opening database:", err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to database:", err.Error())
	} 
	return db

	
}

