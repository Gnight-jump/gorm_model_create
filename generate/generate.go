/**
生成models的主要功能
*/
package generate

import (
	"gorm.io/gorm"
	"modelgenerator/conf"
	"modelgenerator/dbConn"
	"modelgenerator/util"
	"strings"
)

// 初始化需要传入一个model
func Genertate(config conf.Conf) {
	tableNames := config.TargetModel.TableName
	// 获取表单信息
	db := dbConn.GetConnect()
	tables := getTables(db, tableNames, config.Dbase.DBName) //生成所有表信息
	for _, table := range tables {
		// 获取字段
		fields := getFields(db, table.Name)
		// 生成model
		generateModel(&config.TargetModel, table, fields)
	}
}

// 获取具体表单
func getTables(db *gorm.DB, tableNames []string, dbName string) []Table {

	// 字符串拼接生成表名范围
	tableNamesStr := "'" + strings.Join(tableNames, "','") + "'"

	// 获取指定表信息
	var tables []Table
	if tableNamesStr == "" {
		db.Raw("SELECT TABLE_NAME as Name,TABLE_COMMENT as Comment FROM information_schema.TABLES " +
			"WHERE table_schema='" + dbName + "';").Find(&tables)
	} else {
		db.Raw("SELECT TABLE_NAME as Name,TABLE_COMMENT as Comment FROM information_schema.TABLES " +
			"WHERE TABLE_NAME IN (" + tableNamesStr + ") AND " +
			"table_schema='" + dbName + "';").Find(&tables)
	}
	return tables
}

// 获取字段的详情信息
func getFields(db *gorm.DB, tableName string) []Field {
	var fields []Field
	db.Raw("show FULL COLUMNS from " + tableName + ";").Find(&fields)
	return fields
}

//生成Model
func generateModel(gModel *conf.GModel, table Table, fields []Field) {

	var builder strings.Builder
	builder.WriteString("package " + gModel.PackageName + "\n\n")

	// 表注释
	if len(table.Comment) > 0 {
		builder.WriteString("// " + table.Comment + "\n")
	}

	// 生成结构体
	builder.WriteString("type " + util.Marshal(table.Name) + " struct {\n")

	// 文件内容填充
	for _, field := range fields {
		fieldName := field.Field
		/**
		字段名 字段类型 `json:"字段名" gorm:"column:字段名"` //注释
		*/
		builder.WriteString("\t" + util.Marshal(fieldName) + "\t" + getFiledType(field) + "\t" +
			"`" + getFieldJson(field) + "`\t" + getFieldComment(field) + "\n")
	}
	builder.WriteString("}\n")

	// 函数名称返回自身
	/**
	func (e *结构体名) TableName() string {
	    return 结构体名
	}
	*/
	builder.WriteString("func (e *" + util.Marshal(table.Name) +
		") TableName() string { \n    return \"" + table.Name + "\"\n}")

	// 文件生成
	fileName := gModel.StorePath + util.Marshal(table.Name) + ".go"
	util.MakeGoModel(gModel.StorePath, fileName, builder.String(), gModel.ModelCover)
}
