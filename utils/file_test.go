package utils

import (
	"reflect"
	"testing"
)

type JSONData struct {
	Data string `json:"data"`
}

func TestJSON(t *testing.T) {
	type args struct {
		filepath string
		data     *JSONData
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestJSON",
			args: args{
				filepath: "test_output/file/test.json",
				data:     &JSONData{"Test"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := WriteJSON(tt.args.filepath, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("WriteJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err := ReadJSON(tt.args.filepath, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("ReadJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestReadFileInZip(t *testing.T) {
	type args struct {
		zipPath        string
		targetFileName string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "TestReadFileInZip",
			args: args{
				zipPath:        "test_input/file/file.zip",
				targetFileName: "file/test.json",
			},
			want:    "{\n \"data\": \"Test\"\n}",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadFileInZip(tt.args.zipPath, tt.args.targetFileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadFileInZip() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ReadFileInZip() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReadFileFromTarGz(t *testing.T) {
	tarGzFile := "test_input/parsing/arxiv_2103.06299.tar.gz"
	fileName := "bibliography.bib"

	fileContents, err := ReadFileFromTarGz(tarGzFile, fileName)
	if err != nil {
		t.Error(err)
	}

	t.Log("File Contents:", string(fileContents), "\n")
}

type TestData struct {
	Name  string
	Age   int
	Email string
}

func TestSaveLoadGob(t *testing.T) {
	testFile := "test_output/file/test.gob"

	// Create a test data struct
	testData := TestData{
		Name:  "John Doe",
		Age:   30,
		Email: "johndoe@example.com",
	}

	// Serialize the test data to the file
	err := SaveGob(&testData, testFile)
	if err != nil {
		t.Fatalf("Error serializing data: %v", err)
	}

	// Deserialize the test data from the file
	var loadedData TestData
	err = LoadGob(testFile, &loadedData)
	if err != nil {
		t.Fatalf("Error deserializing data: %v", err)
	}

	// Compare the original and loaded data
	if !reflect.DeepEqual(testData, loadedData) {
		t.Errorf("Loaded data does not match the original data")
	}

}
