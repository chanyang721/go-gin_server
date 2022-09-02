package r

import (
	"net/http"
	"time"
	"ts-s/mo"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SR(a *gin.Engine) {
	sr := a.Group("/store")

	// TODO: 상점 정보 조회
	sr.GET("/:id", func(c *gin.Context) {
		// 1. 상점 id
		id := c.Params.ByName("id")
		Id, _ := primitive.ObjectIDFromHex(id)

		lat, _ := decimal.NewFromString("41.40338")
		lng, _ := decimal.NewFromString("2.17403")

		loc, err := time.LoadLocation("Asia/Seoul")
		if err != nil {
			panic(err)
		}

		// 2. 상점 정보 조회
		s := mo.S{
			Id:            Id,
			Name:          "이찬양 마켓",
			Phone:         "010-0000-0000",
			StoreType:     "PLATFORM",
			Category:      "상점 카테고리1",
			PlatformType:  "11번가",
			ShopName:      "11번가에서 1등한 매장",
			ShopAddress:   "한국 주소",
			CompanyNumber: "000-00-00000",
			Latitude:      lat,
			Longitude:     lng,
			Radius:        1111,
			CreatedAt:     time.Now().In(loc),
			UpdatedAt:     time.Now().In(loc),
		}

		// 3. 상점 정보 리턴
		c.JSON(http.StatusOK, gin.H{
			"item": s,
		})

	})

	// TODO: 플랫폼의 상점 리스트 조회
	sr.GET("/list", func(c *gin.Context) {
		//! TODO: 상점 카테고리, 상점 유형에 따라 필터링
		lat, _ := decimal.NewFromString("41.40338")
		lat1, _ := decimal.NewFromString("84.40338")
		lng, _ := decimal.NewFromString("2.17403")
		lng1, _ := decimal.NewFromString("123.17403")

		loc, err := time.LoadLocation("Asia/Seoul")
		if err != nil {
			panic(err)
		}

		// 2. 상점 리스트 조회
		s1 := mo.S{
			Id:            primitive.NewObjectID(),
			Name:          "이찬양 마켓",
			Phone:         "010-0000-0000",
			StoreType:     "PLATFORM",
			Category:      "상점 카테고리1",
			PlatformType:  "11번가",
			ShopName:      "11번가에서 1등한 매장",
			ShopAddress:   "한국 주소",
			CompanyNumber: "000-00-00000",
			Latitude:      lat,
			Longitude:     lng,
			Radius:        1111,
			CreatedAt:     time.Now().In(loc),
			UpdatedAt:     time.Now().In(loc),
		}

		s2 := mo.S{
			Id:            primitive.NewObjectID(),
			Name:          "이찬양 마켓 분점 2",
			Phone:         "010-0000-0000",
			StoreType:     "PLATFORM",
			Category:      "상점 카테고리432",
			PlatformType:  "G마켓",
			ShopName:      "11번가에서 2등한 매장",
			ShopAddress:   "한국 어딘가",
			CompanyNumber: "000-00-00000",
			Latitude:      lat1,
			Longitude:     lng1,
			Radius:        1111,
			CreatedAt:     time.Now().In(loc),
			UpdatedAt:     time.Now().In(loc),
		}

		// 3. 상점 리스트 리턴
		c.JSON(http.StatusOK, gin.H{
			"items": []mo.S{s1, s2},
		})
	})

}
