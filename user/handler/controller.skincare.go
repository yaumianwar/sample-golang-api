package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yaumianwar/sample-golang-api/shared"
	"github.com/yaumianwar/sample-golang-api/user"
	"github.com/yaumianwar/sample-golang-api/user/form"
	"net/http"
	"strconv"
)

func GetAllSkincareByUser(svc *user.Service) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		// convert id param to uint
		userID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
		if err != nil {
			resp := shared.NewErrorResponse(err)
			ctx.JSON(resp.Status, resp)
			return
		}

		// call service to return single user data by id
		mUser, err := svc.GetSingleUser(userID)
		if err != nil {
			resp := shared.NewErrorResponse(err)
			ctx.JSON(resp.Status, resp)
			return
		}

		// call service to return all skincare data by user id
		mSkincares, err := svc.GetAllSkincareByUser(mUser)
		if err != nil {
			resp := shared.NewErrorResponse(err)
			ctx.JSON(resp.Status, resp)
			return
		}

		// return application/json response data
		ctx.JSON(http.StatusOK, gin.H{
			"data":   mSkincares,
			"status": gin.H{"message": "data fetched"},
		})
	}
}

func GetSingleSkincare(svc *user.Service) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		// convert id param to uint
		userID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
		if err != nil {
			resp := shared.NewErrorResponse(err)
			ctx.JSON(resp.Status, resp)
			return
		}

		skincareID, err := strconv.ParseUint(ctx.Param("skincareId"), 10, 64)
		if err != nil {
			resp := shared.NewErrorResponse(err)
			ctx.JSON(resp.Status, resp)
			return
		}

		// call service to return single user data by id
		mUser, err := svc.GetSingleUser(userID)
		if err != nil {
			resp := shared.NewErrorResponse(err)
			ctx.JSON(resp.Status, resp)
			return
		}

		// call service to return skincare data by user id and id
		mSkincare, err := svc.GetSingleSkincare(mUser.ID, skincareID)
		if err != nil {
			resp := shared.NewErrorResponse(err)
			ctx.JSON(resp.Status, resp)
			return
		}

		// return application/json response data
		ctx.JSON(http.StatusOK, gin.H{
			"data":   mSkincare,
			"status": gin.H{"message": "data fetched"},
		})
	}
}

func CreateSkincare(svc *user.Service) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		// convert id param to uint
		userID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
		if err != nil {
			resp := shared.NewErrorResponse(err)
			ctx.JSON(resp.Status, resp)
			return
		}

		// call service to return single user data by id
		mUser, err := svc.GetSingleUser(userID)
		if err != nil {
			resp := shared.NewErrorResponse(err)
			ctx.JSON(resp.Status, resp)
			return
		}

		// extract form data
		var skincareForm struct {
			Data form.SkincareForm `json:"data"`
		}
		if err := ctx.Bind(&skincareForm); err != nil {
			resp := shared.NewErrorResponse(err)
			ctx.JSON(resp.Status, resp)
			return
		}

		// call service to create skincare data
		mSkincare, err := svc.CreateSkincare(mUser, skincareForm.Data)
		if err != nil {
			resp := shared.NewErrorResponse(err)
			ctx.JSON(resp.Status, resp)
			return
		}

		// return application/json response data
		ctx.JSON(http.StatusCreated, gin.H{
			"data":   mSkincare,
			"status": gin.H{"message": "data created"},
		})
	}
}

func UpdateSkincare(svc *user.Service) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		// convert id param to uint
		userID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
		if err != nil {
			resp := shared.NewErrorResponse(err)
			ctx.JSON(resp.Status, resp)
			return
		}

		skincareID, err := strconv.ParseUint(ctx.Param("skincareId"), 10, 64)
		if err != nil {
			resp := shared.NewErrorResponse(err)
			ctx.JSON(resp.Status, resp)
			return
		}

		// call service to return single user data by id
		mUser, err := svc.GetSingleUser(userID)
		if err != nil {
			resp := shared.NewErrorResponse(err)
			ctx.JSON(resp.Status, resp)
			return
		}

		// call service to return skincare data by user id and id
		mSkincare, err := svc.GetSingleSkincare(mUser.ID, skincareID)
		if err != nil {
			resp := shared.NewErrorResponse(err)
			ctx.JSON(resp.Status, resp)
			return
		}

		// extract form data
		var skincareForm struct {
			Data form.SkincareForm `json:"data"`
		}
		if err := ctx.Bind(&skincareForm); err != nil {
			resp := shared.NewErrorResponse(err)
			ctx.JSON(resp.Status, resp)
			return
		}

		// call service to create skincare data
		mSkincare, err = svc.UpdateSkincare(mSkincare, skincareForm.Data)
		if err != nil {
			resp := shared.NewErrorResponse(err)
			ctx.JSON(resp.Status, resp)
			return
		}

		// return application/json response data
		ctx.JSON(http.StatusOK, gin.H{
			"data":   mSkincare,
			"status": gin.H{"message": "data updated"},
		})
	}
}

func DeleteSkincare(svc *user.Service) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		// convert id param to uint
		userID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
		if err != nil {
			resp := shared.NewErrorResponse(err)
			ctx.JSON(resp.Status, resp)
			return
		}

		skincareID, err := strconv.ParseUint(ctx.Param("skincareId"), 10, 64)
		if err != nil {
			resp := shared.NewErrorResponse(err)
			ctx.JSON(resp.Status, resp)
			return
		}

		// call service to return single user data by id
		mUser, err := svc.GetSingleUser(userID)
		if err != nil {
			resp := shared.NewErrorResponse(err)
			ctx.JSON(resp.Status, resp)
			return
		}

		// call service to return skincare data by user id and id
		mSkincare, err := svc.GetSingleSkincare(mUser.ID, skincareID)
		if err != nil {
			resp := shared.NewErrorResponse(err)
			ctx.JSON(resp.Status, resp)
			return
		}

		// call service to delete skincare data
		err = svc.DeleteSkincare(mSkincare)
		if err != nil {
			resp := shared.NewErrorResponse(err)
			ctx.JSON(resp.Status, resp)
			return
		}

		// return application/json response data
		ctx.JSON(http.StatusOK, gin.H{
			"status": gin.H{"message": "ok"},
		})
	}
}
