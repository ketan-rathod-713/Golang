package api

import "task6MuxGorm/app"

type Api struct {
	App *app.App
	// Other Services
}

func NewApi(app *app.App) (myApi *Api, err error) {
	myApi = &Api{
		App: app,
	}

	return myApi, nil
}
