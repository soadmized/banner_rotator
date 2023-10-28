package stat

import "context"

type Repository interface {
	GetStat(ctx context.Context, slotID, bannerID, groupID string) (*Stat, error)
	AddShow(ctx context.Context, slotID, bannerID, groupID string) error
	AddClick(ctx context.Context, slotID, bannerID, groupID string) error
}

type Service struct {
	Repo Repository
}

func (s *Service) GetStat(ctx context.Context, slotID, bannerID, groupID string) (*Stat, error) {
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
