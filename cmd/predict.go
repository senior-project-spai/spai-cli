package cmd

import (
	"fmt"

	"github.com/senior-project-spai/spai-cli/service"
	"github.com/spf13/cobra"
)

// predictCmd represents the predict command
var predictCmd = &cobra.Command{
	Use:   "predict",
	Short: "predict an image (put image into SPAI)",
	Args:  cobra.RangeArgs(1, 2),

	RunE: func(cmd *cobra.Command, args []string) error {
		var (
			err     error
			imageID string
		)
		switch len(args) {
		case 1:
			imageID, err = service.PredictImage(args[0], "")
		case 2:
			imageID, err = service.PredictImage(args[0], args[1])
		default:
			return fmt.Errorf("predictCmd: number of arguments is not 1 or 2")
		}

		if err != nil {
			return err
		}

		// ImageID from Response
		fmt.Println("ImageID:", imageID)

		return nil
	},
}

func init() {
	imagesCmd.AddCommand(predictCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// predictCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// predictCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
