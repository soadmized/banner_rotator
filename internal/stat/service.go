package stat

import (
	"context"

	"github.com/pkg/errors"
	"github.com/soadmized/banners_rotator/internal/bandit"
	"github.com/soadmized/banners_rotator/internal/banner"
)

//go:generate  mockery --tags="mock" --filename repo.go --name Repository
type Repository interface {
	GetStat(ctx context.Context, slotID, bannerID, groupID string) (*Stat, error)
	AddShow(ctx context.Context, slotID, bannerID, groupID string) error
	AddClick(ctx context.Context, slotID, bannerID, groupID string) error
	AddBanner(ctx context.Context, slotID, bannerID string) error
	RemoveBanner(ctx context.Context, slotID, bannerID string) error
	GetBannerIDs(ctx context.Context, slotID string) ([]banner.ID, error)
}

type Service struct {
	Repo Repository
}

func (s *Service) getStat(ctx context.Context, slotID, bannerID, groupID string) (*Stat, error) {
	stat, err := s.Repo.GetStat(ctx, slotID, bannerID, groupID)
	if err != nil {
		return nil, err
	}

	return stat, nil
}

func (s *Service) AddShow(ctx context.Context, slotID, bannerID, groupID string) error {
	err := s.Repo.AddShow(ctx, slotID, bannerID, groupID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) AddClick(ctx context.Context, slotID, bannerID, groupID string) error {
	err := s.Repo.AddClick(ctx, slotID, bannerID, groupID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) AddBanner(ctx context.Context, slotID, bannerID string) error {
	err := s.Repo.AddBanner(ctx, slotID, bannerID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) RemoveBanner(ctx context.Context, slotID, bannerID string) error {
	err := s.Repo.RemoveBanner(ctx, slotID, bannerID)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) bannerClicksByGroupID(ctx context.Context, slotID, groupID string) (map[banner.ID]int, error) {
	banners, err := s.Repo.GetBannerIDs(ctx, slotID)
	if err != nil {
		return nil, err
	}

	bannerStat := make(map[banner.ID]int, len(banners))

	for _, v := range banners {
		stat, err := s.getStat(ctx, slotID, string(v), groupID)
		if err != nil {
			return nil, errors.Wrap(err, "get banner stat by group id")
		}

		bannerStat[v] = stat.Clicks
	}

	return bannerStat, nil
}

// PickBanner implements multi armed bandit algorithm.
func (s *Service) PickBanner(ctx context.Context, slotID, groupID string) (banner.ID, error) {
	bannerStat, err := s.bannerClicksByGroupID(ctx, slotID, groupID)
	if err != nil {
		return "", err
	}

	banners, err := s.Repo.GetBannerIDs(ctx, slotID)
	if err != nil {
		return "", err
	}

	b := bandit.New(bannerStat, banners)

	bannerID := b.Pick()

	err = s.AddShow(ctx, slotID, string(bannerID), groupID)
	if err != nil {
		return "", errors.Wrap(err, "add click after picking banner")
	}

	return bannerID, nil
}
