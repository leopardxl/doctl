package doctl

import "testing"

var mimeTypeTests = []struct {
	filePath string
	mimeType string
}{
	{"testdata/filepath_test1.txt", "text/plain"},
}

var objectNameTests = []struct {
	filePath   string
	objectName string
}{
	{"/this/isa/fake/file/path", "path"},
	{"this/is/a/fake/relative/path/test", "test"},
	{"justafile", "justafile"},
	{"/this/isa/directory/", ""},
}

func TestMimeType(t *testing.T) {
	for _, test := range mimeTypeTests {
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

func TestObjectName(t *testing.T) {
	for _, test := range objectNameTests {
		result := ObjectName(test.filePath)
		if result != test.objectName {
			t.Errorf(`ObjectName("%s"): "%s", want "%s"`,
				test.filePath, result, test.objectName)
		}
	}
}
