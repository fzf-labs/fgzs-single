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

func newFileUpload(db *gorm.DB, opts ...gen.DOOption) fileUpload {
	_fileUpload := fileUpload{}

	_fileUpload.fileUploadDo.UseDB(db, opts...)
	_fileUpload.fileUploadDo.UseModel(&model.FileUpload{})

	tableName := _fileUpload.fileUploadDo.TableName()
	_fileUpload.ALL = field.NewAsterisk(tableName)
	_fileUpload.ID = field.NewInt64(tableName, "id")
	_fileUpload.FileCategory = field.NewString(tableName, "file_category")
	_fileUpload.FileName = field.NewString(tableName, "file_name")
	_fileUpload.OriginalFileName = field.NewString(tableName, "original_file_name")
	_fileUpload.Storage = field.NewString(tableName, "storage")
	_fileUpload.Path = field.NewString(tableName, "path")
	_fileUpload.Ext = field.NewString(tableName, "ext")
	_fileUpload.Size = field.NewInt64(tableName, "size")
	_fileUpload.Sha1 = field.NewString(tableName, "sha1")
	_fileUpload.Status = field.NewInt32(tableName, "status")
	_fileUpload.CreatedAt = field.NewTime(tableName, "created_at")
	_fileUpload.UpdatedAt = field.NewTime(tableName, "updated_at")
	_fileUpload.DeletedAt = field.NewField(tableName, "deleted_at")

	_fileUpload.fillFieldMap()

	return _fileUpload
}

type fileUpload struct {
	fileUploadDo fileUploadDo

	ALL              field.Asterisk
	ID               field.Int64
	FileCategory     field.String // 文件分类
	FileName         field.String // 文件新名称
	OriginalFileName field.String // 文件原名称
	Storage          field.String // 存储方式
	Path             field.String // 文件路径
	Ext              field.String // 文件类型
	Size             field.Int64  // 文件大小
	Sha1             field.String // 文件sha1值
	Status           field.Int32  // 状态(1 正常 2冻结)
	CreatedAt        field.Time   // 创建时间
	UpdatedAt        field.Time   // 更新时间
	DeletedAt        field.Field  // 删除时间

	fieldMap map[string]field.Expr
}

func (f fileUpload) Table(newTableName string) *fileUpload {
	f.fileUploadDo.UseTable(newTableName)
	return f.updateTableName(newTableName)
}

func (f fileUpload) As(alias string) *fileUpload {
	f.fileUploadDo.DO = *(f.fileUploadDo.As(alias).(*gen.DO))
	return f.updateTableName(alias)
}

func (f *fileUpload) updateTableName(table string) *fileUpload {
	f.ALL = field.NewAsterisk(table)
	f.ID = field.NewInt64(table, "id")
	f.FileCategory = field.NewString(table, "file_category")
	f.FileName = field.NewString(table, "file_name")
	f.OriginalFileName = field.NewString(table, "original_file_name")
	f.Storage = field.NewString(table, "storage")
	f.Path = field.NewString(table, "path")
	f.Ext = field.NewString(table, "ext")
	f.Size = field.NewInt64(table, "size")
	f.Sha1 = field.NewString(table, "sha1")
	f.Status = field.NewInt32(table, "status")
	f.CreatedAt = field.NewTime(table, "created_at")
	f.UpdatedAt = field.NewTime(table, "updated_at")
	f.DeletedAt = field.NewField(table, "deleted_at")

	f.fillFieldMap()

	return f
}

func (f *fileUpload) WithContext(ctx context.Context) *fileUploadDo {
	return f.fileUploadDo.WithContext(ctx)
}

func (f fileUpload) TableName() string { return f.fileUploadDo.TableName() }

func (f fileUpload) Alias() string { return f.fileUploadDo.Alias() }

func (f *fileUpload) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := f.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (f *fileUpload) fillFieldMap() {
	f.fieldMap = make(map[string]field.Expr, 13)
	f.fieldMap["id"] = f.ID
	f.fieldMap["file_category"] = f.FileCategory
	f.fieldMap["file_name"] = f.FileName
	f.fieldMap["original_file_name"] = f.OriginalFileName
	f.fieldMap["storage"] = f.Storage
	f.fieldMap["path"] = f.Path
	f.fieldMap["ext"] = f.Ext
	f.fieldMap["size"] = f.Size
	f.fieldMap["sha1"] = f.Sha1
	f.fieldMap["status"] = f.Status
	f.fieldMap["created_at"] = f.CreatedAt
	f.fieldMap["updated_at"] = f.UpdatedAt
	f.fieldMap["deleted_at"] = f.DeletedAt
}

func (f fileUpload) clone(db *gorm.DB) fileUpload {
	f.fileUploadDo.ReplaceConnPool(db.Statement.ConnPool)
	return f
}

func (f fileUpload) replaceDB(db *gorm.DB) fileUpload {
	f.fileUploadDo.ReplaceDB(db)
	return f
}

type fileUploadDo struct{ gen.DO }

func (f fileUploadDo) Debug() *fileUploadDo {
	return f.withDO(f.DO.Debug())
}

func (f fileUploadDo) WithContext(ctx context.Context) *fileUploadDo {
	return f.withDO(f.DO.WithContext(ctx))
}

func (f fileUploadDo) ReadDB() *fileUploadDo {
	return f.Clauses(dbresolver.Read)
}

func (f fileUploadDo) WriteDB() *fileUploadDo {
	return f.Clauses(dbresolver.Write)
}

func (f fileUploadDo) Session(config *gorm.Session) *fileUploadDo {
	return f.withDO(f.DO.Session(config))
}

func (f fileUploadDo) Clauses(conds ...clause.Expression) *fileUploadDo {
	return f.withDO(f.DO.Clauses(conds...))
}

func (f fileUploadDo) Returning(value interface{}, columns ...string) *fileUploadDo {
	return f.withDO(f.DO.Returning(value, columns...))
}

func (f fileUploadDo) Not(conds ...gen.Condition) *fileUploadDo {
	return f.withDO(f.DO.Not(conds...))
}

func (f fileUploadDo) Or(conds ...gen.Condition) *fileUploadDo {
	return f.withDO(f.DO.Or(conds...))
}

func (f fileUploadDo) Select(conds ...field.Expr) *fileUploadDo {
	return f.withDO(f.DO.Select(conds...))
}

func (f fileUploadDo) Where(conds ...gen.Condition) *fileUploadDo {
	return f.withDO(f.DO.Where(conds...))
}

func (f fileUploadDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *fileUploadDo {
	return f.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (f fileUploadDo) Order(conds ...field.Expr) *fileUploadDo {
	return f.withDO(f.DO.Order(conds...))
}

func (f fileUploadDo) Distinct(cols ...field.Expr) *fileUploadDo {
	return f.withDO(f.DO.Distinct(cols...))
}

func (f fileUploadDo) Omit(cols ...field.Expr) *fileUploadDo {
	return f.withDO(f.DO.Omit(cols...))
}

func (f fileUploadDo) Join(table schema.Tabler, on ...field.Expr) *fileUploadDo {
	return f.withDO(f.DO.Join(table, on...))
}

func (f fileUploadDo) LeftJoin(table schema.Tabler, on ...field.Expr) *fileUploadDo {
	return f.withDO(f.DO.LeftJoin(table, on...))
}

func (f fileUploadDo) RightJoin(table schema.Tabler, on ...field.Expr) *fileUploadDo {
	return f.withDO(f.DO.RightJoin(table, on...))
}

func (f fileUploadDo) Group(cols ...field.Expr) *fileUploadDo {
	return f.withDO(f.DO.Group(cols...))
}

func (f fileUploadDo) Having(conds ...gen.Condition) *fileUploadDo {
	return f.withDO(f.DO.Having(conds...))
}

func (f fileUploadDo) Limit(limit int) *fileUploadDo {
	return f.withDO(f.DO.Limit(limit))
}

func (f fileUploadDo) Offset(offset int) *fileUploadDo {
	return f.withDO(f.DO.Offset(offset))
}

func (f fileUploadDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *fileUploadDo {
	return f.withDO(f.DO.Scopes(funcs...))
}

func (f fileUploadDo) Unscoped() *fileUploadDo {
	return f.withDO(f.DO.Unscoped())
}

func (f fileUploadDo) Create(values ...*model.FileUpload) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Create(values)
}

func (f fileUploadDo) CreateInBatches(values []*model.FileUpload, batchSize int) error {
	return f.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (f fileUploadDo) Save(values ...*model.FileUpload) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Save(values)
}

func (f fileUploadDo) First() (*model.FileUpload, error) {
	if result, err := f.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.FileUpload), nil
	}
}

func (f fileUploadDo) Take() (*model.FileUpload, error) {
	if result, err := f.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.FileUpload), nil
	}
}

func (f fileUploadDo) Last() (*model.FileUpload, error) {
	if result, err := f.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.FileUpload), nil
	}
}

func (f fileUploadDo) Find() ([]*model.FileUpload, error) {
	result, err := f.DO.Find()
	return result.([]*model.FileUpload), err
}

func (f fileUploadDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.FileUpload, err error) {
	buf := make([]*model.FileUpload, 0, batchSize)
	err = f.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (f fileUploadDo) FindInBatches(result *[]*model.FileUpload, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return f.DO.FindInBatches(result, batchSize, fc)
}

func (f fileUploadDo) Attrs(attrs ...field.AssignExpr) *fileUploadDo {
	return f.withDO(f.DO.Attrs(attrs...))
}

func (f fileUploadDo) Assign(attrs ...field.AssignExpr) *fileUploadDo {
	return f.withDO(f.DO.Assign(attrs...))
}

func (f fileUploadDo) Joins(fields ...field.RelationField) *fileUploadDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Joins(_f))
	}
	return &f
}

func (f fileUploadDo) Preload(fields ...field.RelationField) *fileUploadDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Preload(_f))
	}
	return &f
}

func (f fileUploadDo) FirstOrInit() (*model.FileUpload, error) {
	if result, err := f.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.FileUpload), nil
	}
}

func (f fileUploadDo) FirstOrCreate() (*model.FileUpload, error) {
	if result, err := f.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.FileUpload), nil
	}
}

func (f fileUploadDo) FindByPage(offset int, limit int) (result []*model.FileUpload, count int64, err error) {
	result, err = f.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = f.Offset(-1).Limit(-1).Count()
	return
}

func (f fileUploadDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = f.Count()
	if err != nil {
		return
	}

	err = f.Offset(offset).Limit(limit).Scan(result)
	return
}

func (f fileUploadDo) Scan(result interface{}) (err error) {
	return f.DO.Scan(result)
}

func (f fileUploadDo) Delete(models ...*model.FileUpload) (result gen.ResultInfo, err error) {
	return f.DO.Delete(models)
}

func (f *fileUploadDo) withDO(do gen.Dao) *fileUploadDo {
	f.DO = *do.(*gen.DO)
	return f
}
