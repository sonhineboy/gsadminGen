package svr

func GetRepositorySub() string {

	return `package repositorys

import (
	"github.com/sonhineboy/gsadmin/service/app/models"
	"github.com/sonhineboy/gsadmin/service/app/requests"
	"github.com/sonhineboy/gsadmin/service/global"
	"gorm.io/gorm"
)

type {{.Name | Title}}Repository struct {
	db *gorm.DB
}

// New{{.Name | Title}}Repository 实例化
func New{{.Name | Title}}Repository() *{{.Name | Title}}Repository {
	return &{{.Name | Title}}Repository{
		db: global.Db,
	}
}

//FindById 根据id 查询信息
func (re *{{.Name | Title}}Repository) FindById(id int) (models.{{.Name | Title}}, error) {
	var (
		model models.{{.Name | Title}}
	)
	tx := re.db.First(&model, id)

	return model, global.GormTans(tx.Error)
}

//UpdateById 根据id 更新信息
func (re *{{.Name | Title}}Repository) UpdateById(id int, data requests.{{.Name | Title}}Request) (int64, error) {
	var (
		model models.{{.Name | Title}}
	)
	tx := re.db.Model(&model).Where("id = ?", id).Updates(models.{{.Name | Title}}{
		{{range .Fields}}{{ .Name | Title}}:	data.{{ .Name | Title}}
		{{end}}
	})
	return tx.RowsAffected, tx.Error
}

//DelByIds 根据id 删除数据
func (re *{{.Name | Title}}Repository) DelByIds(ids []int) (int64, error) {
	var (
		model models.{{.Name | Title}}
	)
	tx := re.db.Delete(&model, ids)
	return tx.RowsAffected, tx.Error
}

//Page 返回分页数据
func (re *{{.Name | Title}}Repository) Page(where map[string]interface{}, page int, pageSize int, sortField string) map[string]interface{} {
	var (
		total  int64
		data   []models.{{.Name | Title}}
		offSet int
	)
	db := global.Db.Model(&models.{{.Name | Title}}{})

	if where != nil && len(where) > 0 {
		db.Where(where)
	}
	db.Count(&total)

	if page <= 0 {
		page = 1
	}
	offSet = (page - 1) * pageSize
	db.Preload("Menus").Limit(pageSize).Order(sortField + " desc" + ",id desc").Offset(offSet)
	db.Find(&data)
	return global.Pages(page, pageSize, int(total), data)
}

//Insert 写入数据
func (re *{{.Name | Title}}Repository) Insert(data requests.{{.Name | Title}}Request) (model models.{{.Name | Title}}, err error) {

	model = models.{{.Name | Title}}{
		{{range .Fields}}{{ .Name | Title}}:	data.{{ .Name | Title}}
		{{end}}
	}

	result := re.db.Create(&model)
	err = result.Error
	return model, err
}
`
}
