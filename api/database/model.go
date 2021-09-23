package database

type SqliteTable struct {
	Type     string
	Name     string
	TblName  string
	Rootpage string
	Sql      string
}
type MysqlTable struct {
	TABLE_CATALOG   string
	TABLE_SCHEMA    string
	TABLE_NAME      string
	TABLE_TYPE      string
	ENGINE          string
	VERSION         string
	ROW_FORMAT      string
	TABLE_ROWS      string
	AVG_ROW_LENGTH  string
	DATA_LENGTH     string
	MAX_DATA_LENGTH string
	INDEX_LENGTH    string
	DATA_FREE       string
	AUTO_INCREMENT  string
	CREATE_TIME     string
	UPDATE_TIME     string
	CHECK_TIME      string
	TABLE_COLLATION string
	CHECKSUM        string
	CREATE_OPTIONS  string
	TABLE_COMMENT   string
}
type Table struct {
}
