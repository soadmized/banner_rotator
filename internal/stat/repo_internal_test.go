package stat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_docToStat(t *testing.T) {
	tests := []struct {
		name string
		d    statDoc
		want Stat
	}{
		{
			name: "positive",
			d: statDoc{
				Clicks: 2,
				Shows:  2,
			},
			want: Stat{
				Clicks: 2,
				Shows:  2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := docToStat(tt.d)
			assert.Equal(t, tt.want, got)
		})
	}
}
