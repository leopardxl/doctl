package doctl

import "testing"

var filePathTests = []struct {
	filePath string
	mimeType string
}{
	{"testdata/filepath_test1.txt", "text/plain"},
}

func TestFilePath(t *testing.T) {
	for _, test := range filePathTests {
		result, err := MimeType(test.filePath)
		if result != test.mimeType {
			t.Errorf("MimeType: Got %s. Expected %s from file %s",
				result, test.mimeType, test.filePath)
		}
		if err != nil {
			t.Errorf(`MimeType("%s"): %v`, test.filePath, err)
		}
	}
}
