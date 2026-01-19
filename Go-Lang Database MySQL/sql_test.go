package GolangDBMySQL

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestExecSql(t *testing.T){
	db:=GetConnection()
	defer db.Close()

	ctx := context.Background()
	scriptsql:="INSERT INTO customer(id, name) VALUES('c003', 'Keshinryan')"

	_, err := db.ExecContext(ctx, scriptsql)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success insert new customer")
}

func TestQuerySql(t *testing.T){
	db:=GetConnection()
	defer db.Close()

	ctx := context.Background()
	scriptsql:="SELECT id, name FROM customer"

	rows, err := db.QueryContext(ctx, scriptsql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		err = rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("ID:", id)
		fmt.Println("Name:", name)
	}
}

func TestQuerySqlComplex(t *testing.T){
	db:=GetConnection()
	defer db.Close()

	ctx := context.Background()
	scriptsql:="SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"

	rows, err := db.QueryContext(ctx, scriptsql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var birthDate sql.NullTime 
		var createdAt time.Time
		var married bool
		err = rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}
		fmt.Println("==================================")
		fmt.Println("ID:", id)
		fmt.Println("Name:", name)
		if email.Valid {
			fmt.Println("Email:", email.String)
		}
		fmt.Println("Balance:", balance)
		fmt.Println("Rating:", rating)
		if birthDate.Valid{
			fmt.Println("Birth Date:", birthDate.Time)
		}
		fmt.Println("Marrried:", married)
		fmt.Println("Created At:", createdAt)
		fmt.Println("==================================")
	}
}

func TestSqlInjection(t *testing.T){
	db:=GetConnection()
	defer db.Close()

	ctx := context.Background()
	username:="admin'; #"
	password:="salah"

	scriptsql:="SELECT username FROM user WHERE username='" + username + "' AND password='" + password + "' LIMIT 1 ;"
	rows, err := db.QueryContext(ctx, scriptsql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Login Success", username)
	}else {
		fmt.Println("Login Failed")
	}
}

func TestSqlInjectionSafe(t *testing.T){
	db:=GetConnection()
	defer db.Close()

	ctx := context.Background()
	username:="admin"
	password:="admin"

	scriptsql:="SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1 "
	rows, err := db.QueryContext(ctx, scriptsql, username, password)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if rows.Next() {
		var username string
		rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Login Success", username)
	}else {
		fmt.Println("Login Failed")
	}
}

func TestExecSqlParam(t *testing.T){
	db:=GetConnection()
	defer db.Close()

	ctx := context.Background()
	username:="user01"
	password:="user01"

	scriptsql:="INSERT INTO user(username, password) VALUES(?, ?)"

	_, err := db.ExecContext(ctx, scriptsql, username, password)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success insert new customer")
}

func TestAutoIncrement(t *testing.T){
	db:=GetConnection()
	defer db.Close()

	ctx := context.Background()

	email:="jason@gmail.com"
	comment:="Hello Jason Here"

	scriptsql:="INSERT INTO comments(email, comment) VALUES(?, ?)"

	result, err := db.ExecContext(ctx, scriptsql, email, comment)
	if err != nil {
		panic(err)
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println("Success insert new comment with id", lastInsertId)
}

func TestPrepareStatement(t *testing.T){
	db:=GetConnection()
	defer db.Close()

	ctx := context.Background()

	scriptsql:="INSERT INTO comments(email, comment) VALUES(?, ?)"

	stmt, err := db.PrepareContext(ctx,scriptsql)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	for i:=0; i<10; i++ {
		email:= "user" + strconv.Itoa(i) + "@gmail.com"
		comment	:= "Hello this is comment number " + strconv.Itoa(i)

		result, err := stmt.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}
		id , err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("Comment id", id, "success inserted")
	}
}

func TestTransaction(t *testing.T){
	db:=GetConnection()
	defer db.Close()

	ctx := context.Background()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	scriptsql:="INSERT INTO comments(email, comment) VALUES(?, ?)"
	for i:=0; i<10; i++ {
		email:= "user" + strconv.Itoa(i) + "@gmail.com"
		comment	:= "Hello this is comment number " + strconv.Itoa(i)

		result, err := tx.ExecContext(ctx, scriptsql ,email, comment)
		if err != nil {
			panic(err)
		}
		id , err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		fmt.Println("Comment id", id, "success inserted")
	}

	err = tx.Rollback()
	if err != nil {
		panic(err)
	}
}