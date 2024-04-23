package initializer

import (
	"database/sql"
	"fmt"
	"log"
)

type ConnectDB struct {
	db *sql.DB
}

// instance of sql.DB
var DB *sql.DB

// in same package i can use type from different file.go
// for exmple => *ConfigDB which is located in loadEnv.go
func StartConnectDB(config *ConfigDB) {

	// HOW TO CONNECT TO MULTIPLE DATABASE in golang?
	// what PATTERN TO USE??

	// var DB *sql.DB
	var err error

	// PROBLEM ==> KNP ERROR SAAT PAKE APP.env ?????
	// TimeZone=Asia/Shanghai  ==> set waktu/jam saat ini
	fmt.Println("test: ", config.DBHost, config.DBPort, config.DBUsername, config.DBPassword, config.DBName)
	conn_credential := fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai`, config.DBHost, config.DBPort, config.DBUsername, config.DBPassword, config.DBName)
	// conn_credential := fmt.Sprintf(`host=%s port=%d user=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Shanghai`, host, port, user, password, dbname, sslmode_val)
	DB, err = sql.Open("postgres", conn_credential)

	// defer DB.Close()  // NOTE ===> CLOSE DISINI AKAN GAGALKAN SAAT CONTROLLER FUNCTION dipanggil

	if err != nil {
		// use log.Fatal() because this will
		// os.Exit()
		log.Fatal(err.Error(), "error connection to db")
	}
	err = DB.Ping()
	if err != nil {
		log.Fatal("ping connection is failed ", err)
	}

	fmt.Println("connection is started!")
}
