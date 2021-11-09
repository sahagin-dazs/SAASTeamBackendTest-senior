package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"saasteamtest/saasbackend/internal"
	"github.com/spf13/cobra"

	"saasteamtest/saasbackend/logging"
	log "github.com/sirupsen/logrus"

	
)

func main() {
	log.SetFormatter(&log.TextFormatter{})
	logging.SetLogOutput()
	level, err := log.ParseLevel("debug")
	if err != nil {
		fmt.Println("parse log level:", err)
		os.Exit(1)
	}
	log.SetLevel(level)

	if err := rootCmd.Execute(); err != nil {
		log.Error("Error calling rootCmd.Execute():", err)
		os.Exit(1)
	}
	if err := recover(); err != nil {
		fmt.Println("Error:", err)
	}
}

var rootCmd = &cobra.Command{
	Use:          "saasbackend",
	Short:        "Serve the SAAS Backend Test API",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		logging.LogInfo(ctx, "start api", "SAAS Backend Test running on port 8000")
		log.Fatal(http.ListenAndServe(
			fmt.Sprintf(":8000"),
			internal.RouterInitializer(),
		))
		return nil
	},
}

