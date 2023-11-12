package slot_test

import (
	"context"
	"testing"

	"github.com/pkg/errors"
	"github.com/soadmized/banners_rotator/internal/slot"
	"github.com/soadmized/banners_rotator/internal/slot/mocks"
	"github.com/stretchr/testify/require"
)

func TestService_Get(t *testing.T) {
	ctx := context.Background()
	id := slot.ID("slot1")
	want := slot.Slot{
		ID:          id,
		Description: "desc slot 1",
	}

	repo := mocks.NewRepository(t)
	repo.On("Get", ctx, id).Return(&want, nil)

	srv := slot.Service{Repo: repo}

	got, err := srv.Get(ctx, id)
	require.NoError(t, err)
	require.Equal(t, &want, got)

}

func TestService_GetError(t *testing.T) {
	ctx := context.Background()
	id := slot.ID("slot1")

	repo := mocks.NewRepository(t)
	repo.On("Get", ctx, id).Return(nil, errors.New("get slot"))

	srv := slot.Service{Repo: repo}

	_, err := srv.Get(ctx, id)
	require.Error(t, err)
}

func TestService_Create(t *testing.T) {
	ctx := context.Background()
	id := slot.ID("slot1")
	desc := "slot 1 desc"

	repo := mocks.NewRepository(t)
	repo.On("Create", ctx, id, desc).Return(nil)

	srv := slot.Service{Repo: repo}

	err := srv.Create(ctx, id, desc)
	require.NoError(t, err)
}

func TestService_CreateError(t *testing.T) {
	ctx := context.Background()
	id := slot.ID("slot1")
	desc := "slot 1 desc"

	repo := mocks.NewRepository(t)
	repo.On("Create", ctx, id, desc).Return(errors.New("create slot"))

	srv := slot.Service{Repo: repo}

	err := srv.Create(ctx, id, desc)
	require.Error(t, err)
}
