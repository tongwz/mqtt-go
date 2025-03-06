package types

import (
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// CommonModelStruct 通用结构体
type CommonModelStruct interface {
	TableName() string
	GetID() int
}

// CommonWhereStruct 通用where查询
type CommonWhereStruct interface {
	Where(*gorm.DB) *gorm.DB
	TableName() string
}

type CommonDao interface {
	Insert(tx *gorm.DB, entity CommonModelStruct) (err error)
	InsertBackID(tx *gorm.DB, entity CommonModelStruct) (err error)
	Update(tx *gorm.DB, entity CommonModelStruct, updateInfo map[string]any, where CommonWhereStruct) (err error)
	SimpleUpdate(tx *gorm.DB, updateInfo map[string]any, where CommonWhereStruct) (err error)
	SimpleListDiyReTurn(tx *gorm.DB, entity CommonModelStruct) (err error)
	PluckList(tx *gorm.DB, entity CommonModelStruct) (err error)
	Detail(tx *gorm.DB, entity CommonModelStruct) (err error)
	SimpleDetail(tx *gorm.DB, entity CommonModelStruct) (err error)
	SimpleDetailForUpdate(tx *gorm.DB, entity CommonModelStruct) (err error)
	SimpleDetailDiyReturn(tx *gorm.DB, entity CommonModelStruct) (err error)
	Delete(tx *gorm.DB, entity CommonModelStruct) (err error)
}

// BaseGenericDao 一个是model 一个where
type BaseGenericDao[T CommonModelStruct, TT CommonWhereStruct] struct{}

// Insert 插入
func (b *BaseGenericDao[T, TT]) Insert(tx *gorm.DB, entity T) error {
	return tx.Table(entity.TableName()).Create(&entity).Error
}

// InsertBackID 插入数据返回id
func (b *BaseGenericDao[T, TT]) InsertBackID(tx *gorm.DB, entity T) (int, error) {
	err := tx.Table(entity.TableName()).Create(&entity).Error
	return entity.GetID(), err
}

// batchSize一次插入多少数据
func (b *BaseGenericDao[T, TT]) InsertBatch(tx *gorm.DB, entity []T, batchSize int) error {
	if len(entity) == 0 {
		return errors.New("empty entity")
	}
	return tx.Table(entity[0].TableName()).CreateInBatches(&entity, batchSize).Error
}

// Update 更新
func (b *BaseGenericDao[T, TT]) Update(tx *gorm.DB, entity T, updateInfo map[string]any, where TT) error {
	return where.Where(tx.Table(entity.TableName())).Updates(updateInfo).Error
}

// 简化更新
func (b *BaseGenericDao[T, TT]) SimpleUpdate(tx *gorm.DB, updateInfo map[string]any, where TT) error {
	return where.Where(tx.Table(where.TableName())).Updates(updateInfo).Error
}

// 简化列表
func (b *BaseGenericDao[T, TT]) SimpleList(tx *gorm.DB, where TT) ([]T, error) {
	list := make([]T, 0)
	err := where.Where(tx.Table(where.TableName())).Find(&list).Error
	return list, err
}

// 列表 返回自定义列表 list要传引用
func (b *BaseGenericDao[T, TT]) SimpleListDiyReturn(tx *gorm.DB, where TT, list any) error {
	return where.Where(tx.Table(where.TableName())).Find(list).Error
}

// 由于有些查询效率过低因此先查到单独id列一类的数据之后再做二次id查询
func (b *BaseGenericDao[T, TT]) PluckList(tx *gorm.DB, where TT, list any, column string) error {
	return where.Where(tx.Table(where.TableName())).Pluck(column, list).Error
}

// 详情
func (b *BaseGenericDao[T, TT]) Detail(tx *gorm.DB, entity T, where TT) (T, error) {
	var info T
	err := where.Where(tx.Table(entity.TableName())).First(&info).Error
	return info, err
}

// 简化详情
func (b *BaseGenericDao[T, TT]) SimpleDetail(tx *gorm.DB, where TT) (T, error) {
	var info T
	err := where.Where(tx.Table(where.TableName())).Take(&info).Error
	return info, err
}

// 简化详情 for update
func (b *BaseGenericDao[T, TT]) SimpleDetailForUpdate(tx *gorm.DB, where TT) (T, error) {
	var info T
	err := where.Where(tx.Table(where.TableName()).Clauses(clause.Locking{Strength: "UPDATE"})).Take(&info).Error
	return info, err
}

// 自定义返回详情 info 传指针
func (b *BaseGenericDao[T, TT]) SimpleDetailDiyReturn(tx *gorm.DB, where TT, info any) error {
	return where.Where(tx.Table(where.TableName())).Take(info).Error
}

// 查询某个条件数据数量
func (b *BaseGenericDao[T, TT]) Count(tx *gorm.DB, where TT, count *int64) error {
	return where.Where(tx.Table(where.TableName())).Count(count).Error
}

// 查询字段之和 totalAmount = float64 或者 int int64都有可能
func (b *BaseGenericDao[T, TT]) Sum(tx *gorm.DB, where TT, column string, totalAmount any) error {
	return where.Where(tx.Table(where.TableName())).Select("COALESCE(SUM(" + column + "), 0)").Scan(totalAmount).Error
}

// 删除
func (b *BaseGenericDao[T, TT]) Delete(tx *gorm.DB, entity T, where TT) error {
	return where.Where(tx.Table(entity.TableName())).Delete(&entity).Error
}

func (b *BaseGenericDao[T, TT]) SimpleDelete(tx *gorm.DB, where TT) error {
	var entity T
	return where.Where(tx.Table(where.TableName())).Delete(&entity).Error
}
