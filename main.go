package main

import (
	"fmt"
	actions "instasafe/action"
	"instasafe/env"
	"instasafe/logs"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"
)

var (
	apiHost string         // API host
	apiPort string         // API port
	err     error          // error variable for error handling
	g       errgroup.Group // Initialize the error group for goroutines
)

func main() {
	//fmt.Println(env.Env["OWNER_NAME"])
	apiHost = env.Env["API_HOST"]
	apiPort = env.Env["API_PORT"]

	// Start the API server in a goroutine
	g.Go(func() error {
		return startAPIServer() // Start the API server
	})
	// Wait for the API server to finish
	if err = g.Wait(); err != nil {
		log.Fatalln(err) // Exit with error
	}

}

// startAPIServer starts the API server
func startAPIServer() (err error) {
	var (
		router    *gin.Engine  // gin router
		apiServer *http.Server // HTTP server
	)

	gin.SetMode(gin.ReleaseMode) // Set the gin mode to release
	gin.DisableConsoleColor()    // Disable the console color

	router = gin.Default()              // Initialize the router
	router.RedirectFixedPath = true     // Enable the fixed path redirect
	router.RedirectTrailingSlash = true // Enable the trailing slash redirect

	// Add the gin middleware
	router.Use(generateTransId())
	// API Endpoints
	router.GET("/statistics", actions.Stat)
	router.POST("/transactions", actions.Transactions)
	router.DELETE("/transactions", actions.Delete)

	// apierver contains instance details for the server
	apiServer = &http.Server{
		Addr:         apiHost + ":" + apiPort, // Set the address (host:port)
		Handler:      router,                  // Set the router
		ReadTimeout:  60 * time.Second,        // Set the read timeout
		WriteTimeout: 60 * time.Second,        // Set the write timeout
	}

	fmt.Println("Starting API server on " + apiHost + ":" + apiPort)
	logs.Info("", "starting api server on "+apiHost+":"+apiPort, "")
	// Listen and serve on the host and port specified
	return apiServer.ListenAndServe()
}

// generateTransId is used to generate a unique transaction id for each request
func generateTransId() gin.HandlerFunc {
	return func(c *gin.Context) {
		UUID, _ := uuid.NewRandom() // Generate a new UUID
		c.Set("transId", UUID.String())
	}
}
