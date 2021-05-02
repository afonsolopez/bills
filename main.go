package main

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"path"
	"strings"

	"github.com/afonsolopez/bills/controllers"
	"github.com/afonsolopez/bills/models"
	"github.com/afonsolopez/bills/setup"

	"github.com/webview/webview"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var err error

//go:embed reactjs/build
var f embed.FS

func main() {

	setup.DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	setup.DB.AutoMigrate(&models.Bill{}, models.Company{}, models.Date{}, models.Tag{})

	// Handle to ./reactjs/build folder on root path
	http.HandleFunc("/", rootHandler)

	// Handle to showMessage func on /hello path
	http.HandleFunc("/hello", controllers.ShowMessage)

	// Handle to GetAllBills func on /hello path
	http.HandleFunc("/all", controllers.GetAllBills)

	// Handle to GetLastMonthBills func on /hello path
	http.HandleFunc("/last", controllers.GetLastMonthBills)

	// Handle to GetMonthBills func on /getMonthBills path
	http.HandleFunc("/getMonthBills", controllers.GetMonthBills)

	// Handle to GetMonthBillsByTag func on /getMonthBillsByTag path
	http.HandleFunc("/getMonthBillsByTag", controllers.GetMonthBillsByTag)

	// Handle to GetAllCompanies func on /getAllCompanies path
	http.HandleFunc("/getAllCompanies", controllers.GetAllCompanies)

	// Handle to GetAllTags func on /getAllTags path
	http.HandleFunc("/getAllTags", controllers.GetAllTags)

	// Handle to CreateNewBill func on /hello path
	http.HandleFunc("/createNewBill", controllers.CreateNewBill)

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
