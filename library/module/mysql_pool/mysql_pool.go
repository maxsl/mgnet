package mysql_pool

import (
	"github.com/goodkele/mgnet/library/mysql"
)

type mysql_pool struct {
	
}

func NewMysqlPool() *mysql_pool {
	return new(mysql_pool)
}

func (this *mysql_pool) Get() * mysql.Connection{
	
}