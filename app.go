// +build wasm

// The UI is running only on a web browser. Therefore, the build instruction
// above is to compile the code below only when the program is built for the
// WebAssembly (wasm) architecture.

package main

import (
	"time"

	models "github.com/clydotron/go-app-test/model"
	"github.com/clydotron/go-app-test/ui"
	"github.com/clydotron/go-app-test/views"
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

type DataStore struct {
	MI []models.MachineInfo
	MV *views.Machines
}

func (d *DataStore) update(mi []models.MachineInfo) {
	d.MI = mi
	d.MV.UpdateM(&mi)
}

// The main function is the entry point of the UI. It is where components are
// associated with URL paths and where the UI is started.
func main() {
	//app.Log("setting up routes")

	taskRepo := models.NewTaskInfoRepo()

	mi := []models.MachineInfo{
		models.MachineInfo{
			Name:   "Machine 1",
			Role:   "Manager",
			Status: "running",
			Tasks: []models.TaskInfo{
				models.TaskInfo{
					Name:        "Redis",
					Tag:         "3.2.1",
					ContainerID: "bunch of hex",
					State:       "running",
					Updated:     time.Now(),
				},
			},
		},
	}

	DS := &DataStore{
		MV: &views.Machines{MI: mi},
		MI: mi,
	}

	//setup the timer:

	app.Route("/", &ui.Updater{}) // hello component is associated with URL path "/".
	app.Route("/workcation", &workcation{})
	app.Route("/machines", DS.MV)
	app.RouteWithRegexp("^/node.*", &ui.Node{})
	app.RouteWithRegexp("^/task.*", ui.NewTaskDetail(taskRepo))

	// app.Route("/machine", &views.Machines{MI: []models.MachineInfo{
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
	// 			models.TaskInfo{
	// 				Name:        "Mongo",
	// 				Tag:         "4.2.1",
	// 				ContainerID: "bunch of hex",
	// 				State:       "stopped",
	// 				Updated:     time.Now(),
	// 			},
	// 		},
	// 	},
	// }})
	app.Run() // Launches the PWA.
}

// return app.Div().OnClick(c.onClick)
// }

// type imageFrame struct {
// 	app.Compo
// 	clicked bool
// }

// func (ifx *imageFrame) Render() app.UI {
// 	cvalue := "p-12 "
// 	if ifx.clicked {
// 		cvalue += "bg-green-400"
// 	} else {
// 		cvalue += "bg-purple-600"
// 	}

// 	return app.Div().OnClick(ifx.onClick).
// 		Class(cvalue).
// 		Body(
// 			app.H3().Text("ClickMe!x"),
// 			app.Img().Src("/web/img/logo.svg").Class("p-2").Class("h-12"),
// 		)
// }

// // }

// func (c *imageFrame) onClick(ctx app.Context, e app.Event) {
// 	app.Log("click")
// 	c.clicked = !c.clicked
// 	c.Update()

// }

// type testList struct {
// 	app.Compo
// }

// func (c *testList) Render() app.UI {
// 	data := []string{
// 		"hello",
// 		"go-app",
// 		"is",
// 		"sexy",
// 	}

// 	return app.Ul().Body(
// 		app.Range(data).Slice(func(i int) app.UI {
// 			return app.Li().Text(data[i])
// 		}),
// 	)
// }
