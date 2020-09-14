package cmd

import (
	"fmt"
	"time"

	"github.com/senior-project-spai/spai-cli/service"
	"github.com/spf13/cobra"
)

// imagesCmd represents the images command
var imagesCmd = &cobra.Command{
	Use:   "images",
	Short: "Action for images",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		image, err := service.GetImage(args[0])
		if err != nil {
			return err
		}

		// Print result
		fmt.Println("ID:", image.ID)
		fmt.Println("Path:", image.Path)
		fmt.Println("Input Timestamp:", time.Unix(0, image.Timestamp.Image*1000000))

		// Age
		fmt.Print("Age Timestamp: \t\t\t")
		switch image.Timestamp.Age.Valid {
		case true:
			fmt.Println(time.Unix(0, image.Timestamp.Age.Int64*1000000))
		case false:
			fmt.Println("Null")
		}

		// Gender
		fmt.Print("Gender Timestamp: \t\t")
		switch image.Timestamp.Gender.Valid {
		case true:
			fmt.Println(time.Unix(0, image.Timestamp.Gender.Int64*1000000))
		case false:
			fmt.Println("Null")
		}

		// Emotion
		fmt.Print("Emotion Timestamp: \t\t")
		switch image.Timestamp.Emotion.Valid {
		case true:
			fmt.Println(time.Unix(0, image.Timestamp.Emotion.Int64*1000000))
		case false:
			fmt.Println("Null")
		}

		// Face Recognition
		fmt.Print("Face Recognition Timestamp: \t")
		switch image.Timestamp.FaceRecognition.Valid {
		case true:
			fmt.Println(time.Unix(0, image.Timestamp.FaceRecognition.Int64*1000000))
		case false:
			fmt.Println("Null")
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(imagesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// imagesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// imagesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
