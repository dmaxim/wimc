package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dmaxim/wimc/cloudResource"
	"github.com/dmaxim/wimc/database"
	_ "github.com/go-sql-driver/mysql"
)

func resourceTypeHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("resource type endpoint called"))
}

func middelwareHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("before handler; middleware start")
		start := time.Now()
		handler.ServeHTTP(writer, request)
		fmt.Printf("middlware finished; %s", time.Since(start))
	})
}

const apiBasePath = "/api"

func main() {
	database.SetupDatabase()
	cloudResource.SetupRoutes(apiBasePath)
	http.HandleFunc("/resourcetypes", resourceTypeHandler)
	http.ListenAndServe(":5000", nil)
}
