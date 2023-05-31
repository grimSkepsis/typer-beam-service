package main

import (
	"log"
	"net/http"
	"os"
	"typebeast-service/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

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

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
