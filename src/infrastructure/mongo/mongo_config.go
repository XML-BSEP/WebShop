package mongo

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	_ "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	_ "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	_ "go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"time"
)

func init_viper() {
	viper.SetConfigFile(`src/configurations/mongo.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func GetDbName() string{
	init_viper()
	return viper.GetString(`database`)
}
func GetDocker() bool{
	init_viper()
	return viper.GetBool(`docker`)
}
func NewMongoClient() (*mongo.Client, *context.Context){
	init_viper()
	var mongo_uri string
	if os.Getenv("DOCKER_ENV") != "" {
		mongo_uri = viper.GetString(`mongodb_uri_docker`)
	}else{
		mongo_uri = viper.GetString(`mongodb_uri_localhost`)
	}

	clientOptions := options.Client().ApplyURI(mongo_uri)
	client, err := mongo.NewClient(clientOptions)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return client, &ctx
}