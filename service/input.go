package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

// PredictImage predict an image by specific an image path
func PredictImage(path string, id string) (ImageID string, err error) {

	// Date for sending
	data := new(bytes.Buffer)
	multipartWriter := multipart.NewWriter(data)

	// image field
	imagePartWriter, err := multipartWriter.CreateFormFile("image", filepath.Base(path))
	if err != nil {
		return "", fmt.Errorf("PredictImage: %w", err)
	}
	imageFileReader, err := os.Open(path)
	defer imageFileReader.Close()
	if err != nil {
		return "", fmt.Errorf("PredictImage: %w", err)
	}
	if _, err := io.Copy(imagePartWriter, imageFileReader); err != nil {
		return "", fmt.Errorf("PredictImage: %w", err)
	}

	// image_id field (optional)
	if id != "" {
		if err := multipartWriter.WriteField("image_id", id); err != nil {
			return "", fmt.Errorf("PredictImage: %w", err)
		}
	}

	multipartWriter.Close()

	// Create request
	req, err := http.NewRequest("POST", "http://image-input-api.spai.svc/_api/image", data)
	if err != nil {
		return "", fmt.Errorf("PredictImage: %w", err)
	}
	req.Header.Set("Content-Type", multipartWriter.FormDataContentType())

	// Submit the request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("PredictImage: %w", err)
	}
	defer func() {
		io.Copy(ioutil.Discard, res.Body)
		res.Body.Close()
	}()

	// Check the response
	if res.StatusCode < 200 || res.StatusCode > 299 {
		return "", fmt.Errorf("PredictImage: response status code: %d", res.StatusCode)
	}

	// Parse body from json
	parsedBody := new(struct {
		ImageID string `json:"image_id"`
	})
	if err := json.NewDecoder(res.Body).Decode(parsedBody); err != nil {
		return "", fmt.Errorf("PredictImage: %w", err)
	}

	return parsedBody.ImageID, nil
}
