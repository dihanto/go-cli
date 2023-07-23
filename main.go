package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/chzyer/readline"
	_ "github.com/lib/pq"
)

type Student struct {
	Name, Gender, Major string
	Age                 int
}

var std []Student

func main() {

	for {
		if password() {
			service()
		}
	}
}

func menu() int {
	var input int

	fmt.Println("1. Create Student")
	fmt.Println("2. Show Students")
	fmt.Println("3. Exit")
	fmt.Println("Enter the number to choose menu :")

	fmt.Scan(&input)

	return input
}

func password() bool {
	var password string

	fmt.Println("Please enter the password to use this app")
	fmt.Scan(&password)
	return password == "jangkrik"
}

func service() {
	for {
		input := menu()
		if input == 1 {
			var name, major, gender string

			name = strings.ToUpper(inputString("What is the name of student?"))
			age, _ := strconv.Atoi(inputInt("What is the age of student?"))
			aged := int(age)
			major = strings.ToUpper(inputString("What is the major of student?"))
			gender = genderInput()

			student := Student{
				Name:   name,
				Age:    aged,
				Major:  major,
				Gender: gender,
			}
			std = append(std, student)

			db := initDatabaseConnection()
			query := "INSERT INTO students (name, age, major, gender) VALUES ($1,$2,$3,$4)"
			_, err := db.Exec(query, student.Name, student.Age, student.Major, student.Gender)
			if err != nil {
				log.Println(err)
			}

		} else if input == 2 {

			db := initDatabaseConnection()
			query := "SELECT name, age, major, gender FROM students"
			rows, err := db.Query(query)
			if err != nil {
				log.Println(err)
			}

			var students []Student
			for rows.Next() {
				var student Student
				err = rows.Scan(&student.Name, &student.Age, &student.Major, &student.Gender)
				if err != nil {
					log.Println(err)
				}
				students = append(students, student)
			}

			if len(students) == 0 {
				fmt.Println("Student not found")
				break
			}

			fmt.Println(students)
		} else if input == 3 {
			os.Exit(1)
		}
	}
}

func inputInt(msg string) string {
	var input string
	// var err error
	for {
		fmt.Println(msg)
		fmt.Scan(&input)

		_, err := strconv.ParseInt(input, 10, 32)

		if err != nil {
			fmt.Println("input not valid, please enter the age of student")
			continue
		}
		break
	}
	return input
}

func genderInput() string {
	for {
		data := inputString("What is the gender of student :")
		if data == "male" {
			return fmt.Sprintln("Male")
		} else if data == "female" {
			return fmt.Sprintln("Female")
		} else {
			return "Please type male/female"
		}
	}
}

func initDatabaseConnection() *sql.DB {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=cli sslmode=disable")
	if err != nil {
		fmt.Sprintln("Failed to connect database")
	}
	return db
}

func inputString(msg string) (result string) {
	rl, err := readline.New(msg)
	if err != nil {
		fmt.Println("Error creating readline instance:", err)
		return inputString(msg) // Retry if there's an error
	}
	defer rl.Close()

	for {
		line, err := rl.Readline()
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		result = strings.TrimSpace(line)

		if len(result) == 0 {
			fmt.Println("input is empty, please type something")
			continue
		}

		_, err = strconv.Atoi(result)
		if err == nil {
			fmt.Println("input not valid, please read the instruction carefully.")
			continue
		} else if result == "true" || result == "false" {
			fmt.Println("input not valid, please read the instruction carefully.")
			continue
		}

		return result
	}
}
