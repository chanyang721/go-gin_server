package r

import (
	"net/http"
	"time"
	"ts-s/mo"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// 라우터 -> 핸들러 -> 유효성검사 -> 서비스 -> 레포
func GR(a *gin.Engine) {
	gr := a.Group("/gift")

	{
		//? TODO: 선물 조회
		gr.GET("/:id", func(c *gin.Context) {
			// 1. 입력된 아이디의 선물 조회
			// 2. 조회되면 리턴
			// 3. 정보가 없으면 에러 리턴
			loc, err := time.LoadLocation("Asia/Seoul")
			if err != nil {
				panic(err)
			}

			id := c.Params.ByName("id")
			Id, _ := primitive.ObjectIDFromHex(id)

			u := mo.U{
				Id:       primitive.NewObjectID(),
				UserType: mo.UE("NORMAL"),
			}

			g := mo.G{
				Id:        Id,
				Sender:    primitive.NewObjectID(),
				Receiver:  u.Id,
				Use:       false,
				Confirm:   false,
				GiftType:  "REAL",
				Category:  "선물카테고리1",
				QrCode:    "QR코드",
				CreatedAt: time.Now().In(loc),
				UpdatedAt: time.Now().In(loc),
			}

			c.JSON(http.StatusOK, gin.H{
				"item": g,
			})

		})

		//? TODO: 사용자의 선물 리스트 조회
		//? TODO: 1. 사용자가 이벤트를 보지 않은 선물 리스트 조회
		gr.GET("/list/:event", func(c *gin.Context) {
			// receiver와 사용자의 id가 같은 선물 리스트 리턴
			e := c.Params.ByName("event")

			loc, err := time.LoadLocation("Asia/Seoul")
			if err != nil {
				panic(err)
			}

			p := c.Params.ByName("giftId")
			g1Id, _ := primitive.ObjectIDFromHex(p)

			su1 := mo.U{
				Id:        primitive.NewObjectID(),
				UserType:  mo.UE("NORMAL"),
				CreatedAt: time.Now().In(loc),
			}

			su2 := mo.U{
				Id:       primitive.NewObjectID(),
				UserType: mo.UE("PARTNER"),
			}

			ru := mo.U{
				Id:       primitive.NewObjectID(),
				UserType: mo.UE("NORMAL"),
			}

			g1 := mo.G{
				Id:        g1Id,
				Sender:    su1.Id,
				Receiver:  ru.Id,
				Use:       false,
				Confirm:   true,
				GiftType:  "REAL",
				Category:  "선물카테고리1",
				QrCode:    "QR코드",
				CreatedAt: time.Now().In(loc),
				UpdatedAt: time.Now().In(loc),
			}

			g2 := mo.G{
				Id:        primitive.NewObjectID(),
				Sender:    su2.Id,
				Receiver:  ru.Id,
				Use:       false,
				Confirm:   false,
				GiftType:  "TICKET",
				Category:  "선물카테고리2",
				QrCode:    "QR코드",
				CreatedAt: time.Now().In(loc),
				UpdatedAt: time.Now().In(loc),
			}

			g3 := mo.G{
				Id:        primitive.NewObjectID(),
				Sender:    su2.Id,
				Receiver:  ru.Id,
				Use:       false,
				Confirm:   true,
				GiftType:  "TICKET",
				Category:  "선물카테고리2",
				QrCode:    "QR코드",
				CreatedAt: time.Now().In(loc),
				UpdatedAt: time.Now().In(loc),
			}

			if e == "1" {
				gb := []mo.G{g1}

				c.JSON(http.StatusOK, gin.H{
					"items": gb,
				})
			}

			if e == "0" {

				gb := []mo.G{g2}

				c.JSON(http.StatusOK, gin.H{
					"items": gb,
				})
			}

			if e == "2" {
				gb := []mo.G{g1, g2, g3}

				c.JSON(http.StatusOK, gin.H{
					"items": gb,
				})
			}

		})

		//? TODO: 좌표 주변의 선물 리스트 조회
		gr.GET("/list/from/:x/:y", func(c *gin.Context) {
			// 1. 좌표값 params
			// 2. 좌표값 반경 Nm에 포함되는 선물 리스트 리턴
			//! 좌표, 상품 정보 유효성 검사 진행
			// validation

			loc, err := time.LoadLocation("Asia/Seoul")
			if err != nil {
				panic(err)
			}

			//! 1. 좌표값은 params로 받음
			x := c.Params.ByName("x")
			y := c.Params.ByName("y")

			//! 2. 상품 정보는 body로 받음
			// i := mo.G{}
			// c.ShouldBindJSON(&i)

			//! 3. 입력된 좌표 변환
			// 플러터 좌표 데이터 타입 https://developers.kakao.com/docs/latest/ko/kakaonavi/flutter
			// 레포에서 생성된 좌표가 포함된 선물 정보
			// lat, _ := decimal.NewFromString("37.5194529")
			lat, _ := decimal.NewFromString(x)
			lng, _ := decimal.NewFromString(y)

			su1 := mo.U{
				Id:        primitive.NewObjectID(),
				UserType:  mo.UE("NORMAL"),
				CreatedAt: time.Now().In(loc),
			}

			su2 := mo.U{
				Id:       primitive.NewObjectID(),
				UserType: mo.UE("PARTNER"),
			}

			ru := mo.U{
				Id:       primitive.NewObjectID(),
				UserType: mo.UE("NORMAL"),
			}

			g1 := mo.G{
				Id:        primitive.NewObjectID(),
				Sender:    su1.Id,
				Receiver:  ru.Id,
				Use:       false,
				Confirm:   false,
				GiftType:  "REAL",
				Category:  "선물카테고리1",
				QrCode:    "QR코드",
				Latitude:  lat,
				Longitude: lng,
			}

			g2 := mo.G{
				Id:        primitive.NewObjectID(),
				Sender:    su2.Id,
				Receiver:  ru.Id,
				Use:       false,
				Confirm:   true,
				GiftType:  "TICKET",
				Category:  "선물카테고리2",
				QrCode:    "QR코드",
				Latitude:  lat,
				Longitude: lng,
			}

			gb := []mo.G{g1, g2}

			c.JSON(http.StatusOK, gin.H{
				"items": gb,
			})

		})

		//? TODO: 직접 전달 선물 생성
		gr.POST("/direct/:to", func(c *gin.Context) {
			su := mo.U{
				Id:       primitive.NewObjectID(),
				Name:     "이찬양",
				UserType: mo.UE("NORMAL"),
			}
			//TODO: 인풋 상품 정보 스키마 생성

			//! 1. params to는 수령자의 Id
			id := c.Params.ByName("to")
			Id, _ := primitive.ObjectIDFromHex(id)

			//! 2. 상품 정보는 body에 있음
			// i := mo.G{}
			// c.ShouldBindJSON(&i)

			//! 4. to가 회원이면 Receiver에 할당 후 선물 생성

			//! 5. to가 비회원이면 Receiver 없이 선물 생성

			// 과정 후 생성된 선물 데이터
			i := mo.G{
				Id:       primitive.NewObjectID(),
				Sender:   su.Id,
				Receiver: Id,
				Use:      false,
				Confirm:  false,
				GiftType: "REAL",
				Category: "선물카테고리1",
				QrCode:   "QR코드",
			}

			c.JSON(http.StatusOK, gin.H{
				"item":   i,
				"sender": su,
			})

		})

		//? TODO: 좌표 전달 선물 생성
		gr.POST("/indirect/:x/:y", func(c *gin.Context) {
			su := mo.U{
				Id:       primitive.NewObjectID(),
				UserType: mo.UE("NORMAL"),
			}
			//! 좌표, 상품 정보 유효성 검사 진행
			// validation

			//! 1. 좌표값은 params로 받음
			x := c.Params.ByName("x")
			y := c.Params.ByName("y")

			//! 2. 상품 정보는 body로 받음
			// i := mo.G{}
			// c.ShouldBindJSON(&i)

			//! 3. 입력된 좌표 변환
			// 플러터 좌표 데이터 타입 https://developers.kakao.com/docs/latest/ko/kakaonavi/flutter
			// 레포에서 생성된 좌표가 포함된 선물 정보
			// lat, _ := decimal.NewFromString("37.5194529")
			lat, _ := decimal.NewFromString(x)
			lng, _ := decimal.NewFromString(y)

			//! 4. 선물 생성
			i := mo.G{
				Id:        primitive.NewObjectID(),
				Sender:    su.Id,
				Receiver:  primitive.NewObjectID(),
				Use:       false,
				Confirm:   false,
				GiftType:  "REAL",
				Category:  "선물카테고리1",
				QrCode:    "QR코드",
				Latitude:  lat,
				Longitude: lng,
			}

			//! 5. 선물 리턴
			c.JSON(http.StatusOK, gin.H{
				"item": i,
			})

		})

		// FIXME: 일반 사용자가 준 선물 생성 - 삭제
		gr.POST("/", func(c *gin.Context) {
			loc, err := time.LoadLocation("Asia/Seoul")
			if err != nil {
				panic(err)
			}
			// ! 1. 선물 정보 body
			// i := mo.G{}
			// c.ShouldBindJSON(&i)

			// !  2.
			p := c.Params.ByName("giftId")
			g1Id, _ := primitive.ObjectIDFromHex(p)

			su1 := mo.U{
				Id:        primitive.NewObjectID(),
				UserType:  mo.UE("NORMAL"),
				CreatedAt: time.Now().In(loc),
			}

			su2 := mo.U{
				Id:       primitive.NewObjectID(),
				UserType: mo.UE("PARTNER"),
			}

			ru := mo.U{
				Id:       primitive.NewObjectID(),
				UserType: mo.UE("NORMAL"),
			}

			g1 := mo.G{
				Id:       g1Id,
				Sender:   su1.Id,
				Receiver: ru.Id,
				Use:      false,
				Confirm:  false,
				GiftType: "REAL",
				Category: "선물카테고리1",
				QrCode:   "QR코드",
			}

			g2 := mo.G{
				Id:       primitive.NewObjectID(),
				Sender:   su2.Id,
				Receiver: ru.Id,
				Use:      false,
				Confirm:  true,
				GiftType: "TICKET",
				Category: "선물카테고리2",
				QrCode:   "QR코드",
			}

			gb := []mo.G{g1, g2}

			c.JSON(http.StatusOK, gin.H{
				"giftBox": gb,
			})

		})

		//? TODO: 선물 정보 수정
		//? TODO: 1. 선물(실물)의 배송 정보(주소) 수정
		//? TODO: 2. 좌표에 수령자가 결정되지 않은 있는 선물 획득
		gr.PATCH("/", func(c *gin.Context) {
			// 1. 선물 id 조회
			// 2. 선물 id가 없으면 에러
			// 3. 선물이 있으면 선물 정보 수정
			// 4. 선물 정보 리턴
			su := mo.U{
				Id:       primitive.NewObjectID(),
				UserType: mo.UE("NORMAL"),
			}

			id := c.Params.ByName("id")
			Id, _ := primitive.ObjectIDFromHex(id)

			i := mo.G{}
			c.ShouldBindJSON(&i)

			g := mo.G{
				Id:              Id,
				Sender:          su.Id,
				Receiver:        primitive.NewObjectID(),
				Use:             i.Use,
				Confirm:         i.Confirm,
				GiftType:        i.GiftType,
				Category:        i.Category,
				QrCode:          i.QrCode,
				Latitude:        i.Latitude,
				Longitude:       i.Longitude,
				ReservationTime: i.ReservationTime,
			}

			c.JSON(http.StatusOK, gin.H{
				"item": g,
			})

		})

		// TODO: 선물 소비 시, 삭제 업데이트
		gr.PUT("/delete", func(c *gin.Context) {
			// 1. 선물을 사용 완료 요청
			// 2. deletedAt: 업데이트
			loc, err := time.LoadLocation("Asia/Seoul")
			if err != nil {
				panic(err)
			}

			su := mo.U{
				Id:       primitive.NewObjectID(),
				UserType: mo.UE("NORMAL"),
			}

			i := mo.G{}
			c.ShouldBindJSON(&i)

			g := mo.G{
				Id:        i.Id,
				Sender:    su.Id,
				Receiver:  primitive.NewObjectID(),
				Use:       false,
				Confirm:   false,
				GiftType:  "REAL",
				Category:  "선물카테고리1",
				QrCode:    "QR코드",
				CreatedAt: time.Now().In(loc),
				UpdatedAt: time.Now().In(loc),
				IsDeleted: true,
			}

			c.JSON(http.StatusOK, gin.H{
				"item": g,
			})

		})

	}
}
