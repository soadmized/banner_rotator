package banner

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_docToBanner(t *testing.T) {
	tests := []struct {
		name string
		d    bannerDoc
		want *Banner
	}{
		{
			name: "positive",
			d: bannerDoc{
				ID:          "banner1",
				Description: "banner 1",
			},
			want: &Banner{
				ID:          "banner1",
				Description: "banner 1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := docToBanner(tt.d)
			assert.Equal(t, tt.want, got)
		})
	}
}
