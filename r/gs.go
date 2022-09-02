package r

import (
	"net/http"
	"time"
	"ts-s/mo"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GsR(a *gin.Engine) {
	gsr := a.Group("/goods")

	// TODO: 상품 정보 조회
	gsr.GET("/:id", func(c *gin.Context) {
		// 1.
		id := c.Params.ByName("id")
		Id, _ := primitive.ObjectIDFromHex(id)

		loc, err := time.LoadLocation("Asia/Seoul")
		if err != nil {
			panic(err)
		}

		// i := mo.G{}
		// c.ShouldBindJSON(&i)

		gs := mo.Gs{
			Id:        Id,
			StoreId:   primitive.NewObjectID(),
			Name:      "상품 이름",
			Price:     "10000",
			GoodsType: "TICKET",
			Category:  "카테고리",
			Quantity:  123,
			CreatedAt: time.Now().In(loc),
			UpdatedAt: time.Now().In(loc),
		}

		c.JSON(http.StatusOK, gin.H{
			"items": gs,
		})

	})

	//TODO: 플랫폼 상점의 상품 리스트 조회
	gsr.GET("/list/:storeId", func(c *gin.Context) {
		sId := c.Params.ByName("storeId")
		SId, _ := primitive.ObjectIDFromHex(sId)

		loc, err := time.LoadLocation("Asia/Seoul")
		if err != nil {
			panic(err)
		}

		gs1 := mo.Gs{
			Id:        primitive.NewObjectID(),
			StoreId:   SId,
			Name:      "상품 이름",
			Price:     "10000",
			GoodsType: "TICKET",
			Category:  "카테고리",
			Quantity:  123,
			CreatedAt: time.Now().In(loc),
			UpdatedAt: time.Now().In(loc),
		}

		gs2 := mo.Gs{
			Id:        primitive.NewObjectID(),
			StoreId:   SId,
			Name:      "상품 이름",
			Price:     "10000",
			GoodsType: "TICKET",
			Category:  "카테고리",
			Quantity:  123,
			CreatedAt: time.Now().In(loc),
			UpdatedAt: time.Now().In(loc),
		}

		gsl := []mo.Gs{gs1, gs2}

		c.JSON(http.StatusOK, gin.H{
			"items": gsl,
		})

	})

}
