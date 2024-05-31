package svr

func GetControllerSub() string {

	return `package system

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sonhineboy/gsadmin/service/app/models"
	"github.com/sonhineboy/gsadmin/service/app/repositorys"
	"github.com/sonhineboy/gsadmin/service/app/requests"
	"github.com/sonhineboy/gsadmin/service/global"
	"github.com/sonhineboy/gsadmin/service/global/response"
	"strconv"
)

type {{.Name | Title}}Controller struct{}

func (controller *{{.Name | Title}}Controller) Index(ctx *gin.Context) {

	var (
		params global.List
		re := repositorys.New{{.Name | Title}}Repository()
	)
	_ = ctx.ShouldBind(&params)
	response.Success(ctx, "ok", re.Page(params.Where, params.Page, params.PageSize, "created_at"))
}

func (controller *{{.Name | Title}}Controller) Save(ctx *gin.Context) {
	var (
		data  requests.{{.Name |Title}}Request
		err   error
		model models.{{.Name | Title}}
		re    = repositorys.New{{.Name | Title}}Repository()
	)
	err = ctx.ShouldBind(&data)
	if err != nil {
		var v validator.ValidationErrors
		ok := errors.As(err, &v)
		if ok {
			response.Failed(ctx, global.GetError(v, data))
		} else {
			response.Failed(ctx, err.Error())
		}
		return
	}

	model, err = re.Insert(data)
	if err != nil {
		response.Failed(ctx, err.Error())
		return
	}
	response.Success(ctx, "ok", model)
}

func (controller *{{.Name | Title}}Controller) Edit(ctx *gin.Context) {
	var (
		err          error
		id           int
		request      requests.{{.Name |Title}}Request
		re           = repositorys.New{{.Name | Title}}Repository()
		rowsAffected int64
	)

	id, err = strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.Failed(ctx, err.Error())
		return
	}

	err = ctx.ShouldBind(&request)
	if err != nil {
		var v validator.ValidationErrors
		ok := errors.As(err, &v)
		if ok {
			response.Failed(ctx, global.GetError(v, request))
		} else {
			response.Failed(ctx, err.Error())
		}
		return
	}
	rowsAffected, err = re.UpdateById(id, request)
	if err != nil {
		response.Failed(ctx, err.Error())
		return
	}
	response.Success(ctx, "ok", gin.H{"rows_Affected": rowsAffected})
}

func (controller *{{.Name | Title}}Controller) Delete(ctx *gin.Context) {

	var (
		ids          requests.Delete{{.Name |Title}}Request
		err          error
		rowsAffected int64
		re           = repositorys.New{{.Name | Title}}Repository()
	)

	err = ctx.ShouldBind(&ids)
	if err != nil {
		var v validator.ValidationErrors
		ok := errors.As(err, &v)
		if ok {
			response.Failed(ctx, global.GetError(v, ids))
		} else {
			response.Failed(ctx, err.Error())
		}
		return
	}
	rowsAffected, err = re.DelByIds(ids.Ids)
	if err != nil {
		response.Failed(ctx, err.Error())
		return
	}

	response.Success(ctx, "ok", gin.H{"rows_Affected": rowsAffected})
	return
}

func (controller *{{.Name | Title}}Controller) Get(ctx *gin.Context) {

	var (
		err   error
		id    int
		model models.{{.Name | Title}}
		re    = repositorys.New{{.Name | Title}}Repository()
	)

	id, err = strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.Failed(ctx, err.Error())
		return
	}

	model, err = re.FindById(id)
	if err != nil {
		response.Failed(ctx, err.Error())
		return
	}
	response.Success(ctx, "ok", model)
}
`
}