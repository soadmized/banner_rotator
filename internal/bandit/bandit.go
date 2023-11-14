package bandit

import (
	"math/rand"

	"github.com/soadmized/banners_rotator/internal/banner"
)

const epsilon = 0.3

type Bandit struct {
	rewards map[banner.ID]int // armID: rewards
	arms    []banner.ID
}

func (b *Bandit) Pick() banner.ID {
	probability := rand.Float64() //nolint:gosec

	if probability <= (1 - epsilon) {
		// exploit
		return b.armWithMaxReward()
	}

	// explore
	return b.randomArm()
}

func (b *Bandit) armWithMaxReward() banner.ID {
	var maxK banner.ID

	maxV := 0

	for k, v := range b.rewards {
		if v > maxV {
			maxV = v
			maxK = k
		}
	}

	return maxK
}

func (b *Bandit) randomArm() banner.ID {
	armID := rand.Intn(len(b.arms)) //nolint:gosec

	return b.arms[armID]
}

func New(rewards map[banner.ID]int, arms []banner.ID) *Bandit {
	bandit := Bandit{
		rewards: rewards,
		arms:    arms,
	}

	return &bandit
}
