package banner_test

import (
	"context"
	"testing"

	"github.com/pkg/errors"
	"github.com/soadmized/banners_rotator/internal/banner"
	"github.com/soadmized/banners_rotator/internal/banner/mocks"
	"github.com/stretchr/testify/require"
)

func TestService_Get(t *testing.T) {
	ctx := context.Background()
	bannerID := banner.ID("banner1")
	want := banner.Banner{
		ID:          bannerID,
		Description: "desc banner 1",
	}

	repo := mocks.NewRepository(t)
	repo.On("Get", ctx, bannerID).Return(&want, nil)

	srv := banner.Service{Repo: repo}

	got, err := srv.Get(ctx, bannerID)
	require.NoError(t, err)
	require.Equal(t, &want, got)

}

func TestService_GetError(t *testing.T) {
	ctx := context.Background()
	bannerID := banner.ID("banner1")

	repo := mocks.NewRepository(t)
	repo.On("Get", ctx, bannerID).Return(nil, errors.New("get banner"))

	srv := banner.Service{Repo: repo}

	_, err := srv.Get(ctx, bannerID)
	require.Error(t, err)
}

func TestService_Create(t *testing.T) {
	ctx := context.Background()
	bannerID := banner.ID("banner1")
	desc := "banner desc"

	repo := mocks.NewRepository(t)
	repo.On("Create", ctx, bannerID, desc).Return(nil)

	srv := banner.Service{Repo: repo}

	err := srv.Create(ctx, bannerID, desc)
	require.NoError(t, err)
}

func TestService_CreateError(t *testing.T) {
	ctx := context.Background()
	bannerID := banner.ID("banner1")
	desc := "banner desc"

	repo := mocks.NewRepository(t)
	repo.On("Create", ctx, bannerID, desc).Return(errors.New("create banner"))

	srv := banner.Service{Repo: repo}

	err := srv.Create(ctx, bannerID, desc)
	require.Error(t, err)
}
