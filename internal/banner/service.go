package banner

import "context"

//go:generate  mockery --tags="mock" --filename repo.go --name Repository
type Repository interface {
	Get(ctx context.Context, id ID) (*Banner, error)
	Create(ctx context.Context, id ID, desc string) error
}

type Service struct {
	Repo Repository
}

func (s *Service) Get(ctx context.Context, id ID) (*Banner, error) {
	banner, err := s.Repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return banner, nil
}

func (s *Service) Create(ctx context.Context, id ID, desc string) error {
	err := s.Repo.Create(ctx, id, desc)
	if err != nil {
		return err
	}

	return nil
}
