package svr

func GetControllerSub() string {

	return `package {{.Pkg}}

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/sonhineboy/gsadmin/service/app/models"
	"github.com/sonhineboy/gsadmin/service/app/repositorys"
	"github.com/sonhineboy/gsadmin/service/app/requests"
	"github.com/sonhineboy/gsadmin/service/global"
	"github.com/sonhineboy/gsadmin/service/global/response"
	"strconv"
)

type {{.Name | Transform}}Controller struct{}

func (controller *{{.Name | Transform}}Controller) Index(ctx *gin.Context) {

	var (
		params global.List
		re = repositorys.New{{.Name | Transform}}Repository()
	)
	_ = ctx.ShouldBindBodyWith(&params, binding.JSON)
	response.Success(ctx, "ok", re.Page(params.Where, params.Page, params.PageSize, "created_at"))
}

func (controller *{{.Name | Transform}}Controller) Save(ctx *gin.Context) {
	var (
		data  requests.{{.Name |Transform}}Request
		err   error
		model models.{{.Name | Transform}}
		re    = repositorys.New{{.Name | Transform}}Repository()
	)
	err = ctx.ShouldBindBodyWith(&data,binding.JSON)
	if err != nil {
		response.Failed(ctx, global.GetError(err, data))
		return
	}

	model, err = re.Insert(data)
	if err != nil {
		response.Failed(ctx, err.Error())
		return
	}
	response.Success(ctx, "ok", model)
}

func (controller *{{.Name | Transform}}Controller) Edit(ctx *gin.Context) {
	var (
		err          error
		id           int
		request      requests.{{.Name |Transform}}Request
		re           = repositorys.New{{.Name | Transform}}Repository()
		rowsAffected int64
	)

	id, err = strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.Failed(ctx, err.Error())
		return
	}

	err = ctx.ShouldBindBodyWith(&request, binding.JSON)
	if err != nil {
		response.Failed(ctx, global.GetError(err, request))
		return
	}
	rowsAffected, err = re.UpdateById(id, request)
	if err != nil {
		response.Failed(ctx, err.Error())
		return
	}
	response.Success(ctx, "ok", gin.H{"rows_Affected": rowsAffected})
}

func (controller *{{.Name | Transform}}Controller) Delete(ctx *gin.Context) {

	var (
		ids          requests.Delete{{.Name |Transform}}Request
		err          error
		rowsAffected int64
		re           = repositorys.New{{.Name | Transform}}Repository()
	)

	err = ctx.ShouldBindBodyWith(&ids, binding.JSON)
	if err != nil {
		response.Failed(ctx, global.GetError(err, ids))
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

func (controller *{{.Name | Transform}}Controller) Get(ctx *gin.Context) {

	var (
		err   error
		id    int
		model models.{{.Name | Transform}}
		re    = repositorys.New{{.Name | Transform}}Repository()
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
