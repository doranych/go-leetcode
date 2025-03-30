package partition_labels

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_partitionLabels(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want []int
	}{
		{"", "ababcbacadefegdehijhklij", []int{9, 7, 8}},
		{"", "eccbbbbdec", []int{10}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, partitionLabels(tt.s))
		})
	}
}
