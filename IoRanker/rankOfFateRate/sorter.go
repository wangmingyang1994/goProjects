package rankOfFateRate

type Sorter interface {
	SortRank(*[]Person, int, int)
}

type Bubble struct{}

func (bubble Bubble) SortRank(persons *[]Person, start int, end int) {
	//冒泡排序
	for i := 0; i < len(*persons)-1; i++ {
		for i := 0; i < len(*persons)-1; i++ {
			for j := i; j < len(*persons); j++ {
				if (*persons)[i].FatRate > (*persons)[j].FatRate {
					(*persons)[i], (*persons)[j] = (*persons)[j], (*persons)[i]
				}
			}
		}
	}
}

type Quick struct{}

func (quick Quick) SortRank(persons *[]Person, start int, end int) {
	//快排
	pivotIdx := (start + end) / 2
	pivotV := (*persons)[pivotIdx].FatRate
	l, r := start, end
	for l <= r {
		for (*persons)[l].FatRate < pivotV {
			l++
		}
		for (*persons)[r].FatRate > pivotV {
			r--
		}
		if l >= r {
			break
		}

		(*persons)[l], (*persons)[r] = (*persons)[r], (*persons)[l]
		l++
		r--
	}
	if l == r {
		l++
		r--
	}
	if r > start {
		quick.SortRank(persons, start, r)
	}
	if l < end {
		quick.SortRank(persons, l, end)
	}

}
