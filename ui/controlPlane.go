package ui

import "github.com/maxence-charriere/go-app/v7/pkg/app"

// ControlPlane ...
type ControlPlane struct {
	app.Compo
}

// Render ...
func (c *ControlPlane) Render() app.UI {
	return app.Div().
		Class("bg-purple-400").
		Body(
			app.Text("ControlPlane"),
		)
}
