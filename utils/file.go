package utils

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func ReadJSON(filepath string, data any) error {
	file, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, data)
	if err != nil {
		return err
	}

	return nil
}

func WriteJSON(filePath string, data any) error {
	err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
	if err != nil {
		return err
	}

	jsonFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	byteValue, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}

	_, err = jsonFile.Write(byteValue)
	if err != nil {
		return err
	}

	return nil
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func ReadFileInZip(zipPath string, targetFileName string) (string, error) {
	// Open the zip file for reading
	zipFile, err := zip.OpenReader(zipPath)
	if err != nil {
		return "", err
	}
	defer zipFile.Close()

	// Iterate through each file in the zip archive
	for _, file := range zipFile.File {
		// Check if the file name matches the target file name
		if file.Name == targetFileName {
			// Open the file inside the archive
			zipFile, err := file.Open()
			if err != nil {
				return "", err
			}
			defer zipFile.Close()

			// Read the contents of the file
			data, err := io.ReadAll(zipFile)
			if err != nil {
				return "", err
			}

			// Convert the byte slice to a string
			content := string(data)

			return content, nil
		}
	}

	return "", fmt.Errorf("file '%s' not found in zip: '%s'", targetFileName, zipPath)
}

// Returns first found
func ReadFileFromTarGz(tarGzFile, extension string) (string, error) {
	// Open the tar.gz file
	file, err := os.Open(tarGzFile)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Create a gzip reader
	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		return "", err
	}
	defer gzipReader.Close()

	// Create a tar reader
	tarReader := tar.NewReader(gzipReader)

	// Iterate over the files in the tar reader
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}

		// Check if the current file matches the desired file name
		if strings.HasSuffix(header.Name, extension) {
			// Read the contents of the file
			fileContents, err := io.ReadAll(tarReader)
			if err != nil {
				return "", err
			}
			// Convert the byte slice to a string
			content := string(fileContents)

			return content, nil
		}
	}

	return "", fmt.Errorf("file '%s' not found in tar.gz: '%s'", extension, tarGzFile)
}

func SaveGob(data interface{}, filename string) error {

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	gzipWriter := gzip.NewWriter(file)
	defer gzipWriter.Close()

	encoder := gob.NewEncoder(gzipWriter)

	err = encoder.Encode(data)
	if err != nil {
		return err
	}

	return nil
}

func LoadGob(filename string, data interface{}) error {

	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		return err
	}
	defer gzipReader.Close()

	decoder := gob.NewDecoder(gzipReader)

	err = decoder.Decode(data)
	if err != nil {
		return err
	}

	return nil
}
