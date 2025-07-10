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

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new todo",
	Long:  `Add will create a new todo item to the list`,
	Run:   addRun,
}

func addRun(cmd *cobra.Command, args []string) {
	// fmt.Println("add called")
	// items := []todo.Item{}
	items, err := todo.ReadItems("/Users/vipul/.tridos.json")
	if err != nil {
		log.Printf("%v", err)
	}

	for _, x := range args {
		items = append(items, todo.Item{Text: x})
	}

	err = todo.SaveItems("/Users/vipul/.tridos.json", items)
	if err != nil {
		fmt.Printf("%v", err)
	}
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
