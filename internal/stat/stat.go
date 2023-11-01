package stat

import (
	"banners_rotator/internal/banner"
	"banners_rotator/internal/demogroup"
	"banners_rotator/internal/slot"
)

// SlotID
//   |
//   BannerID1
//   |  |
//   |  GroupID1
//   |  | |
//   |  | Stat
//	 |  |   |
//   |  |   Clicks
//	 |  |   |
//   |  |   Shows
//   |  |
//   |  GroupID2
//   |    |
//   |    ...
//   |
//   BannerID2
//     |
//     ... and so on

type SlotStat struct {
	ID         slot.ID
	BannerStat BannerStat
}

type BannerStat map[banner.ID]GroupStat

type GroupStat map[demogroup.ID]Stat

type Stat struct {
	Clicks int
	Shows  int
}
