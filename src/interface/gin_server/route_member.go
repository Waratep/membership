package gin_server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Waratep/membership/src/entity/member"
	"github.com/Waratep/membership/src/use_case"
	"github.com/gin-gonic/gin"
)

type CreateMemberRequest struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	Province    string `json:"province"`
	District    string `json:"district"`
	Subdistrict string `json:"subdistrict"`
	Postcode    string `json:"postcode"`
	BirthDate   string `json:"birth_date"`
}

type CreateMemberResponse struct {
	ID int64 `json:"id"`
}

func (g GinServer) addRouteMember(base gin.IRoutes) {
	base.POST("/member", g.CreateMember())
	base.PATCH("/member", g.UpdateMember())
	base.GET("/member", g.GetMembers())
	base.GET("/member/:id", g.GetMember())
	base.DELETE("/member", g.DeleteMember())

}

func (m CreateMemberRequest) toMemberEntity() member.Member {
	return member.Member{
		FirstName: m.FirstName,
		LastName:  m.LastName,
		Phone:     m.Phone,
		Email:     m.Email,
		Address:   m.Address,
	}
}

func (g GinServer) CreateMember() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var m CreateMemberRequest
		err := json.NewDecoder(ctx.Request.Body).Decode(&m)
		if err != nil {
			log.Println("Error decode json request body")

			g.errorHandle(*ctx, use_case.ErrorDataTransform)
			return
		}

		ID, err := g.useCase.CreateMember(*ctx, m.toMemberEntity())
		if err != nil {
			log.Println("Error create member", err)
			g.errorHandle(*ctx, err)

			return
		} else {
			ctx.JSON(http.StatusCreated, CreateMemberResponse{
				ID: ID,
			})
			return
		}
	}
}

func (g GinServer) UpdateMember() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Welcome Gin UpdateMember")
	}
}

func (g GinServer) GetMembers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Welcome Gin GetMembers")
	}
}

func (g GinServer) GetMember() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Welcome Gin GetMembers")
	}
}

func (g GinServer) DeleteMember() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Welcome Gin GetMembers")
	}
}
