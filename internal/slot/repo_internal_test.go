package slot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_docToGroup(t *testing.T) {
	tests := []struct {
		name string
		d    slotDoc
		want *Slot
	}{
		{
			name: "positive",
			d: slotDoc{
				ID:          "slot1",
				Description: "slot 1",
			},
			want: &Slot{
				ID:          "slot1",
				Description: "slot 1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := docToSlot(tt.d)
			assert.Equal(t, tt.want, got)
		})
	}
}
