package api

import (
	"net/http"

	"github.com/labstack/echo"
)

type Operation struct {
	Key       string `json:"key"`
	Type      string `json:"type"`
	Name      string `json:"name"`
	Status    string `json:"status"`
	UpdatedAt string `json:"UpdatedAt"`
	Memo      string `json:"memo"`
}

var advancedOperation1 = []Operation{
	{
		Key:       "op1",
		Type:      "订购关系生效",
		Name:      "曲丽丽",
		Status:    "agree",
		UpdatedAt: "2017-10-03  19:23:12",
		Memo:      "-",
	},
	{
		Key:       "op2",
		Type:      "财务复审",
		Name:      "付小小",
		Status:    "reject",
		UpdatedAt: "2017-10-03  19:23:12",
		Memo:      "不通过原因",
	},
	{
		Key:       "op3",
		Type:      "部门初审",
		Name:      "周毛毛",
		Status:    "agree",
		UpdatedAt: "2017-10-03  19:23:12",
		Memo:      "-",
	},
	{
		Key:       "op4",
		Type:      "提交订单",
		Name:      "林东东",
		Status:    "agree",
		UpdatedAt: "2017-10-03  19:23:12",
		Memo:      "很棒",
	},
	{
		Key:       "op5",
		Type:      "创建订单",
		Name:      "汗牙牙",
		Status:    "agree",
		UpdatedAt: "2017-10-03  19:23:12",
		Memo:      "-",
	},
}

var advancedOperation2 = []Operation{
	{
		Key:       "op1",
		Type:      "订购关系生效",
		Name:      "曲丽丽",
		Status:    "agree",
		UpdatedAt: "2017-10-03  19:23:12",
		Memo:      "-",
	},
}

var advancedOperation3 = []Operation{
	{
		Key:       "op1",
		Type:      "创建订单",
		Name:      "汗牙牙",
		Status:    "agree",
		UpdatedAt: "2017-10-03  19:23:12",
		Memo:      "-",
	},
}

func GetProfileAdvancedData() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"advancedOperation1": advancedOperation1,
			"advancedOperation2": advancedOperation2,
			"advancedOperation3": advancedOperation3,
		})
	}
}

type Application struct {
	ID           string `json:"id"`
	Status       string `json:"status"`
	OrderNo      string `json:"orderNo"`
	ChildOrderNo string `json:"childOrderNo"`
}

type UserInfo struct {
	Name     string `json:"name"`
	Tel      string `json:"tel"`
	Delivery string `json:"delivery"`
	Addr     string `json:"addr"`
	Remark   string `json:"remark"`
}

type Progress struct {
	Key      string `json:"key"`
	Time     string `json:"time"`
	Rate     string `json:"rate"`
	Status   string `json:"status"`
	Operator string `json:"operator"`
	Cost     string `json:"cost"`
}

var basicProgress = []Progress{
	{
		Key:      "1",
		Time:     "2017-10-01 14:10",
		Rate:     "联系客户",
		Status:   "processing",
		Operator: "取货员 ID1234",
		Cost:     "5mins",
	},
	{
		Key:      "2",
		Time:     "2017-10-01 14:05",
		Rate:     "取货员出发",
		Status:   "success",
		Operator: "取货员 ID1234",
		Cost:     "1h",
	},
	{
		Key:      "3",
		Time:     "2017-10-01 13:05",
		Rate:     "取货员接单",
		Status:   "success",
		Operator: "取货员 ID1234",
		Cost:     "5mins",
	},
	{
		Key:      "4",
		Time:     "2017-10-01 13:00",
		Rate:     "申请审批通过",
		Status:   "success",
		Operator: "系统",
		Cost:     "1h",
	},
	{
		Key:      "5",
		Time:     "2017-10-01 12:00",
		Rate:     "发起退货申请",
		Status:   "success",
		Operator: "用户",
		Cost:     "5mins",
	},
}

type Good struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Barcode string `json:"barcode"`
	Price   string `json:"price"`
	Num     string `json:"num"`
	Amount  string `json:"amount"`
}

var basicGoods = []Good{
	{
		ID:      "1234561",
		Name:    "矿泉水 550ml",
		Barcode: "12421432143214321",
		Price:   "2.00",
		Num:     "1",
		Amount:  "2.00",
	},
	{
		ID:      "1234562",
		Name:    "凉茶 300ml",
		Barcode: "12421432143214322",
		Price:   "3.00",
		Num:     "2",
		Amount:  "6.00",
	},
	{
		ID:      "1234563",
		Name:    "好吃的薯片",
		Barcode: "12421432143214323",
		Price:   "7.00",
		Num:     "4",
		Amount:  "28.00",
	},
	{
		ID:      "1234564",
		Name:    "特别好吃的蛋卷",
		Barcode: "12421432143214324",
		Price:   "8.50",
		Num:     "3",
		Amount:  "25.50",
	},
}

func GetProfileBasic() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id := ctx.QueryParam("id")
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"userInfo": UserInfo{
				Name:     "random",
				Tel:      "18100000000",
				Delivery: "菜鸟物流",
				Addr:     "浙江省杭州市西湖区万塘路18号",
				Remark:   "备注",
			},
			"application": Application{
				ID:           id,
				Status:       "已取货",
				OrderNo:      "random",
				ChildOrderNo: "random",
			},
			"basicGoods":    basicGoods,
			"basicProgress": basicProgress,
		})
	}
}
