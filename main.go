// +build !wasm

// The server is a classic Go program that can run on various architecture but
// not on WebAssembly. Therefore, the build instruction above is to exclude the
// code below from being built on the wasm architecture.

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"log"
	"net/http"
	"strconv"

	"github.com/clydotron/go-app-test/client"
	"github.com/clydotron/go-app-test/models"
	"github.com/maxence-charriere/go-app/v7/pkg/app"

	//"github.com/wcharczuk/go-chart" //exposes "chart"
	"github.com/wcharczuk/go-chart/v2"
)

func lager(format string, v ...interface{}) {
	fmt.Println(format, v)

}

type ClientX struct {
	cc *client.ClusterClient
	ct *models.ClusterTracker
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func testImageX2(w http.ResponseWriter, r *http.Request) {
	m := image.NewRGBA(image.Rect(0, 0, 240, 240))
	blue := color.RGBA{0, 0, 255, 255}
	draw.Draw(m, m.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src)

	var img image.Image = m
	writeImagePNG(w, &img)
}

func testImageX(w http.ResponseWriter, r *http.Request) {

	pie := chart.PieChart{
		Width:  512,
		Height: 512,
		Values: []chart.Value{
			{Value: 5, Label: "Blue"},
			{Value: 5, Label: "Green"},
			{Value: 4, Label: "Gray"},
			{Value: 4, Label: "Orange"},
			{Value: 3, Label: "Deep Blue"},
			{Value: 3, Label: "??"},
			{Value: 1, Label: "!!"},
		},
	}

	buffer := bytes.NewBuffer([]byte{})
	pie.Render(chart.PNG, buffer)

	// if err := png.Encode(buffer, *img); err != nil {
	// 	log.Println("unable to encode PNG image.")
	// }

	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	if _, err := w.Write(buffer.Bytes()); err != nil {
		log.Println("unable to write image.")
	}
}

func testStackedBarChart(w http.ResponseWriter, r *http.Request) {

	sbc := chart.StackedBarChart{
		Title: "Test",
		Background: chart.Style{
			Padding: chart.Box{
				Top: 40,
			},
		},
		Height: 512,
		Bars: []chart.StackedBar{
			{
				Name: "This",
				Values: []chart.Value{
					{Value: 5, Label: "Blue"},
					{Value: 5, Label: "Green"},
					{Value: 4, Label: "Gray"},
					{Value: 3, Label: "Orange"},
					{Value: 3, Label: "Test"},
					{Value: 2, Label: "??"},
					{Value: 1, Label: "!!"},
				},
			},
			{
				Name: "Test",
				Values: []chart.Value{
					{Value: 10, Label: "Blue"},
					{Value: 5, Label: "Green"},
					{Value: 1, Label: "Gray"},
				},
			},
			{
				Name: "Test 2",
				Values: []chart.Value{
					{Value: 10, Label: "Blue"},
					{Value: 6, Label: "Green"},
					{Value: 4, Label: "Gray"},
				},
			},
		},
	}

	buffer := new(bytes.Buffer)
	sbc.Render(chart.PNG, buffer)
	writeImageToPNG(w, buffer)
}

func writeImageToPNG(w http.ResponseWriter, buffer *bytes.Buffer) {

	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	if _, err := w.Write(buffer.Bytes()); err != nil {
		log.Println("unable to write image.")
	}
}
func writeImagePNG(w http.ResponseWriter, img *image.Image) {

	buffer := new(bytes.Buffer)
	if err := png.Encode(buffer, *img); err != nil {
		log.Println("unable to encode PNG image.")
	}

	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	if _, err := w.Write(buffer.Bytes()); err != nil {
		log.Println("unable to write image.")
	}
}
func writeImageJPEG(w http.ResponseWriter, img *image.Image) {

	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, *img, nil); err != nil {
		log.Println("unable to encode image.")
	}

	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	if _, err := w.Write(buffer.Bytes()); err != nil {
		log.Println("unable to write image.")
	}
}

// apiCluster -- this is called on its own go routine, so we can block
func (cx *ClientX) apiCluster(w http.ResponseWriter, r *http.Request) {
	fmt.Println("api request: cluster")

	setupResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	if (*r).Method != "GET" {
		fmt.Println("Incorrect method:", (*r).Method)
		return
	}
	//make sure this is a get
	//get the latest cluster info from the store and json it...

	cx.ct.UpdateStatus()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cx.ct.CI)
}

// The main function is the entry of the server. It is where the HTTP handler
// that serves the UI is defined and where the server is started.
//
// Note that because main.go and app.go are built for different architectures,
// this main() function is not in conflict with the one in
// app.go.
func main() {

	// could move all of this into init
	client := client.NewClusterClient()
	err := client.Connect("localhost:50051")
	if err != nil {
		log.Fatalln("### Client failed to connect:", err)
	}
	defer client.Close()

	clusterTracker := models.NewClusterTracker(client)
	clusterTracker.InitWithFakeData()
	clusterTracker.Start()
	defer clusterTracker.Stop()

	cx := &ClientX{
		cc: client,
		ct: clusterTracker,
	}
	// sequence of events:
	// create client (gRPC connection to server)
	// create the cluster tracker - responsible for making the HealthCheck gRPC call and [optionally] maintaining a snaphot of the cluster (so can diff)
	// ClientX(Api) - implements the http.HandleFunc call to respond to api requests

	// app.Handler is a standard HTTP handler that serves the UI and its
	// resources to make it work in a web browser.
	//
	// It implements the http.Handler interface so it can seamlessly be used
	// with the Go HTTP standard library.
	http.Handle("/", &app.Handler{
		Name:        "Hello",
		Description: "Experimental",
		Styles: []string{
			"/web/tailwind.css", "/web/test.css", // Include .css file.
		},
	})

	http.HandleFunc("/api/v1/cluster", cx.apiCluster)
	http.HandleFunc("/api/v1/blue", testImageX)
	http.HandleFunc("/api/v1/red", testStackedBarChart)

	fmt.Println("up and running...")

	err = http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
