package service

type Repository interface {
	Save(short, original string) error
	Get(short string) (string, error)
}