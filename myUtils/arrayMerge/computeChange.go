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

//func ApplyBeforeChange(before []int64, change []int64)(after []int64, err error){
//	after := make([]int, 0)
//
//	sort.Sort(Int64Slice(before))
//	sort.Sort(AbsInt64Slice(change))
//
//
//	for i,j:=0,0;i<len(before)&&j<len(change);{
//		if before[i] < AbsInt64(change[j]){
//
//		}else if before[i] == AbsInt64(change[j]){
//
//		}else{
//
//		}
//	}
//}
