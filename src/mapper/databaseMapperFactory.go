package mapper

import (
	"BGP/src/config"
	"BGP/src/exception"
	"BGP/src/util"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func InitConnect() (E error) {
	defer func() {
		r := recover()
		if r != nil {
			E = exception.NewSystemError("initConnect-mapper", util.Strval(r))
		}
	}()

	// mongodb
	clientOptions := options.Client().ApplyURI(config.Database.DatabaseType + "://" +
		config.Database.DatabaseUser + ":" + config.Database.DatabasePassword +
		"@" + config.Database.DatabaseHost + ":" + config.Database.DatabasePort)
	println(config.Database.DatabaseType + "://" +
		config.Database.DatabaseUser + ":" + config.Database.DatabasePassword +
		"@" + config.Database.DatabaseHost + ":" + config.Database.DatabasePort)

	var err error
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return exception.NewDataBaseError("initConnect-mapper", "mongodb connect error")
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return exception.NewDataBaseError("initConnect-mapper", "mongodb ping error")
	}
	util.Loglevel(util.Info, "initConnect-mapper", "mongodb connect success")

	return nil
}

func CloseConnect() (E error) {
	defer func() {
		r := recover()
		if r != nil {
			E = exception.NewSystemError("CloseConnect-mapper", util.Strval(r))
		}
	}()

	err := client.Disconnect(context.Background())
	if err != nil {
		return exception.NewDataBaseError("CloseConnect-mapper", "mongodb disconnect error")
	}
	util.Loglevel(util.Info, "CloseConnect-mapper", "mongodb disconnect success")

	return nil
}
