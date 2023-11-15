package bandit

import (
	"fmt"
	"testing"

	"github.com/soadmized/banners_rotator/internal/banner"
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

	fmt.Println("\n======\nRESULT 1 = ", result, "\n====== ")

	// now banner2 get more clicks then banner3
	rewards["banner2"] = 4

	result["banner1"] = 0
	result["banner2"] = 0
	result["banner3"] = 0

	for i := 0; i < 100; i++ {
		bannerID := bandit.Pick()
		result[bannerID]++
	}

	require.Greater(t, result["banner2"], result["banner1"]+result["banner3"])

	fmt.Println("\n======\nRESULT 2 = ", result, "\n====== ")
}
