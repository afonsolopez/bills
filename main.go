package main

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"path"
	"strings"
	"time"

	"github.com/afonsolopez/bills/controllers"
	"github.com/afonsolopez/bills/setup"

	_ "github.com/mattn/go-sqlite3"

	"github.com/webview/webview"
)

var err error

// Compiles all react data to binary

//go:embed reactjs/build
var f embed.FS

func main() {

	// Setup the database connection
	setup.DB, err = sql.Open("sqlite3",
		"./data.db")
	if err != nil {
		log.Fatal(err)
	}

	// SQL command to create all database tables if no exists
	createDB := `
	CREATE TABLE IF NOT EXISTS bills (id integer,
		company_id integer,
		title text,
		price real,
		date_id integer,
		tag_id integer,
		PRIMARY KEY (id),
		CONSTRAINT fk_bills_date FOREIGN KEY (date_id) REFERENCES dates(id),
		CONSTRAINT fk_bills_tag FOREIGN KEY (tag_id) REFERENCES tags(id),
		CONSTRAINT fk_bills_company FOREIGN KEY (company_id) REFERENCES companies(id));
		
		CREATE TABLE IF NOT EXISTS companies (id integer,
		name text,
		PRIMARY KEY (id));
		
		CREATE TABLE IF NOT EXISTS dates (id integer,
		time_stamp datetime,
		PRIMARY KEY (id));
		
		CREATE TABLE IF NOT EXISTS tags (id integer,
		name text,
		PRIMARY KEY (id));
	`

	// Create a context
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	// Execute the database/tables creation SQL command
	mount, err := setup.DB.ExecContext(ctx, createDB)
	if err != nil {
		log.Fatal(err)
	}
	// Count each row affectd by the database/tables creation
	rows, err := mount.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Rows affected when creating table: %d", rows)

	// Define the root app path./reactjs/build
	http.HandleFunc("/", rootHandler)

	// Get all companies
	http.HandleFunc("/companies", controllers.GetAllCompanies) // [GET]

	// Gell all tags
	http.HandleFunc("/tags", controllers.GetAllTags) // [GET]

	// Get all months with registerd bill
	http.HandleFunc("/months", controllers.GetAllMonths) // [GET]

	// Returns the total of money spent on that month and days left also on that month
	http.HandleFunc("/big-numbers", controllers.GetBigNumbers) // [GET/POST]

	// Latest registered bills
	http.HandleFunc("/latest", controllers.GetLatestBills) // [GET]

	// Get all bills grouped by same day
	http.HandleFunc("/by-day", controllers.GetMonthBillsByDay) // [GET]

	// Get all bills grouped by same tag
	http.HandleFunc("/by-tag", controllers.GetMonthBillsByTag) // [GET/POST]

	// Get all bills from a month
	http.HandleFunc("/month", controllers.GetThisMonthBills) // [GET/POST]

	// Create a new bill
	http.HandleFunc("/create", controllers.CreateNewBill) // [POST]

	// Delete a bill
	http.HandleFunc("/delete", controllers.DeleteBill) // [DELETE]

	// Running a server as goroutine (for non-blocking working)
	go http.ListenAndServe(":8000", nil)

	// Window setup
	debug := true // Devtools on/off
	w := webview.New(debug)
	//	defer w.Destroy()
	w.SetTitle("Bills - 0.1a")             // App name
	w.SetSize(800, 600, webview.HintFixed) // Window size 800x600px (w/o resize)
	w.Navigate("http://localhost:8000")    // Running on: http://localhost:8000
	w.Run()
}

// Handler to make possible to render SPA's
// https://github.com/akmittal/go-embed/blob/main/main_extended.go
func rootHandler(rw http.ResponseWriter, req *http.Request) {
	upath := req.URL.Path
	if !strings.HasPrefix(upath, "/") {
		upath = "/" + upath
		req.URL.Path = upath
	}
	upath = path.Clean(upath)
	fsys := fs.FS(f)
	fmt.Println(upath)
	contentStatic, _ := fs.Sub(fsys, "reactjs/build")
	if _, err := contentStatic.Open(strings.TrimLeft(upath, "/")); err != nil { // If file not found server index/html from root
		req.URL.Path = "/"
	}
	http.FileServer(http.FS(contentStatic)).ServeHTTP(rw, req)
}
