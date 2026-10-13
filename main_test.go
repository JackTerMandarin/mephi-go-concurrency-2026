package main

import (
	"reflect"
	"testing"
)

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name       string
		size       int
		wantLength int
	}{
		{"negative size", -5, 0},
		{"zero size", 0, 0},
		{"one element", 1, 1},
		{"many elements", 1000, 1000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := generateRandomElements(tt.size)
			if len(data) != tt.wantLength {
				t.Fatalf("length = %d, want %d", len(data), tt.wantLength)
			}
			for i, value := range data {
				if value <= 0 {
					t.Errorf("data[%d] = %d, want a positive number", i, value)
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"nil", nil, 0},
		{"empty", []int{}, 0},
		{"one element", []int{42}, 42},
		{"maximum first", []int{9, 3, 2}, 9},
		{"maximum middle", []int{2, 9, 3}, 9},
		{"maximum last", []int{2, 3, 9}, 9},
		{"equal elements", []int{7, 7, 7}, 7},
		{"zeros", []int{0, 0, 0}, 0},
		{"negative elements", []int{-8, -2, -5}, -2},
		{"mixed signs", []int{-5, 0, 12, -3}, 12},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.data); got != tt.want {
				t.Errorf("maximum(%v) = %d, want %d", tt.data, got, tt.want)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"nil", nil, 0},
		{"empty", []int{}, 0},
		{"one element", []int{42}, 42},
		{"short slice", []int{2, 9, 3}, 9},
		{"short negative slice", []int{-8, -2, -5}, -2},
		{"eight elements", []int{1, 8, 3, 4, 5, 6, 7, 2}, 8},
		{"maximum in remainder", []int{1, 2, 3, 4, 5, 6, 7, 8, 99}, 99},
		{"equal elements", []int{7, 7, 7, 7, 7, 7, 7, 7}, 7},
		{"zeros", make([]int, 16), 0},
		{"negative elements", []int{-8, -2, -5, -9, -4, -6, -7, -3, -1}, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxChunks(tt.data); got != tt.want {
				t.Errorf("maxChunks(%v) = %d, want %d", tt.data, got, tt.want)
			}
		})
	}
}

func TestMaxChunksAllPositions(t *testing.T) {
	// Проверяем границы частей и остаток, перемещая максимум по слайсу.
	for _, size := range []int{8, 9, 15, 16, 17, 1003} {
		data := make([]int, size)
		for i := range data {
			data[i] = 100
			if got := maxChunks(data); got != 100 {
				t.Fatalf("size %d, maximum at %d: got %d, want 100", size, i, got)
			}
			data[i] = 0
		}
	}
}

func TestMaximumDoesNotChangeData(t *testing.T) {
	for name, findMaximum := range map[string]func([]int) int{
		"maximum":   maximum,
		"maxChunks": maxChunks,
	} {
		t.Run(name, func(t *testing.T) {
			data := []int{4, 1, 8, 2, 9, 3, 7, 5, 6}
			original := append([]int(nil), data...)
			findMaximum(data)
			if !reflect.DeepEqual(data, original) {
				t.Errorf("input changed: got %v, want %v", data, original)
			}
		})
	}
}
