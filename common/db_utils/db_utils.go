package dbutils

func GetDbDns(
	host string,
	port string,
	dbName string,
	user string,
	password string,
	ssl string,
) string {
	return "host=" + host + " port=" + port + " dbname=" + dbName + " user=" + user + " password=" + password + " sslmode=" + ssl
}
