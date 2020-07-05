package main

import (
	"log"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	e := echo.New()

	// start graphql
	schema := NewSchema()
	graphqlSchema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: schema.Query(),
	})

	if err != nil {
		logrus.Fatal(err)
	}

	graphQLHandler := handler.New(&handler.Config{
		Schema:   &graphqlSchema,
		GraphiQL: true,
		Pretty:   true,
	})

	group := e.Group("/graphql")
	group.POST("", echo.WrapHandler(graphQLHandler))
	// end graphql

	log.Fatal(e.Start(viper.GetString("server.address")))
}
