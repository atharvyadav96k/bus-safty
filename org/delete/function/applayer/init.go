package applayer

import "github.com/atharvyadav96k/gcp/app"

func Init() *app.App {
	app := application.Init(application{})
	app.InitEnvironmentVariables()
	if err := app.InitFirestore(app.Env.GetSecret("PROJECT_ID")); err != nil {
		panic(err)
	}
	return app
}
