package r

import (
	"net/http"
	"time"
	"ts-s/mo"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GR(a *gin.Engine) {
	gr := a.Group("/gift")

	{
		gr.GET("/")

		// 상품 구입 후, 13번 수령자(사용자, 비사용자)의 정보를 입력하여 선물을 생성 후 전달한다
		gr.POST("/", func(c *gin.Context) {
			// 라우터 -> 핸들러 -> 서비스 -> 레포
			// i := mo.G{}
			// c.ShouldBindJSON(&i)

			// 레포에서 생성된 선물 정보
			i := mo.G{
				Id:       primitive.NewObjectID(),
				Sender:   primitive.NewObjectID(),
				Receiver: primitive.NewObjectID(),
				Use:      false,
				Confirm:  false,
				GiftType: "REAL",
				Category: "선물카테고리1",
				QRcode:   "QR코드",
			}

			c.JSON(http.StatusOK, gin.H{
				"gift": i,
			})

		})

		/**
		* 일반 사용자가 준 선물을 선물함에 넣기
		 */
		gr.POST("/:giftId", func(c *gin.Context) {
			loc, err := time.LoadLocation("Asia/Seoul")
			if err != nil {
				panic(err)
			}

			p := c.Params.ByName("giftId")
			g1Id, _ := primitive.ObjectIDFromHex(p)

			su1 := mo.U{
				Id:        primitive.NewObjectID(),
				UserType:  "NORMAL",
				CreatedAt: time.Now().In(loc),
			}

			su2 := mo.U{
				Id:        primitive.NewObjectID(),
				UserType:  "PARTNER",
				CreatedAt: time.Now().In(loc),
			}

			ru := mo.U{
				Id:        primitive.NewObjectID(),
				UserType:  "NORMAL",
				CreatedAt: time.Now().In(loc),
			}

			g1 := mo.G{
				Id:        g1Id,
				Sender:    su1.Id,
				Receiver:  ru.Id,
				Use:       false,
				Confirm:   false,
				GiftType:  "REAL",
				Category:  "선물카테고리1",
				QRcode:    "QR코드",
				CreatedAt: time.Now().In(loc),
			}

			g2 := mo.G{
				Id:        primitive.NewObjectID(),
				Sender:    su2.Id,
				Receiver:  ru.Id,
				Use:       false,
				Confirm:   true,
				GiftType:  "TICKET",
				Category:  "선물카테고리2",
				QRcode:    "QR코드",
				CreatedAt: time.Now().In(loc),
			}

			gb := []mo.G{g1, g2}

			c.JSON(http.StatusOK, gin.H{
				"giftBox": gb,
			})
		})

		//TODO
		gr.POST("/:")

		//TODO! 좌표에 있는 선물을 선물함에 넣기

	}
}
