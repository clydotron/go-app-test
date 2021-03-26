package views

import (
	"fmt"

	"github.com/clydotron/go-app-test/models"
	"github.com/clydotron/go-app-test/ui"
	"github.com/clydotron/go-app-test/utils"
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

// Events ...
type Events struct {
	app.Compo
	events []models.EventInfo

	eb     *utils.EventBus
	dataCh utils.DataChannel
	doneCh chan bool
	topic  string
}

//constructor?

func NewEventsView(eb *utils.EventBus) *Events {
	e := &Events{
		eb:     eb,
		doneCh: make(chan bool),
		dataCh: make(chan utils.DataEvent),
		topic:  "event",
	}
	return e
}

func (c *Events) handleEvent(d utils.DataEvent) {
	if d.Topic == c.topic {
		app.Dispatch(func() {
			ei, ok := d.Data.(*models.EventInfo)
			if ok {
				fmt.Println("Events.handleEvent:", ei.ID)
				c.events = append(c.events, *ei)
				c.Update()
			}
		})
	}
}

func (c *Events) OnMount(ctx app.Context) {
	fmt.Println("Events onMount >start<")
	defer fmt.Println("Events onMount >end<")

	// need a way to get all of the events up until now?

	c.eb.Subscribe(c.topic, c.dataCh)

	go func() {
		for {
			select {
			case d := <-c.dataCh:
				c.handleEvent(d)
			case <-c.doneCh:
				fmt.Println("Events: done.")
				return
			}
		}
	}()

}
func (c *Events) OnDismount() {
	defer fmt.Println("Events dismounted")

	c.doneCh <- true
	c.eb.Unsubscribe(c.topic, c.dataCh)

}

func (c *Events) clearEvents(ctx app.Context, e app.Event) {
	app.Dispatch(func() {
		c.events = []models.EventInfo{}
		c.Update()
	})
}

// Render ...
func (c *Events) Render() app.UI {

	return app.Div().Class("h-screen w-screen").
		Body(
			&ui.NavBar{},
			app.Div().Class("pt-20 bg-purple-100 flex flex-col px-2").
				Body(
					app.Table().
						Body(
							app.Button().Class("rounded bg-indigo-200 p-1").Text("Clear").OnClick(c.clearEvents),
							app.Tr().Class("bg-gray-200").Body(
								//bg color, change alignment
								app.Th().Class("text-left").Text("Date"),
								app.Th().Class("text-left").Text("Event"),
								app.Th().Class("text-left").Text("ID"),
							),
							app.Range(c.events).Slice(func(i int) app.UI {
								e := c.events[i]
								bgcolor := "bg-blue-200"
								if i%2 == 1 {
									bgcolor = "bg-blue-100"
								}
								return app.Tr().Class(bgcolor).
									Body(
										app.Td().Text(e.TimeStamp.Format("02 Jan 2006 at 15:04:05")),
										app.Td().Text(e.Name),
										app.Td().Text(e.ID),
									)

							}),
						),
				),
		)
}

/*
<div class="flex flex-col">
  <div class="-my-2 overflow-x-auto sm:-mx-6 lg:-mx-8">
    <div class="py-2 align-middle inline-block min-w-full sm:px-6 lg:px-8">
      <div class="shadow overflow-hidden border-b border-gray-200 sm:rounded-lg">
        <table class="min-w-full divide-y divide-gray-200">
          <thead class="bg-gray-50">
            <tr>
              <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Name
              </th>
              <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Title
              </th>
              <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Status
              </th>
              <th scope="col" class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                Role
              </th>
              <th scope="col" class="relative px-6 py-3">
                <span class="sr-only">Edit</span>
              </th>
            </tr>
          </thead>
*/
