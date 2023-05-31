package main

import (
	"log"
	"net/http"
	"os"
	"typebeast-service/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/clerkinc/clerk-sdk-go/clerk"
	"go.uber.org/zap"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	logger.Info("Hello, world!")

	// token := os.Getenv("INFLUXDB_TOKEN")
	// url := "http://localhost:8086"
	// client := influxdb2.NewClient(url, token)

	// //influx stuff
	// org := "grim co"
	// bucket := "WPM"
	// writeAPI := client.WriteAPIBlocking(org, bucket)
	// for value := 0; value < 5; value++ {
	// 	tags := map[string]string{
	// 		"tagname1": "tagvalue1",
	// 	}
	// 	fields := map[string]interface{}{
	// 		"field1": value,
	// 	}
	// 	point := write.NewPoint("measurement1", tags, fields, time.Now())
	// 	time.Sleep(1 * time.Second) // separate points by 1 second

	// 	if err := writeAPI.WritePoint(context.Background(), point); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }

	// ///
	// queryAPI := client.QueryAPI(org)
	// query := `from(bucket: "WPM")
	//         |> range(start: -10m)
	//         |> filter(fn: (r) => r._measurement == "measurement1")`
	// results, err := queryAPI.Query(context.Background(), query)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// // for results.Next() {
	// // 	fmt.Println(results.Record())
	// // }
	// if err := results.Err(); err != nil {
	// 	log.Fatal(err)
	// }
	// ///
	// query = `from(bucket: "WPM")
	// |> range(start: -10m)
	// |> filter(fn: (r) => r._measurement == "measurement1")
	// |> mean()`
	// results, err = queryAPI.Query(context.Background(), query)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for results.Next() {
	// 	fmt.Println(results.Record())
	// }
	// if err := results.Err(); err != nil {
	// 	log.Fatal(err)
	// }
	///

	//CLERK STUFF
	apiKey := os.Getenv("CLERK_API_KEY")

	client, err := clerk.NewClient(apiKey)
	if err != nil {
		// handle error
	}

	// List all users for current application
	users, err := client.Users().ListAll(clerk.ListAllUsersParams{})

	logger.Info("Users!", zap.Reflect("users", users))

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
