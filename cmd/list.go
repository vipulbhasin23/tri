/*
Copyright © 2025 Vipul B

*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/vipulbhasin23/tri/todo"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all todos",
	Long: `List will display a list of all the todos saved to the list.`,
	Run: func(cmd *cobra.Command, args []string) {
		items, err := todo.ReadItems("/Users/vipul/.tridos.json")

		if err != nil {
			log.Printf("%v", err)
		}
		fmt.Println(items)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
