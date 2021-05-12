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

//go:embed reactjs/build
var f embed.FS

func main() {

	// Databse connection
	// var err error
	setup.DB, err = sql.Open("sqlite3",
		"./data.db")
	if err != nil {
		log.Fatal(err)
	}

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

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	mount, err := setup.DB.ExecContext(ctx, createDB)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := mount.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Rows affected when creating table: %d", rows)

	// Handle to ./reactjs/build folder on root path
	http.HandleFunc("/", rootHandler)

	http.HandleFunc("/companies", controllers.GetAllCompanies)

	http.HandleFunc("/tags", controllers.GetAllTags)

	http.HandleFunc("/months", controllers.GetAllMonths)

	http.HandleFunc("/big-numbers", controllers.GetBigNumbers)

	http.HandleFunc("/latest", controllers.GetLatestBills)

	http.HandleFunc("/by-day", controllers.GetMonthBillsByDay)

	http.HandleFunc("/by-tag", controllers.GetMonthBillsByTag)

	http.HandleFunc("/month", controllers.GetThisMonthBills)

	http.HandleFunc("/create", controllers.CreateNewBill)

	http.HandleFunc("/delete", controllers.DeleteBill)

	// Run server at port 8000 as goroutine
	// for non-block working
	go http.ListenAndServe(":8000", nil)

	// Let's open window app with:
	debug := true
	w := webview.New(debug)
	//	defer w.Destroy()
	w.SetTitle("Bills - Beta")             //  - name: Golang App
	w.SetSize(800, 600, webview.HintFixed) //  - sizes: 800x600 px
	w.Navigate("http://localhost:8000")    //  - address: http://localhost:8000
	w.Run()
}

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
