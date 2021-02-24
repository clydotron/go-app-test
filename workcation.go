package main

import "github.com/maxence-charriere/go-app/v7/pkg/app"

type workcation struct {
	app.Compo
}

func (c *workcation) Render() app.UI {

	return app.Div().
		Class("grid lg:grid-cols-2 bg-purple-200").
		Body(
			app.Div().Class("px-8 py-12 max-w-md mx-auto sm:max-w-xl lg:max-w-full").Body(
				app.Img().Src("/web/img/logo.svg").Class("p-2 h-12").Alt("Workcation"),
				app.Img().Src("/web/img/beach-work.jpg").Class("mt-6 rounded-lg shadow-xl sm:mt-8 sm:h-64 sm:w-full sm:object-cover sm:object-center lg:hidden").Alt("on beach"),
				app.H1().Class("mt-6 text-2xl font-bold text.gray-900 lg:text-3xl").Body(
					app.Text("You can work from anywhere. "),
					app.Br().Class("hidden lg:inline"),
					app.Span().Text("Take advantage of it.").Class("text-indigo-500"),
				),
				app.P().
					Text("Workcation helps ypu find work-friendly rentals in beautiful locations so you can enjoy some nice weather, even when you are not on vacation.").
					Class("mt-2 text-gray-600 sm:text-xl sm:mt-4"),
				app.Div().Class("mt-4").Body(
					app.A().Href("#").Text("Book your escape").Class("inline-block px-5 py-2 bg-indigo-500 focus:outline-none focus:ring-offset-2 focus:ring-opacity-50 focus:ring-indigo-500 focus:ring hover:bg-indigo-400 hover:-translate-y-0.5 transform transition text-white rounded-lg shadow-lg uppercase tracking-wider text-semibold sm:mt-6"),
				),
			),
			app.Div().Class("hidden relative lg:block").Body(
				app.Img().Src("/web/img/beach-work.jpg").Class("absolute inset-0 x-full h-full object-cover object-center").Alt("on beach"),
			),
		)
}

//npx tailwindcss-cli buildcss/tailwind.css -o web/tailwindx.css
