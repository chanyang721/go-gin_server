package r

import (
	"net/http"
	"ts-s/mo"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GR(a *gin.Engine) {
	gr := a.Group("/gift")

	gr.GET("/")
	{

		/**
		* 일반 사용자가 준 선물을 선물함에 넣기
		 */
		gr.POST("/:giftId", func(c *gin.Context) {
			su1 := mo.U{
				Id:       primitive.NewObjectID(),
				UserType: "NORMAL",
			}

			su2 := mo.U{
				Id:       primitive.NewObjectID(),
				UserType: "PARTNER",
			}

			ru := mo.U{
				Id:       primitive.NewObjectID(),
				UserType: "NORMAL",
			}

			g1 := mo.G{
				Id:        primitive.NewObjectID(),
				Sender:    su1.Id,
				Recipient: ru.Id,
				Use:       false,
				GiftType:  "REAL",
				Category:  "선물카테고리1",
				QRcode:    "QR코드",
			}

			g2 := mo.G{
				Id:        primitive.NewObjectID(),
				Sender:    su2.Id,
				Recipient: ru.Id,
				Use:       false,
				GiftType:  "TICKET",
				Category:  "선물카테고리2",
				QRcode:    "QR코드",
			}

			gb := []mo.G{g1, g2}

			c.JSON(http.StatusOK, gin.H{
				"giftBox": gb,
			})
		})

		//TODO 비사용자가 카톡으로 받은 선물을 선물함에 넣기

		//TODO! 좌표에 있는 선물을 선물함에 넣기

	}
}
