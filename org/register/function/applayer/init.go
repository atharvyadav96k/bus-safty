package applayer

import "github.com/atharvyadav96k/gcp/app"

func Init() *app.App {
	appInstance := new(app.App).Init()
	if err := appInstance.Init(); err != nil {
		panic(err)
	}
	return appInstance
}
