package api

import (
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

type Point struct {
	X string `json:"x"`
	Y int64  `json:"y"`
}

type SearchData struct {
	Index   int    `json:"index"`
	Keyword string `json:"keyword"`
	Count   int    `json:"count"`
	Range   int    `json:"range"`
	Status  int    `json:"status"`
}

type OfflineData struct {
	Name string `json:"name"`
	CVR  int    `json:"cvr"`
}

type OfflineChartData struct {
	X  int64 `json:"x"`
	Y1 int64 `json:"y1"`
	Y2 int64 `json:"y2"`
}

type RadarOriginData struct {
	Name       string `json:"name"`
	Ref        int    `json:"ref"`
	Koubei     int    `json:"koubei"`
	Output     int    `json:"output"`
	Contribute int    `json:"contribute"`
	Hot        int    `json:"hot"`
}

type RadarData struct {
	Name  string `json:"name"`
	Label string `json:"label"`
	Value int    `json:"value"`
}

func GetChartData() echo.HandlerFunc {
	format := "2006-01-02"
	begin := time.Now()
	day := 24 * time.Hour
	var visitData []Point
	fakeY := []int64{7, 5, 4, 2, 4, 7, 5, 6, 5, 9, 6, 3, 1, 5, 3, 6, 5}
	for i, v := range fakeY {
		x := begin.Add(day * time.Duration(i)).Format(format)
		visitData = append(visitData, Point{X: x, Y: v})
	}
	var visitData2 []Point
	fakeY2 := []int64{1, 6, 4, 8, 3, 7, 2}
	for i, v := range fakeY2 {
		x := begin.Add(day * time.Duration(i)).Format(format)
		visitData2 = append(visitData2, Point{X: x, Y: v})
	}
	var salesData []Point
	for i := 0; i < 12; i++ {
		salesData = append(salesData, Point{
			X: fmt.Sprintf("%d month", i+1),
			Y: int64(math.Floor(rand.Float64()*1000)) + 200,
		})
	}
	var searchData []SearchData
	for i := 0; i < 12; i++ {
		searchData = append(searchData, SearchData{
			Index:   i + 1,
			Keyword: fmt.Sprintf("search keyword-%d", i),
			Count:   int(math.Floor(rand.Float64() * 1000)),
			Range:   int(math.Floor(rand.Float64() * 100)),
			Status:  int(math.Floor(rand.Float64()*10)) % 2,
		})
	}
	salesTypeData := []Point{
		{
			X: "Household appliances",
			Y: 4544,
		},
		{
			X: "Edible wine",
			Y: 3321,
		},
		{
			X: "Health care",
			Y: 3113,
		},
		{
			X: "Clothing bags",
			Y: 2341,
		},
		{
			X: "Maternal and child products",
			Y: 1231,
		},
		{
			X: "other",
			Y: 1231,
		},
	}
	salesTypeDataOnline := []Point{

		{
			X: "Household appliances",
			Y: 244,
		},
		{
			X: "Edible wine",
			Y: 321,
		},
		{
			X: "Health care",
			Y: 311,
		},
		{
			X: "Clothing bags",
			Y: 41,
		},
		{
			X: "Maternal and child products",
			Y: 121,
		},
		{
			X: "other",
			Y: 111,
		},
	}
	salesTypeDataOffline := []Point{
		{
			X: "Household appliances",
			Y: 99,
		},
		{
			X: "Edible wine",
			Y: 188,
		},
		{
			X: "Health care",
			Y: 344,
		},
		{
			X: "Clothing bags",
			Y: 255,
		},
		{
			X: "other",
			Y: 65,
		},
	}
	var offlineData []OfflineData
	for i := 0; i < 10; i++ {
		offlineData = append(offlineData, OfflineData{
			Name: fmt.Sprintf("Stores %d", i),
			CVR:  (rand.Int() * 9) / 10,
		})
	}
	var offlineChartData []OfflineChartData
	for i := 0; i < 10; i++ {
		offlineChartData = append(offlineChartData, OfflineChartData{
			X:  begin.Unix() + 1000*60*30*int64(i),
			Y1: int64(rand.Int()*100) + 10,
			Y2: int64(rand.Int()*100) + 10,
		})
	}
	radarOriginData := []RadarOriginData{
		{
			Name:       "personal",
			Ref:        10,
			Koubei:     8,
			Output:     4,
			Contribute: 5,
			Hot:        7,
		},
		{
			Name:       "team",
			Ref:        3,
			Koubei:     9,
			Output:     6,
			Contribute: 3,
			Hot:        1,
		},
		{
			Name:       "department",
			Ref:        4,
			Koubei:     1,
			Output:     6,
			Contribute: 5,
			Hot:        7,
		},
	}
	var radarData []RadarData
	for _, v := range radarOriginData {
		radarData = append(radarData, RadarData{
			Name:  "ref",
			Label: "Reference",
			Value: v.Ref,
		})
		radarData = append(radarData, RadarData{
			Name:  "koubei",
			Label: "Word of mouth",
			Value: v.Koubei,
		})
		radarData = append(radarData, RadarData{
			Name:  "output",
			Label: "Output",
			Value: v.Output,
		})
		radarData = append(radarData, RadarData{
			Name:  "contribute",
			Label: "Contribute",
			Value: v.Contribute,
		})
		radarData = append(radarData, RadarData{
			Name:  "hot",
			Label: "Hot",
			Value: v.Hot,
		})
	}
	return func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"visitData":            visitData,
			"visitData2":           visitData2,
			"salesData":            salesData,
			"searchData":           searchData,
			"offlineData":          offlineData,
			"offlineChartData":     offlineChartData,
			"salesTypeData":        salesTypeData,
			"salesTypeDataOnline":  salesTypeDataOnline,
			"salesTypeDataOffline": salesTypeDataOffline,
			"radarData":            radarData,
		})
	}
}
