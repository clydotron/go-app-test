package ui

import (
	"fmt"
	"net/url"
	"strings"

	models "github.com/clydotron/go-app-test/model"
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

// TaskDetail ...
type TaskDetail struct {
	app.Compo

	TI   models.TaskInfo
	repo *models.TaskInfoRepo
}

func NewTaskDetail(repo *models.TaskInfoRepo) *TaskDetail {
	var td TaskDetail
	td.repo = repo
	return &td
}

// OnNav ...
func (c *TaskDetail) OnNav(ctx app.Context, url *url.URL) {

	//@todo is there a better way to do this? could i encode ?if=x
	ep := url.EscapedPath()
	s := strings.Split(ep, "/")
	id := s[len(s)-1]

	fmt.Println("TaskID:", id)

	ti, err := c.repo.GetTaskInfo(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.TI = *ti
	fmt.Println(ti)

	// is this called on the correct thread?
}

func (c *TaskDetail) OnMount(ctx app.Context) {
	fmt.Println("TaskDetail mounted")
}

func (c *TaskDetail) OnDismount(ctx app.Context) {
	fmt.Println("TaskDetail OnDismount")
}

// Render ...
// box with border:
// width = that of container - height (?)
func (c *TaskDetail) Render() app.UI {
	return app.Div().Text("TaskDetailX")
}
