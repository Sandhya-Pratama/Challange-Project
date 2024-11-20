package tests

import (
	"testing"

	"github.com/Sandhya-Pratama/TEST-BACK-END_Sandhya/data"
)

func TestLoadData(t *testing.T) {
	data, err := data.LoadData("../data.txt")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(data) == 0 {
		t.Fatalf("Expected data to be loaded, got empty slice")
	}

	expectedCode := "rv1"
	if data[0].Code != expectedCode {
		t.Errorf("Expected first data code to be %v, got %v", expectedCode, data[0].Code)
	}
}
