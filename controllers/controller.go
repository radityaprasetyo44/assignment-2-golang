package controllers

import (
	"assignment2/models"
	"assignment2/repositories"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type InDB struct {
	DB *gorm.DB
}

func (idb *InDB) GetOrders(c *gin.Context) {
	var (
		orders []models.OrderData
	)

	payloadOrder := repositories.QueryPayload{
		OutputOrders: orders,
	}
	err := repositories.FindOrder(&payloadOrder)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_message": "failed to find order",
		})
		return
	}

	for i, value := range payloadOrder.OutputOrders {
		payloadItem := repositories.QueryPayload{
			Query: map[string]interface{}{
				"order_id": value.OrderId,
			},
			OutputItems: []models.Item{},
		}
		errItem := repositories.FindItem(&payloadItem)
		if errItem == nil {
			payloadOrder.OutputOrders[i].Items = payloadItem.OutputItems
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"result": payloadOrder.OutputOrders,
	})
}

func (idb *InDB) CreateOrder(c *gin.Context) {
	var (
		body models.OrderRequest
	)

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	orderId := "ODR1"
	queryPayload := repositories.QueryPayload{
		OutputOrder: models.Order{},
	}
	errFind := repositories.FindLastOrder(&queryPayload)
	if errFind == nil {
		conv, _ := strconv.Atoi(queryPayload.OutputOrder.OrderId[3:])
		convString := strconv.Itoa(conv + 1)
		orderId = "ODR" + convString
	}

	payload := repositories.Command{
		DocumentOrder: models.Order{
			OrderId:      orderId,
			CustomerName: body.CustomerName,
			OrderedAt:    body.OrderedAt,
		},
	}
	errOrder := repositories.CreateOrder(&payload)
	if errOrder != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{
			"error_message": "failed to create order",
		})
		return
	}

	var items []models.Item
	for _, value := range body.Items {
		itemId := "ITM1"
		queryPayload := repositories.QueryPayload{
			OutputItem: models.Item{},
		}
		errFind := repositories.FindLastItem(&queryPayload)
		if errFind == nil {
			conv, _ := strconv.Atoi(queryPayload.OutputItem.ItemId[3:])
			convString := strconv.Itoa(conv + 1)
			itemId = "ITM" + convString
		}

		payload = repositories.Command{
			DocumentItem: models.Item{
				ItemId:      itemId,
				OrderId:     orderId,
				ItemCode:    value.ItemCode,
				Quantity:    value.Quantity,
				Description: value.Description,
			},
		}
		errItem := repositories.CreateItem(&payload)
		if errItem != nil {
			log.Println("Failed to create item")
		} else {
			items = append(items, payload.DocumentItem)
		}
	}

	payload.DocumentOrder.OrderId = orderId
	payload.DocumentOrder.CustomerName = body.CustomerName
	payload.DocumentOrder.OrderedAt = body.OrderedAt

	var data models.OrderData
	byteOrder, _ := json.Marshal(payload.DocumentOrder)
	json.Unmarshal(byteOrder, &data)
	data.Items = items

	c.JSON(http.StatusCreated, gin.H{
		"result": data,
	})
}

func (idb *InDB) DeleteOrder(c *gin.Context) {
	id := c.Param("orderId")
	queryPayload := repositories.QueryPayload{
		Param: id,
	}
	err := repositories.FindByOrderId(&queryPayload)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_message": fmt.Sprintf("failed to find order, id: %v", id),
		})
		return
	}

	payloadDelete := repositories.Command{
		Param: id,
	}
	delOrder := repositories.DeleteOrder(&payloadDelete)
	if delOrder != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{
			"error_message": fmt.Sprintf("failed to delete order, id: %v", id),
		})
		return
	}

	payloadDeleteItem := repositories.Command{
		Param: id,
	}
	delItem := repositories.DeleteItem(&payloadDeleteItem)
	if delItem != nil {
		log.Println("Gagal menghapus item")
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "success",
	})
}

func (idb *InDB) UpdateOrder(c *gin.Context) {
	var (
		body models.OrderRequest
	)

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	id := c.Param("orderId")
	queryPayload := repositories.QueryPayload{
		Param: id,
	}
	err := repositories.FindByOrderId(&queryPayload)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_message": fmt.Sprintf("failed to find order, id: %v", id),
		})
		return
	}

	payload := repositories.Command{
		Param: id,
		DocumentOrder: models.Order{
			CustomerName: body.CustomerName,
			OrderedAt:    body.OrderedAt,
		},
	}
	errOrder := repositories.UpdateOrder(&payload)
	if errOrder != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_message": fmt.Sprintf("failed to update order, id: %v", id),
		})
		return
	}

	var items []models.Item
	for _, value := range body.Items {
		payload = repositories.Command{
			Param: value.LineItemId,
			DocumentItem: models.Item{
				ItemCode:    value.ItemCode,
				Quantity:    value.Quantity,
				Description: value.Description,
			},
		}
		errItem := repositories.UpdateItem(&payload)
		if errItem != nil {
			log.Println("Failed to update item")
		}

		payload.DocumentItem.ItemCode = value.ItemCode
		payload.DocumentItem.Quantity = value.Quantity
		payload.DocumentItem.Description = value.Description
		payload.DocumentItem.ItemId = value.LineItemId
		payload.DocumentItem.OrderId = id

		items = append(items, payload.DocumentItem)
	}

	payload.DocumentOrder.OrderId = id
	payload.DocumentOrder.CustomerName = body.CustomerName
	payload.DocumentOrder.OrderedAt = body.OrderedAt

	var data models.OrderData
	byteOrder, _ := json.Marshal(payload.DocumentOrder)
	json.Unmarshal(byteOrder, &data)
	data.Items = items

	c.JSON(http.StatusOK, gin.H{
		"result": data,
	})
}
