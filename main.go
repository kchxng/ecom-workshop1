package main

import (
	"fmt"
	"github.com/kchxng/ecom-api/configs"
	"github.com/kchxng/ecom-api/configs/db"
	"github.com/kchxng/ecom-api/router"
	"net"
	"net/http"
	"time"
)

func main() {
	configs.InitConfig()
	configs.InitTimeZone()
	db := db.InitDatabase()
	app := router.SetupRoutes(db)
	// app.Run(":8081")

	// Start the server in a goroutine
	go func() {
		listener, err := net.Listen("tcp", ":8081")
		if err != nil {
			fmt.Println("Error starting server:", err)
			return
		}

		// Customize Extract the dynamically assigned port
		PORT := listener.Addr().(*net.TCPAddr).Port
		fmt.Printf("\n[GIN] %s | Server is running on port: %d\n",
			time.Now().Format("2006/01/02 - 15:04:05"),
			PORT,
		)

		if err := http.Serve(listener, app); err != nil {
			fmt.Println("Error serving:", err)
		}
	}()

	// Block the main goroutine to keep the application running
	select {}
}
