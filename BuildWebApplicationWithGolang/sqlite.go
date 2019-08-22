// https://astaxie.gitbooks.io/build-web-application-with-golang/content/ja/05.3.html
/*

First of all, `sqlite3 sample.sqlite` and create the below databases:
CREATE TABLE `userinfo` (
    `uid` INTEGER PRIMARY KEY AUTOINCREMENT,
    `username` VARCHAR(64) NULL,
    `departname` VARCHAR(64) NULL,
    `created` DATE NULL
);

CREATE TABLE `userdeatail` (
    `uid` INT(10) NULL,
    `intro` TEXT NULL,
    `profile` TEXT NULL,
    PRIMARY KEY (`uid`)
);
*/
package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./sample.sqlite")
	if err != nil {
		log.Fatalf("can't open database: %v", err)
	}
	defer db.Close()

	// insert
	stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")
	if err != nil {
		log.Fatalf("can't prepare statement: %v", err)
	}

	res, err := stmt.Exec("astaxie", "研究開発部門", "2012-12-09")
	if err != nil {
		log.Fatalf("can't execute statement: %v", err)
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		log.Fatalf("can't fetch LastInsertId: %v", err)
	}

	fmt.Println(lastID)

	// update
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	if err != nil {
		log.Fatalf("can't prepare statement: %v", err)
	}

	res, err = stmt.Exec("astaxieupdate", lastID)
	if err != nil {
		log.Fatalf("can't execute statement: %v", err)
	}

	affect, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("can't fetch affected row: %v", err)
	}

	fmt.Println(affect)

	// search
	rows, err := db.Query("SELECT * FROM userinfo")
	if err != nil {
		log.Fatalf("can't query: %v", err)
	}

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created time.Time
		err = rows.Scan(&uid, &username, &department, &created)

		if err != nil {
			log.Fatalf("can't scan row: %v", err)
		}

		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	// delete
	stmt, err = db.Prepare("delete from userinfo where uid=?")
	if err != nil {
		log.Fatalf("can't prepare statement: %v", err)
	}

	res, err = stmt.Exec(lastID)
	if err != nil {
		log.Fatalf("can't execute statement: %v", err)
	}

	affect, err = res.RowsAffected()
	if err != nil {
		log.Fatalf("can't fetch affected row: %v", err)
	}

	fmt.Println(affect)
}
