package applayer

import "github.com/atharvyadav96k/gcp/app"

func Init() *app.App {
	appInstance := new(application).Init()
	appInstance.InitEnvironmentVariables()
	if err := appInstance.InitNeon(); err != nil {
		panic(err)
	}
	return appInstance
}
