package main

import (
	"github.com/liampm/journal/src/cli"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "journal",
	Short: "Manage your journal",
	Long:  "TODO",
}

var findWithTagCmd = &cobra.Command{
	Use:   "find",
	Short: "Find with tag",
	Long:  "TODO - this should probably be something like 'journal entry find -t foo'",
	Run: func(cmd *cobra.Command, args []string) {
		cli.ListFilesWithTag(tag)
	},
}

var tag string


func init() {

	findWithTagCmd.Flags().StringVarP(&tag, "tag", "t", "", "The tag to search by")
	rootCmd.AddCommand(findWithTagCmd)
}

func main() {
	rootCmd.Execute()
}
