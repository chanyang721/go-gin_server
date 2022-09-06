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
		/*
		* @description:	선물 조회
		* @Auth:		header: Bearer Token
		* @Params:		id: uuid  [ ex: 631171e69676913c1ad7413e ]
		* @Success:		200, { data: Gift }
		* @Failure:
		 */
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

			su := mo.U{
				Id:       primitive.NewObjectID(),
				UserType: mo.UE("NORMAL"),
			}

			ru := mo.U{
				Id:       primitive.NewObjectID(),
				UserType: mo.UE("NORMAL"),
			}

			gs := mo.Gs{}

			g := mo.G{
				Id:        Id,
				Sender:    su.Id,
				Receiver:  ru.Id,
				EventId:   primitive.NewObjectID(),
				Goods:     gs,
				Use:       false,
				Confirm:   false,
				GiftType:  "REAL",
				Category:  "선물카테고리1",
				QrCode:    "QR코드",
				CreatedAt: time.Now().In(loc),
				UpdatedAt: time.Now().In(loc),
			}

			c.JSON(http.StatusOK, gin.H{
				"data": g,
			})

		})

		/*
		* @description:
		*				1. 사용자의 선물 리스트 조회
		*				2. 사용자가 이벤트를 보지 않은 선물 리스트 조회
		* @Auth:		header: Bearer Token
		* @Params:		event: bool [ ex: "0" or "1" ]
		* @Success:		200, { data: Gift }
		* @Failure:
		 */
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

			gs := mo.Gs{}

			g1 := mo.G{
				Id:        g1Id,
				Sender:    su1.Id,
				Receiver:  ru.Id,
				EventId:   primitive.NewObjectID(),
				Goods:     gs,
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
				EventId:   primitive.NewObjectID(),
				Goods:     gs,
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
				EventId:   primitive.NewObjectID(),
				Goods:     gs,
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
					"data": gb,
				})
			}

			if e == "0" {

				gb := []mo.G{g2}

				c.JSON(http.StatusOK, gin.H{
					"data": gb,
				})
			}

			if e == "2" {
				gb := []mo.G{g1, g2, g3}

				c.JSON(http.StatusOK, gin.H{
					"data": gb,
				})
			}

		})

		/*
		* @description:	좌표 주변의 선물 리스트 조회
		* @Auth:		header: Bearer Token
		* @Params:
		* 				1. lat: decimal [ ex: 41.40338 ]  (위도)
		* 				2. lng: decimal [ ex: 102.17403 ] (경도)
		* @Success:		200, { data: []Gift }
		* @Failure:
		 */
		gr.GET("/list/from/:lat/:lng", func(c *gin.Context) {
			// 1. 좌표값 params
			// 2. 좌표값 반경 Nm에 포함되는 선물 리스트 리턴
			//! 좌표, 상품 정보 유효성 검사 진행
			// validation

			loc, err := time.LoadLocation("Asia/Seoul")
			if err != nil {
				panic(err)
			}

			//! 1. 좌표값은 params로 받음
			lat := c.Params.ByName("lat")
			lng := c.Params.ByName("lng")

			//! 2. 상품 정보는 body로 받음
			// i := mo.G{}
			// c.ShouldBindJSON(&i)

			//! 3. 입력된 좌표 변환
			// 플러터 좌표 데이터 타입 https://developers.kakao.com/docs/latest/ko/kakaonavi/flutter
			// 레포에서 생성된 좌표가 포함된 선물 정보
			// lat, _ := decimal.NewFromString("37.5194529")
			Lat, _ := decimal.NewFromString(lat)
			Lng, _ := decimal.NewFromString(lng)

			su1 := mo.U{
				Id:        primitive.NewObjectID(),
				UserType:  mo.UE("NORMAL"),
				CreatedAt: time.Now().In(loc),
			}

			su2 := mo.U{
				Id:       primitive.NewObjectID(),
				UserType: mo.UE("PARTNER"),
			}

			gs := mo.Gs{}

			g1 := mo.G{
				Id:        primitive.NewObjectID(),
				Sender:    su1.Id,
				EventId:   primitive.NewObjectID(),
				Goods:     gs,
				Use:       false,
				Confirm:   false,
				GiftType:  "REAL",
				Category:  "선물카테고리1",
				QrCode:    "QR코드",
				Latitude:  Lat,
				Longitude: Lng,
			}

			g2 := mo.G{
				Id:        primitive.NewObjectID(),
				Sender:    su2.Id,
				EventId:   primitive.NewObjectID(),
				Goods:     gs,
				Use:       false,
				Confirm:   true,
				GiftType:  "TICKET",
				Category:  "선물카테고리2",
				QrCode:    "QR코드",
				Latitude:  Lat,
				Longitude: Lng,
			}

			c.JSON(http.StatusOK, gin.H{
				"data": []mo.G{g1, g2},
			})

		})

		/*
		* @description:	선물을 카톡 메시지로 직접 전달
		* @Auth:		header: Bearer Token
		* @Params:
		* 				1. id: uuid  [ ex: 631171e69676913c1ad7413e ]
		* 				2. to: phone [ ex: 010-7206-7102 ]
		* @Success:		200, { message: string, data: nil }
		* @Failure:
		 */
		gr.GET("/:id/send/:to", func(c *gin.Context) {
			// su := mo.U{
			// 	Id:       primitive.NewObjectID(),
			// 	Name:     "이찬양",
			// 	UserType: mo.UE("NORMAL"),
			// }
			//TODO: 인풋 상품 정보 스키마 생성

			//! 1. params to는 수령자의 Id
			// phone := c.Params.ByName("to")
			// Phone, _ := primitive.ObjectIDFromHex(phone)
			// ? to에서 사용자 정보(전화번호)를 입력
			// ? 카톡 메시지에는 선물 정보가 포함
			//TODO: 사용자에게 카카오톡 메시지 보내기

			// 과정 후 생성된 선물 데이터

			c.JSON(http.StatusOK, gin.H{
				"message": "카카오톡 메시지 전송 완료",
				"data":    nil,
			})

		})

		/*
		* @description:	받은 선물 수령
		* @Auth:		header: Bearer Token
		* @Body:		Gift
		* @Success:		200, { data: Gift }
		* @Failure:
		 */
		gr.PATCH("/receive", func(c *gin.Context) {
			// 선물 정보, 선물을 받기 위해 들어온 사용자
			su := mo.U{
				Id:       primitive.NewObjectID(),
				Name:     "선물 준 사람 이름",
				UserType: mo.UE("NORMAL"),
			}

			// ru := mo.U{
			// 	Id:       primitive.NewObjectID(),
			// 	Name:     "선물 받은 사람",
			// 	UserType: mo.UE("NORMAL"),
			// }

			i := mo.G{}
			c.ShouldBind(&i)

			gs := mo.Gs{}

			// 1. i.Id 으로 => 선물 조회
			// 2. 선물 받는 사람을 로그인한 사용자로 수정
			g := mo.G{
				Id:       primitive.NewObjectID(),
				Sender:   su.Id,
				EventId:  primitive.NewObjectID(),
				Goods:    gs,
				Use:      false,
				Confirm:  false,
				GiftType: "REAL",
				Category: "선물카테고리1",
				QrCode:   "QR코드",
			}

			c.JSON(http.StatusOK, gin.H{
				"data": g,
			})
		})

		/*
		* @description:	좌표 전달 선물 수정
		* @Auth:		header: Bearer Token
		* @Params:
		* 				1. lat: decimal [ ex: 41.40338 ]  (위도)
		* 				2. lng: decimal [ ex: 102.17403 ] (경도)
		* @Body:		Gift
		* @Success:		200, { data: Gift }
		* @Failure:
		 */
		gr.PATCH("/drop/:lat/:lng", func(c *gin.Context) {
			su := mo.U{
				Id:       primitive.NewObjectID(),
				UserType: mo.UE("NORMAL"),
			}
			//! 좌표, 상품 정보 유효성 검사 진행
			// validation

			//! 1. 좌표값은 params로 받음
			lat := c.Params.ByName("lat")
			lng := c.Params.ByName("lng")

			//! 2. 상품 정보는 body로 받음
			// i := mo.G{}
			// c.ShouldBindJSON(&i)

			//! 3. 입력된 좌표 변환
			// 플러터 좌표 데이터 타입 https://developers.kakao.com/docs/latest/ko/kakaonavi/flutter
			// 레포에서 생성된 좌표가 포함된 선물 정보
			// lat, _ := decimal.NewFromString("37.5194529")
			Lat, _ := decimal.NewFromString(lat)
			Lng, _ := decimal.NewFromString(lng)

			//! 4. 선물 생성
			g := mo.G{
				Id:        primitive.NewObjectID(),
				Sender:    su.Id,
				Use:       false,
				Confirm:   false,
				GiftType:  "REAL",
				Category:  "선물카테고리1",
				QrCode:    "QR코드",
				Latitude:  Lat,
				Longitude: Lng,
			}

			//! 5. 선물 리턴
			c.JSON(http.StatusOK, gin.H{
				"data": g,
			})

		})

		/*
		* @description:	상품 결재 후 선물 생성
		* @Auth:		header: Bearer Token
		* @Body:		Gift
		* @Success:		200, { data: Gift }
		* @Failure:
		 */
		gr.POST("/", func(c *gin.Context) {
			loc, err := time.LoadLocation("Asia/Seoul")
			if err != nil {
				panic(err)
			}

			su := mo.U{
				Id:       primitive.NewObjectID(),
				UserType: mo.UE("NORMAL"),
			}

			// ! 1. 선물 정보 body
			i := mo.G{}
			c.ShouldBindJSON(&i)

			g := mo.G{
				Id:        primitive.NewObjectID(),
				Sender:    su.Id,
				EventId:   i.EventId,
				Goods:     i.Goods,
				Use:       false,
				Confirm:   false,
				GiftType:  "REAL",
				Category:  "선물카테고리1",
				QrCode:    i.QrCode,
				CreatedAt: time.Now().In(loc),
				UpdatedAt: time.Now().In(loc),
			}

			c.JSON(http.StatusOK, gin.H{
				"data": g,
			})

		})

		/*
		* @description:
		*		1. 선물(실물)의 배송 정보(주소) 수정
		* 		2. 좌표에 수령자가 결정되지 않은 있는 선물 획득
		* @Auth:		header: Bearer Token
		* @Body:		Gift
		* @Success:		200, { data: Gift }
		* @Failure:
		 */
		gr.PATCH("/", func(c *gin.Context) {
			// 1. 선물 id 조회
			// 2. 선물 id가 없으면 에러
			// 3. 선물이 있으면 선물 정보 수정
			// 4. 선물 정보 리턴
			su := mo.U{
				Id:       primitive.NewObjectID(),
				UserType: mo.UE("NORMAL"),
			}

			ru := mo.U{
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
				Receiver:        ru.Id,
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
				"data": g,
			})

		})

		/*
		* @description: 선물 소비 시, use, isDeleted 업데이트
		* @Auth:		header: Bearer Token
		* @Body:		Gift
		* @Success:		200, { data: Gift }
		* @Failure:
		 */
		// TODO:
		gr.PATCH("/used/:id", func(c *gin.Context) {
			// 1. 선물을 사용 완료 요청
			// 2. use, isDeleted: 업데이트
			loc, err := time.LoadLocation("Asia/Seoul")
			if err != nil {
				panic(err)
			}

			su := mo.U{
				Id:       primitive.NewObjectID(),
				UserType: mo.UE("NORMAL"),
			}
			ru := mo.U{
				Id:       primitive.NewObjectID(),
				UserType: mo.UE("NORMAL"),
			}

			i := mo.G{}
			c.ShouldBindJSON(&i)

			g := mo.G{
				Id:        i.Id,
				Sender:    su.Id,
				Receiver:  ru.Id,
				Use:       true,
				Confirm:   false,
				GiftType:  "REAL",
				Category:  "선물카테고리1",
				QrCode:    "QR코드",
				CreatedAt: time.Now().In(loc),
				UpdatedAt: time.Now().In(loc),
				IsDeleted: true,
			}

			c.JSON(http.StatusOK, gin.H{
				"data": g,
			})

		})

	}
}
