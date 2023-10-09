package utils

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func DownloadFile(filePath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s for url '%s'", resp.Status, url)
	}

	err = EnsureFolderExistForFile(filePath)
	if err != nil {
		return err
	}

	// Create the file
	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	//TODO check for consistency/hashsum

	return nil
}

func MakeMultipartPostRequest(url string, formData map[string]string, fileParamName, filePath string) (*http.Response, error) {
	// Create a new multipart buffer
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Add the form fields to the multipart form
	for key, value := range formData {
		_ = writer.WriteField(key, value)
	}

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create a new form file field
	filePart, err := writer.CreateFormFile(fileParamName, filePath)
	if err != nil {
		return nil, err
	}

	// Copy the file content to the form file field
	_, err = io.Copy(filePart, file)
	if err != nil {
		return nil, err
	}

	// Close the multipart writer
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	// Create a new HTTP request with the multipart form
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}

	// Set the Content-Type header to multipart/form-data
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Send the request and get the response
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func SaveResponseBodyToFile(response *http.Response, filePath string) error {
	// Create a new file
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Copy the response body to the file
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}
