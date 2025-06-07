package context

import "context"

var ctx context.Context

func SetContext(c context.Context) {
	ctx = c
}

func GetContext() context.Context {
	return ctx
}
