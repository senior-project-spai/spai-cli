/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/senior-project-spai/spai-cli/service"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List images in SPAI",

	Args: cobra.NoArgs,

	RunE: func(cmd *cobra.Command, args []string) error {
		images, err := service.ListImages()
		if err != nil {
			return err
		}

		for _, image := range images {
			isDone := image.Timestamp.Age.Valid && image.Timestamp.Gender.Valid && image.Timestamp.Emotion.Valid && image.Timestamp.FaceRecognition.Valid
			imageTime := time.Unix(0, image.Timestamp.Image*1000000)
			fmt.Printf("ID: %s, Time: %v, Done: %v\n", image.ID, imageTime, isDone)
		}
		return nil
	},
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

	// TODO: Add limit rows flag
}
