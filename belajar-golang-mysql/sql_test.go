package belajar_golang_database

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	_ "time"
)
func TestExecSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()
	script := "INSERT INTO customer(id, nama) VALUES ('kambing','Kambing')"
	_,err := db.ExecContext(ctx,script)
	if err != nil {
		panic(err)
	}
	fmt.Println("berhasil insert customer baru")
}

func TestQuarySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()
	script := "SELECT id,nama,email,balance,rating,birth_date,married,created_at FROM customer"
	rows,err := db.QueryContext(ctx,script)
	if err != nil {
		panic(err)
	}
	for rows.Next(){
		var id,nama,email sql.NullString //sql.null mengembalikan tipe data dan valid misalnya String , Valid   
		var balance sql.NullInt32
		var rating sql.NullFloat64
		var birthDate,createdAt sql.NullTime
		var married sql.NullBool
		err := rows.Scan(&id,&nama,&email,&balance,&rating,&birthDate,&married,&createdAt)
		if err != nil{
			panic(err)
		}
		fmt.Println("===============")
		fmt.Println("ID         :", id.String)
		fmt.Println("Nama       :", nama)
		fmt.Println("Email      :", email)
		fmt.Println("Balance    :", balance)
		fmt.Println("Rating     :", rating)
		fmt.Println("Birth Date :", birthDate)
		fmt.Println("Married    :", married)
		fmt.Println("Created At :", createdAt)
		fmt.Println("===============")
		fmt.Println()

	}
	defer rows.Close()
	
}

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	username := "dimas"
	password := "rahasia123"
	ctx := context.Background()
	script := "SELECT username FROM user WHERE username = ? AND password = ?"
	rows,err := db.QueryContext(ctx,script,username,password)
	if err != nil {
		panic(err)
	}
	if rows.Next(){
		var username string
		err := rows.Scan(&username)
		if err != nil{
			panic(err)
		}
		fmt.Println("===============")
		fmt.Println("login sukses    :", username)
		fmt.Println("===============")
	}else{
		fmt.Println("===============")
		fmt.Println("login gagal    ")
		fmt.Println("===============")
	}
	defer rows.Close()
}

func TestExecSqlPar(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()
	username := "kambing"
	password := "apalahkambing"
	script := "INSERT INTO user(username, password) VALUES (?,?)"
	_,err := db.ExecContext(ctx,script,username,password)
	if err != nil {
		panic(err)
	}
	fmt.Println("berhasil insert customer baru")
}

func TestAutoIncrement(t *testing.T){
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()
	email := "kambing@gmail.com"
	comment := "tescomment"
	script := "INSERT INTO comments (email, comment) values (?, ?)"
	result,err := db.ExecContext(ctx,script,email,comment)

	if err != nil {
		panic(err)
	}
	insertId,err := result.LastInsertId()
	if err != nil{
		panic(err)
	}
	fmt.Println("berhasil memasukan komen baru dengan id = ",insertId)
}

//digunakan saat parameter yang digunakan selalu sama namun isi datanya berbeda beda
func TestPrepareStatment(t *testing.T){
	db := GetConnection()
	defer db.Close()
	// email := "kambing@gmail.com"
	// comment := "tescomment"
	ctx := context.Background()
	script := "INSERT INTO comments(email, comment) values (?, ?)"
	statment,err := db.PrepareContext(ctx,script)
	if err != nil {
		panic(err)
	}
	defer statment.Close()
	for i:= 0;i<10;i++{
		email := "dimas"+strconv.Itoa(i)+"@gmail.com"
		comment := "komentar ke "+strconv.Itoa(i)

		result,err := statment.ExecContext(ctx,email,comment)
		if err != nil{
			panic(err)
		}
		id,err := result.LastInsertId()
		if err != nil{
			panic(err)
		}
		fmt.Println("komen berhasil ditambahkan dengan id = " , id)
	}
}

func TestTransaction(t *testing.T){
	db := GetConnection()
	defer db.Close()
	ctx := context.Background()
	script := "INSERT INTO comments(email, comment) values (?, ?)"

	tx, err := db.Begin()
	if err != nil{
		panic(err)
	}
	for i:= 0;i<10;i++{
		email := "dimas"+strconv.Itoa(i)+"@gmail.com"
		comment := "komentar ke "+strconv.Itoa(i)

		result,err := tx.ExecContext(ctx,script,email,comment)
		if err != nil{
			panic(err)
		}
		id,err := result.LastInsertId()
		if err != nil{
			panic(err)
		}
		fmt.Println("komen berhasil ditambahkan dengan id = " , id)
	}

	err = tx.Commit()//bisa menggunakan rollback agar datanya tidak masuk ke db
	if err != nil{
			panic(err)
		}
}