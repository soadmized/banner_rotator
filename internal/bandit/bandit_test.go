package bandit

import (
	"banners_rotator/internal/banner"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBandit_Pick(t *testing.T) {
	rewards := map[banner.ID]int{
		"banner1": 2,
		"banner2": 2,
		"banner3": 3, // banner3 should be picked mostly
	}

	arms := []banner.ID{
		"banner1",
		"banner2",
		"banner3",
	}

	result := map[banner.ID]int{ // how many each banner was picked
		"banner1": 0,
		"banner2": 0,
		"banner3": 0,
	}

	bandit := New(rewards, arms)

	for i := 0; i < 100; i++ {
		bannerID := bandit.Pick()
		result[bannerID]++
	}

	require.Greater(t, result["banner3"], result["banner1"]+result["banner2"])

	fmt.Println("\n======\nRESULT = ", result, "\n====== ")
}
