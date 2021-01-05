package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"userRepository/internal/vipers"
)

//Datastore holds Database connection and Methods
type Datastore  struct {
	Db *sqlx.DB
}

//initialize Datastore with Db
func DbConnect() (*Datastore,error) {
	//get configs for database connection
	dbConf, err := vipers.GetDbconfigs()
	if err != nil {
		log.Fatal(err)
	}
	dbInstance := &Datastore{}
	dbInstance.Db, err = sqlx.Connect(dbConf.Drivername, dbConf.Username+":"+dbConf.Password+"@tcp("+dbConf.Host+":"+dbConf.Port+")/"+dbConf.DbName)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	log.Println("Database : connected successfully")
	return dbInstance, nil
}