package compare_version_numbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_compareVersion(t *testing.T) {
	tests := []struct {
		name     string
		version1 string
		version2 string
		want     int
	}{

		{name: "", version1: "1.2", version2: "1.10", want: -1},
		{name: "", version1: "1.01", version2: "1.001", want: 0},
		{name: "", version1: "1.0", version2: "1.0.0.0", want: 0},
		{name: "", version1: "2.0", version2: "2.0.1", want: -1},
		{name: "", version1: "3.4.5", version2: "3.4.5", want: 0},
		{name: "", version1: "0.1", version2: "0.0.1", want: 1},
		{name: "", version1: "1.0.0", version2: "1", want: 0},
		{name: ".", version1: "1.0.1", version2: "1.0.0", want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := compareVersion(tt.version1, tt.version2)
			assert.Equal(t, tt.want, got)
		})
	}
}
