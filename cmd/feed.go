/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/aabdelrhmann1/rssagg/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
)

// feedCmd represents the feed command
var feedCmd = &cobra.Command{
	Use:   "feed",
	Short: "This shows u a feed",
	Long:  `shows a feed for you from all users`,
	Run: func(cmd *cobra.Command, args []string) {
		godotenv.Load(".env")
		dbURL := os.Getenv("DB_URL")
		if dbURL == "" {
			log.Fatal("dbURL is not found in the environment")
		}

		conn, err := sql.Open("postgres", dbURL)
		if err != nil {
			log.Fatal("can't connect to database:", err)
		}
		defer conn.Close()
		feeds, err := database.New(conn).GetFeeds(context.Background())
		if err != nil {
			log.Fatal("can't catch from database:", err)
		}

		if len(feeds) == 0 {
			fmt.Printf("no feeds to show")
			return
		}
		fmt.Printf("ID\tNAME\tURL\n")
		for _, feed := range feeds {
			fmt.Printf("%s\t%s\t%s\n", feed.ID, feed.Name, feed.Url)
		}
	},
}

func init() {
	rootCmd.AddCommand(feedCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// feedCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// feedCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
