package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

var db *sql.DB

type Page struct {
	Title string
	Body  []byte
}
type todo struct {
	id      int
	task    string
	is_done string
}

const (
	// Initialize connection constants.
	HOST     = "localhost"
	PORT     = 5432
	DATABASE = "postgre"
	USER     = "postgre"
	PASSWORD = "pass"
)

func main() {

	// Capture connection properties.
	var connectionString = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", HOST, PORT, USER, PASSWORD, DATABASE)
	fmt.Println(connectionString)
	// Get a database handle.
	var err error
	db, err = sql.Open("postgres", connectionString)
	checkError(err)

	pingErr := db.Ping()
	checkError(pingErr)
	fmt.Println("Connected!")

	// Drop previous table of same name if one exists.
	_, err = db.Exec("DROP TABLE IF EXISTS to_do;")
	fmt.Println("Finished dropping table (if existed)")

	// Create table.
	_, err = db.Exec("CREATE TABLE to_do (id serial PRIMARY KEY, task VARCHAR(50), is_done VARCHAR(50));")
	checkError(err)
	fmt.Println("Finished creating table")

	// Insert some data into table.
	sql_statement := "INSERT INTO to_do (task, is_done) VALUES ($1, $2);"
	_, err = db.Exec(sql_statement, "task1", true)
	checkError(err)
	_, err = db.Exec(sql_statement, "task2", false)
	checkError(err)
	_, err = db.Exec(sql_statement, "task3", false)
	checkError(err)
	fmt.Println("Inserted 3 rows of data")

	http.HandleFunc("/", handler)
	http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, Welcome the tod od app %s!", r.URL.Path[1:])
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	t := todo{}

	read_statement := "SELECT * from to_do;"
	rows, err := db.Query(read_statement)
	checkError(err)
	defer rows.Close()

	var todo_list []todo

	for rows.Next() {
		switch err := rows.Scan(&t.id, &t.task, &t.is_done); err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned")
		case nil:
			fmt.Printf("Data row = (%d, %s, %d)\n", t.id, t.task, t.is_done)
			t = todo{
				id:      t.id,
				task:    t.task,
				is_done: t.is_done,
			}
			todo_list = append(todo_list, t)
		default:
			checkError(err)
		}
	}
	for _, t := range todo_list {
		fmt.Fprintf(w, "task name: %s, is done: %s \n", t.task, t.is_done)
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
