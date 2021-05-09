package main

import (
	"database/sql"
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"path"
	"strings"

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
	var err error
	setup.DB, err = sql.Open("sqlite3",
		"./test.db")
	if err != nil {
		log.Fatal(err)
	}

	// Handle to ./reactjs/build folder on root path
	http.HandleFunc("/", rootHandler)

	http.HandleFunc("/companies", controllers.GetAllCompanies)

	http.HandleFunc("/tags", controllers.GetAllTags)

	http.HandleFunc("/big-numbers", controllers.GetBigNumbers)

	http.HandleFunc("/latest", controllers.GetLatestBills)

	http.HandleFunc("/by-day", controllers.GetMonthBillsByDay)

	http.HandleFunc("/by-tag", controllers.GetMonthBillsByTag)

	http.HandleFunc("/month", controllers.GetThisMonthBills)

	http.HandleFunc("/create", controllers.CreateNewBill)

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
