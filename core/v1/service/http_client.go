package service

type HttpClient interface {
	Get(url string,header map[string]string)(error, []byte)
}
