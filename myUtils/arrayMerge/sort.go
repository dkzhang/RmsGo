package arrayMerge

type AbsInt64Slice []int64

func (p AbsInt64Slice) Len() int           { return len(p) }
func (p AbsInt64Slice) Less(i, j int) bool { return AbsInt64(p[i]) < AbsInt64(p[j]) }
func (p AbsInt64Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type Int64Slice []int64

func (p Int64Slice) Len() int           { return len(p) }
func (p Int64Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Int64Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func AbsInt64(x int64) int64 {
	if x < 0 {
		return -x
	} else {
		return x
	}
}
