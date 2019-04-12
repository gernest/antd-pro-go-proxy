package api

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

var titles = []string{
	"Alipay",
	"Angular",
	"Ant Design",
	"Ant Design Pro",
	"Bootstrap",
	"React",
	"Vue",
	"Webpack",
}

var avatars = []string{
	"https://gw.alipayobjects.com/zos/rmsportal/WdGqmHpayyMjiEhcKoVE.png", // Alipay
	"https://gw.alipayobjects.com/zos/rmsportal/zOsKZmFRdUtvpqCImOVY.png", // Angular
	"https://gw.alipayobjects.com/zos/rmsportal/dURIMkkrRFpPgTuzkwnB.png", // Ant Design
	"https://gw.alipayobjects.com/zos/rmsportal/sfjbOqnsXXJgNCjCzDBL.png", // Ant Design Pro
	"https://gw.alipayobjects.com/zos/rmsportal/siCrBXXhmvTQGWPNLBow.png", // Bootstrap
	"https://gw.alipayobjects.com/zos/rmsportal/kZzEzemZyKLKFsojXItE.png", // React
	"https://gw.alipayobjects.com/zos/rmsportal/ComBAopevLwENQdKWiIn.png", // Vue
	"https://gw.alipayobjects.com/zos/rmsportal/nxkuOJlFJuAUhzlMTCEe.png", // Webpack
}

var avatars2 = []string{
	"https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png",
	"https://gw.alipayobjects.com/zos/rmsportal/cnrhVkzwxjPwAaCfPbdc.png",
	"https://gw.alipayobjects.com/zos/rmsportal/gaOngJwsRYRaVAuXXcmB.png",
	"https://gw.alipayobjects.com/zos/rmsportal/ubnKSIfAJTxIgXOKlciN.png",
	"https://gw.alipayobjects.com/zos/rmsportal/WhxKECPNujWoWEFNdnJE.png",
	"https://gw.alipayobjects.com/zos/rmsportal/jZUIxmJycoymBprLOUbT.png",
	"https://gw.alipayobjects.com/zos/rmsportal/psOgztMplJMGpVEqfcgF.png",
	"https://gw.alipayobjects.com/zos/rmsportal/ZpBqSxLxVEXfcUNoPKrz.png",
	"https://gw.alipayobjects.com/zos/rmsportal/laiEnJdGHVOhJrUShBaJ.png",
	"https://gw.alipayobjects.com/zos/rmsportal/UrQsqscbKEpNuJcvBZBu.png",
}

var covers = []string{
	"https://gw.alipayobjects.com/zos/rmsportal/uMfMFlvUuceEyPpotzlq.png",
	"https://gw.alipayobjects.com/zos/rmsportal/iZBVOIhGJiAnhplqjvZW.png",
	"https://gw.alipayobjects.com/zos/rmsportal/iXjVmWVHbCJAyqvDxdtx.png",
	"https://gw.alipayobjects.com/zos/rmsportal/gLaIAoVWTtLbBWZNYEMg.png",
}

var desc = []string{
	"那是一种内在的东西， 他们到达不了，也无法触及的",
	"希望是一个好东西，也许是最好的，好东西是不会消亡的",
	"生命就像一盒巧克力，结果往往出人意料",
	"城镇中有那么多的酒馆，她却偏偏走进了我的酒馆",
	"那时候我只会想自己想要什么，从不想自己拥有什么",
}

var user = []string{
	"付小小",
	"曲丽丽",
	"林东东",
	"周星星",
	"吴加好",
	"朱偏右",
	"鱼酱",
	"乐哥",
	"谭小仪",
	"仲尼",
}

type Member struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}
type Group struct {
	ID             string   `json:"id"`
	Owner          string   `json:"owner"`
	Title          string   `json:"title"`
	Avatar         string   `json:"avatar"`
	Cover          string   `json:"cover"`
	Status         string   `json:"status"`
	Percent        float64  `json:"percent"`
	Logo           string   `json:"logo"`
	CreatedAt      int64    `json:"createdAt"`
	UpdatedAt      int64    `json:"updatedAt"`
	SubDescription string   `json:"subDescription"`
	Description    string   `json:"description"`
	ActiveUser     int64    `json:"activeUser"`
	NewUser        int64    `json:"newUser"`
	Star           int64    `json:"star"`
	Like           int64    `json:"Like"`
	Message        int64    `json:"message"`
	Content        string   `json:"content"`
	Members        []Member `json:"members"`
}

func fakeList(count int) (result []Group) {
	fc := func(c int) string {
		x := c / 4
		if x%2 == 0 {
			return covers[c%4]
		}
		return covers[3-(c%4)]
	}
	stat := []string{"active", "exception", "normal"}
	d := "In the development process of intermediate products, different design specifications and implementation methods will appear, but there are often many similar pages and components, and these similar components will be separated into a set of standard specifications."
	ts := time.Now()
	c := "The paragraph indicates: ant gold clothing design platform ant.design, with minimal workload, seamless access to ant gold clothing ecology, providing experience solutions across design and development. Ant ant design platform ant.design, with minimal workload, seamless access to ant gold clothing ecology, providing experience solutions across design and development"
	elapse := 2 * time.Hour
	for i := 0; i < count; i++ {
		result = append(result, Group{
			ID:             fmt.Sprintf("fake-list-%d", i),
			Owner:          user[i%10],
			Title:          titles[i%8],
			Avatar:         avatars[i%8],
			Cover:          fc(i),
			Status:         stat[i%3],
			Percent:        math.Ceil(rand.Float64()*50) + 50,
			Logo:           avatars[i%8],
			CreatedAt:      ts.Add(elapse * time.Duration(i) * -1).Unix(),
			UpdatedAt:      ts.Add(elapse * time.Duration(i) * -1).Unix(),
			SubDescription: desc[i%5],
			Description:    d,
			ActiveUser:     int64(math.Ceil(rand.Float64()*100000)) + 100000,
			NewUser:        int64(math.Ceil(rand.Float64()*1000)) + 1000,
			Star:           int64(math.Ceil(rand.Float64()*100)) + 100,
			Like:           int64(math.Ceil(rand.Float64()*100)) + 100,
			Message:        int64(math.Ceil(rand.Float64()*10)) + 10,
			Content:        c,
			Members: []Member{
				{
					Avatar: "https://gw.alipayobjects.com/zos/rmsportal/ZiESqWwCXBRQoaPONSJe.png",
					Name:   "Qu Lili",
					ID:     "member1",
				},
				{
					Avatar: "https://gw.alipayobjects.com/zos/rmsportal/tBOxZPlITHqwlGjsJWaF.png",
					Name:   "Wang Zhaojun",
					ID:     "member2",
				},
				{
					Avatar: "https://gw.alipayobjects.com/zos/rmsportal/sBxjgqiuHMGRkIjqlQCd.png",
					Name:   "Dong Nana",
					ID:     "member3",
				},
			},
		})
	}
	return
}

var sourceData []Group

func GetFakeList() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		count, _ := strconv.Atoi(ctx.QueryParam("count"))
		if count == 0 {
			count = 20
		}
		r := fakeList(count)
		sourceData = r
		return ctx.JSON(http.StatusOK, r)
	}
}

type postFakeReq struct {
	Method string `json:"method"`
	Group
}

func PostFakeList() echo.HandlerFunc {
	filter := func(in []Group, fn func(*Group) bool) (out []Group) {
		for _, v := range in {
			if fn(&v) {
				out = append(out, v)
			}
		}
		return
	}
	return func(ctx echo.Context) error {
		var r postFakeReq
		json.NewDecoder(ctx.Request().Body).Decode(&r)
		result := sourceData
		switch r.Method {
		case "delete":
			result = filter(result, func(g *Group) bool {
				return g.ID != r.ID
			})
		case "update":
			for k := range result {
				if result[k].ID == r.ID {
					result[k] = r.Group
				}
			}
		case "post":
			r.ID = fmt.Sprintf("fake-list-%d", len(result))
			r.CreatedAt = time.Now().Unix()
			result = append([]Group{r.Group}, result...)
		}
		return ctx.JSON(http.StatusOK, result)
	}
}

func GetFakeCaptcha() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, "captcha-xxx")
	}
}

func PostForms() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, map[string]string{
			"message": "OK",
		})
	}
}

type Item struct {
	Name string `json:"name"`
	Link string `json:"link"`
}

type Activity struct {
	ID        string `json:"id"`
	UpdatedAt int64  `json:"updatedAt"`
	User      *User  `json:"user"`
	Group     Item   `json:"group"`
	Project   Item   `json:"project"`
	Comment   Item   `json:"comment,omitempty"`
	Template  string `json:"template"`
}

func GetActivity() echo.HandlerFunc {
	now := time.Now().Unix()
	activities := []Activity{
		{
			ID:        "trend-1",
			UpdatedAt: now,
			User: &User{
				Name:   "曲丽丽",
				Avatar: avatars2[0],
			},
			Group: Item{
				Name: "高逼格设计天团",
				Link: "http://github.com/",
			},
			Project: Item{
				Name: "六月迭代",
				Link: "http://github.com/",
			},
			Template: "在 @{group} 新建项目 @{project}",
		},
		{
			ID:        "trend-2",
			UpdatedAt: now,
			User: &User{
				Name:   "付小小",
				Avatar: avatars2[1],
			},
			Group: Item{
				Name: "高逼格设计天团",
				Link: "http://github.com/",
			},
			Project: Item{
				Name: "六月迭代",
				Link: "http://github.com/",
			},
			Template: "在 @{group} 新建项目 @{project}",
		},
		{
			ID:        "trend-3",
			UpdatedAt: now,
			User: &User{
				Name:   "林东东",
				Avatar: avatars2[2],
			},
			Group: Item{
				Name: "中二少女团",
				Link: "http://github.com/",
			},
			Project: Item{
				Name: "六月迭代",
				Link: "http://github.com/",
			},
			Template: "在 @{group} 新建项目 @{project}",
		},
		{
			ID:        "trend-4",
			UpdatedAt: now,
			User: &User{
				Name:   "周星星",
				Avatar: avatars2[4],
			},
			Project: Item{
				Name: "5 月日常迭代",
				Link: "http://github.com/",
			},
			Template: "将 @{project} 更新至已发布状态",
		},
		{
			ID:        "trend-5",
			UpdatedAt: now,
			User: &User{
				Name:   "朱偏右",
				Avatar: avatars2[3],
			},
			Project: Item{
				Name: "工程效能",
				Link: "http://github.com/",
			},
			Comment: Item{
				Name: "留言",
				Link: "http://github.com/",
			},
			Template: "在 @{project} 发布了 @{comment}",
		},
		{
			ID:        "trend-6",
			UpdatedAt: now,
			User: &User{
				Name:   "乐哥",
				Avatar: avatars2[5],
			},
			Group: Item{
				Name: "程序员日常",
				Link: "http://github.com/",
			},
			Project: Item{
				Name: "品牌迭代",
				Link: "http://github.com/",
			},
			Template: "在 @{group} 新建项目 @{project}",
		},
	}

	return func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, activities)
	}
}

type Notice struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Logo        string `json:"logo"`
	Description string `json:"description"`
	UpdatedAt   int64  `json:"updatedAt"`
	Member      string `json:"member"`
	HREF        string `json:"href"`
	MemberLink  string `json:"memberLink"`
}

type TagListItem struct {
	Name  string `json:"name"`
	Type  int    `json:"type"`
	Value int    `json:"value"`
}

func GetNotice() echo.HandlerFunc {
	now := time.Now().Unix()
	notices := []Notice{
		{
			ID:          "xxx1",
			Title:       titles[0],
			Logo:        avatars[0],
			Description: "It’s an inner thing that they can’t reach and can’t reach.",
			UpdatedAt:   now,
			Member:      "Scientific moving bricks",
			HREF:        "",
			MemberLink:  "",
		},
		{
			ID:          "xxx2",
			Title:       titles[1],
			Logo:        avatars[1],
			Description: "Hope is a good thing, maybe the best, good things will not die.",
			UpdatedAt:   now,
			Member:      "Wu Yanzu",
			HREF:        "",
			MemberLink:  "",
		},
		{
			ID:          "xxx3",
			Title:       titles[2],
			Logo:        avatars[2],
			Description: "There are so many pubs in the town, but she walked into my pub.",
			UpdatedAt:   now,
			Member:      "Secondary 2 Girls Group",
			HREF:        "",
			MemberLink:  "",
		},
		{
			ID:          "xxx4",
			Title:       titles[3],
			Logo:        avatars[3],
			Description: "At that time, I only thought about what I wanted, and I didn’t want to have what I had.",
			UpdatedAt:   now,
			Member:      "Programmer everyday",
			HREF:        "",
			MemberLink:  "",
		},
		{
			ID:          "xxx5",
			Title:       titles[4],
			Logo:        avatars[4],
			Description: "Winter is coming",
			UpdatedAt:   now,
			Member:      "High-force design",
			HREF:        "",
			MemberLink:  "",
		},
		{
			ID:          "xxx6",
			Title:       titles[5],
			Logo:        avatars[5],
			Description: "Life is like a box of chocolates, and the results are often unexpected.",
			UpdatedAt:   now,
			Member:      "Lie to learn computer",
			HREF:        "",
			MemberLink:  "",
		},
	}
	return func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, notices)
	}
}

func GetTags() echo.HandlerFunc {
	tags := []TagListItem{
		{
			Name:  "梅州市",
			Value: 19,
			Type:  1,
		},
		{
			Name:  "云浮市",
			Value: 45,
			Type:  1,
		},
		{
			Name:  "吉安市",
			Value: 58,
			Type:  2,
		},
		{
			Name:  "佛山市",
			Value: 85,
			Type:  0,
		},
		{
			Name:  "玉树藏族自治州",
			Value: 86,
			Type:  0,
		},
		{
			Name:  "保定市",
			Value: 50,
			Type:  0,
		},
		{
			Name:  "南投县",
			Value: 46,
			Type:  1,
		},
		{
			Name:  "茂名市",
			Value: 5,
			Type:  1,
		},
		{
			Name:  "南平市",
			Value: 42,
			Type:  0,
		},
		{
			Name:  "营口市",
			Value: 33,
			Type:  2,
		},
		{
			Name:  "阳泉市",
			Value: 24,
			Type:  1,
		},
		{
			Name:  "重庆市",
			Value: 16,
			Type:  1,
		},
		{
			Name:  "固原市",
			Value: 82,
			Type:  2,
		},
		{
			Name:  "锦州市",
			Value: 85,
			Type:  1,
		},
		{
			Name:  "梧州市",
			Value: 75,
			Type:  2,
		},
		{
			Name:  "株洲市",
			Value: 26,
			Type:  2,
		},
		{
			Name:  "中卫市",
			Value: 19,
			Type:  1,
		},
		{
			Name:  "自贡市",
			Value: 6,
			Type:  1,
		},
		{
			Name:  "黄南藏族自治州",
			Value: 84,
			Type:  2,
		},
		{
			Name:  "长春市",
			Value: 59,
			Type:  2,
		},
		{
			Name:  "鹤壁市",
			Value: 64,
			Type:  1,
		},
		{
			Name:  "昭通市",
			Value: 63,
			Type:  2,
		},
		{
			Name:  "大同市",
			Value: 83,
			Type:  1,
		},
		{
			Name:  "烟台市",
			Value: 8,
			Type:  1,
		},
		{
			Name:  "海外",
			Value: 76,
			Type:  1,
		},
		{
			Name:  "北京市",
			Value: 9,
			Type:  2,
		},
		{
			Name:  "玉树藏族自治州",
			Value: 42,
			Type:  1,
		},
		{
			Name:  "青岛市",
			Value: 25,
			Type:  1,
		},
		{
			Name:  "平凉市",
			Value: 81,
			Type:  1,
		},
		{
			Name:  "北京市",
			Value: 23,
			Type:  0,
		},
		{
			Name:  "台中市",
			Value: 43,
			Type:  0,
		},
		{
			Name:  "商洛市",
			Value: 30,
			Type:  1,
		},
		{
			Name:  "天津市",
			Value: 61,
			Type:  1,
		},
		{
			Name:  "来宾市",
			Value: 52,
			Type:  1,
		},
		{
			Name:  "澳门半岛",
			Value: 9,
			Type:  0,
		},
		{
			Name:  "十堰市",
			Value: 12,
			Type:  1,
		},
		{
			Name:  "金门县",
			Value: 13,
			Type:  0,
		},
		{
			Name:  "龙岩市",
			Value: 84,
			Type:  0,
		},
		{
			Name:  "鞍山市",
			Value: 40,
			Type:  1,
		},
		{
			Name:  "海东市",
			Value: 96,
			Type:  1,
		},
		{
			Name:  "贵港市",
			Value: 61,
			Type:  1,
		},
		{
			Name:  "金华市",
			Value: 57,
			Type:  2,
		},
		{
			Name:  "大庆市",
			Value: 59,
			Type:  2,
		},
		{
			Name:  "喀什地区",
			Value: 41,
			Type:  1,
		},
		{
			Name:  "海东市",
			Value: 21,
			Type:  0,
		},
		{
			Name:  "南平市",
			Value: 87,
			Type:  2,
		},
		{
			Name:  "海东市",
			Value: 54,
			Type:  1,
		},
		{
			Name:  "泸州市",
			Value: 73,
			Type:  1,
		},
		{
			Name:  "上海市",
			Value: 77,
			Type:  0,
		},
		{
			Name:  "银川市",
			Value: 71,
			Type:  1,
		},
		{
			Name:  "平顶山市",
			Value: 29,
			Type:  1,
		},
		{
			Name:  "天津市",
			Value: 18,
			Type:  1,
		},
		{
			Name:  "海口市",
			Value: 51,
			Type:  1,
		},
		{
			Name:  "宁波市",
			Value: 45,
			Type:  2,
		},
		{
			Name:  "庆阳市",
			Value: 62,
			Type:  1,
		},
		{
			Name:  "延安市",
			Value: 73,
			Type:  1,
		},
		{
			Name:  "锦州市",
			Value: 92,
			Type:  1,
		},
		{
			Name:  "海外",
			Value: 33,
			Type:  1,
		},
		{
			Name:  "沧州市",
			Value: 92,
			Type:  1,
		},
		{
			Name:  "白山市",
			Value: 60,
			Type:  0,
		},
		{
			Name:  "澎湖县",
			Value: 43,
			Type:  1,
		},
		{
			Name:  "黄南藏族自治州",
			Value: 51,
			Type:  1,
		},
		{
			Name:  "潮州市",
			Value: 39,
			Type:  2,
		},
		{
			Name:  "厦门市",
			Value: 36,
			Type:  1,
		},
		{
			Name:  "长春市",
			Value: 62,
			Type:  1,
		},
		{
			Name:  "平凉市",
			Value: 94,
			Type:  2,
		},
		{
			Name:  "宜昌市",
			Value: 90,
			Type:  0,
		},
		{
			Name:  "北海市",
			Value: 28,
			Type:  0,
		},
		{
			Name:  "吴忠市",
			Value: 79,
			Type:  1,
		},
		{
			Name:  "达州市",
			Value: 40,
			Type:  1,
		},
		{
			Name:  "新界",
			Value: 29,
			Type:  1,
		},
		{
			Name:  "威海市",
			Value: 7,
			Type:  2,
		},
		{
			Name:  "阿里地区",
			Value: 93,
			Type:  1,
		},
		{
			Name:  "双鸭山市",
			Value: 66,
			Type:  0,
		},
		{
			Name:  "三亚市",
			Value: 70,
			Type:  2,
		},
		{
			Name:  "通辽市",
			Value: 8,
			Type:  1,
		},
		{
			Name:  "鹤岗市",
			Value: 98,
			Type:  0,
		},
		{
			Name:  "宜春市",
			Value: 75,
			Type:  2,
		},
		{
			Name:  "铁岭市",
			Value: 17,
			Type:  1,
		},
		{
			Name:  "泉州市",
			Value: 28,
			Type:  2,
		},
		{
			Name:  "三沙市",
			Value: 79,
			Type:  2,
		},
		{
			Name:  "北京市",
			Value: 75,
			Type:  0,
		},
		{
			Name:  "巴中市",
			Value: 13,
			Type:  2,
		},
		{
			Name:  "晋中市",
			Value: 92,
			Type:  0,
		},
		{
			Name:  "庆阳市",
			Value: 35,
			Type:  1,
		},
		{
			Name:  "牡丹江市",
			Value: 36,
			Type:  0,
		},
		{
			Name:  "铜川市",
			Value: 83,
			Type:  2,
		},
		{
			Name:  "莆田市",
			Value: 88,
			Type:  1,
		},
		{
			Name:  "澎湖县",
			Value: 58,
			Type:  0,
		},
		{
			Name:  "三门峡市",
			Value: 59,
			Type:  1,
		},
		{
			Name:  "柳州市",
			Value: 45,
			Type:  1,
		},
		{
			Name:  "荆州市",
			Value: 58,
			Type:  1,
		},
		{
			Name:  "铜陵市",
			Value: 14,
			Type:  0,
		},
		{
			Name:  "海南藏族自治州",
			Value: 11,
			Type:  1,
		},
		{
			Name:  "天津市",
			Value: 97,
			Type:  0,
		},
		{
			Name:  "嘉义市",
			Value: 38,
			Type:  2,
		},
		{
			Name:  "聊城市",
			Value: 71,
			Type:  1,
		},
		{
			Name:  "阿里地区",
			Value: 79,
			Type:  2,
		},
		{
			Name:  "辽阳市",
			Value: 75,
			Type:  1,
		},
		{
			Name:  "海北藏族自治州",
			Value: 86,
			Type:  0,
		},
	}
	return func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"list": tags,
		})
	}
}
