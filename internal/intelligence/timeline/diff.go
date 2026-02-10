package timeline

func Diff(a, b Snapshot) bool {
	return a.Hash != b.Hash
}
