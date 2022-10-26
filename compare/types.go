package compare

type Comparer interface {
	Compare(file1, file2 string) string
}
