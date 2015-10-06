package orm


type OrmJson struct {
	Key string
	Value string
}

type OrmConfig struct {
	Table string			// 表名
	Columns []string		// 字段名
	ColumnsType []string	// 字段类型
	PrimaryKey []int		// 主键字段
	JsonColumns map[int]int	// json字段
	JsonColumnsType map[int][]OrmJson	// json字段类型
}

var (
	Config []OrmConfig = []OrmConfig{
		// 用户表
		{
			"user",
			[]string{"UserId", "PtId", "Name", "Desc"},
			[]string{"int64", "string", "string", "string"},
			[]int{0},
			map[int]int{3:1024},
			map[int][]OrmJson{
				3 : []OrmJson{
						{"Name", "string"}, 
						{"Age", "int"} }}}}
)