package postgres


var maxIdleConns = 0
var maxOpenConns = 0

func SetMaxIdleConns(conns int) {
	maxIdleConns = conns
	Database.SetMaxIdleConns(maxIdleConns)
}

func SetMaxOpenConns(conns int) {
	maxOpenConns = conns
	Database.SetMaxOpenConns(maxOpenConns)
}
