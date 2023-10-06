package main

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"baseProject/logger"
)

const webPort = ":80" // 要記得 ":" ，很重要阿，沒有的話就不會通

type Config struct{}

func main() {
	logger.Init()
	// s := api.NewServer()

	client, err := GetMongoDBConnection("mongodb://admin:password@pt-mongo:27017")
	if err != nil {
		logger.L.Fatal(err.Error())
	}

	// 關閉 MongoDB 連線
	defer func() {
		if err = client.Disconnect(context.Background()); err != nil {
			logger.L.Fatal(err.Error())
		}
	}()

	// eRepo := _eRepo.NewMongoEmployeeRepository(client)
	// eUsecase := _eUsecase.NewEmployeeUsecase(eRepo)
	// _eDelivery.NewEmployeeHandler(s.Router, eUsecase)

	// esRepo := _esRepo.NewMongoESRepo(client)
	// esUsecase := _esUsecase.NewEAUsecase(esRepo)
	// _esDelivery.NewESHandler(s.Router, esUsecase)

	// err = s.Router.Run(webPort) // 這邊一啟動後，就不能再加 router 了，所以放在最後面
	// if err != nil {
	// 	logger.L.Panic(err.Error())

	// }
}

func GetMongoDBConnection(uri string) (*mongo.Client, error) {
	// 設置客戶端選項
	clientOptions := options.Client().ApplyURI(uri)

	// 建立客戶端
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	// 建立連線上下文
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 連線 MongoDB
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	// 確保連線關閉
	// 不需要 defer 關閉連線

	// 測試連線
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	// 回傳連線
	return client, nil
}
