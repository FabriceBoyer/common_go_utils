package utils

import (
	"os"
	"reflect"
	"testing"
)

// func TestMain(m *testing.M) {
// 	os.MkdirAll("test_output", os.ModePerm)
// }

type JSONData struct {
	Data string `json:"data"`
}

func TestJSON(t *testing.T) {
	type args struct {
		filepath    string
		data        *JSONData
		prettyPrint bool
		compressed  bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestJSON",
			args: args{
				filepath:    "test_output/file/test.json",
				data:        &JSONData{"Test"},
				prettyPrint: false,
				compressed:  false,
			},
			wantErr: false,
		},
		{
			name: "Test Gz JSON ",
			args: args{
				filepath:    "test_output/file/test.json.gz",
				data:        &JSONData{"Test"},
				prettyPrint: true,
				compressed:  true,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := WriteJSON(tt.args.filepath, tt.args.data, tt.args.prettyPrint, tt.args.compressed); (err != nil) != tt.wantErr {
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

func TestWriteStringArrayToFile(t *testing.T) {
	// Example string array
	strArray := []string{"Hello", "World", "Knowledge"}

	// Temporary test file
	tmpFile := "test_output/file/string_array.txt"

	// Call the function to write the string array to a file
	err := WriteStringArrayToFile(strArray, tmpFile)
	if err != nil {
		t.Fatalf("Failed to write string array to file: %v", err)
	}

	// Read the contents of the test file
	fileData, err := os.ReadFile(tmpFile)
	if err != nil {
		t.Fatalf("Failed to read test file: %v", err)
	}

	// Expected file contents
	expectedData := "Hello\nWorld\nKnowledge\n"

	// Check if the file contents match the expected data
	if !reflect.DeepEqual(string(fileData), expectedData) {
		t.Errorf("Unexpected file contents.\nExpected: %s\nActual: %s", expectedData, fileData)
	}

}

func TestAddExtensionIfNotExist(t *testing.T) {
	type args struct {
		filePath  string
		extension string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Already exists",
			args: args{
				filePath:  "test_output/file/test.txt",
				extension: ".txt",
			},
			want: "test_output/file/test.txt",
		},
		{
			name: "Missing first extension",
			args: args{
				filePath:  "test_output/file/test",
				extension: ".txt",
			},
			want: "test_output/file/test.txt",
		},
		{
			name: "Missing second extension",
			args: args{
				filePath:  "test_output/file/test.txt",
				extension: ".gz",
			},
			want: "test_output/file/test.txt.gz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddExtensionIfNotExist(tt.args.filePath, tt.args.extension); got != tt.want {
				t.Errorf("AddExtensionIfNotExist() = %v, want %v", got, tt.want)
			}
		})
	}
}
