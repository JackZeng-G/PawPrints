package seed

import (
	"pawprints-server/internal/model"
	"gorm.io/gorm"
)

func SeedPetData(db *gorm.DB) error {
	categories := []model.PetCategory{
		{Name: "猫", Icon: "cat", SortOrder: 1},
		{Name: "狗", Icon: "dog", SortOrder: 2},
		{Name: "兔子", Icon: "rabbit", SortOrder: 3},
		{Name: "仓鼠", Icon: "hamster", SortOrder: 4},
		{Name: "鸟", Icon: "bird", SortOrder: 5},
		{Name: "鱼", Icon: "fish", SortOrder: 6},
		{Name: "爬行动物", Icon: "reptile", SortOrder: 7},
	}

	for _, cat := range categories {
		if err := db.FirstOrCreate(&cat, model.PetCategory{Name: cat.Name}).Error; err != nil {
			return err
		}
	}

	var catCat, dogCat, rabbitCat, hamsterCat, birdCat, fishCat, reptileCat model.PetCategory
	db.Where("name = ?", "猫").First(&catCat)
	db.Where("name = ?", "狗").First(&dogCat)
	db.Where("name = ?", "兔子").First(&rabbitCat)
	db.Where("name = ?", "仓鼠").First(&hamsterCat)
	db.Where("name = ?", "鸟").First(&birdCat)
	db.Where("name = ?", "鱼").First(&fishCat)
	db.Where("name = ?", "爬行动物").First(&reptileCat)

	breeds := []model.PetBreed{
		{CategoryID: catCat.ID, Name: "布偶猫", ExpectedLifespan: "12-15年", Size: "大型"},
		{CategoryID: catCat.ID, Name: "英国短毛猫", ExpectedLifespan: "12-17年", Size: "中型"},
		{CategoryID: catCat.ID, Name: "美国短毛猫", ExpectedLifespan: "15-20年", Size: "中型"},
		{CategoryID: catCat.ID, Name: "暹罗猫", ExpectedLifespan: "12-15年", Size: "中型"},
		{CategoryID: catCat.ID, Name: "波斯猫", ExpectedLifespan: "12-17年", Size: "中型"},
		{CategoryID: catCat.ID, Name: "橘猫", ExpectedLifespan: "15-20年", Size: "中型"},
		{CategoryID: catCat.ID, Name: "狸花猫", ExpectedLifespan: "15-20年", Size: "中型"},
		{CategoryID: dogCat.ID, Name: "柯基犬", ExpectedLifespan: "12-15年", Size: "小型"},
		{CategoryID: dogCat.ID, Name: "金毛寻回犬", ExpectedLifespan: "10-12年", Size: "大型"},
		{CategoryID: dogCat.ID, Name: "拉布拉多犬", ExpectedLifespan: "10-12年", Size: "大型"},
		{CategoryID: dogCat.ID, Name: "哈士奇", ExpectedLifespan: "12-14年", Size: "中型"},
		{CategoryID: dogCat.ID, Name: "泰迪/贵宾犬", ExpectedLifespan: "12-15年", Size: "小型"},
		{CategoryID: dogCat.ID, Name: "柴犬", ExpectedLifespan: "12-15年", Size: "中型"},
		{CategoryID: dogCat.ID, Name: "边境牧羊犬", ExpectedLifespan: "12-15年", Size: "中型"},
		{CategoryID: dogCat.ID, Name: "法国斗牛犬", ExpectedLifespan: "10-12年", Size: "小型"},
		{CategoryID: rabbitCat.ID, Name: "荷兰垂耳兔", ExpectedLifespan: "8-12年", Size: "小型"},
		{CategoryID: rabbitCat.ID, Name: "安哥拉兔", ExpectedLifespan: "7-12年", Size: "大型"},
		{CategoryID: rabbitCat.ID, Name: "迷你兔", ExpectedLifespan: "8-12年", Size: "小型"},
		{CategoryID: hamsterCat.ID, Name: "叙利亚仓鼠", ExpectedLifespan: "2-3年", Size: "小型"},
		{CategoryID: hamsterCat.ID, Name: "侏儒仓鼠", ExpectedLifespan: "2-3年", Size: "小型"},
		{CategoryID: birdCat.ID, Name: "虎皮鹦鹉", ExpectedLifespan: "5-10年", Size: "小型"},
		{CategoryID: birdCat.ID, Name: "文鸟", ExpectedLifespan: "7-10年", Size: "小型"},
		{CategoryID: birdCat.ID, Name: "金丝雀", ExpectedLifespan: "10-15年", Size: "小型"},
		{CategoryID: fishCat.ID, Name: "金鱼", ExpectedLifespan: "5-10年", Size: "小型"},
		{CategoryID: fishCat.ID, Name: "锦鲤", ExpectedLifespan: "25-35年", Size: "大型"},
		{CategoryID: fishCat.ID, Name: "热带鱼", ExpectedLifespan: "3-5年", Size: "小型"},
		{CategoryID: reptileCat.ID, Name: "乌龟", ExpectedLifespan: "20-50年", Size: "小型"},
		{CategoryID: reptileCat.ID, Name: "蜥蜴", ExpectedLifespan: "5-15年", Size: "小型"},
		{CategoryID: reptileCat.ID, Name: "蛇", ExpectedLifespan: "10-20年", Size: "小型"},
	}

	for _, breed := range breeds {
		if err := db.FirstOrCreate(&breed, model.PetBreed{Name: breed.Name}).Error; err != nil {
			return err
		}
	}

	return nil
}
