package gmongo

type Option interface {
	apply(op *options)
}

type options struct {
	dbName, url string
}

type opFunc func(op *options)

func (o opFunc) apply(op *options) {
	o(op)
}

func WithDbName(dbname string) Option {
	return opFunc(func(op *options) {
		op.dbName = dbname
	})
}

func WithUrl(url string) Option {
	return opFunc(func(op *options) {
		op.url = url
	})
}
