package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/tong-server-ja/rabbitmq"
)

func Create(c *gin.Context) {
	rawData := map[string]interface{}{}
	corrID := generateUUID()
	c.ShouldBindJSON(&rawData)

	messages, _ := json.Marshal(rawData)

	response := rabbitmq.Client(messages, "tong_service", "tong_service.tong.create_user", corrID)
	fmt.Printf(" [gate-way] Create Result: %s", string(response))

	// c.JSON(200, response)
	c.Data(200, "application/json; charset=utf-8", response)
	c.Abort()
}

func GetAll(c *gin.Context) {
	rawData := map[string]interface{}{}
	corrID := generateUUID()
	c.ShouldBindJSON(&rawData)

	messages, _ := json.Marshal(rawData)

	response := rabbitmq.Client(messages, "tong_service", "tong_service.tong.get_all_users", corrID)
	fmt.Printf(" [gate-way] Get-All Result: %s", string(response))

	c.Data(200, "application/json; charset=utf-8", response)
	c.Abort()
}

func GetByID(c *gin.Context) {
	rawData := map[string]interface{}{}
	corrID := generateUUID()
	c.ShouldBindJSON(&rawData)
	rawData["search_id"] = GetUserID(c.Param("id"))

	messages, _ := json.Marshal(rawData)

	response := rabbitmq.Client(messages, "tong_service", "tong_service.tong.get_user_by_id", corrID)
	fmt.Printf(" [gate-way] Get-By-ID : %s", string(response))

	c.Data(200, "application/json; charset=utf-8", response)
	c.Abort()
}

func Delete(c *gin.Context) {
	rawData := map[string]interface{}{}
	corrID := generateUUID()
	c.ShouldBindJSON(&rawData)
	rawData["search_id"] = GetUserID(c.Param("id"))

	messages, _ := json.Marshal(rawData)

	response := rabbitmq.Client(messages, "tong_service", "tong_service.tong.delete_user_by_id", corrID)
	fmt.Printf(" [gate-way] Get-By-ID : %s", string(response))

	c.Data(200, "application/json; charset=utf-8", response)
	c.Abort()
}

func UnDelete(c *gin.Context) {
	rawData := map[string]interface{}{}
	corrID := generateUUID()
	c.ShouldBindJSON(&rawData)
	rawData["search_id"] = GetUserID(c.Param("id"))

	messages, _ := json.Marshal(rawData)

	response := rabbitmq.Client(messages, "tong_service", "tong_service.tong.undelete_user_by_id", corrID)
	fmt.Printf(" [gate-way] Get-By-ID : %s", string(response))

	c.Data(200, "application/json; charset=utf-8", response)
	c.Abort()
}

func Updates(c *gin.Context) {
	rawData := map[string]interface{}{}
	corrID := generateUUID()
	c.ShouldBindJSON(&rawData)

	messages, _ := json.Marshal(rawData)

	response := rabbitmq.Client(messages, "tong_service", "tong_service.tong.updates_user_status", corrID)
	fmt.Print(" [gate-way] Updates-User-Status : %s", string(response))

	c.Data(200, "application/json; charset=utf-8", response)
	c.Abort()
}

func ChangePassword(c *gin.Context) {
	rawData := map[string]interface{}{}
	corrID := generateUUID()
	c.ShouldBindJSON(&rawData)

	messages, _ := json.Marshal(rawData)

	response := rabbitmq.Client(messages, "tong_service", "tong_service.tong.change_user_password", corrID)
	fmt.Print(" [gate-way] Change-User-Password : %s", string(response))

	c.Data(200, "application/json; charset=utf-8", response)
	c.Abort()
}

func generateUUID() string {
	var err error
	u1 := uuid.Must(uuid.NewV4(), err)
	return u1.String()
}

func GetUserID(urlParam string) uint32 {
	userID, _ := strconv.Atoi(urlParam)
	return uint32(userID)
}
