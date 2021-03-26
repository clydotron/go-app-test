package views_test

import (
	"fmt"
	"testing"

	"github.com/clydotron/go-app-test/models"
	"github.com/clydotron/go-app-test/utils"
	"github.com/clydotron/go-app-test/views"
)

func TestHandleEvent(t *testing.T) {

	eb := utils.NewEventBus()
	v := views.NewProcessesView(eb)

	//send a PI event, confirm that
	pi := models.ProcessInfo{ID: "p1", Name: "p1"}
	eb.Publish("PI", &pi)

	fmt.Println(v)

}
