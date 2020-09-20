package arrayMerge

import (
	"fmt"
	"sort"
)

func ComputeChange(before []int64, after []int64) (change []int64, increased int, reduced int, err error) {
	// nil check
	if before == nil {
		before = make([]int64, 0)
	}
	if after == nil {
		after = make([]int64, 0)
	}

	// sort array
	sort.Sort(Int64Slice(before))
	sort.Sort(Int64Slice(after))

	// negative check
	if len(before) != 0 && before[0] < 0 {
		return nil, 0, 0,
			fmt.Errorf("all number in before array must be positive")
	}
	if len(after) != 0 && after[0] < 0 {
		return nil, 0, 0,
			fmt.Errorf("all number in after array must be positive")
	}

	// duplicate check
	for x := 0; x < len(before)-1; x++ {
		if before[x] == before[x+1] {
			return nil, 0, 0,
				fmt.Errorf("all number in before array must be no duplicated")
		}
	}
	for x := 0; x < len(after)-1; x++ {
		if after[x] == after[x+1] {
			return nil, 0, 0,
				fmt.Errorf("all number in after array must be no duplicated")
		}
	}

	// init result
	change = make([]int64, 0)
	increased = 0
	reduced = 0

	// two index scan
	i, j := 0, 0
	for i < len(before) && j < len(after) {
		if before[i] < after[j] {
			// remove
			change = append(change, -before[i])
			i++
			reduced++
		} else if before[i] == after[j] {
			// no change
			i++
			j++
		} else {
			// add
			change = append(change, after[j])
			j++
			increased++
		}
	}
	if i == len(before) {
		for jt := j; jt < len(after); jt++ {
			change = append(change, after[jt])
			increased++
		}
	} else {
		for it := i; it < len(before); it++ {
			change = append(change, -before[it])
			reduced++
		}
	}

	return change, increased, reduced, nil
}

func ApplyBeforeChange(before []int64, change []int64) (after []int64, err error) {
	if change == nil || len(change) == 0 {
		// no change
		if before == nil {
			return []int64{}, nil
		} else {
			return before, nil
		}
	}

	after = make([]int64, 0)
	sort.Sort(AbsInt64Slice(change))

	if before == nil || len(before) == 0 {
		for j, c := range change {
			if c > 0 {
				after = append(after, c)
			} else {
				return nil,
					fmt.Errorf("negative value change[%d]=%d is illegal since before array is empty", j, c)
			}
		}
		return after, nil
	}

	sort.Sort(Int64Slice(before))

	// positive&zero check for before array
	if before[0] <= 0 {
		return nil,
			fmt.Errorf("negative/zero value %d in before array is illegal", before[0])
	}
	if change[0] == 0 {
		return nil,
			fmt.Errorf("zero value %d in change array is illegal", change[0])
	}

	// two index scan
	i, j := 0, 0
	for i < len(before) && j < len(change) {
		if before[i] < AbsInt64(change[j]) {
			after = append(after, before[i])
			i++
		} else if before[i] == AbsInt64(change[j]) {
			if change[j] < 0 {
				i++
				j++
				continue
			} else {
				return nil,
					fmt.Errorf("positive value change[%d]=%d is illegal since it is already existed in the before array, change=%v", j, change[j], change)
			}
		} else {
			if change[j] > 0 {
				after = append(after, change[j])
				j++
			} else {
				return nil,
					fmt.Errorf("negative value change[%d]=%d is illegal since it is not exist in the before array", j, change[j])
			}
		}
	}
	if i == len(before) {
		for jt := j; jt < len(change); jt++ {
			if change[jt] > 0 {
				after = append(after, change[jt])
			} else {
				return nil,
					fmt.Errorf("negative value change[%d]=%d is illegal since before array is empty", jt, change[jt])
			}
		}
	} else {
		for it := i; it < len(before); it++ {
			after = append(after, before[it])
		}
	}

	return after, nil
}

func ApplyBeforeSubtract(before []int64, subtract []int64) (after []int64, err error) {
	change := make([]int64, len(subtract))
	for i, s := range subtract {
		if s <= 0 {
			return nil, fmt.Errorf("subtract positive check failed: subtract[%d]=%d", i, s)
		}
		change[i] = -s
	}

	return ApplyBeforeChange(before, change)
}
