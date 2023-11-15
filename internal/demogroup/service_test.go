package demogroup_test

import (
	"context"
	"testing"

	"github.com/pkg/errors"
	"github.com/soadmized/banners_rotator/internal/demogroup"
	"github.com/soadmized/banners_rotator/internal/demogroup/mocks"
	"github.com/stretchr/testify/require"
)

func TestService_Get(t *testing.T) {
	ctx := context.Background()
	id := demogroup.ID("group1")
	want := demogroup.Group{
		ID:          id,
		Description: "desc group 1",
	}

	repo := mocks.NewRepository(t)
	repo.On("Get", ctx, id).Return(&want, nil)

	srv := demogroup.Service{Repo: repo}

	got, err := srv.Get(ctx, id)
	require.NoError(t, err)
	require.Equal(t, &want, got)
}

func TestService_GetError(t *testing.T) {
	ctx := context.Background()
	id := demogroup.ID("group1")

	repo := mocks.NewRepository(t)
	repo.On("Get", ctx, id).Return(nil, errors.New("get demogroup"))

	srv := demogroup.Service{Repo: repo}

	_, err := srv.Get(ctx, id)
	require.Error(t, err)
}

func TestService_Create(t *testing.T) {
	ctx := context.Background()
	id := demogroup.ID("group1")
	desc := "group desc"

	repo := mocks.NewRepository(t)
	repo.On("Create", ctx, id, desc).Return(nil)

	srv := demogroup.Service{Repo: repo}

	err := srv.Create(ctx, id, desc)
	require.NoError(t, err)
}

func TestService_CreateError(t *testing.T) {
	ctx := context.Background()
	id := demogroup.ID("group1")
	desc := "group desc"

	repo := mocks.NewRepository(t)
	repo.On("Create", ctx, id, desc).Return(errors.New("create group"))

	srv := demogroup.Service{Repo: repo}

	err := srv.Create(ctx, id, desc)
	require.Error(t, err)
}
