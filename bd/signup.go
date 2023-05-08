package bd

import (
	"fmt"

	"github.com/fabiorodolfo84/gambituser/models"
	"github.com/fabiorodolfo84/gambituser/tools"
	_ "github.com/go-sql-driver/mysql"
)

func Signup(sig models.Signup) error {
	fmt.Println("Comienza Registro")
	err := DbConnect()
	if err != nil {
		return err
	}

	defer Db.Close()

	sentencia := "INSERT INTO users (User_Email, User_UUID, User_DateAdd) VALUES ('" + sig.UserEmail + "','" + sig.UserUUID + "','" + tools.FechaMySQL() + "')"
	fmt.Println(sentencia)
	_, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("Signup > Ejecution Exitosa")
	return nil

}
