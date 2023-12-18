package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Student struct {
	Id int
	Name string
	Grade float64 
}

func main(){
	db, err := sql.Open("postgres", "dbname=gotest sslmode=disable user=pc")
	
	if err != nil {
		panic(fmt.Sprint(err))
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS
Students (
	id INT PRIMARY KEY,
	name VARCHAR,
	grade FLOAT
)	
`)

	if err != nil {
		panic(fmt.Sprint(err))
	}

	stud := Student{
		Name: "Ricky",
		Id: 4 ,
		Grade: 9.8,
	}

	err = InsertStudentToTable(db, stud)

	if err != nil{
		fmt.Println(err)
	}

	students, err := FindStudentsByID(db)

	if(err != nil){
		fmt.Println(err)
	}
	fmt.Println(students)
}

func InsertStudentToTable(db *sql.DB, student Student) error {
	_, err := db.Exec(`
INSERT INTO Students(id, name, grade)
VALUES ($1, $2, $3)`, student.Id, student.Name, student.Grade)

	return err
}

func FindStudentsByID(db *sql.DB) ([]Student, error) {
	find, err := db.Query(`
	SELECT *
	FROM Students
	`)

	if err != nil {
		fmt.Println("err")
		return nil, err
	}

	var students []Student;

	for find.Next() {
		student := Student{}
		err := find.Scan(&student.Id, &student.Name, &student.Grade)

		if (err != nil){
			return nil, err
		}

		students = append(students, student)
	}

	return students, nil
}