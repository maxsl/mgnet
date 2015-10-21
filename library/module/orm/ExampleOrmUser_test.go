package orm

import (
	"github.com/goodkele/mgnet/library/mysql"	
	"os"
	"strconv"
	"testing"
)

var TestConnParam mysql.ConnectionParams

func env(key, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	return v
}

func init() {
	TestConnParam.Host = env("TEST_MYSQL_HOST", "192.168.1.21")
	TestConnParam.Port, _ = strconv.Atoi(env("TEST_MYSQL_PORT", "3306"))
	TestConnParam.DbName = env("TEST_MYSQL_DBNAME", "mgnet")
	TestConnParam.Uname = env("TEST_MYSQL_UNAME", "root")
	TestConnParam.Pass = env("TEST_MYSQL_PASS", "123456")
}

func Test_All(t *testing.T) {
	param := TestConnParam

	conn, err := mysql.Connect(param)
	if err != nil {
		t.Errorf("Connect failed. err %d .", err)
	}

	user := &ExampleOrmUser{1, "pt1", "name1", ExampleUserDesc{"DescName1", 1}, false}
	user.Insert(conn)
	
	user2 := &ExampleOrmUser{2, "pt2", "name2", ExampleUserDesc{"DescName2", 2}, false}
	user2.Insert(conn)

	user3 := &ExampleOrmUser{3, "pt3", "name3", ExampleUserDesc{"DescName3", 3}, false}
	user3.Insert(conn)
	
	user4 := &ExampleOrmUser{1, "", "", ExampleUserDesc{"", 0}, false}
	user4.LoadDb(conn)
	user4.Name = "SaveDb"
	user4.SaveDb(conn)
	
	user5 := &ExampleOrmUser{3, "", "", ExampleUserDesc{"", 0}, false}
	user5.LoadDb(conn)
	user3.Delete(conn)
	
	conn.Close()
}


