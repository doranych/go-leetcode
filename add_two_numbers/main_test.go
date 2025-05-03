package add_two_numbers

import (
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	tests := []struct {
		name string
		l1   *ListNode
		l2   *ListNode
		want *ListNode
	}{
		{"", &ListNode{0, nil}, &ListNode{0, nil}, &ListNode{0, nil}},
		{"", &ListNode{9, &ListNode{9, &ListNode{9, nil}}}, &ListNode{1, nil}, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{1, nil}}}}},
		{"", &ListNode{1, &ListNode{8, nil}}, &ListNode{0, nil}, &ListNode{1, &ListNode{8, nil}}},
		{"", &ListNode{1, &ListNode{8, nil}}, &ListNode{0, &ListNode{1, nil}}, &ListNode{1, &ListNode{9, nil}}},
		{"", &ListNode{1, &ListNode{8, nil}}, &ListNode{0, &ListNode{9, nil}}, &ListNode{1, &ListNode{7, &ListNode{1, nil}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.l1, tt.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
