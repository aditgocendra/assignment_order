package controllers

import (
	"assignment_order/database"
	"assignment_order/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateOrder godoc
// @Summary     Create Order Data
// @Description Note : To create order data, you need to fill in some fields, forget the [OrderID, OrderedAt, ItemID] fields because they will be filled automatically by the system
// @Tags        order
// @Accept      json
// @Produce     json
// @Param       request body     models.Orders true "create order"
// @Success     201     {object} models.Orders
// @Router      /orders [post]
func CreateOrder(ctx *gin.Context) {
	db := database.GetDB()
	var newOrder models.Orders

	// Check request body
	if err := ctx.ShouldBindJSON(&newOrder); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	// Insert data to database
	err := db.Create(&newOrder).Error

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	
	// Return response body
	ctx.JSON(http.StatusCreated, gin.H{
		"Order":newOrder,
	})

}

// GetOrder godoc
// @Summary     Read All Order Data
// @Description Reading all data with no params
// @Tags        order
// @Accept      json
// @Produce     json
// @Success     201 {object} models.Orders
// @Router      /orders [get]
func GetOrder(ctx *gin.Context) {
	db := database.GetDB()
	var orderData []models.Orders

	// Get all data to database
	err := db.Preload("Items").Find(&orderData).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	
	// Return response body
	ctx.JSON(http.StatusCreated, gin.H{
		"OrderData":orderData,
	})
}

// UpdateOrder godoc
// @Summary     Update Order Data
// @Description Note : To change the order data, you need to fill in some fields, and forget about [OrderedAt]
// @Tags        order
// @Accept      json
// @Produce     json
// @Param       order_id path     string        true "order_id"
// @Param       request  body     models.Orders true "update order"
// @Success     201      {object} models.Orders
// @Router      /orders/{order_id} [put]
func UpdateOrder(ctx *gin.Context){
	db := database.GetDB()
	var newDataOrder models.Orders

	// Check request body
	if err := ctx.ShouldBindJSON(&newDataOrder); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	id, _ := ctx.Params.Get("id")
	idOrder, _ := strconv.Atoi(id)
	newDataOrder.OrderID = uint(idOrder)
	
	// Update Order
	err := db.Model(&newDataOrder).Updates(&newDataOrder).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	
	// Update Item Order
	for _, v := range newDataOrder.Items {
		err := db.Model(&newDataOrder.Items).Updates(v).Error
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err,
			})
			return
		}
	}

	// Return response body
	ctx.JSON(http.StatusCreated, gin.H{
		"Order":newDataOrder,
	})
}

// DeleteOrder godoc
// @Summary     Delete Order Data
// @Description Delete data order with param[OrderID]
// @Tags        order
// @Accept      json
// @Produce     json
// @Param       order_id path     string true "order_id"
// @Success     201      {object} models.Orders
// @Router      /orders/{order_id} [delete]
func DeleteOrder(ctx *gin.Context){
	db := database.GetDB()

	id, _ := ctx.Params.Get("id")
	idOrder, _ := strconv.Atoi(id)

	var order models.Orders
	var items models.Items


	// Delete item data
	deletedItem := db.Model(items).Where("order_id = ?", idOrder).Delete(items)
	if deletedItem.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": deletedItem.Error,
		})
		return
	}

	if deletedItem.RowsAffected < 1 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Data Not Found",
		})
		return
	}
	
	// Delete order data
	deletedOrder :=  db.Delete(&order, idOrder)
	if deletedOrder.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": deletedOrder.Error,
		})
		return
	}

	if deletedOrder.RowsAffected < 1 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Data Not Found",
		})
		return
	}

	// Return Response Success
	ctx.JSON(http.StatusOK, "Successfuly Deleting Data")
}

