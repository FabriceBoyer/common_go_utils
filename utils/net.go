package utils

import (
	"fmt"
	"io"
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
