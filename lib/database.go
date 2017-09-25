package lib

import (
	// "database/sql"
	"fmt"
	// _ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetDB() *gorm.DB {
	// db, err := gorm.Open("mysql", "bilimger:bilimger@/bilimger?charset=utf8&parseTime=True&loc=Local")
	db, err := gorm.Open("mysql", "bilimger:bilimger@/bilimger")
	// db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	if err != nil {
		fmt.Printf("error %v\n", err)
		panic("failed to connect database")
	}

	// Disable table name's pluralization globally
	db.SingularTable(true) // if set this to true, `User`'s default table name will be `user`, table name setted with `TableName` won't be affected

	return db
}

func GetTable(tableName string) *gorm.DB {
	db := GetDB()
	db.Select(tableName)
	return db
}

/*
func ExcuteSql(sqlString string) *sql.DB {
	db, err := sql.Open("mysql", "bilimger:bilimger@/bilimger")
	if err != nil {
		fmt.Printf("error %v\n", err)
		panic("failed to connect database")
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	rows, err := db.Query("SELECT email FROM user WHERE id=?", 1)
	if err != nil {
		fmt.Printf("error %v\n", err)
	}
	defer rows.Close()

	for rows.Next() {
		var email string
		if err := rows.Scan(&email); err != nil {
			fmt.Printf("error %v\n", err)
		}
		fmt.Printf("%s ----------------------------------------------------------- email\n", email)
	}
	if err := rows.Err(); err != nil {
		fmt.Printf("error %v\n", err)
	}

	return db
}
*/
