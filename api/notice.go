package api

import (
	"net/http"

	"github.com/labstack/echo"
)

type Notices struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Avatar      string `json:"avatar"`
	Description string `json:"description"`
	Datetime    string `json:"datetime"`
	Read        bool   `json:"read"`
	ClickClose  bool   `json:"clickClose"`
	Type        string `json:"type"`
	Extra       string `json:"extra"`
	Status      string `json:"status"`
}

func GetNotices() echo.HandlerFunc {
	notices := []Notices{
		{
			ID:       "000000001",
			Avatar:   "https://gw.alipayobjects.com/zos/rmsportal/ThXAXghbEsBCCSDihZxY.png",
			Title:    "You received 14 new weekly newspapers",
			Datetime: "2017-08-09",
			Type:     "notification",
		},
		{
			ID:       "000000002",
			Avatar:   "https://gw.alipayobjects.com/zos/rmsportal/OKJXDXrmkNshAMvwtvhu.png",
			Title:    "Your recommended Quinnie has passed the third round of interviews.",
			Datetime: "2017-08-08",
			Type:     "notification",
		},
		{
			ID:       "000000003",
			Avatar:   "https://gw.alipayobjects.com/zos/rmsportal/kISTdvpyTAhtGxpovNWd.png",
			Title:    "This template can distinguish between multiple notification types",
			Datetime: "2017-08-07",
			Read:     true,
			Type:     "notification",
		},
		{
			ID:       "000000004",
			Avatar:   "https://gw.alipayobjects.com/zos/rmsportal/GvqBnKhFgObvnSGkDsje.png",
			Title:    "The left icon is used to distinguish between different types",
			Datetime: "2017-08-07",
			Type:     "notification",
		},
		{
			ID:       "000000005",
			Avatar:   "https://gw.alipayobjects.com/zos/rmsportal/ThXAXghbEsBCCSDihZxY.png",
			Title:    "Do not exceed two lines of content, automatically cut off when exceeded",
			Datetime: "2017-08-07",
			Type:     "notification",
		},
		{
			ID:          "000000006",
			Avatar:      "https://gw.alipayobjects.com/zos/rmsportal/fcHMVNCjPOsbUGdEduuv.jpeg",
			Title:       "Qu Lili commented on you",
			Description: "Description information description information description information",
			Datetime:    "2017-08-07",
			Type:        "message",
			ClickClose:  true,
		},
		{
			ID:          "000000007",
			Avatar:      "https://gw.alipayobjects.com/zos/rmsportal/fcHMVNCjPOsbUGdEduuv.jpeg",
			Title:       "Zhu is right, I replied to you.",
			Description: "This template is used to remind people who have interacted with you.",
			Datetime:    "2017-08-07",
			Type:        "message",
			ClickClose:  true,
		},
		{
			ID:          "000000008",
			Avatar:      "https://gw.alipayobjects.com/zos/rmsportal/fcHMVNCjPOsbUGdEduuv.jpeg",
			Title:       "title",
			Description: "This template is used to remind people who have interacted with you.",
			Datetime:    "2017-08-07",
			Type:        "message",
			ClickClose:  true,
		},
		{
			ID:          "000000009",
			Title:       "mission name",
			Description: "The mission needs to be started before 2017-01-12 20:00",
			Extra:       "has not started",
			Status:      "todo",
			Type:        "event",
		},
		{
			ID:          "000000010",
			Title:       "Third-party emergency code change",
			Description: "Guan Lin submitted on 2017-01-06, need to complete the code change task before 2017-01-07",
			Extra:       "Expire immediately",
			Status:      "urgent",
			Type:        "event",
		},
		{
			ID:          "000000011",
			Title:       "Information security exam",
			Description: "Appointed Zhuer to complete the update and release before 2017-01-09",
			Extra:       "It took 8 days",
			Status:      "doing",
			Type:        "event",
		},
		{
			ID:          "000000012",
			Title:       "ABCD version released",
			Description: "Guan Lin submitted on 2017-01-06, need to complete the code change task before 2017-01-07",
			Extra:       "processing",
			Status:      "processing",
			Type:        "event",
		},
	}
	return func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, notices)
	}
}
