package arguments

type Arguments struct {
	err         error
	programName string
	copyright   string
	args        map[string]ArgumentDescriptor[int64, uint64, float64, string, bool]
}
