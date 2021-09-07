package template

const HeaderTmpl = `
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package {{.}}

import(
	"database/sql"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gen/helper"
)
`

const FuncTmpl = `
/*
{{.Doc}}*/
func ({{.S}} {{.MethodStruct}}){{.MethodName}}({{.GetParamInTmpl}})({{.GetResultParamInTmpl}}){
	{{if .HasSqlData}}params := map[string]interface{}{ {{range $index,$data:=.SqlData}}
		"{{$data}}":{{$data}},{{end}}
	}

	{{end}}var generateSQL string{{range $line:=.SqlTmplList}}{{$line}}
	{{end}}

	{{if .HasNeedNewResult}}result ={{if .ResultData.IsMap}}make{{else}}new{{end}}({{if ne .ResultData.Package ""}}{{.ResultData.Package}}.{{end}}{{.ResultData.Type}}){{end}}
	{{.ExecuteResult}} = {{.S}}.UnderlyingDB().{{.GormOption}}(generateSQL{{if .HasSqlData}},params{{end}}){{if not .ResultData.IsNull}}.Find({{if .HasGotPoint}}&{{end}}{{.ResultData.Name}}){{end}}.Error
	return
}

`

const BaseStruct = `
type {{.NewStructName}} struct {
	gen.DO

	{{range $p :=.Members}}{{$p.Name}}  field.{{$p.NewType}}
	{{end}}
}

func New{{.StructName}}(db *gorm.DB) *{{.NewStructName}} {
	_{{.NewStructName}} := new({{.NewStructName}})

	_{{.NewStructName}}.UseDB(db)
	_{{.NewStructName}}.UseModel({{.StructInfo.Package}}.{{.StructInfo.Type}}{})

	{{if .HasMember}}tableName := _{{.NewStructName}}.TableName(){{end}}
	{{range $p :=.Members}} _{{$.NewStructName}}.{{$p.Name}} = field.New{{$p.NewType}}(tableName, "{{$p.ColumnName}}")
	{{end}}
	
	return _{{.NewStructName}}
}

`

const BaseGormFunc = `
func ({{.S}} {{.NewStructName}}) Debug() *{{.NewStructName}} {
	{{.S}}.DO = *{{.S}}.DO.Debug().(*gen.DO)
	return &{{.S}}
}

func ({{.S}} {{.NewStructName}}) WithContext(ctx context.Context) *{{.NewStructName}} {
	{{.S}}.DO = *{{.S}}.DO.WithContext(ctx).(*gen.DO)
	return &{{.S}}
}

func ({{.S}} {{.NewStructName}}) Hints(hs ...gen.Hint) *{{.NewStructName}} {
	{{.S}}.DO = *{{.S}}.DO.Hints(hs...).(*gen.DO)
	return &{{.S}}
}

func ({{.S}} {{.NewStructName}}) Not(conds ...gen.Condition) *{{.NewStructName}} {
	{{.S}}.DO = *{{.S}}.DO.Not(conds...).(*gen.DO)
	return &{{.S}}
}

func ({{.S}} {{.NewStructName}}) Or(conds ...gen.Condition) *{{.NewStructName}} {
	{{.S}}.DO = *{{.S}}.DO.Or(conds...).(*gen.DO)
	return &{{.S}}
}

func ({{.S}} {{.NewStructName}}) Select(conds ...field.Expr) *{{.NewStructName}} {
	{{.S}}.DO = *{{.S}}.DO.Select(conds...).(*gen.DO)
	return &{{.S}}
}

func ({{.S}} {{.NewStructName}}) Where(conds ...gen.Condition) *{{.NewStructName}} {
	{{.S}}.DO = *{{.S}}.DO.Where(conds...).(*gen.DO)
	return &{{.S}}
}

func ({{.S}} {{.NewStructName}}) Order(conds ...field.Expr) *{{.NewStructName}} {
	{{.S}}.DO = *{{.S}}.DO.Order(conds...).(*gen.DO)
	return &{{.S}}
}

func ({{.S}} {{.NewStructName}}) Distinct(cols ...field.Expr) *{{.NewStructName}} {
	{{.S}}.DO = *{{.S}}.DO.Distinct(cols...).(*gen.DO)
	return &{{.S}}
}

func ({{.S}} {{.NewStructName}}) Omit(cols ...field.Expr) *{{.NewStructName}} {
	{{.S}}.DO = *{{.S}}.DO.Omit(cols...).(*gen.DO)
	return &{{.S}}
}

func ({{.S}} {{.NewStructName}}) Join(table schema.Tabler, on ...gen.Condition) *{{.NewStructName}} {
	{{.S}}.DO = *{{.S}}.DO.Join(table, on...).(*gen.DO)
	return &{{.S}}
}

func ({{.S}} {{.NewStructName}}) LeftJoin(table schema.Tabler, on ...gen.Condition) *{{.NewStructName}} {
	{{.S}}.DO = *{{.S}}.DO.LeftJoin(table, on...).(*gen.DO)
	return &{{.S}}
}

func ({{.S}} {{.NewStructName}}) RightJoin(table schema.Tabler, on ...gen.Condition) *{{.NewStructName}} {
	{{.S}}.DO = *{{.S}}.DO.RightJoin(table, on...).(*gen.DO)
	return &{{.S}}
}

func ({{.S}} {{.NewStructName}}) Group(col field.Expr) *{{.NewStructName}} {
	{{.S}}.DO = *{{.S}}.DO.Group(col).(*gen.DO)
	return &{{.S}}
}

func ({{.S}} {{.NewStructName}}) Having(conds ...gen.Condition) *{{.NewStructName}} {
	{{.S}}.DO = *{{.S}}.DO.Having(conds...).(*gen.DO)
	return &{{.S}}
}

func ({{.S}} {{.NewStructName}}) Limit(limit int) *{{.NewStructName}} {
	{{.S}}.DO = *{{.S}}.DO.Limit(limit).(*gen.DO)
	return &{{.S}}
}

func ({{.S}} {{.NewStructName}}) Offset(offset int) *{{.NewStructName}} {
	{{.S}}.DO = *{{.S}}.DO.Offset(offset).(*gen.DO)
	return &{{.S}}
}

func ({{.S}} {{.NewStructName}}) Scopes(funcs ...func(gen.Dao) gen.Dao) *{{.NewStructName}} {
	{{.S}}.DO = *{{.S}}.DO.Scopes(funcs...).(*gen.DO)
	return &{{.S}}
}

func ({{.S}} {{.NewStructName}}) Unscoped() *{{.NewStructName}} {
	{{.S}}.DO = *{{.S}}.DO.Unscoped().(*gen.DO)
	return &{{.S}}
}

func ({{.S}} {{.NewStructName}}) Create(values ...*{{.StructInfo.Package}}.{{.StructInfo.Type}}) error {
	if len(values) == 0 {
		return nil
	}
	return {{.S}}.DO.Create(values)
}

func ({{.S}} {{.NewStructName}}) CreateInBatches(values []*{{.StructInfo.Package}}.{{.StructInfo.Type}}, batchSize int) error {
	return {{.S}}.DO.CreateInBatches(values, batchSize)
}

func ({{.S}} {{.NewStructName}}) Save(values ...*{{.StructInfo.Package}}.{{.StructInfo.Type}}) error {
	if len(values) == 0 {
		return nil
	}
	return {{.S}}.DO.Save(values)
}

func ({{.S}} {{.NewStructName}}) First() (*{{.StructInfo.Package}}.{{.StructInfo.Type}}, error) {
	if result, err := {{.S}}.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*{{.StructInfo.Package}}.{{.StructInfo.Type}}), nil
	}
}

func ({{.S}} {{.NewStructName}}) Take() (*{{.StructInfo.Package}}.{{.StructInfo.Type}}, error) {
	if result, err := {{.S}}.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*{{.StructInfo.Package}}.{{.StructInfo.Type}}), nil
	}
}

func ({{.S}} {{.NewStructName}}) Last() (*{{.StructInfo.Package}}.{{.StructInfo.Type}}, error) {
	if result, err := {{.S}}.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*{{.StructInfo.Package}}.{{.StructInfo.Type}}), nil
	}
}

func ({{.S}} {{.NewStructName}}) Find() ([]*{{.StructInfo.Package}}.{{.StructInfo.Type}}, error) {
	result, err := {{.S}}.DO.Find()
	return result.([]*{{.StructInfo.Package}}.{{.StructInfo.Type}}), err
}

func ({{.S}} {{.NewStructName}}) FindInBatches(result []*{{.StructInfo.Package}}.{{.StructInfo.Type}}, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return {{.S}}.DO.FindInBatches(&result, batchSize, fc)
}

func ({{.S}} {{.NewStructName}}) FindByPage(offset int, limit int) (result []*{{.StructInfo.Package}}.{{.StructInfo.Type}}, count int64, err error) {
	count, err = {{.S}}.Count()
	if err != nil {
		return
	}

	result, err = {{.S}}.Offset(offset).Limit(limit).Find()
	return
}

func ({{.S}} {{.NewStructName}}) Model(result *{{.StructInfo.Package}}.{{.StructInfo.Type}}) *{{.NewStructName}} {
	{{.S}}.DO = *{{.S}}.DO.Model(result).(*gen.DO)
	return &{{.S}}
}

func ({{.S}} {{.NewStructName}}) Begin(opts ...*sql.TxOptions) *{{.NewStructName}} {
	{{.S}}.DO = *{{.S}}.DO.Begin(opts...).(*gen.DO)
	return &{{.S}}
}

func ({{.S}} {{.NewStructName}}) Commit() *{{.NewStructName}} {
	{{.S}}.DO = *{{.S}}.DO.Commit().(*gen.DO)
	return &{{.S}}
}

func ({{.S}} {{.NewStructName}}) RollBack() *{{.NewStructName}} {
	{{.S}}.DO = *{{.S}}.DO.Commit().(*gen.DO)
	return &{{.S}}
}

func ({{.S}} {{.NewStructName}}) SavePoint(name string) *{{.NewStructName}} {
	{{.S}}.DO = *{{.S}}.DO.SavePoint(name).(*gen.DO)
	return &{{.S}}
}

func ({{.S}} {{.NewStructName}}) RollBackTo(name string) *{{.NewStructName}} {
	{{.S}}.DO = *{{.S}}.DO.RollBackTo(name).(*gen.DO)
	return &{{.S}}
}
`

const UseTmpl = `
type Query struct{
	db *gorm.DB

	{{range $name,$d :=.Data}}{{$d.StructName}} *{{$d.NewStructName}}
	{{end}}
}

func (q *Query) Transaction(fc func(db *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.withTx(tx)) }, opts...)
}

func (q Query) Begin(opts ...*sql.TxOptions) *Query {
	q.db = q.db.Begin(opts...)
	return &q
}

func (q Query) Commit() *Query {
	q.db = q.db.Commit()
	return &q
}

func (q Query) Rollback() *Query {
	q.db = q.db.Rollback()
	return &q
}

func (q Query) SavePoint(name string) *Query {
	q.db = q.db.SavePoint(name)
	return &q
}

func (q Query) RollbackTo(name string) *Query {
	q.db = q.db.RollbackTo(name)
	return &q
}

func (q Query) withTx(tx *gorm.DB) *Query {
	q.db = tx
	return &q
}

func Use(db *gorm.DB) *Query {
	return &Query{
		db: db,
		{{range $name,$d :=.Data}}{{$d.StructName}}: New{{$d.StructName}}(db),
		{{end}}
	}
}
`

// ModelTemplate used as a variable because it cannot load template file after packed, params still can pass file
const ModelTemplate = `
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
package {{.StructInfo.Package}}

import "time"

const TableName{{.StructName}} = "{{.TableName}}"

// {{.TableName}}
type {{.StructName}} struct {
    {{range .Members}}
	{{if .MultilineComment -}}
	/*
{{.ColumnComment}}
    */
	{{end -}}
    {{.Name}} {{.ModelType}} ` + "`json:\"{{.JSONTag}}\" gorm:\"column:{{.GORMTag}}\"{{.NewTag}}` " +
	"{{if not .MultilineComment}}{{if .ColumnComment}}// {{.ColumnComment}}{{end}}{{end}}" +
	`{{end}}
}

// TableName .
func (*{{.StructName}}) TableName() string {
    return TableName{{.StructName}}
}
`
