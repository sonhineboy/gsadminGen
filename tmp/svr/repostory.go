package svr

func GetRepositorySub() string {

	return `package repositorys

import (
	"github.com/sonhineboy/gsadmin/service/app/models"
	"github.com/sonhineboy/gsadmin/service/app/requests"
	"github.com/sonhineboy/gsadmin/service/global"
	"gorm.io/gorm"
)

type {{.Name | Transform}}Repository struct {
	db *gorm.DB
}

// New{{.Name | Transform}}Repository 实例化
func New{{.Name | Transform}}Repository() *{{.Name | Transform}}Repository {
	return &{{.Name | Transform}}Repository{
		db: global.Db,
	}
}

//FindById 根据id 查询信息
func (re *{{.Name | Transform}}Repository) FindById(id int) (models.{{.Name | Transform}}, error) {
	var (
		model models.{{.Name | Transform}}
	)
	tx := re.db.First(&model, id)

	return model, global.GormTans(tx.Error)
}

//UpdateById 根据id 更新信息
func (re *{{.Name | Transform}}Repository) UpdateById(id int, data requests.{{.Name | Transform}}Request) (int64, error) {
	var (
		model models.{{.Name | Transform}}
	)
	tx := re.db.Model(&model).Where("id = ?", id).Updates(models.{{.Name | Transform}}{
		{{range .Fields}}
		{{ .Name | Transform}}:	data.{{ .Name | Transform}},
		{{end}}
	})
	return tx.RowsAffected, tx.Error
}

//DelByIds 根据id 删除数据
func (re *{{.Name | Transform}}Repository) DelByIds(ids []int) (int64, error) {
	var (
		model models.{{.Name | Transform}}
	)
	tx := re.db.Delete(&model, ids)
	return tx.RowsAffected, tx.Error
}

//Page 返回分页数据
func (re *{{.Name | Transform}}Repository) Page(where map[string]interface{}, page int, pageSize int, sortField string) map[string]interface{} {
	var (
		total  int64
		data   []models.{{.Name | Transform}}
		offSet int
	)
	db := global.Db.Model(&models.{{.Name | Transform}}{})

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
func (re *{{.Name | Transform}}Repository) Insert(data requests.{{.Name | Transform}}Request) (model models.{{.Name | Transform}}, err error) {

	model = models.{{.Name | Transform}}{
		{{range .Fields}}{{ .Name | Transform}}:	data.{{ .Name | Transform}},
		{{end}}
	}

	result := re.db.Create(&model)
	err = result.Error
	return model, err
}
`
}
