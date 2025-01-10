package string_compression

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_compress(t *testing.T) {
	type args struct {
		chars []byte
	}
	tests := []struct {
		name      string
		args      args
		want      int
		wantSlice []byte
	}{
		{"", args{[]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}}, 6, []byte{'a', '2', 'b', '2', 'c', '3'}},
		{"", args{[]byte{'a'}}, 1, []byte{'a'}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compress(tt.args.chars); got != tt.want {
				t.Errorf("compress() = %v, want %v", got, tt.want)
				require.Equal(t, tt.wantSlice, tt.args.chars[:got])
			}
		})
	}
}
