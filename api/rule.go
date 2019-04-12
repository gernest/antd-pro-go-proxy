package api

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo"
)

type TableListItem struct {
	Key       int       `json:"key"`
	Disabled  bool      `json:"disabled"`
	Href      string    `json:"href"`
	Avatar    string    `json:"avatar"`
	Name      string    `json:"name"`
	Title     string    `json:"title"`
	Owner     string    `json:"owner"`
	Desc      string    `json:"desc"`
	CallNo    int64     `json:"callNo"`
	Status    int       `json:"status"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"createdAt"`
	Progress  int64     `json:"progress"`
}

func (i TableListItem) Less(j TableListItem, field, order string) bool {
	switch field {
	case "status":
		if order == "descend" {
			return j.Status < i.Status
		}
		return i.Status < j.Status
	case "updatedAt":
		if order == "descend" {
			return j.UpdatedAt.Before(i.UpdatedAt)
		}
		return i.UpdatedAt.Before(j.UpdatedAt)
	}
	return true
}

func setupTableList() []TableListItem {
	var tb []TableListItem
	now := time.Now()
	day := 24 * time.Hour
	av := []string{
		"https://gw.alipayobjects.com/zos/rmsportal/eeHMaZBwmTvLdIwMfBpg.png",
		"https://gw.alipayobjects.com/zos/rmsportal/udxAbMEhpwthVVcjLXik.png",
	}
	for i := 0; i < 46; i++ {
		tb = append(tb, TableListItem{
			Key:       i,
			Disabled:  i%6 == 0,
			Href:      "https://ant.design",
			Avatar:    av[i%2],
			Name:      fmt.Sprintf("TradeCode %d", i),
			Title:     fmt.Sprintf("a task name %d", i),
			Owner:     "Qu Lili",
			Desc:      "This is a description",
			CallNo:    int64(math.Floor(rand.Float64() * 10000)),
			Status:    int(math.Floor(rand.Float64()*10)) % 4,
			UpdatedAt: now.Add(day * time.Duration((i/2)+1)),
			CreatedAt: now.Add(day * time.Duration((i/2)+1)),
			Progress:  int64(math.Ceil(rand.Float64() * 1000)),
		})
	}
	return tb
}

var tableListDataSource = setupTableList()

func filterTableList(s []TableListItem, fn func(TableListItem) bool) []TableListItem {
	var o []TableListItem
	for _, v := range s {
		if fn(v) {
			o = append(o, v)
		}
	}
	return o
}

func GetRule() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		dataSource := make([]TableListItem, len(tableListDataSource))
		copy(dataSource, tableListDataSource)
		sorter := ctx.QueryParam("sorter")
		if sorter != "" {
			s := strings.Split(sorter, "_")
			sort.Slice(dataSource, func(i, j int) bool {
				prev, next := dataSource[i], dataSource[j]
				return prev.Less(next, s[0], s[1])
			})
		}
		status := ctx.QueryParam("status")
		if status != "" {
			var s []int
			for _, v := range strings.Split(status, ",") {
				a, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s = append(s, a)
			}
			dataSource = filterTableList(dataSource, func(v TableListItem) bool {
				for _, a := range s {
					if v.Status == a {
						return true
					}
				}
				return false
			})
		}
		name := ctx.QueryParam("name")
		if name != "" {
			dataSource = filterTableList(dataSource, func(v TableListItem) bool {
				return strings.Contains(v.Name, name)
			})
		}
		size := 10
		ps := ctx.QueryParam("pageSize")
		if ps != "" {
			a, err := strconv.Atoi(ps)
			if err != nil {
				return err
			}
			size = a
		}
		current := 1
		cp := ctx.QueryParam("currentPage")
		if cp != "" {
			a, err := strconv.Atoi(cp)
			if err != nil {
				return err
			}
			current = a
		}
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"list": dataSource,
			"pagination": map[string]interface{}{
				"total":    len(dataSource),
				"pageSize": size,
				"current":  current,
			},
		})

	}
}

func PostRule() echo.HandlerFunc {
	get := GetRule()
	type RuleReq struct {
		TableListItem
		Method string `json:"method"`
	}
	av := []string{
		"https://gw.alipayobjects.com/zos/rmsportal/eeHMaZBwmTvLdIwMfBpg.png",
		"https://gw.alipayobjects.com/zos/rmsportal/udxAbMEhpwthVVcjLXik.png",
	}
	return func(ctx echo.Context) error {
		var req RuleReq
		err := json.NewDecoder(ctx.Request().Body).Decode(&req)
		if err != nil {
			return err
		}
		switch req.Method {
		case "delete":
			tableListDataSource = filterTableList(tableListDataSource, func(v TableListItem) bool {
				return v.Key == req.Key
			})
		case "post":
			now := time.Now()
			i := int(math.Ceil(rand.Float64() * 10000))
			tableListDataSource = append(tableListDataSource, TableListItem{
				Key:       i,
				Disabled:  i%6 == 0,
				Href:      "https://ant.design",
				Avatar:    av[i%2],
				Name:      fmt.Sprintf("TradeCode %d", i),
				Title:     fmt.Sprintf("a task name %d", i),
				Owner:     "Qu Lili",
				Desc:      req.Desc,
				CallNo:    int64(math.Floor(rand.Float64() * 10000)),
				Status:    int(math.Floor(rand.Float64()*10)) % 4,
				UpdatedAt: now.Add(24 * time.Hour * time.Duration((i/2)+1)),
				CreatedAt: now.Add(24 * time.Hour * time.Duration((i/2)+1)),
				Progress:  int64(math.Ceil(rand.Float64() * 1000)),
			})
		case "update":
			for k := range tableListDataSource {
				if req.Key == tableListDataSource[k].Key {
					tableListDataSource[k] = req.TableListItem
					break
				}
			}
		}
		return get(ctx)
	}
}
