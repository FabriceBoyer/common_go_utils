package utils

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDownloadFile(t *testing.T) {
	output_dir := "test_output/net"
	os.MkdirAll(output_dir, os.ModePerm)

	file_name := "robots.txt"
	file_path := filepath.Join(output_dir, file_name)
	url := "http://google.com/" + file_name

	err := DownloadFile(file_path, url)
	if err != nil {
		t.Error(err)
	}
}
