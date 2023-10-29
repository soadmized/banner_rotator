package api

import (
	"banners_rotator/internal/banner"
	"banners_rotator/internal/demogroup"
	"banners_rotator/internal/slot"
	"banners_rotator/internal/stat"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"io"
	"net/http"
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
	GetStat(ctx context.Context, slotID, bannerID, groupID string) (*stat.Stat, error)
	AddShow(ctx context.Context, slotID, bannerID, groupID string) error
	AddClick(ctx context.Context, slotID, bannerID, groupID string) error
	AddBanner(ctx context.Context, slotID, bannerID string) error
	RemoveBanner(ctx context.Context, slotID, bannerID string) error
}

const (
	removeBannerPath = "remove_banner"
	addBannerPath    = "add_banner"
	addClickPath     = "add_click"
	pickBannerPath   = "pick_banner"
)

type Api struct {
	Router *gin.Engine

	BannerSrv BannerService
	SlotSrv   SlotService
	GroupSrv  DemoGroupService
	StatSrv   StatService
}

type reqBody struct {
	BannerID string `json:"banner_id,omitempty"`
	SlotID   string `json:"slot_id,omitempty"`
	GroupID  string `json:"group_id,omitempty"`
}

func (a *Api) RemoveBanner(ctx *gin.Context) {
	body, err := decodeBody(ctx.Request.Body)
	if err != nil {
		ctx.Writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	err = a.StatSrv.RemoveBanner(ctx, body.SlotID, body.BannerID)
	if err != nil {
		ctx.Writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	ctx.Writer.WriteHeader(http.StatusOK)
}

func (a *Api) AddBanner(ctx *gin.Context) {
	body, err := decodeBody(ctx.Request.Body)
	if err != nil {
		ctx.Writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	err = a.StatSrv.AddBanner(ctx, body.SlotID, body.BannerID)
	if err != nil {
		ctx.Writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	ctx.Writer.WriteHeader(http.StatusOK)
}

func (a *Api) AddClick(ctx *gin.Context) {
	body, err := decodeBody(ctx.Request.Body)
	if err != nil {
		ctx.Writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	err = a.StatSrv.AddClick(ctx, body.SlotID, body.BannerID, body.GroupID)
	if err != nil {
		ctx.Writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	ctx.Writer.WriteHeader(http.StatusOK)
}

// TODO
func (a *Api) PickBanner(ctx *gin.Context) {
	_, err := decodeBody(ctx.Request.Body)
	if err != nil {
		ctx.Writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	// multi armed bandit method here

	ctx.JSON(http.StatusOK, struct {
		id string
	}{id: "banner42"})
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

func (a *Api) RegisterHandlers() {
	a.Router.POST(addBannerPath, a.AddBanner)
	a.Router.POST(addClickPath, a.AddClick)
	a.Router.POST(removeBannerPath, a.RemoveBanner)
	a.Router.POST(pickBannerPath, a.PickBanner)
}
