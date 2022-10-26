package compare

import "errors"

func NewComparer(s string) (Comparer, error) {
	switch s {
	case "bytes.Equal":
		return NewBytesEqualComparer()
	}

	return nil, errors.New("unknown Comparer type: " + s)
}
