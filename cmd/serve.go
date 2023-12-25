/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/shaik80/SalahTimingsBackend/config"
	"github.com/shaik80/SalahTimingsBackend/internal/server"
	"github.com/shaik80/SalahTimingsBackend/utils/logger"
	"github.com/spf13/cobra"

	database "github.com/shaik80/SalahTimingsBackend/internal/db"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "to start a server",
	Long:  ``,
	Run:   ServeFunc,
}

func ServeFunc(cmd *cobra.Command, args []string) {
	// Load configuration
	if err := config.Load(); err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Get configuration values
	addr := config.GetConfig().Server.Address
	port := config.GetConfig().Server.Port

	_, err := database.Initialize()
	if err != nil {
		log.Fatalf("Error starting db: %v", err)
	} else {
		logger.Logger.Printf("DB successfully connect")
	}

	// Start the server
	app := server.New()
	log.Printf("Server is running on %s:%d", addr, port)
	if err := app.Listen(fmt.Sprintf("%s:%d", addr, port)); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
