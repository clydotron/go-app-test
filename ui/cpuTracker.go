package ui

import (
	"bytes"
	"encoding/base64"
	"fmt"

	"github.com/clydotron/go-app-test/utils"
	"github.com/maxence-charriere/go-app/v7/pkg/app"
	"github.com/wcharczuk/go-chart/v2"
	//exposes "chart"
)

type CpuTracker struct {
	app.Compo

	id  string
	eb  *utils.EventBus
	sub *utils.EventBusSubscriber

	history  []int
	encChart string
}

func NewCpuTracker(id string, eb *utils.EventBus) *CpuTracker {

	tt := &CpuTracker{id: id,
		eb:  eb,
		sub: utils.NewEventBusSubscriber("PI", eb),
	}
	return tt
}

func (c *CpuTracker) handleEvent(d utils.DataEvent) {
	go c.updateChart()
}

func (c *CpuTracker) updateChart() {

	// chart will always be n pixels wide
	// show data for

	// graph := chart.Chart{
	// 	Series: []chart.Series{
	// 		chart.ContinuousSeries{
	// 			XValues: []float64{1.0, 2.0, 3.0, 4.0},
	// 			YValues: []float64{1.0, 2.0, 3.0, 4.0},
	// 		},
	// 	},
	// }

	// buffer := bytes.NewBuffer([]byte{})
	// err := graph.Render(chart.PNG, buffer)

	// do the heavy lifting off the main thread. Update the encoded image once we are ready
	//fmt.Println("updating chart!")
	graph := chart.Chart{
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: chart.Seq{Sequence: chart.NewLinearSequence().WithStart(1.0).WithEnd(100.0)}.Values(),
				YValues: chart.Seq{Sequence: chart.NewRandomSequence().WithLen(100).WithMin(50).WithMax(150)}.Values(),
			},
		},
	}

	buffer := bytes.NewBuffer([]byte{})
	graph.Render(chart.PNG, buffer)
	enc := base64.StdEncoding.EncodeToString(buffer.Bytes())

	app.Dispatch(func() {
		c.encChart = enc
		c.Update()
	})
}

func (c *CpuTracker) OnMount(ctx app.Context) {

	fmt.Println("CpuTracker.OnMount: ")

	c.sub.Start(c.handleEvent)
}

// OnDismount ...
func (c *CpuTracker) OnDismount() {
	defer fmt.Println("CpuTracker dismounted")

	c.sub.Stop()
}

func (c *CpuTracker) Render() app.UI {
	h := ""
	if c.encChart == "" {
		h = " hidden"
		//return nil
	}

	enc := "data:image/png;base64, " + c.encChart

	return app.Div().
		Class("bg-purple-300 h-screen w-screen").
		Body(
			app.Text("Chart"),
			app.Img().Src(enc).Class("p-2"+h).Alt("orange"),
		)

}
