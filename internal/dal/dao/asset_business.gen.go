// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"fgzs-single/internal/dal/model"
)

func newAssetBusiness(db *gorm.DB, opts ...gen.DOOption) assetBusiness {
	_assetBusiness := assetBusiness{}

	_assetBusiness.assetBusinessDo.UseDB(db, opts...)
	_assetBusiness.assetBusinessDo.UseModel(&model.AssetBusiness{})

	tableName := _assetBusiness.assetBusinessDo.TableName()
	_assetBusiness.ALL = field.NewAsterisk(tableName)
	_assetBusiness.ID = field.NewInt64(tableName, "id")
	_assetBusiness.Key = field.NewString(tableName, "key")
	_assetBusiness.UID = field.NewString(tableName, "uid")
	_assetBusiness.CommonOne = field.NewString(tableName, "common_one")
	_assetBusiness.CommonTwo = field.NewString(tableName, "common_two")
	_assetBusiness.CommonThree = field.NewString(tableName, "common_three")
	_assetBusiness.BusinessType = field.NewString(tableName, "business_type")
	_assetBusiness.AssetType = field.NewInt32(tableName, "asset_type")
	_assetBusiness.AssetOpt = field.NewInt32(tableName, "asset_opt")
	_assetBusiness.Amount = field.NewInt64(tableName, "amount")
	_assetBusiness.Record = field.NewField(tableName, "record")
	_assetBusiness.Status = field.NewInt32(tableName, "status")
	_assetBusiness.CreatedAt = field.NewTime(tableName, "created_at")
	_assetBusiness.UpdatedAt = field.NewTime(tableName, "updated_at")
	_assetBusiness.DeletedAt = field.NewField(tableName, "deleted_at")

	_assetBusiness.fillFieldMap()

	return _assetBusiness
}

type assetBusiness struct {
	assetBusinessDo assetBusinessDo

	ALL          field.Asterisk
	ID           field.Int64
	Key          field.String // 流水号
	UID          field.String // 用户ID
	CommonOne    field.String // 关联的业务的自定义ID
	CommonTwo    field.String // 关联的业务的自定义ID
	CommonThree  field.String // 关联的业务的自定义ID
	BusinessType field.String // 业务类型
	AssetType    field.Int32  // 资产类型
	AssetOpt     field.Int32  // 操作类型(1增加，-1减少)
	Amount       field.Int64  // 金额
	Record       field.Field  // 记录
	Status       field.Int32  // 状态
	CreatedAt    field.Time   // 创建时间
	UpdatedAt    field.Time   // 更新时间
	DeletedAt    field.Field  // 删除时间

	fieldMap map[string]field.Expr
}

func (a assetBusiness) Table(newTableName string) *assetBusiness {
	a.assetBusinessDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a assetBusiness) As(alias string) *assetBusiness {
	a.assetBusinessDo.DO = *(a.assetBusinessDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *assetBusiness) updateTableName(table string) *assetBusiness {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt64(table, "id")
	a.Key = field.NewString(table, "key")
	a.UID = field.NewString(table, "uid")
	a.CommonOne = field.NewString(table, "common_one")
	a.CommonTwo = field.NewString(table, "common_two")
	a.CommonThree = field.NewString(table, "common_three")
	a.BusinessType = field.NewString(table, "business_type")
	a.AssetType = field.NewInt32(table, "asset_type")
	a.AssetOpt = field.NewInt32(table, "asset_opt")
	a.Amount = field.NewInt64(table, "amount")
	a.Record = field.NewField(table, "record")
	a.Status = field.NewInt32(table, "status")
	a.CreatedAt = field.NewTime(table, "created_at")
	a.UpdatedAt = field.NewTime(table, "updated_at")
	a.DeletedAt = field.NewField(table, "deleted_at")

	a.fillFieldMap()

	return a
}

func (a *assetBusiness) WithContext(ctx context.Context) *assetBusinessDo {
	return a.assetBusinessDo.WithContext(ctx)
}

func (a assetBusiness) TableName() string { return a.assetBusinessDo.TableName() }

func (a assetBusiness) Alias() string { return a.assetBusinessDo.Alias() }

func (a *assetBusiness) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *assetBusiness) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 15)
	a.fieldMap["id"] = a.ID
	a.fieldMap["key"] = a.Key
	a.fieldMap["uid"] = a.UID
	a.fieldMap["common_one"] = a.CommonOne
	a.fieldMap["common_two"] = a.CommonTwo
	a.fieldMap["common_three"] = a.CommonThree
	a.fieldMap["business_type"] = a.BusinessType
	a.fieldMap["asset_type"] = a.AssetType
	a.fieldMap["asset_opt"] = a.AssetOpt
	a.fieldMap["amount"] = a.Amount
	a.fieldMap["record"] = a.Record
	a.fieldMap["status"] = a.Status
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["updated_at"] = a.UpdatedAt
	a.fieldMap["deleted_at"] = a.DeletedAt
}

func (a assetBusiness) clone(db *gorm.DB) assetBusiness {
	a.assetBusinessDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a assetBusiness) replaceDB(db *gorm.DB) assetBusiness {
	a.assetBusinessDo.ReplaceDB(db)
	return a
}

type assetBusinessDo struct{ gen.DO }

func (a assetBusinessDo) Debug() *assetBusinessDo {
	return a.withDO(a.DO.Debug())
}

func (a assetBusinessDo) WithContext(ctx context.Context) *assetBusinessDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a assetBusinessDo) ReadDB() *assetBusinessDo {
	return a.Clauses(dbresolver.Read)
}

func (a assetBusinessDo) WriteDB() *assetBusinessDo {
	return a.Clauses(dbresolver.Write)
}

func (a assetBusinessDo) Session(config *gorm.Session) *assetBusinessDo {
	return a.withDO(a.DO.Session(config))
}

func (a assetBusinessDo) Clauses(conds ...clause.Expression) *assetBusinessDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a assetBusinessDo) Returning(value interface{}, columns ...string) *assetBusinessDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a assetBusinessDo) Not(conds ...gen.Condition) *assetBusinessDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a assetBusinessDo) Or(conds ...gen.Condition) *assetBusinessDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a assetBusinessDo) Select(conds ...field.Expr) *assetBusinessDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a assetBusinessDo) Where(conds ...gen.Condition) *assetBusinessDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a assetBusinessDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *assetBusinessDo {
	return a.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (a assetBusinessDo) Order(conds ...field.Expr) *assetBusinessDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a assetBusinessDo) Distinct(cols ...field.Expr) *assetBusinessDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a assetBusinessDo) Omit(cols ...field.Expr) *assetBusinessDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a assetBusinessDo) Join(table schema.Tabler, on ...field.Expr) *assetBusinessDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a assetBusinessDo) LeftJoin(table schema.Tabler, on ...field.Expr) *assetBusinessDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a assetBusinessDo) RightJoin(table schema.Tabler, on ...field.Expr) *assetBusinessDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a assetBusinessDo) Group(cols ...field.Expr) *assetBusinessDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a assetBusinessDo) Having(conds ...gen.Condition) *assetBusinessDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a assetBusinessDo) Limit(limit int) *assetBusinessDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a assetBusinessDo) Offset(offset int) *assetBusinessDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a assetBusinessDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *assetBusinessDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a assetBusinessDo) Unscoped() *assetBusinessDo {
	return a.withDO(a.DO.Unscoped())
}

func (a assetBusinessDo) Create(values ...*model.AssetBusiness) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a assetBusinessDo) CreateInBatches(values []*model.AssetBusiness, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a assetBusinessDo) Save(values ...*model.AssetBusiness) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a assetBusinessDo) First() (*model.AssetBusiness, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.AssetBusiness), nil
	}
}

func (a assetBusinessDo) Take() (*model.AssetBusiness, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.AssetBusiness), nil
	}
}

func (a assetBusinessDo) Last() (*model.AssetBusiness, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.AssetBusiness), nil
	}
}

func (a assetBusinessDo) Find() ([]*model.AssetBusiness, error) {
	result, err := a.DO.Find()
	return result.([]*model.AssetBusiness), err
}

func (a assetBusinessDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AssetBusiness, err error) {
	buf := make([]*model.AssetBusiness, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a assetBusinessDo) FindInBatches(result *[]*model.AssetBusiness, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a assetBusinessDo) Attrs(attrs ...field.AssignExpr) *assetBusinessDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a assetBusinessDo) Assign(attrs ...field.AssignExpr) *assetBusinessDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a assetBusinessDo) Joins(fields ...field.RelationField) *assetBusinessDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a assetBusinessDo) Preload(fields ...field.RelationField) *assetBusinessDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a assetBusinessDo) FirstOrInit() (*model.AssetBusiness, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.AssetBusiness), nil
	}
}

func (a assetBusinessDo) FirstOrCreate() (*model.AssetBusiness, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.AssetBusiness), nil
	}
}

func (a assetBusinessDo) FindByPage(offset int, limit int) (result []*model.AssetBusiness, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a assetBusinessDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a assetBusinessDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a assetBusinessDo) Delete(models ...*model.AssetBusiness) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *assetBusinessDo) withDO(do gen.Dao) *assetBusinessDo {
	a.DO = *do.(*gen.DO)
	return a
}
