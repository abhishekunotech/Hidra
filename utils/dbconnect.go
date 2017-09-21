package utils

import (
	"database/sql"
	"fmt"
//	"net/http"
	_ "github.com/go-sql-driver/mysql"
)

func DBConnect(){

	db, err := sql.Open("mysql", "root:redhat@/access_manager")
	if err != nil {
		fmt.Println(err.Error())
	}
        
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
	}
	//stmt, err := db.Prepare("INSERT Users SET id=?,username=?,password=?,component=?")
	
	rows, err := db.Query("SELECT * FROM Users")
	if err != nil {
                fmt.Println(err.Error())
        }
	 
	for rows.Next() {
            var id int
            var username string
            var password string
            var component string
	    var email string
            err = rows.Scan(&id, &username, &password, &component,&email)
            if err != nil {
                fmt.Println(err.Error())
	    }
	
	    fmt.Println(id)
            fmt.Println(username)
            fmt.Println(password)
            fmt.Println(component)
	    fmt.Println(email)
        }

}
