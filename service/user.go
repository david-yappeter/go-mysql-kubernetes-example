package service

import (
	"context"
	"fmt"
	"myapp/config"
	"myapp/graph/model"
)

func UserCreate(ctx context.Context, input model.NewUser) (*model.User, error) {
	db := config.ConnectGorm()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	user := model.User{
		Name: input.Name,
	}
	if err := db.Model(&model.User{}).Create(&user).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &user, nil
}

func UserUpdate(ctx context.Context, input model.UpdateUser) (*string, error) {
	db := config.ConnectGorm()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	if err := db.Model(&model.User{}).Where("id = ?", input.ID).Updates(map[string]interface{}{
		"name": input.Name,
	}).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	return nil, nil
}

func UserDelete(ctx context.Context, id int) (*string, error) {
	db := config.ConnectGorm()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	if err := db.Model(&model.User{}).Delete(&model.User{}, id).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	return nil, nil
}

func UserGetByID(ctx context.Context, id int) (*model.User, error) {
	db := config.ConnectGorm()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var user model.User
	if err := db.Model(&model.User{}).Take(&user, id).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &user, nil
}

func UserGetAll(ctx context.Context) ([]*model.User, error) {
	db := config.ConnectGorm()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	var users []*model.User
	if err := db.Model(&model.User{}).Find(&users).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	return users, nil
}
