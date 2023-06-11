package service

type URLShortener interface {
	Create(sourceURL string) (int, error)
	Get(shortURL string) error
}
