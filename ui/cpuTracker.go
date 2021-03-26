package ui

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/url"

	"github.com/clydotron/go-app-test/utils"
	"github.com/maxence-charriere/go-app/v7/pkg/app"
	"github.com/wcharczuk/go-chart/v2"
	//exposes "chart"
)

const ebTarget = "PI"

type CpuTracker struct {
	app.Compo

	id string
	//ci models.CPUInfo
	eb     *utils.EventBus
	dataCh utils.DataChannel
	doneCh chan bool

	history  []int
	encChart string
}

func NewCpuTracker(id string, eb *utils.EventBus) *CpuTracker {

	fmt.Println("NewCpuTracker --")
	tt := &CpuTracker{id: id, eb: eb}

	return tt
}

func (c *CpuTracker) handleEvent(d utils.DataEvent) {

	if d.Topic == ebTarget {
		go c.updateChart()
	}
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
	fmt.Println("updating chart!")
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

func (c *CpuTracker) OnNav(ctx app.Context, url *url.URL) {

}

func (c *CpuTracker) OnMount(ctx app.Context) {

	fmt.Println("CpuTracker.OnMount: ")

	if c.dataCh == nil {
		fmt.Println("CpuTracker: initializing")
		c.doneCh = make(chan bool)
		c.dataCh = make(chan utils.DataEvent)
		c.eb.Subscribe(ebTarget, c.dataCh)

		go func() {
			for {
				select {
				case d := <-c.dataCh:
					c.handleEvent(d)
				case <-c.doneCh:
					fmt.Println("CpuTracker: all done!")
					return
				}
			}
		}()
	}
}

// OnDismount ...
func (c *CpuTracker) OnDismount() {
	defer fmt.Println("CpuTracker dismounted")

	c.doneCh <- true
	c.eb.Unsubscribe(ebTarget, c.dataCh)

	c.dataCh = nil
	c.doneCh = nil
}

func (c *CpuTracker) Render() app.UI {
	// if no data, return immediately
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
