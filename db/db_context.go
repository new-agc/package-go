package agcdb

const (
	//DB_HOST       = "tcp(127.0.0.1:3306)"
	DB_HOST       = "tcp(10.0.7.200:3306)"
	DB_NAME       = "agc_test"
	DB_USER       = "root"
	DB_PASS       = "ketilinux"
	DB_CONNECTION = DB_USER + ":" + DB_PASS + "@" + DB_HOST + "/" + DB_NAME + "?charset=utf8"
)