package service

import (
	"context"
	"fmt"
	"myapp/config"
	"myapp/graph/model"
)

func ItemCreate(ctx context.Context, input model.NewItem) (*model.Item, error) {
	db := config.ConnectGorm()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	item := model.Item{
		Name:   input.Name,
		UserID: input.UserID,
	}
	if err := db.Model(&model.Item{}).Create(&item).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &item, nil
}

func ItemUpdate(ctx context.Context, input model.UpdateItem) (*string, error) {
	db := config.ConnectGorm()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	if err := db.Model(&model.Item{}).Where("id = ?", input.ID).Updates(map[string]interface{}{
		"name":    input.Name,
		"user_id": input.UserID,
	}).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	return nil, nil
}

func ItemDelete(ctx context.Context, id int) (*string, error) {
	db := config.ConnectGorm()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	if err := db.Model(&model.Item{}).Delete(&model.Item{}, id).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	return nil, nil
}

func ItemGetByID(ctx context.Context, id int) (*model.Item, error) {
	db := config.ConnectGorm()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var item model.Item
	if err := db.Model(&model.Item{}).Take(&item, id).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &item, nil
}

func ItemGetAll(ctx context.Context) ([]*model.Item, error) {
	db := config.ConnectGorm()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var item []*model.Item
	if err := db.Model(&model.Item{}).Find(&item).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	return item, nil
}

func ItemGetByUserID(ctx context.Context, userID int) ([]*model.Item, error) {
	db := config.ConnectGorm()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var item []*model.Item
	if err := db.Model(&model.Item{}).Where("user_id = ?", userID).Find(&item).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	return item, nil
}

func ItemGetByArrayUserID(ctx context.Context, userIds []int) ([]*model.Item, error) {
	db := config.ConnectGorm()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var items []*model.Item
	if err := db.Model(&model.Item{}).Where("user_id IN (?)", userIds).Find(&items).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	return items, nil
}

func ItemDataloaderByUserID(ids []int) ([][]*model.Item, []error) {
	resp, err := ItemGetByArrayUserID(context.Background(), ids)
	if err != nil {
		fmt.Println(err)
		return nil, []error{err}
	}

	itemById := map[int][]*model.Item{}
	for _, val := range resp {
		itemById[val.UserID] = append(itemById[val.UserID], val)
	}

	items := make([][]*model.Item, len(ids))
	for i, id := range ids {
		items[i] = itemById[id]
	}

	return items, nil
}
