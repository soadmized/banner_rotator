package demogroup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_docToGroup(t *testing.T) {
	tests := []struct {
		name string
		d    groupDoc
		want *Group
	}{
		{
			name: "positive",
			d: groupDoc{
				ID:          "group1",
				Description: "group 1",
			},
			want: &Group{
				ID:          "group1",
				Description: "group 1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := docToGroup(tt.d)
			assert.Equal(t, tt.want, got)
		})
	}
}
