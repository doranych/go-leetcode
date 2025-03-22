package find_all_possible_recipes_from_given_supplies

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findAllRecipes(t *testing.T) {
	tests := []struct {
		name        string
		recipes     []string
		ingredients [][]string
		supplies    []string
		want        []string
	}{
		{"",
			[]string{"bread"},
			[][]string{{"yeast", "flour"}},
			[]string{"yeast", "flour", "corn"},
			[]string{"bread"},
		},
		{"",
			[]string{"bread", "sandwich"},
			[][]string{{"yeast", "flour"}, {"bread", "meat"}},
			[]string{"yeast", "flour", "meat"},
			[]string{"bread", "sandwich"},
		},
		{"",
			[]string{"bread", "sandwich", "burger"},
			[][]string{{"yeast", "flour"}, {"bread", "meat"}, {"sandwich", "meat", "bread"}},
			[]string{"yeast", "flour", "meat"},
			[]string{"bread", "sandwich", "burger"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, findAllRecipes(tt.recipes, tt.ingredients, tt.supplies))
		})
	}
}
