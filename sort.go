package main

type arr []int32

func (a arr) Sort() []int32 {
	var resArr = make([]int32, len(a))

	for i, l := range a {
		binarySort(resArr, l, int32(i))
	}

	return append(a[:0], resArr...)
}

func binarySort(arr []int32, newEl int32, i int32) {
	if i == 0 {
		arr[0] = newEl
		return
	}

	if i == 1 && arr[0] > newEl {
		arr[0], arr[1] = newEl, arr[0]
		return
	}

	if newEl >= arr[i-1] {
		arr[i] = newEl
		return
	}

	var h int32 = 0
	var m = i / 2

	for {
		switch true {
		case m == h:
			goto RESULT
		case arr[m] > newEl:
			m = (m + h) / 2
		case arr[m] < newEl:
			h = m + 1
			m = (m + i) / 2

		case arr[m] == newEl:
			goto RESULT
		}
	}

RESULT:
	if arr[m] < newEl {
		m += 1
	}

	var newArr []int32

	newArr = append([]int32{}, arr[:m]...)
	newArr = append(newArr, newEl)
	newArr = append(newArr, arr[m:len(arr)-1]...)
	newArr = append(arr[:0], newArr...)

	return
}
