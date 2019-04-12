package api

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
)

type Tag struct {
	Key   string `json:"ley"`
	Label string `json:"label"`
}

type User struct {
	ID          string     `json:"userid"`
	Name        string     `json:"name"`
	Avatar      string     `json:"avatar"`
	Email       string     `json:"email"`
	Signature   string     `json:"signature"`
	Title       string     `json:"title"`
	Group       string     `json:"Group"`
	Tags        []Tag      `json:"tags"`
	NotifyCount int64      `json:"notifyCount"`
	UnreadCount int64      `json:"unreadCount"`
	Country     string     `json:"country"`
	Geographic  Geographic `json:"geographic"`
	Address     string     `json:"address"`
	Phone       string     `json:"phone"`
}

type Geographic struct {
	Province Tag `json:"province"`
	City     Tag `json:"city"`
}

type Profile struct {
	Key     string `json:"key"`
	Name    string `json:"name"`
	Age     uint   `json:"age"`
	Address string `json:"address"`
}

type LoginForm struct {
	Username string
	Password string
	Type     interface{}
}

type LoginResponse struct {
	Status           string      `json:"status"`
	CurrentAuthority string      `json:"currentAuthority"`
	Type             interface{} `json:"type"`
}

func CurrentUser() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, User{
			Name:      "Serati Ma",
			Avatar:    "https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png",
			ID:        "00000001",
			Email:     "antdesign@alipay.com",
			Signature: "Be tolerant to diversity, tolerance is a virtue",
			Title:     "Interaction expert",
			Group:     "Ant gold clothing - a certain business group - a certain platform department - a certain technical department - UED",
			Tags: []Tag{
				{
					Key:   "0",
					Label: "Very thoughtful",
				},
				{
					Key:   "1",
					Label: "Focus on design",
				},
				{
					Key:   "2",
					Label: "Spicy",
				},
				{
					Key:   "3",
					Label: "Big long legs",
				},
				{
					Key:   "4",
					Label: "Chuanmeizi",
				},
				{
					Key:   "5",
					Label: "Haina Baichuan",
				},
			},
			NotifyCount: 12,
			UnreadCount: 11,
			Country:     "China",
			Geographic: Geographic{
				Province: Tag{
					Label: "Zhejiang Province",
					Key:   "330000",
				},
				City: Tag{
					Label: "Hangzhou",
					Key:   "330100",
				},
			},
			Phone: "0752-268888888",
		})
	}
}

func Users() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, []Profile{
			{
				Key:     "1",
				Name:    "John Brown",
				Age:     32,
				Address: "New York No. 1 Lake Park",
			},
			{
				Key:     "2",
				Name:    "Jim Green",
				Age:     42,
				Address: "London No. 1 Lake Park",
			},
			{
				Key:     "3",
				Name:    "Joe Black",
				Age:     32,
				Address: "Sidney No. 1 Lake Park",
			},
		})
	}
}

func Login() echo.HandlerFunc {
	pass := "ant.design"
	return func(ctx echo.Context) error {
		var r LoginForm
		err := json.NewDecoder(ctx.Request().Body).Decode(&r)
		if err != nil {
			//TODO handle error
		}
		if r.Password == pass && r.Username == "admin" {
			return ctx.JSON(http.StatusOK, LoginResponse{
				Status:           "ok",
				Type:             r.Type,
				CurrentAuthority: "admin",
			})
		}
		if r.Password == pass && r.Username == "user" {
			return ctx.JSON(http.StatusOK, LoginResponse{
				Status:           "ok",
				Type:             r.Type,
				CurrentAuthority: "user",
			})
		}

		return ctx.JSON(http.StatusBadRequest, LoginResponse{
			Status:           "error",
			Type:             r.Type,
			CurrentAuthority: "guest",
		})
	}
}

func Register() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, map[string]string{
			"status":           "ok",
			"currentAuthority": "user",
		})
	}
}
