package bootstrap

import "go.mongodb.org/mongo-driver/mongo"

type Application struct {
	Env   *Env
	Mongo *mongo.Client
}

func NewApp() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Mongo = NewDb(app.Env)
	return *app
}

func (app *Application) CloseDbConnection() {
	CloseMongoDBConnection(app.Mongo)
}
