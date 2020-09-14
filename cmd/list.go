package cmd

import (
	"fmt"
	"time"

	"github.com/senior-project-spai/spai-cli/service"
	"github.com/spf13/cobra"
)

var limit uint

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List images in SPAI",

	Args: cobra.NoArgs,

	RunE: func(cmd *cobra.Command, args []string) error {

		images, err := service.ListImages(limit)
		if err != nil {
			return err
		}

		for _, image := range images {

			imageTime := time.Unix(0, image.Timestamp.Image*1000000)
			fmt.Printf("ID: %s, Time: %v", image.ID, imageTime)

			// Print complete time if all result have timestamp
			if isDone := image.Timestamp.Age.Valid && image.Timestamp.Gender.Valid && image.Timestamp.Emotion.Valid && image.Timestamp.FaceRecognition.Valid; isDone {
				fmt.Printf(", Done Time: %v\n", time.Unix(0, findMax([]int64{image.Timestamp.Age.Int64, image.Timestamp.Gender.Int64, image.Timestamp.Emotion.Int64, image.Timestamp.FaceRecognition.Int64})*1000000))
			} else {
				fmt.Println()
			}
		}
		return nil
	},
}

func findMax(a []int64) (max int64) {
	max = a[0]
	for _, value := range a {
		if value > max {
			max = value
		}
	}
	return max
}

func init() {
	imagesCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	listCmd.Flags().UintVarP(&limit, "limit", "n", 0, "Limit number of results")
	// TODO: Add limit rows flag
}
