/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/aabdelrhmann1/rssagg/internal/database"
	"github.com/aabdelrhmann1/rssagg/internal/server"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the RSS aggregation server",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := godotenv.Load(".env"); err != nil && !errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("load environment: %w", err)
		}

		portString := os.Getenv("PORT")
		if portString == "" {
			return fmt.Errorf("PORT is not found in the environment")
		}

		dbURL := os.Getenv("DB_URL")
		if dbURL == "" {
			return fmt.Errorf("DB_URL is not found in the environment")
		}

		conn, err := sql.Open("postgres", dbURL)
		if err != nil {
			return fmt.Errorf("open database: %w", err)
		}
		defer conn.Close()

		db := database.New(conn)
		srv := server.New(db)
		return srv.Run(":" + portString)
	},
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
