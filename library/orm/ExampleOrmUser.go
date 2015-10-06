package orm

import (
	"github.com/goodkele/mgnet/library/mysql"
	"encoding/json"
	"fmt"
)

type ExampleUserDesc struct {
	Name string
	Age int
}

type ExampleOrmUser struct {
	UserId int64
	PtId string
	Name string
	Desc ExampleUserDesc

	isLoad bool	// 是否已载入
}

// 加载
func (this *ExampleOrmUser) LoadDb(conn *mysql.Connection) {

	res,err := conn.QueryTable(fmt.Sprintf("SELECT `UserId`, `PtId`, `Name`, `Desc` FROM `user` WHERE UserId='%d'", this.UserId))
	if err != nil {
	}
	rows := res.Rows()
	if len(rows) > 0 {
		this.UserId = rows[0][0].Int64()
		this.PtId = rows[0][1].String()
		this.Name = rows[0][2].String()
		err = json.Unmarshal([]byte(rows[0][3].String()), &this.Desc)
		if err != nil {
		}

		this.isLoad = true
	}

}

// 更新
func (this *ExampleOrmUser) SaveDb(conn *mysql.Connection) {

	jsonDesc, err := json.Marshal(this.Desc)
	if err != nil {
	}

	// fixme 检测长度报警
	_,err = conn.Execute(fmt.Sprintf("UPDATE `user` SET  `PtId`='%s', `Name`='%s', `Desc`='%s' WHERE `UserId`='%d'", this.PtId, this.Name, jsonDesc, this.UserId))
	if err != nil {
	}
}

// 新增
func (this *ExampleOrmUser) Insert(conn *mysql.Connection) {

	jsonDesc, err := json.Marshal(this.Desc)
	if err != nil {
	}

	// fixme 检测长度报警
	_,err = conn.Execute(fmt.Sprintf("INSERT INTO `user`(`UserId`, `PtId`, `Name`, `Desc`) VALUES('%d', '%s', '%s', '%s')", this.UserId, this.PtId, this.Name, jsonDesc))
	if err != nil {
	}

	this.isLoad = true

}

// 删除
func (this *ExampleOrmUser) Delete(conn *mysql.Connection) {

	_,err := conn.Execute(fmt.Sprintf("DELETE FROM `user` WHERE UserId='%d'", this.UserId))
	if err != nil {
	}

}
