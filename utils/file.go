package utils

import (
	"archive/tar"
	"archive/zip"
	"bufio"
	"bytes"
	"compress/gzip"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func ReadJSON(filename string, data any) error {

	compressed := strings.HasSuffix(filename, ".gz")

	if compressed {

		// Open the gzip file for reading
		file, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer file.Close()

		// Create a gzip reader on top of the file
		gzipReader, err := gzip.NewReader(file)
		if err != nil {
			return err
		}
		defer gzipReader.Close()

		// Create a JSON decoder on top of the gzip reader
		jsonDecoder := json.NewDecoder(gzipReader)

		// Decode the JSON data into the provided data structure
		err = jsonDecoder.Decode(data)
		if err != nil {
			return err
		}

	} else {

		file, err := os.ReadFile(filename)
		if err != nil {
			return err
		}

		err = json.Unmarshal(file, data)
		if err != nil {
			return err
		}
	}

	return nil
}

func WriteJSON(filePath string, data any, prettyPrint bool, compress bool) error {

	err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
	if err != nil {
		return err
	}

	extension := ".json"
	if compress {
		extension += ".gz"
	}

	jsonFile, err := os.Create(AddExtensionIfNotExist(filePath, extension))
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	// Convert the data to JSON
	var jsonData []byte
	if prettyPrint {
		jsonData, err = json.MarshalIndent(data, "", " ")
	} else {
		jsonData, err = json.Marshal(data)
	}
	if err != nil {
		return err
	}

	if compress {
		var compressedData bytes.Buffer

		// Create a gzip writer on top of the buffer
		gzipWriter := gzip.NewWriter(&compressedData)
		_, err = gzipWriter.Write(jsonData)
		if err != nil {
			return err
		}
		gzipWriter.Close() // must close it now, not defer it

		_, err = jsonFile.Write(compressedData.Bytes())
		if err != nil {
			return err
		}

	} else {
		_, err = jsonFile.Write(jsonData)
		if err != nil {
			return err
		}
	}

	return nil
}

func AddExtensionIfNotExist(filePath string, extension string) string {
	// Check if the file path already has an extension
	if !strings.HasSuffix(filePath, extension) {
		// Add the extension to the file path
		filePath += extension
	}
	return filePath
}

func RemoveExtension(filePath string) string {
	fileName := path.Base(filePath)
	fileExt := path.Ext(fileName)
	fileNameWithoutExt := fileName[:len(fileName)-len(fileExt)]
	return fileNameWithoutExt
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func EnsureFolderExistForFile(filePath string) error {
	err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
	if err != nil {
		return err
	}
	return nil
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
		if path.Base(file.Name) == path.Base(targetFileName) {
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

func SaveGob(data any, filename string) error {

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

func LoadGob(filename string, data any) error {

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

func WriteStringArrayToFile(strArray []string, filename string) error {
	// Create a new file
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a buffered writer to write to the file
	writer := bufio.NewWriter(file)

	// Write each string in the array to the file
	for _, str := range strArray {
		_, err := writer.WriteString(str + "\n")
		if err != nil {
			return err
		}
	}

	// Flush the buffer to ensure all data is written
	err = writer.Flush()
	if err != nil {
		return err
	}

	return nil
}
