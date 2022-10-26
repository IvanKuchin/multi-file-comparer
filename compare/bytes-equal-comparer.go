package compare

import (
	"bytes"
	"io/ioutil"
)

type BytesEqualComparer struct {
}

func (obj *BytesEqualComparer) Compare(file1, file2 string) string {
	result := ""

	f1, err1 := ioutil.ReadFile(file1)

	if err1 != nil {
		result += "ERROR (" + file1 + ")" + err1.Error() + "\n"
	}

	f2, err2 := ioutil.ReadFile(file2)

	if err2 != nil {
		result += "ERROR (" + file2 + ")" + err2.Error() + "\n"
	}

	if !bytes.Equal(f1, f2) {
		result += "diff " + file1 + " " + file2 + "\n"
	}

	return result
}

func NewBytesEqualComparer() (Comparer, error) {
	return &BytesEqualComparer{}, nil
}
