package todocontroller

import (
	db "FP1-Hacktiv8/infra/db"
	"FP1-Hacktiv8/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary Get all todos
// @Description Get a list of todos
// @Produce  json
// @Success 200 {object} []models.Todos
// @Router /todos [get]
func Index(c *gin.Context) {
	var todos []models.Todos

	db.DB.Find(&todos)
	c.JSON(http.StatusOK, gin.H{"todos": todos})
}

// @Summary Get a todos by ID
// @Description Get a todo by ID
// @Produce  json
// @Param id path string true "Todo ID"
// @Success 200 {object} models.Todos
// @Router /todos/{id} [get]
func Show(c *gin.Context) {
	var todos models.Todos

	id := c.Param("id")

	if err := db.DB.First(&todos, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "todolist tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"todos": todos})

}

// @Summary Create a todos
// @Description Create a new todos
// @Produce  json
// @Accept json
// @Param input body models.Todos true "Todo input"
// @Success 200 {object} models.Todos
// @Router /todos [post]
func Create(c *gin.Context) {
	var todos models.Todos

	if err := c.ShouldBindJSON(&todos); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	db.DB.Create(&todos)
	c.JSON(http.StatusOK, gin.H{
		"message":  "todolist berhasil dimasukkan",
		"todolist": todos,
	})

}

// @Summary Update a todos by ID
// @Description Update a todos by ID
// @Produce  json
// @Accept json
// @Param id path string true "Todo ID"
// @Param input body models.Todos true "Todo input"
// @Success 200 {object} models.Todos
// @Router /todos/{id} [put]
func Update(c *gin.Context) {
	var todos models.Todos

	id := c.Param("id")

	if err := c.ShouldBindJSON(&todos); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if db.DB.Model(&todos).Where("id = ?", id).Updates(&todos).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Update Gagal"})
	}

	c.JSON(http.StatusOK, gin.H{"message": "todolist berhasil diperbarui"})
}

// @Summary Delete a todos by ID
// @Description Delete a todos by ID
// @Produce  json
// @Param id path string true "Todo ID"
// @Success 200 {object} models.Todos
// @Router /todos/{id} [delete]
func Delete(c *gin.Context) {
	var todos models.Todos

	id := c.Param("id")

	if err := db.DB.Delete(&todos, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("todolist %s berhasil dihapus", id)})
}
