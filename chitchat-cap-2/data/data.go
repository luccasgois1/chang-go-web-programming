package data

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	_ "github.com/lib/pq"
)

type ConfigDB struct {
	Host     string
	User     string
	Password string
	Name     string
	Type     string
}

var configDB ConfigDB
var Db *sql.DB
var pathConfigFile string

// This function initialized the Db variable when the application starts
func init() {
	loadConfigDB()
	var err error
	Db, err = sql.Open(configDB.Type, fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		configDB.Host,
		configDB.User,
		configDB.Password,
		configDB.Name))
	if err != nil {
		log.Fatalln("unable to stablish connection to db:", err)
	}
}

// Load the configurations data for the DB
func loadConfigDB() {
	getPathConfigFile()
	file, err := os.Open(pathConfigFile)
	if err != nil {
		log.Fatalln("unable to load:", pathConfigFile, err)
	}
	decoder := json.NewDecoder(file)
	configDB = ConfigDB{}
	err = decoder.Decode(&configDB)
	if err != nil {
		log.Fatalln("unable to decode:", pathConfigFile, err)
	}

	overwriteConfigDBWithEnvVariables()
}

func getPathConfigFile() {
	// This gets the path of the current file
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatalln("not able to get executable path")
	}
	execDir := filepath.Dir(filename)
	pathConfigFile = filepath.Join(execDir, "config_db.json")
}

// Overwrites the default value if environment variables are given
func overwriteConfigDBWithEnvVariables() {
	if os.Getenv("POSTGRES_HOST") != "" {
		configDB.Host = os.Getenv("POSTGRES_HOST")
	}
	if os.Getenv("POSTGRES_USER") != "" {
		configDB.User = os.Getenv("POSTGRES_USER")
	}
	if os.Getenv("POSTGRES_PASSWORD") != "" {
		configDB.Password = os.Getenv("POSTGRES_PASSWORD")
	}
}
