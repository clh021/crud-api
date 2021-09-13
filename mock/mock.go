package mock

import (
	"gorm.io/gorm"
)

type mockManager struct {
	db *gorm.DB
}

func NewMock(db *gorm.DB) mockManager {
	return mockManager{db}
}

func (m mockManager) ReSetAll() {
	m.initTestTable()
	m.initTestData()
}
func (m mockManager) initTestTable() {
	m.db.AutoMigrate(&TestTable{})
}
func (m mockManager) initTestData() {
	data := []TestTable{
		{
			Tag:     "歌词",
			Name:    "小周",
			Content: `1983年小巷12月晴朗，夜的第七章，打字机继续推向，接近事实的那下一行`,
			Address: "湖北省武汉市",
			Contact: "18012345678",
		},
		{
			Tag:     "歌词",
			Name:    "小杰",
			Content: `石楠烟斗的雾 飘向枯萎的树，沉默的对我哭诉，贝克街旁的圆形广场，盔甲骑士臂上，鸢尾花的徽章 微亮`,
			Address: "湖北省武汉市",
			Contact: "18012345678",
		},
		{
			Tag:     "歌词",
			Name:    "小伦",
			Content: `无人马车声响，深夜的拜访，邪恶 在维多利亚的月光下，血色的开场，消失的手枪 焦黑的手杖 融化的蜡像，谁不在场 珠宝箱上 符号的假象，矛盾通往 他堆砌的死巷，证据被完美埋葬，那嘲弄苏格兰警场 的嘴角上扬`,
			Address: "湖北省武汉市",
			Contact: "18012345678",
		},
		{
			Tag:     "歌词",
			Name:    "小夜",
			Content: `女声：如果邪恶是华丽残酷的乐章（那么正义是深沉无奈的惆怅）女声：它的终场我会亲手写上（那我就点亮在灰烬中的微光）`,
			Address: "湖北省武汉市",
			Contact: "18012345678",
		},
		{
			Tag:     "歌词",
			Name:    "小七",
			Content: `女声：晨曦的光 风干最后一行忧伤（那么雨滴 会洗净黑暗的高墙）女声：黑色的墨 染上安详`,
			Address: "湖北省武汉市",
			Contact: "18012345678",
		},
	}

	for _, d := range data {
		d.Base = Base{
			ID:      NewUUID(),
			Created: NewUnixtime(),
		}
		// m.db.Debug().FirstOrInit(&d, TestTable{
		// 	Tag:     d.Tag,
		// 	Name:    d.Name,
		// 	Content: d.Content,
		// 	Address: d.Address,
		// 	Contact: d.Contact,
		// })
		m.db.Debug().FirstOrCreate(&d, TestTable{
			Tag:     d.Tag,
			Name:    d.Name,
			Content: d.Content,
			Address: d.Address,
			Contact: d.Contact,
		})
	}
}
