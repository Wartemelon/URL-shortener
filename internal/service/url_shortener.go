package service

type URLShortener interface {
	Shorten(url string) (string, error)
	Resolve(short string) (string, error)
}
