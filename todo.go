// todo.go
package main

import (
    "database/sql"
    "go-echo-vue/handlers"
    "github.com/labstack/echo"
    "github.com/labstack/echo/engine/standard"
    "net/http"
    _ "github.com/mattn/go-sqlite3"
)

func main() {

	db := initDB("storage.db")
    migrate(db)

	// create new instance of echo
	e := echo.New()

	e.File("/", "public/index.html")
    e.GET("/tasks", handlers.GetTasks(db))
    e.PUT("/tasks", handlers.PutTask(db))
    e.DELETE("/tasks/:id", handlers.DeleteTask(db))
	// Start as a web server
    e.Run(standard.New(":8080"))

    // this is boundary...
    // e := echo.New()

	e.GET("/tasks", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
    // e.Logger.Fatal(e.Start(":1323"))
}

// todo.go
	func initDB(filepath string) *sql.DB {
	    db, err := sql.Open("sqlite3", filepath)

	    // Here we check for any db errors then exit
	    if err != nil {
	        panic(err)
	    }

	    // If we don't get any errors but somehow still don't get a db connection
	    // we exit as well
	    if db == nil {
	        panic("db nil")
	    }
	    return db
	}

	func migrate(db *sql.DB) {
	    sql := `
	    CREATE TABLE IF NOT EXISTS tasks(
	        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	        name VARCHAR NOT NULL
	    );
	    `

	    _, err := db.Exec(sql)
	    // Exit if something goes wrong with our SQL statement above
	    if err != nil {
	        panic(err)
	    }
	}