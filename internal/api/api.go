package api

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/soadmized/banners_rotator/internal/banner"
	"github.com/soadmized/banners_rotator/internal/demogroup"
	"github.com/soadmized/banners_rotator/internal/slot"
)

type BannerService interface {
	Get(ctx context.Context, id banner.ID) (*banner.Banner, error)
	Create(ctx context.Context, id banner.ID, desc string) error
}

type DemoGroupService interface {
	Get(ctx context.Context, id demogroup.ID) (*demogroup.Group, error)
	Create(ctx context.Context, id demogroup.ID, desc string) error
}

type SlotService interface {
	Get(ctx context.Context, id slot.ID) (*slot.Slot, error)
	Create(ctx context.Context, id slot.ID, desc string) error
}

type StatService interface {
	AddClick(ctx context.Context, slotID, bannerID, groupID string) error
	AddBanner(ctx context.Context, slotID, bannerID string) error
	RemoveBanner(ctx context.Context, slotID, bannerID string) error
	PickBanner(ctx context.Context, slotID, groupID string) (banner.ID, error)
}

const (
	removeBannerPath = "remove_banner"
	addBannerPath    = "add_banner"
	addClickPath     = "add_click"
	pickBannerPath   = "pick_banner"
)

type API struct {
	Router *gin.Engine

	BannerSrv BannerService
	SlotSrv   SlotService
	GroupSrv  DemoGroupService
	StatSrv   StatService
}

type reqBody struct {
	BannerID string `json:"bannerId,omitempty"`
	SlotID   string `json:"slotId,omitempty"`
	GroupID  string `json:"groupId,omitempty"`
}

func (a *API) RemoveBanner(ctx *gin.Context) {
	body, err := decodeBody(ctx.Request.Body)
	if err != nil {
		ctx.Writer.WriteHeader(http.StatusInternalServerError)
		log.Print(err)

		return
	}

	err = a.StatSrv.RemoveBanner(ctx, body.SlotID, body.BannerID)
	if err != nil {
		ctx.Writer.WriteHeader(http.StatusInternalServerError)
		log.Print(err)

		return
	}

	ctx.Writer.WriteHeader(http.StatusOK)
}

func (a *API) AddBanner(ctx *gin.Context) {
	body, err := decodeBody(ctx.Request.Body)
	if err != nil {
		ctx.Writer.WriteHeader(http.StatusInternalServerError)
		log.Print(err)

		return
	}

	err = a.StatSrv.AddBanner(ctx, body.SlotID, body.BannerID)
	if err != nil {
		ctx.Writer.WriteHeader(http.StatusInternalServerError)
		log.Print(err)

		return
	}

	ctx.Writer.WriteHeader(http.StatusOK)
}

func (a *API) AddClick(ctx *gin.Context) {
	body, err := decodeBody(ctx.Request.Body)
	if err != nil {
		ctx.Writer.WriteHeader(http.StatusInternalServerError)
		log.Print(err)

		return
	}

	err = a.StatSrv.AddClick(ctx, body.SlotID, body.BannerID, body.GroupID)
	if err != nil {
		ctx.Writer.WriteHeader(http.StatusInternalServerError)
		log.Print(err)

		return
	}

	ctx.Writer.WriteHeader(http.StatusOK)
}

func (a *API) PickBanner(ctx *gin.Context) {
	body, err := decodeBody(ctx.Request.Body)
	if err != nil {
		ctx.Writer.WriteHeader(http.StatusInternalServerError)
		log.Print(err)

		return
	}

	bannerID, err := a.StatSrv.PickBanner(ctx, body.SlotID, body.GroupID)
	if err != nil {
		ctx.Writer.WriteHeader(http.StatusInternalServerError)
		log.Print(err)

		return
	}

	resp := response{ID: bannerID}

	ctx.JSON(http.StatusOK, resp)
}

func decodeBody(body io.ReadCloser) (reqBody, error) {
	decoder := json.NewDecoder(body)
	res := reqBody{}

	err := decoder.Decode(&res)
	if err != nil {
		return reqBody{}, errors.Wrap(err, "decode req body to model")
	}

	return res, nil
}

func (a *API) RegisterHandlers() {
	a.Router.POST(addBannerPath, a.AddBanner)
	a.Router.POST(addClickPath, a.AddClick)
	a.Router.POST(removeBannerPath, a.RemoveBanner)
	a.Router.POST(pickBannerPath, a.PickBanner)
}

type response struct {
	ID banner.ID `json:"id"`
}
