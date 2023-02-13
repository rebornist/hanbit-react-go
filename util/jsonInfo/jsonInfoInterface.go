package jsoninfo

// ConfigInterface 인터페이스 선언
type JSONConfigInterface interface {
	GetJSONInfo(name string) ([]byte, error)
}
