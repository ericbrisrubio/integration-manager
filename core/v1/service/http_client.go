package service

// HttpClient HttpClient related operations.
type HttpClient interface {
	Get(url string, header map[string]string) ([]byte, error)
	Post(url string, header map[string]string, body []byte) ([]byte, error)
	Delete(url string, header map[string]string) error
	Put(url string, header map[string]string, body []byte) ([]byte, error)
}
