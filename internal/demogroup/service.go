package demogroup

import "context"

// TODO: mockery
type Repository interface {
	Get(ctx context.Context, id ID) (*Group, error)
	Create(ctx context.Context, id ID, desc string) error
}

type Service struct {
	Repo Repository
}

func (s *Service) Get(ctx context.Context, id ID) (*Group, error) {
	group, err := s.Repo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return group, nil
}

func (s *Service) Create(ctx context.Context, id ID, desc string) error {
	err := s.Repo.Create(ctx, id, desc)
	if err != nil {
		return err
	}

	return nil
}
