package api

import (
	"github.com/gin-gonic/gin"
	log "github.com/sillyhatxu/microlog"
	"net/http"
	"word-api/config"
	"word-api/dto"
	"word-api/response"
	"word-api/service"
)

func InitialAPI() {
	log.Info("---------- initial api start ----------")
	router := gin.Default()
	stockRouterGroup := router.Group("/stock-internal-api/stocks")
	{
		stockRouterGroup.POST("/add", addStock)
		stockRouterGroup.PUT("/consume", consume)
		stockRouterGroup.GET("/:id", getStock)
		stockRouterGroup.DELETE("/", deleteStock)
	}
	router.Run(config.Conf.Http.Listen)
}

func addStock(context *gin.Context) {
	var requestBody dto.Products
	err := context.ShouldBindJSON(&requestBody)
	if err != nil {
		context.JSON(http.StatusOK, response.ErrorParamsValidate(nil, err.Error()))
		return
	}
	serviceErr := service.AddSKUs(requestBody.ProductArray)
	if serviceErr != nil {
		context.JSON(http.StatusOK, response.Error(nil, serviceErr.Error()))
		return
	}
	context.JSON(http.StatusOK, response.Success(nil))
}

func consume(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"code":    http.StatusOK,
			"data":    "This is success.",
			"message": "Success",
		},
	)
}

func updateStock(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"code":    http.StatusOK,
			"data":    "This is success.",
			"message": "Success",
		},
	)
}

func deleteStock(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"code":    http.StatusOK,
			"data":    "This is success.",
			"message": "Success",
		},
	)
}

func getStock(c *gin.Context) {
	//c.JSON(
	//	http.StatusOK,
	//	gin.H{
	//		"code":  http.StatusOK,
	//		"data": "{'id':123,'name' : 'test'}",
	//		"message": "Success",
	//	},
	//)

	var msg struct {
		Name    string `json:"user"`
		Message string
		Number  int
	}
	msg.Name = "Lena"
	msg.Message = "hey"
	msg.Number = 123
	// Note that msg.Name becomes "user" in the JSON
	// Will output  :   {"user": "Lena", "Message": "hey", "Number": 123}
	c.JSON(http.StatusOK, msg)
}

func queryStock(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"code":    http.StatusOK,
			"data":    "This is success.",
			"message": "Success",
		},
	)
}
