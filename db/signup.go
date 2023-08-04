package db

import (
	"fmt"
	"github.com/devchrisar/uptrend95User/models"
	"github.com/devchrisar/uptrend95User/tools"
	_ "github.com/go-sql-driver/mysql"
)

func SignUp(sign models.Signup) error {
	fmt.Println(" > SignUp() called")
	err := Connect()
	if err != nil {
		return err
	}
	defer Db.Close()
	query := fmt.Sprintf(`
    INSERT INTO users (User_Email, User_UUID, User_DateAdd) VALUES ('%v' , '%v', '%v')`,
		sign.UserEmail,
		sign.UserUUID,
		tools.MySqlTimestamp(),
	)
	fmt.Println(" > SignUp() query: " + query)
	_, err = Db.Exec(query)
	if err != nil {
		fmt.Println(" > SignUp() error: " + err.Error())
		return err
	}
	fmt.Println(" > SignUp() successful")
	return nil
}
