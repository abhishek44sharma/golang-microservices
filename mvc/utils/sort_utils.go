package utils

import "sort"

func BubbleSort(elements []int) {
	keepRunnelementsg := true
	for keepRunnelementsg {
		keepRunnelementsg = false

		for i := 0; i < len(elements)-1; i++ {
			if elements[i] > elements[i+1] {
				elements[i], elements[i+1] = elements[i+1], elements[i]
				keepRunnelementsg = true
			}
		}
	}
	return
}

func Sort(els []int) {
	if len(els) < 1000 {
		BubbleSort(els)
		return
	}
	sort.Ints(els)
}
