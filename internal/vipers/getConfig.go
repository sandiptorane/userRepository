package vipers

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func GetPort()(port string) {
	log.Println("fetching port configs")
	v := viper.New()
	v.SetConfigName("config") // config file name
	v.SetConfigType("yaml")
	v.AddConfigPath("./configs") // config file path
	v.AddConfigPath("../configs")
	v.AddConfigPath(".")

	err := v.ReadInConfig()
	if err != nil {
		log.Fatal("Unable to fetch port")
	}
	port = v.GetString("userrepo.port")
	return port
}

type DbConfig struct {
	Host string
	Drivername string
	Username string
	Password string
	DbName string
	Port string
}

func GetDbconfigs() (DbConfig,error){
	// Config
	log.Println("fetching db configs")
	v := viper.New()
	v.SetConfigName("config") // config file name
	v.SetConfigType("yaml")
	v.AddConfigPath("./configs") // config file path
	v.AddConfigPath("../configs")
	v.AddConfigPath(".")
	err := v.ReadInConfig()
	if err != nil {
		return DbConfig{},err
	}

	drivername := v.GetString("db.drivername")
	username := v.GetString("db.username")
	password := v.GetString("db.password")
	host := v.GetString("db.host")
	port :=v.GetString("db.port")
	dbname := v.GetString("db.dbName")

	return DbConfig{
		Host : host,
		Port : port,
		Drivername: drivername,
		Username: username,
		Password: password,
		DbName: dbname,
	},nil
}
