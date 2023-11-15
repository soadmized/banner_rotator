package stat_test

import (
	"context"
	"testing"

	"github.com/soadmized/banners_rotator/internal/banner"
	"github.com/soadmized/banners_rotator/internal/stat"
	"github.com/soadmized/banners_rotator/internal/stat/mocks"
	"github.com/stretchr/testify/require"
)

const (
	slotID   = "slot1"
	bannerID = "banner1"
	groupID  = "group1"
)

func TestService_AddShow(t *testing.T) {
	ctx := context.Background()

	repo := mocks.NewRepository(t)
	repo.On("AddShow", ctx, slotID, bannerID, groupID).Return(nil)

	srv := stat.Service{Repo: repo}

	err := srv.AddShow(ctx, slotID, bannerID, groupID)
	require.NoError(t, err)
}

func TestService_AddClick(t *testing.T) {
	ctx := context.Background()

	repo := mocks.NewRepository(t)
	repo.On("AddClick", ctx, slotID, bannerID, groupID).Return(nil)

	srv := stat.Service{Repo: repo}

	err := srv.AddClick(ctx, slotID, bannerID, groupID)
	require.NoError(t, err)
}

func TestService_AddBanner(t *testing.T) {
	ctx := context.Background()

	repo := mocks.NewRepository(t)
	repo.On("AddBanner", ctx, slotID, bannerID).Return(nil)

	srv := stat.Service{Repo: repo}

	err := srv.AddBanner(ctx, slotID, bannerID)
	require.NoError(t, err)
}

func TestService_RemoveBanner(t *testing.T) {
	ctx := context.Background()

	repo := mocks.NewRepository(t)
	repo.On("RemoveBanner", ctx, slotID, bannerID).Return(nil)

	srv := stat.Service{Repo: repo}

	err := srv.RemoveBanner(ctx, slotID, bannerID)
	require.NoError(t, err)
}

func TestService_PickBanner(t *testing.T) {
	ctx := context.Background()

	stats := stat.Stat{
		Clicks: 2,
		Shows:  2,
	}

	repo := mocks.NewRepository(t)
	repo.On("GetBannerIDs", ctx, slotID).Return([]banner.ID{banner.ID(bannerID)}, nil)
	repo.On("GetStat", ctx, slotID, bannerID, groupID).Return(&stats, nil)
	repo.On("AddShow", ctx, slotID, bannerID, groupID).Return(nil)

	srv := stat.Service{Repo: repo}

	id, err := srv.PickBanner(ctx, slotID, groupID)
	require.NoError(t, err)
	require.Equal(t, banner.ID(bannerID), id)
}
