package jsoninfo

// ConfigInterface 인터페이스 선언
type JSONConfigService interface {
	GetJSONInfo(name string) ([]byte, error)
}
