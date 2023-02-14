package dblayer

// Connection info for the DB
type DB struct {
	cate string
}

// DB Connection 초기화
func NewDBConnection(dbCate string) (DBService, error) {
	return &DB{cate: dbCate}, nil
}
