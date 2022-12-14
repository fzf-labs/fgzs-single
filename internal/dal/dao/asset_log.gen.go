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

func newAssetLog(db *gorm.DB, opts ...gen.DOOption) assetLog {
	_assetLog := assetLog{}

	_assetLog.assetLogDo.UseDB(db, opts...)
	_assetLog.assetLogDo.UseModel(&model.AssetLog{})

	tableName := _assetLog.assetLogDo.TableName()
	_assetLog.ALL = field.NewAsterisk(tableName)
	_assetLog.ID = field.NewInt64(tableName, "id")
	_assetLog.UID = field.NewString(tableName, "uid")
	_assetLog.Key = field.NewString(tableName, "key")
	_assetLog.BusinessType = field.NewString(tableName, "business_type")
	_assetLog.AssetType = field.NewInt32(tableName, "asset_type")
	_assetLog.AssetOpt = field.NewInt32(tableName, "asset_opt")
	_assetLog.Amount = field.NewInt64(tableName, "amount")
	_assetLog.BeforeAmount = field.NewInt64(tableName, "before_amount")
	_assetLog.AfterAmount = field.NewInt64(tableName, "after_amount")
	_assetLog.CreatedAt = field.NewTime(tableName, "created_at")
	_assetLog.UpdatedAt = field.NewTime(tableName, "updated_at")
	_assetLog.DeletedAt = field.NewField(tableName, "deleted_at")

	_assetLog.fillFieldMap()

	return _assetLog
}

type assetLog struct {
	assetLogDo assetLogDo

	ALL          field.Asterisk
	ID           field.Int64
	UID          field.String // 用户ID
	Key          field.String // 流水号
	BusinessType field.String // 业务类型
	AssetType    field.Int32  // 资产类型
	AssetOpt     field.Int32  // 操作类型(1增加，-1减少)
	Amount       field.Int64  // 变动的金额
	BeforeAmount field.Int64  // 变更前的金额
	AfterAmount  field.Int64  // 变更后的金额
	CreatedAt    field.Time   // 创建时间
	UpdatedAt    field.Time   // 更新时间
	DeletedAt    field.Field  // 删除时间

	fieldMap map[string]field.Expr
}

func (a assetLog) Table(newTableName string) *assetLog {
	a.assetLogDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a assetLog) As(alias string) *assetLog {
	a.assetLogDo.DO = *(a.assetLogDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *assetLog) updateTableName(table string) *assetLog {
	a.ALL = field.NewAsterisk(table)
	a.ID = field.NewInt64(table, "id")
	a.UID = field.NewString(table, "uid")
	a.Key = field.NewString(table, "key")
	a.BusinessType = field.NewString(table, "business_type")
	a.AssetType = field.NewInt32(table, "asset_type")
	a.AssetOpt = field.NewInt32(table, "asset_opt")
	a.Amount = field.NewInt64(table, "amount")
	a.BeforeAmount = field.NewInt64(table, "before_amount")
	a.AfterAmount = field.NewInt64(table, "after_amount")
	a.CreatedAt = field.NewTime(table, "created_at")
	a.UpdatedAt = field.NewTime(table, "updated_at")
	a.DeletedAt = field.NewField(table, "deleted_at")

	a.fillFieldMap()

	return a
}

func (a *assetLog) WithContext(ctx context.Context) *assetLogDo { return a.assetLogDo.WithContext(ctx) }

func (a assetLog) TableName() string { return a.assetLogDo.TableName() }

func (a assetLog) Alias() string { return a.assetLogDo.Alias() }

func (a *assetLog) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *assetLog) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 12)
	a.fieldMap["id"] = a.ID
	a.fieldMap["uid"] = a.UID
	a.fieldMap["key"] = a.Key
	a.fieldMap["business_type"] = a.BusinessType
	a.fieldMap["asset_type"] = a.AssetType
	a.fieldMap["asset_opt"] = a.AssetOpt
	a.fieldMap["amount"] = a.Amount
	a.fieldMap["before_amount"] = a.BeforeAmount
	a.fieldMap["after_amount"] = a.AfterAmount
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["updated_at"] = a.UpdatedAt
	a.fieldMap["deleted_at"] = a.DeletedAt
}

func (a assetLog) clone(db *gorm.DB) assetLog {
	a.assetLogDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a assetLog) replaceDB(db *gorm.DB) assetLog {
	a.assetLogDo.ReplaceDB(db)
	return a
}

type assetLogDo struct{ gen.DO }

func (a assetLogDo) Debug() *assetLogDo {
	return a.withDO(a.DO.Debug())
}

func (a assetLogDo) WithContext(ctx context.Context) *assetLogDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a assetLogDo) ReadDB() *assetLogDo {
	return a.Clauses(dbresolver.Read)
}

func (a assetLogDo) WriteDB() *assetLogDo {
	return a.Clauses(dbresolver.Write)
}

func (a assetLogDo) Session(config *gorm.Session) *assetLogDo {
	return a.withDO(a.DO.Session(config))
}

func (a assetLogDo) Clauses(conds ...clause.Expression) *assetLogDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a assetLogDo) Returning(value interface{}, columns ...string) *assetLogDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a assetLogDo) Not(conds ...gen.Condition) *assetLogDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a assetLogDo) Or(conds ...gen.Condition) *assetLogDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a assetLogDo) Select(conds ...field.Expr) *assetLogDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a assetLogDo) Where(conds ...gen.Condition) *assetLogDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a assetLogDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *assetLogDo {
	return a.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (a assetLogDo) Order(conds ...field.Expr) *assetLogDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a assetLogDo) Distinct(cols ...field.Expr) *assetLogDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a assetLogDo) Omit(cols ...field.Expr) *assetLogDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a assetLogDo) Join(table schema.Tabler, on ...field.Expr) *assetLogDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a assetLogDo) LeftJoin(table schema.Tabler, on ...field.Expr) *assetLogDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a assetLogDo) RightJoin(table schema.Tabler, on ...field.Expr) *assetLogDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a assetLogDo) Group(cols ...field.Expr) *assetLogDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a assetLogDo) Having(conds ...gen.Condition) *assetLogDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a assetLogDo) Limit(limit int) *assetLogDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a assetLogDo) Offset(offset int) *assetLogDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a assetLogDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *assetLogDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a assetLogDo) Unscoped() *assetLogDo {
	return a.withDO(a.DO.Unscoped())
}

func (a assetLogDo) Create(values ...*model.AssetLog) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a assetLogDo) CreateInBatches(values []*model.AssetLog, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a assetLogDo) Save(values ...*model.AssetLog) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a assetLogDo) First() (*model.AssetLog, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.AssetLog), nil
	}
}

func (a assetLogDo) Take() (*model.AssetLog, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.AssetLog), nil
	}
}

func (a assetLogDo) Last() (*model.AssetLog, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.AssetLog), nil
	}
}

func (a assetLogDo) Find() ([]*model.AssetLog, error) {
	result, err := a.DO.Find()
	return result.([]*model.AssetLog), err
}

func (a assetLogDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AssetLog, err error) {
	buf := make([]*model.AssetLog, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a assetLogDo) FindInBatches(result *[]*model.AssetLog, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a assetLogDo) Attrs(attrs ...field.AssignExpr) *assetLogDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a assetLogDo) Assign(attrs ...field.AssignExpr) *assetLogDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a assetLogDo) Joins(fields ...field.RelationField) *assetLogDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a assetLogDo) Preload(fields ...field.RelationField) *assetLogDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a assetLogDo) FirstOrInit() (*model.AssetLog, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.AssetLog), nil
	}
}

func (a assetLogDo) FirstOrCreate() (*model.AssetLog, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.AssetLog), nil
	}
}

func (a assetLogDo) FindByPage(offset int, limit int) (result []*model.AssetLog, count int64, err error) {
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

func (a assetLogDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a assetLogDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a assetLogDo) Delete(models ...*model.AssetLog) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *assetLogDo) withDO(do gen.Dao) *assetLogDo {
	a.DO = *do.(*gen.DO)
	return a
}
