package repositories

import (
	"assignment2/configs"
	"assignment2/models"
	"errors"
)

type Command struct {
	Param         string
	Table         string
	Model         models.OrderData
	DocumentOrder models.Order
	DocumentItem  models.Item
}

func CreateOrder(payload *Command) error {
	db := configs.DBInit()
	err := db.Create(&payload.DocumentOrder).Error
	if err != nil {
		return errors.New("failed to create")
	}

	return nil
}

func CreateItem(payload *Command) error {
	db := configs.DBInit()
	err := db.Create(&payload.DocumentItem).Error
	if err != nil {
		return errors.New("failed to create")
	}

	return nil
}

func DeleteOrder(payload *Command) error {
	db := configs.DBInit()
	err := db.Unscoped().Where("order_id = ?", payload.Param).Delete(&payload.DocumentOrder).Error
	if err != nil {
		return errors.New("failed to delete")
	}

	return nil
}

func DeleteItem(payload *Command) error {
	db := configs.DBInit()
	err := db.Unscoped().Where("order_id = ?", payload.Param).Delete(&payload.DocumentItem).Error
	if err != nil {
		return errors.New("failed to delete")
	}

	return nil
}

func UpdateOrder(payload *Command) error {
	db := configs.DBInit()
	err := db.Table("orders").Where("order_id = ?", payload.Param).Updates(&payload.DocumentOrder).Error
	if err != nil {
		return errors.New("failed to update")
	}

	return nil
}

func UpdateItem(payload *Command) error {
	db := configs.DBInit()
	err := db.Table("items").Where("item_id = ?", payload.Param).Updates(&payload.DocumentItem).Error
	if err != nil {
		return errors.New("failed to update")
	}

	return nil
}
