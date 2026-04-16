package applyaer

import "github.com/atharvyadav96k/gcp/app"

func Init() *app.App {
	app := application.Init(application{})
	app.InitEnvironmentVariables()
	err := app.InitFirestore()
	if err != nil {
		panic(err)
	}
	return app
}
