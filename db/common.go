package db

import (
	"database/sql"
	"fmt"
	"github.com/devchrisar/uptrend95User/models"
	"github.com/devchrisar/uptrend95User/secretm"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

var SecretModel models.SecretRDSJson
var err error
var Db *sql.DB

func GetSecret() error {
	SecretModel, err = secretm.GetSecret(os.Getenv("SecretName"))
	return err
}

func Connect() error {
	Db, err = sql.Open("mysql", ConnectStr(SecretModel))
	if err != nil {
		fmt.Println(" > Connect() error: " + err.Error())
		return err
	}
	err = Db.Ping()
	if err != nil {
		fmt.Println(" > Connect() error: " + err.Error())
		return err
	}
	fmt.Println(" > Connect() successful")
	return nil
}

func ConnectStr(keys models.SecretRDSJson) string {
	var DbUser, AuthToken, DbEndpoint, DbName string
	DbUser = keys.Username
	AuthToken = keys.Password
	DbEndpoint = keys.Host
	DbName = "uptrend95"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswords=true", DbUser, AuthToken, DbEndpoint, DbName)
	fmt.Println(" > ConnectStr() returning: " + dsn)
	return dsn
}
