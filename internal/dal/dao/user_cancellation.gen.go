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

func newUserCancellation(db *gorm.DB, opts ...gen.DOOption) userCancellation {
	_userCancellation := userCancellation{}

	_userCancellation.userCancellationDo.UseDB(db, opts...)
	_userCancellation.userCancellationDo.UseModel(&model.UserCancellation{})

	tableName := _userCancellation.userCancellationDo.TableName()
	_userCancellation.ALL = field.NewAsterisk(tableName)
	_userCancellation.ID = field.NewInt64(tableName, "id")
	_userCancellation.UID = field.NewString(tableName, "uid")
	_userCancellation.Reason = field.NewString(tableName, "reason")
	_userCancellation.ApplyTime = field.NewTime(tableName, "apply_time")
	_userCancellation.ConfirmTime = field.NewTime(tableName, "confirm_time")
	_userCancellation.ConfirmRemark = field.NewString(tableName, "confirm_remark")
	_userCancellation.Status = field.NewInt32(tableName, "status")
	_userCancellation.CreatedAt = field.NewTime(tableName, "created_at")
	_userCancellation.UpdatedAt = field.NewTime(tableName, "updated_at")
	_userCancellation.DeletedAt = field.NewField(tableName, "deleted_at")

	_userCancellation.fillFieldMap()

	return _userCancellation
}

type userCancellation struct {
	userCancellationDo userCancellationDo

	ALL           field.Asterisk
	ID            field.Int64
	UID           field.String // 用户ID
	Reason        field.String // 申请理由
	ApplyTime     field.Time   // 申请时间
	ConfirmTime   field.Time   // 确认时间
	ConfirmRemark field.String // 确认备注
	Status        field.Int32  // 处理状态（1待处理，2注销通过，3注销驳回）
	CreatedAt     field.Time   // 创建时间
	UpdatedAt     field.Time   // 更新时间
	DeletedAt     field.Field  // 删除时间

	fieldMap map[string]field.Expr
}

func (u userCancellation) Table(newTableName string) *userCancellation {
	u.userCancellationDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u userCancellation) As(alias string) *userCancellation {
	u.userCancellationDo.DO = *(u.userCancellationDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *userCancellation) updateTableName(table string) *userCancellation {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewInt64(table, "id")
	u.UID = field.NewString(table, "uid")
	u.Reason = field.NewString(table, "reason")
	u.ApplyTime = field.NewTime(table, "apply_time")
	u.ConfirmTime = field.NewTime(table, "confirm_time")
	u.ConfirmRemark = field.NewString(table, "confirm_remark")
	u.Status = field.NewInt32(table, "status")
	u.CreatedAt = field.NewTime(table, "created_at")
	u.UpdatedAt = field.NewTime(table, "updated_at")
	u.DeletedAt = field.NewField(table, "deleted_at")

	u.fillFieldMap()

	return u
}

func (u *userCancellation) WithContext(ctx context.Context) *userCancellationDo {
	return u.userCancellationDo.WithContext(ctx)
}

func (u userCancellation) TableName() string { return u.userCancellationDo.TableName() }

func (u userCancellation) Alias() string { return u.userCancellationDo.Alias() }

func (u *userCancellation) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *userCancellation) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 10)
	u.fieldMap["id"] = u.ID
	u.fieldMap["uid"] = u.UID
	u.fieldMap["reason"] = u.Reason
	u.fieldMap["apply_time"] = u.ApplyTime
	u.fieldMap["confirm_time"] = u.ConfirmTime
	u.fieldMap["confirm_remark"] = u.ConfirmRemark
	u.fieldMap["status"] = u.Status
	u.fieldMap["created_at"] = u.CreatedAt
	u.fieldMap["updated_at"] = u.UpdatedAt
	u.fieldMap["deleted_at"] = u.DeletedAt
}

func (u userCancellation) clone(db *gorm.DB) userCancellation {
	u.userCancellationDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u userCancellation) replaceDB(db *gorm.DB) userCancellation {
	u.userCancellationDo.ReplaceDB(db)
	return u
}

type userCancellationDo struct{ gen.DO }

func (u userCancellationDo) Debug() *userCancellationDo {
	return u.withDO(u.DO.Debug())
}

func (u userCancellationDo) WithContext(ctx context.Context) *userCancellationDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userCancellationDo) ReadDB() *userCancellationDo {
	return u.Clauses(dbresolver.Read)
}

func (u userCancellationDo) WriteDB() *userCancellationDo {
	return u.Clauses(dbresolver.Write)
}

func (u userCancellationDo) Session(config *gorm.Session) *userCancellationDo {
	return u.withDO(u.DO.Session(config))
}

func (u userCancellationDo) Clauses(conds ...clause.Expression) *userCancellationDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userCancellationDo) Returning(value interface{}, columns ...string) *userCancellationDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u userCancellationDo) Not(conds ...gen.Condition) *userCancellationDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userCancellationDo) Or(conds ...gen.Condition) *userCancellationDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userCancellationDo) Select(conds ...field.Expr) *userCancellationDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userCancellationDo) Where(conds ...gen.Condition) *userCancellationDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userCancellationDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *userCancellationDo {
	return u.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (u userCancellationDo) Order(conds ...field.Expr) *userCancellationDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userCancellationDo) Distinct(cols ...field.Expr) *userCancellationDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userCancellationDo) Omit(cols ...field.Expr) *userCancellationDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userCancellationDo) Join(table schema.Tabler, on ...field.Expr) *userCancellationDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userCancellationDo) LeftJoin(table schema.Tabler, on ...field.Expr) *userCancellationDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userCancellationDo) RightJoin(table schema.Tabler, on ...field.Expr) *userCancellationDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userCancellationDo) Group(cols ...field.Expr) *userCancellationDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userCancellationDo) Having(conds ...gen.Condition) *userCancellationDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userCancellationDo) Limit(limit int) *userCancellationDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userCancellationDo) Offset(offset int) *userCancellationDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userCancellationDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *userCancellationDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userCancellationDo) Unscoped() *userCancellationDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userCancellationDo) Create(values ...*model.UserCancellation) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userCancellationDo) CreateInBatches(values []*model.UserCancellation, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userCancellationDo) Save(values ...*model.UserCancellation) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userCancellationDo) First() (*model.UserCancellation, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserCancellation), nil
	}
}

func (u userCancellationDo) Take() (*model.UserCancellation, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserCancellation), nil
	}
}

func (u userCancellationDo) Last() (*model.UserCancellation, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserCancellation), nil
	}
}

func (u userCancellationDo) Find() ([]*model.UserCancellation, error) {
	result, err := u.DO.Find()
	return result.([]*model.UserCancellation), err
}

func (u userCancellationDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserCancellation, err error) {
	buf := make([]*model.UserCancellation, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userCancellationDo) FindInBatches(result *[]*model.UserCancellation, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userCancellationDo) Attrs(attrs ...field.AssignExpr) *userCancellationDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userCancellationDo) Assign(attrs ...field.AssignExpr) *userCancellationDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userCancellationDo) Joins(fields ...field.RelationField) *userCancellationDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u userCancellationDo) Preload(fields ...field.RelationField) *userCancellationDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u userCancellationDo) FirstOrInit() (*model.UserCancellation, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserCancellation), nil
	}
}

func (u userCancellationDo) FirstOrCreate() (*model.UserCancellation, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserCancellation), nil
	}
}

func (u userCancellationDo) FindByPage(offset int, limit int) (result []*model.UserCancellation, count int64, err error) {
	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u userCancellationDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u userCancellationDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u userCancellationDo) Delete(models ...*model.UserCancellation) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *userCancellationDo) withDO(do gen.Dao) *userCancellationDo {
	u.DO = *do.(*gen.DO)
	return u
}
