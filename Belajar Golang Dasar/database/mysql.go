package database

var connection string
func init() {
	connection="MySQL Database"
}	
func GetDatabase() string {
	return connection
}