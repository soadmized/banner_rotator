package bandit

import (
	"banners_rotator/internal/banner"
	"math/rand"
)

const epsilon = 0.3

type Bandit struct {
	rewards map[banner.ID]int //armID: rewards
	arms    []banner.ID
}

func (b *Bandit) Pick() banner.ID {
	probability := rand.Float64()

	if probability <= (1 - epsilon) {
		// exploit
		return b.armWithMaxReward()
	} else {
		// explore
		return b.randomArm()
	}
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
	armID := rand.Intn(len(b.arms))

	return b.arms[armID]
}

func New(rewards map[banner.ID]int, arms []banner.ID) *Bandit {
	bandit := Bandit{
		rewards: rewards,
		arms:    arms,
	}

	return &bandit
}
