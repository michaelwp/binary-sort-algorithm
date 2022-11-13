package main

import (
	"reflect"
	"testing"
)

func Test_arr_Sort(t *testing.T) {
	tests := []struct {
		name string
		a    arr
		want []int32
	}{
		{
			name: "test1",
			a:    arr{27, 35, 2, 1, 35, 45, 5, 41, 20, 17, 3, 10},
			want: arr{1, 2, 3, 5, 10, 17, 20, 27, 35, 35, 41, 45},
		},
		{
			name: "test2",
			a:    arr{38, 31, 13, 14, 42, 9, 21, 1, 37, 38, 26, 14, 50, 23},
			want: arr{1, 9, 13, 14, 14, 21, 23, 26, 31, 37, 38, 38, 42, 50},
		},
		{
			name: "test3",
			a:    arr{15, 13, 4, 26, 29, 10, 50, 19, 8, 44, 42, 16, 47, 16},
			want: arr{4, 8, 10, 13, 15, 16, 16, 19, 26, 29, 42, 44, 47, 50},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.a.Sort(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_binarySort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		list := []int32{
			27, 35, 2, 1, 35, 45, 5, 41, 20, 17, 3, 10,
			38, 31, 13, 14, 42, 9, 21, 1, 37, 38, 26, 14, 50, 23,
			29, 23, 22, 4, 15, 9, 37, 40, 43, 10, 45, 43, 25,
			30, 44, 22, 27, 18, 28, 9, 9, 1, 35, 7, 40,
			15, 13, 4, 26, 29, 10, 50, 19, 8, 44, 42, 16, 47, 16,
			8, 27, 40, 43, 34, 28, 2, 37, 37, 39, 33, 11,
			47, 19, 3, 23, 12, 38, 46, 48, 40, 3, 32, 37, 3,
			34, 22, 28, 31, 23, 40, 16, 46, 38, 45,
		}

		arr(list).Sort()
	}
}
