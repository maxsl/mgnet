package orm

import (
	"os"
	"strings"
	"fmt"
	"strconv"
)

// 生成orm代码 
func FlushOrm(dir string) {

	//fileInfo
	_, err := os.Stat(dir);
	if  err != nil {
		fmt.Printf("err:FlushOrm() %v\n", err)
		return
	}
	
	for i:=0; i<len(Config); i++ {
		
		fileName := dir + "/" + "Orm" + strings.Title(Config[i].Table) + ".go"
		
		_, err := os.Stat(fileName);
		if  err == nil {
			continue
		}

		fd, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0755)
	
		FlushHead(Config[0], fd)
		FlushStruct(Config[0], fd)
		
		FlushLoadDb(Config[0], fd)
		FlushSaveDb(Config[0], fd)
		FlushInsert(Config[0], fd)
		FlushDelete(Config[0], fd)
		
		fd.Sync()
		fd.Close()
	}

}

// 输出头部
func FlushHead(config OrmConfig, file *os.File) {
	
	file.WriteString("package orm\n\n")
	file.WriteString("import (\n")
	file.WriteString("\t\"github.com/goodkele/mgnet/library/mysql\"\n")
	file.WriteString("\t\"encoding/json\"\n")
	file.WriteString("\t\"fmt\"\n")
	file.WriteString(")\n")
	
}

// 输出结构体
func FlushStruct(config OrmConfig, file *os.File) {

	file.WriteString("\n")
	for index, jsonColumns := range config.JsonColumnsType {
		file.WriteString("type " + strings.Title(config.Table) + config.Columns[index] + " struct {\n")
		
		for _, value := range jsonColumns {
			file.WriteString("\t")
			file.WriteString(value.Key + " " + value.Value)
			file.WriteString("\n")
		}
		file.WriteString("}\n")
	}

	file.WriteString("\n")
	file.WriteString("type Orm" + strings.Title(config.Table) + " struct {\n")
	for i:=0; i<len(config.Columns); i++ {
		file.WriteString("\t")
		
		// json类型
		if _, ok := config.JsonColumns[i]; ok == true {
			file.WriteString(config.Columns[i] + " " + strings.Title(config.Table + config.Columns[i]))
		} else {
			file.WriteString(config.Columns[i] + " " + config.ColumnsType[i])
		}
		
		file.WriteString("\n")
	}

	file.WriteString("\n")
	file.WriteString("\tisLoad bool\t// 是否已载入\n")
	
	file.WriteString("}\n")

}

// 输出SaveDb
func FlushSaveDb(config OrmConfig, file *os.File) {
	
	file.WriteString("\n")
	file.WriteString("// 更新\n")
	file.WriteString("func (this *OrmUser) SaveDb(conn *mysql.Connection) {\n")
	
	for index, _ := range config.JsonColumnsType {
		file.WriteString("\n")
		file.WriteString("\tjson" + config.Columns[index] + ", err := json.Marshal(this." + config.Columns[index] + ")\n")
		file.WriteString("\tif err != nil {\n")
		file.WriteString("\t}\n")
	}
	
	file.WriteString("\n")
	file.WriteString("\t// fixme 检测长度报警\n")
	
	sql, values := sqlUpdate(config)
	
	file.WriteString("\t_,err = conn.Execute(fmt.Sprintf(\""+ sql +"\", "+ values +"")
	
	file.WriteString("))\n")
	file.WriteString("\tif err != nil {\n")
	file.WriteString("\t}\n")

	file.WriteString("}\n")	

}

// 输出Insert
func FlushInsert(config OrmConfig, file *os.File) {
	
	file.WriteString("\n")
	file.WriteString("// 新增\n")
	file.WriteString("func (this *OrmUser) Insert(conn *mysql.Connection) {\n")
	
	for index, _ := range config.JsonColumnsType {
		file.WriteString("\n")
		file.WriteString("\tjson" + config.Columns[index] + ", err := json.Marshal(this." + config.Columns[index] + ")\n")
		file.WriteString("\tif err != nil {\n")
		file.WriteString("\t}\n")
	}
	
	file.WriteString("\n")
	file.WriteString("\t// fixme 检测长度报警\n")

	sql, values := sqlInsert(config)

	file.WriteString("\t_,err = conn.Execute(fmt.Sprintf(\""+ sql +"\", "+ values +"))\n")

	file.WriteString("\tif err != nil {\n")
	file.WriteString("\t}\n")

	file.WriteString("\n")

	file.WriteString("\tthis.isLoad = true\n")
	
	file.WriteString("\n")

	file.WriteString("}\n")
	
}

// 输出LoadDb
func FlushLoadDb(config OrmConfig, file *os.File) {
	
	file.WriteString("\n")
	file.WriteString("// 加载\n")
	file.WriteString("func (this *OrmUser) LoadDb(conn *mysql.Connection) {\n")

	sql, values := sqlSelect(config)
	file.WriteString("\n")
	file.WriteString("\tres,err := conn.QueryTable(fmt.Sprintf(\""+ sql +"\", "+ values +"))\n")
	file.WriteString("\tif err != nil {\n")
	file.WriteString("\t}\n")

	file.WriteString("\trows := res.Rows()\n")

	file.WriteString("\tif len(rows) > 0 {\n")
	
	for i:=0; i<len(config.Columns); i++ {
		if _, ok := config.JsonColumns[i]; ok == true {
			file.WriteString("\t\terr = json.Unmarshal([]byte(rows[0]["+ strconv.FormatInt(int64(i), 10) +"]."+ calcFormatVerb(config.ColumnsType[i]) +"), &this."+ config.Columns[i] +")\n")
			file.WriteString("\t\tif err != nil {\n")
			file.WriteString("\t\t}\n")
			file.WriteString("\n")
		} else {			
			file.WriteString("\t\tthis." + config.Columns[i] + " = rows[0][" + strconv.FormatInt(int64(i), 10) + "]." + calcFormatVerb(config.ColumnsType[i]))
			file.WriteString("\n")
		}
	}
	
	file.WriteString("\t\tthis.isLoad = true\n")

	file.WriteString("\t}\n")
	
	file.WriteString("\n")
	file.WriteString("}\n")

}

// 输出Delete
func FlushDelete(config OrmConfig, file *os.File) {
	
	file.WriteString("\n")
	file.WriteString("// 删除\n")
	file.WriteString("func (this *OrmUser) Delete(conn *mysql.Connection) {\n")

	sql, values := sqlDelete(config)
	
	file.WriteString("\n")
	file.WriteString("\t_,err := conn.Execute(fmt.Sprintf(\""+ sql +"\", "+ values +"))\n")
	file.WriteString("\tif err != nil {\n")
	file.WriteString("\t}\n")

	file.WriteString("\n")
	file.WriteString("}\n")
	
}



// delete
func sqlDelete(config OrmConfig) (string, string) {
	
	var sql, values string
	columns := []string{}

	sql = "DELETE FROM `" + config.Table + "` WHERE "
	
	for i:=0; i<len(config.PrimaryKey); i++ {
		columns = append(columns, config.Columns[config.PrimaryKey[i]] + "=" + calcSprintfVerb(config.ColumnsType[config.PrimaryKey[i]]))
	}

	sql = sql + strings.Join(columns, ", ")

	// values
	columns = columns[0:0]
	for i:=0; i<len(config.PrimaryKey); i++ {
		columns = append(columns, "this." + config.Columns[config.PrimaryKey[i]])
	}
	
	values = values + strings.Join(columns, ", ")

	return sql, values
	
}

// select
func sqlSelect(config OrmConfig) (string, string) {
	
	var sql, values string
	columns := []string{}

	sql = "SELECT "
	
	for i:=0; i<len(config.Columns); i++ {
		columns = append(columns, "`" + config.Columns[i] + "`")		
	}
	
	sql = sql + strings.Join(columns, ", ")
	
	sql = sql + " FROM `" + config.Table + "` WHERE "
	
	columns = columns[0:0]
	for i:=0; i<len(config.PrimaryKey); i++ {
		columns = append(columns, config.Columns[config.PrimaryKey[i]] + "=" + calcSprintfVerb(config.ColumnsType[config.PrimaryKey[i]]))
	}
	
	sql = sql + strings.Join(columns, ", ")
	
	// values
	columns = columns[0:0]
	for i:=0; i<len(config.PrimaryKey); i++ {
		columns = append(columns, "this." + config.Columns[config.PrimaryKey[i]])
	}
	
	values = values + strings.Join(columns, ", ")
	
	return sql, values

}

// update
func sqlUpdate(config OrmConfig) (string, string) {
	
	var sql, values string
	columns := []string{}

	sql = "UPDATE `"+ config.Table +"` SET "

	for i:=0; i<len(config.Columns); i++ {
		if inArray(i, config.PrimaryKey) {
			continue
		}

		columns = append(columns, "`" + config.Columns[i] + "`=" + calcSprintfVerb(config.ColumnsType[i]))		
	}
	
	sql = sql + " " + strings.Join(columns, ", ")
	
	sql = sql + " WHERE "
	
	columns = columns[0:0]
	for i:=0; i<len(config.PrimaryKey); i++ {
		columns = append(columns, "`" + config.Columns[config.PrimaryKey[i]] + "`=" + calcSprintfVerb(config.ColumnsType[config.PrimaryKey[i]]))
	}

	sql = sql + strings.Join(columns, ", ")
	
	// values
	columns = columns[0:0]
	for i:=0; i<len(config.Columns); i++ {
		if inArray(i, config.PrimaryKey) {
			continue
		}

		// json 类型
		if _, ok := config.JsonColumns[i]; ok == true {
			columns = append(columns, "json" + config.Columns[i])
		} else {			
			columns = append(columns, "this." + config.Columns[i])
		}
	}

	values = values + strings.Join(columns, ", ")
	
	if len(config.PrimaryKey) > 0 {
		values = values + ","
	}
	columns = columns[0:0]
	for i:=0; i<len(config.PrimaryKey); i++ {
		if _, ok := config.JsonColumns[config.PrimaryKey[i]]; ok == true {
			columns = append(columns, "json" + config.Columns[config.PrimaryKey[i]])
		} else {
			columns = append(columns, "this." + config.Columns[config.PrimaryKey[i]])
		}
	}

	values = values + " " + strings.Join(columns, ", ")
	
	return sql, values

}

// insert
func sqlInsert(config OrmConfig) (string, string) {

	var sql, values string
	columns := []string{}
	
	sql = "INSERT INTO `"+ config.Table +"`("

	for i:=0; i<len(config.Columns); i++ {
		columns = append(columns, "`" + config.Columns[i] + "`")
	}
	
	sql = sql + strings.Join(columns, ", ")
	
	sql = sql + ") VALUES("
	
	columns = columns[0:0]
	for i:=0; i<len(config.Columns); i++ {
		// json 类型
		if _, ok := config.JsonColumns[i]; ok == true {
			columns = append(columns, calcSprintfVerb(config.ColumnsType[i]))
		} else {			
			columns = append(columns, calcSprintfVerb(config.ColumnsType[i]))
		}
	}
	
	sql = sql + strings.Join(columns, ", ")

	sql = sql + ")"
	
	// values
	columns = columns[0:0]
	for i:=0; i<len(config.Columns); i++ {
		// json 类型
		if _, ok := config.JsonColumns[i]; ok == true {
			columns = append(columns, "json" + config.Columns[i])
		} else {			
			columns = append(columns, "this." + config.Columns[i])
		}
	}

	values = values + strings.Join(columns, ", ")
	
	return sql, values
	
}



// 计算格式化动作
func calcSprintfVerb(typeName string) string {
	
	var verb string
	switch typeName {
		case "int" :
			fallthrough
		case "int64" :
			fallthrough
		case "uint" :
			fallthrough
		case "uint64" :
			verb = "'%d'"
		
		case "float32" :
		case "float64" :
			verb = "'%f'"
			
		case "string" :
			verb = "'%s'"
	}
	return verb
	
}

func calcFormatVerb(typeName string) string {
	
	var verb string
	switch typeName {
		case "int" :
			verb = "Int()"
		case "int64" :
			verb = "Int64()"
		case "float32" :
			verb = "Float32()"
		case "float64" :
			verb = "Float64()"
		case "string" :
			verb = "String()"
	}
	return verb

}

func inArray(needle int, haystack []int) bool {
	
	for i:=0; i<len(haystack); i++ {
		if needle == haystack[i] {
			return true
		}
	}
	return false
	
}
