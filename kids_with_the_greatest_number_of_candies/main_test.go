package kids_with_the_greatest_number_of_candies

import (
	"reflect"
	"testing"
)

func Test_kidsWithCandies(t *testing.T) {
	tests := []struct {
		name         string
		candies      []int
		extraCandies int
		want         []bool
	}{
		{"", []int{2, 3, 5, 1, 3}, 3, []bool{true, true, true, false, true}},
		{"", []int{4, 2, 1, 1, 2}, 1, []bool{true, false, false, false, false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kidsWithCandies(tt.candies, tt.extraCandies); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kidsWithCandies() = %v, want %v", got, tt.want)
			}
		})
	}
}
