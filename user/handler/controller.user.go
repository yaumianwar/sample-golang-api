package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yaumianwar/sample-golang-api/shared"
	"github.com/yaumianwar/sample-golang-api/user"
	"github.com/yaumianwar/sample-golang-api/user/form"
	"net/http"
	"strconv"
)

func GetAllUsers(svc *user.Service) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		// call service to return all users data
		mUsers, err := svc.GetAllUsers()
		if err != nil {
			resp := shared.NewErrorResponse(err)
			ctx.JSON(resp.Status, resp)
			return
		}

		// return application/json response data
		ctx.JSON(http.StatusOK, gin.H{
			"data":   mUsers,
			"status": gin.H{"message": "data fetched"},
		})
	}
}

func GetSingleUser(svc *user.Service) func(ctx *gin.Context) {
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

		// return application/json response data
		ctx.JSON(http.StatusOK, gin.H{
			"data":   mUser,
			"status": gin.H{"message": "data fetched"},
		})
	}
}

func CreateUser(svc *user.Service) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		// extract form data
		var userForm struct {
			Data form.UserForm `json:"data"`
		}
		if err := ctx.Bind(&userForm); err != nil {
			resp := shared.NewErrorResponse(err)
			ctx.JSON(resp.Status, resp)
			return
		}

		// call service to create user data
		mUser, err := svc.CreateUser(userForm.Data)
		if err != nil {
			resp := shared.NewErrorResponse(err)
			ctx.JSON(resp.Status, resp)
			return
		}

		// return application/json response data
		ctx.JSON(http.StatusCreated, gin.H{
			"data":   mUser,
			"status": gin.H{"message": "data created"},
		})
	}
}

func UpdateUser(svc *user.Service) func(ctx *gin.Context) {
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
		var userForm struct {
			Data form.UserForm `json:"data"`
		}
		if err := ctx.Bind(&userForm); err != nil {
			resp := shared.NewErrorResponse(err)
			ctx.JSON(resp.Status, resp)
			return
		}

		// call service to update user data
		mUser, err = svc.UpdateUser(mUser, userForm.Data)
		if err != nil {
			resp := shared.NewErrorResponse(err)
			ctx.JSON(resp.Status, resp)
			return
		}

		// return application/json response data
		ctx.JSON(http.StatusOK, gin.H{
			"data":   mUser,
			"status": gin.H{"message": "data updated"},
		})
	}
}

func DeleteUser(svc *user.Service) func(ctx *gin.Context) {
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

		// call service to delete user data
		err = svc.DeleteUser(mUser)
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
