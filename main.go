// +build !wasm

// The server is a classic Go program that can run on various architecture but
// not on WebAssembly. Therefore, the build instruction above is to exclude the
// code below from being built on the wasm architecture.

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

func lager(format string, v ...interface{}) {
	fmt.Println(format, v)

}

// type DataStore struct {
// 	MI []models.MachineInfo
// 	MV *views.Machines
// }

// func (d *DataStore) update(mi []models.MachineInfo) {
// 	d.MI = mi
// 	d.MV.UpdateM(&mi)
// }

// // DS ...
// var DS *DataStore

// The main function is the entry of the server. It is where the HTTP handler
// that serves the UI is defined and where the server is started.
//
// Note that because main.go and app.go are built for different architectures,
// this main() function is not in conflict with the one in
// app.go.
func main() {

	// mi := []models.MachineInfo{
	// 	models.MachineInfo{
	// 		Name: "Machine 1",
	// 		Role: "Manager",
	// 		Tasks: []models.TaskInfo{
	// 			models.TaskInfo{
	// 				Name:        "Redis",
	// 				Tag:         "3.2.1",
	// 				ContainerID: "bunch of hex",
	// 				State:       "running",
	// 				Updated:     time.Now(),
	// 			},
	// 		},
	// 	},
	// }

	// DS := &DataStore{
	// 	MV: &views.Machines{MI: mi},
	// 	MI: mi,
	// }
	// fmt.Println(DS)

	/*
	   	ticker := time.NewTicker(500 * time.Millisecond)
	       done := make(chan bool)
	       go func() {
	           for {
	               select {
	               case <-done:
	                   return
	               case t := <-ticker.C:
	                   fmt.Println("Tick at", t)
	               }
	           }
	       }()
	*/
	// app.Handler is a standard HTTP handler that serves the UI and its
	// resources to make it work in a web browser.
	//
	// It implements the http.Handler interface so it can seamlessly be used
	// with the Go HTTP standard library.
	http.Handle("/", &app.Handler{
		Name:        "Hello",
		Description: "An Hello World! example",
		Styles: []string{
			"/web/tailwind.css", // Inlude .css file.
		},
	})

	//http.Handle("/beer")
	//infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	//app.DefaultLogger = lager

	fmt.Println("up and running...")

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
