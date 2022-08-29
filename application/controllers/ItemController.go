package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sanzharanarbay/repository-service-pattern/application/models"
	"github.com/sanzharanarbay/repository-service-pattern/application/services"
	"net/http"
	"strconv"
)

type ItemController struct {
	itemService *services.ItemService
}

func NewItemController(itemService *services.ItemService) *ItemController {
	return &ItemController{
		itemService: itemService,
	}
}

func (h *ItemController) GetItem(c *gin.Context) {
	var item *models.Item
	var err error
	param := c.Param("id")
	idInt, err := strconv.Atoi(param)
	item, err = h.itemService.GetSingleItem(idInt)
	if err != nil {
		fmt.Printf("ERROR - %s", err)
	}
	if item != nil {
		c.JSON(http.StatusOK, item)
	} else {
		c.JSON(http.StatusNotFound, item)
	}

	return
}

func (h *ItemController) GetItemList(c *gin.Context) {
	var itemList *[]models.Item
	var err error
	itemList, err = h.itemService.GetAllItems()
	if err != nil {
		fmt.Printf("ERROR - %s", err)
	}
	c.JSON(http.StatusOK, itemList)
}

func (h *ItemController) CreateItem(c *gin.Context) {
	var itemToCreate models.Item

	if err := c.ShouldBindJSON(&itemToCreate); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "invalid json")
		return
	}
	created, err := h.itemService.InsertItem(&itemToCreate)
	if err != nil {
		fmt.Printf("ERROR - %s", err)
	}
	if created == true {
		fmt.Println("Saved Item Successfully")
	}
	c.JSON(http.StatusCreated, created)
}

func (h *ItemController) DeleteItem(c *gin.Context) {
	param := c.Param("id")
	idInt, err := strconv.Atoi(param)
	deleted, err := h.itemService.DeleteItem(idInt)
	if err != nil {
		fmt.Printf("ERROR - %s", err)
	}
	if deleted == true {
		fmt.Println("Deleted Item Successfully")
	}

	c.JSON(http.StatusOK, deleted)
}

func (h *ItemController) UpdateItem(c *gin.Context) {
	var itemToUpdate models.Item
	param := c.Param("id")
	idInt, err := strconv.Atoi(param)

	if err := c.ShouldBindJSON(&itemToUpdate); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "invalid json")
		return
	}

	updated, err := h.itemService.UpdateItem(&itemToUpdate, idInt)
	if err != nil {
		fmt.Printf("ERROR - %s", err)
	}
	if updated == true {
		fmt.Println("Updated Item Successfully")
	}

	c.JSON(http.StatusCreated, updated)
}
