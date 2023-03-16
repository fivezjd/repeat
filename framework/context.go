/**
 * @Author: realpeanut
 * @Date: 2022/6/18 22:16
 */

package framework

import (
	"context"
	"encoding/json"
	"net/http"
	"sync"
)

type Context struct {
	request        *http.Request
	responseWriter http.ResponseWriter
	hasTimeout     bool
	writerMux      *sync.Mutex
	ctx            context.Context
}

func (ctx *Context) WriterMux() *sync.Mutex { return ctx.writerMux }

func (ctx *Context) BaseContext() context.Context {
	return ctx.request.Context()
}

func (ctx *Context) Done() <-chan struct{} {
	return ctx.BaseContext().Done()
}

func (ctx *Context) Json(status int, obj interface{}) error {
	if ctx.HasTimeout() {
		return nil
	}
	ctx.responseWriter.Header().Set("Content-Type", "application/json")
	ctx.responseWriter.WriteHeader(status)
	byt, err := json.Marshal(obj)
	if err != nil {
		ctx.responseWriter.WriteHeader(500)
		return err
	}
	ctx.responseWriter.Write(byt)
	return nil
}

func (ctx *Context) HasTimeout() bool {
	return ctx.hasTimeout
}

func (ctx *Context) SetHasTimeout() {
	ctx.hasTimeout = true
}

func NewContext(r *http.Request, w http.ResponseWriter) *Context {
	return &Context{
		request:        r,
		responseWriter: w,
		ctx:            r.Context(),
		writerMux:      &sync.Mutex{},
	}
}
