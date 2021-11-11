package service

// HttpClient HttpClient related operations.
type HttpClient interface {
	Get(url string, header map[string]string) (error, []byte)
	Post(url string, header map[string]string, body []byte) (error, []byte)
	Delete(url string, header map[string]string) error
}
