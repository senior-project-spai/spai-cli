package cmd

import (
	"fmt"

	"github.com/senior-project-spai/spai-cli/service"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete all related image data from SPAI with the specific image ID",
	Args:  cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		if err := service.RemoveImageFromDB(args[0]); err != nil {
			return fmt.Errorf("deleteCmd: %w", err)
		}

		if err := service.RemoveImageFromS3(args[0]); err != nil {
			return fmt.Errorf("deleteCmd: %w", err)
		}

		return nil
	},
}

func init() {
	imagesCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
