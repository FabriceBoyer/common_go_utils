package utils

import "testing"

// func TestSetupConfig(t *testing.T) {
// 	err := SetupConfig()
// 	if err != nil {
// 		t.Error(err)
// 	}
// }

func TestSetupLogger(t *testing.T) {
	err := SetupLogger(true, "test.log")
	if err != nil {
		t.Error(err)
	}
}
