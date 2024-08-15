package internal

type Data struct {
	key   string
	value any
}

type HashMapper interface {
	Set(key string, value any)
	Get(key string) (any, bool)
}
