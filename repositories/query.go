package repositories

import (
	"assignment2/configs"
	"assignment2/models"
	"errors"
)

type QueryPayload struct {
	Param        string
	Query        interface{}
	OutputOrder  models.Order
	OutputItem   models.Item
	OutputOrders []models.OrderData
	OutputItems  []models.Item
}

func FindLastOrder(payload *QueryPayload) error {
	db := configs.DBInit()
	result := db.Order("order_id desc").First(&payload.OutputOrder)
	if result.Error != nil || result.RowsAffected == 0 {
		return errors.New("data not found")
	}

	return nil
}

func FindLastItem(payload *QueryPayload) error {
	db := configs.DBInit()
	result := db.Order("item_id desc").First(&payload.OutputItem)
	if result.Error != nil || result.RowsAffected == 0 {
		return errors.New("data not found")
	}

	return nil
}

func FindByOrderId(payload *QueryPayload) error {
	db := configs.DBInit()
	result := db.Where("order_id = ?", payload.Param).First(&payload.OutputOrder)
	if result.Error != nil || result.RowsAffected == 0 {
		return errors.New("data not found")
	}

	return nil
}

func FindOrder(payload *QueryPayload) error {
	db := configs.DBInit()
	result := db.Table("orders").Order("id desc").Find(&payload.OutputOrders)
	if result.Error != nil || result.RowsAffected == 0 {
		return errors.New("data not found")
	}

	return nil
}

func FindItem(payload *QueryPayload) error {
	db := configs.DBInit()
	result := db.Table("items").Where(payload.Query).Find(&payload.OutputItems)
	if result.Error != nil || result.RowsAffected == 0 {
		return errors.New("data not found")
	}

	return nil
}
