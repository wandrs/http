package http

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"time"

	"github.com/unrolled/render"
)

type ResponseWriter interface {
	http.ResponseWriter
	R() *http.Request

	// render functions
	TemplateLookup(t string) *template.Template
	Render(e render.Engine, data interface{})
	Data(status int, v []byte)
	HTML(status int, name string, binding interface{}, htmlOpt ...render.HTMLOptions)
	JSON(status int, v interface{})
	JSONP(status int, callback string, v interface{})
	Text(status int, v string)
	XML(status int, v interface{})
	Error(status int, contents ...string)

	// url path parameters
	Params(p string) string
	ParamsInt(p string) int
	ParamsInt64(p string) int64
	ParamsFloat64(name string) float64
	SetParams(k, v string)

	// query parameters
	Query(key string, defaults ...string) string
	QueryTrim(key string, defaults ...string) string
	QueryStrings(key string, defaults ...[]string) []string
	QueryInt(key string, defaults ...int) int
	QueryInt64(key string, defaults ...int64) int64
	QueryBool(key string, defaults ...bool) bool

	// request.Context() methods
	Deadline() (deadline time.Time, ok bool)
	Done() <-chan struct{}
	Err() error
	Value(key interface{}) interface{}

	// misc render/response functions
	Redirect(location string, status ...int)
	RedirectToFirst(appURL, appSubURL string, location ...string)
	HTMLString(name string, binding interface{}, htmlOpt ...render.HTMLOptions) (string, error)
	ServeContent(name string, r io.ReadSeeker, params ...interface{})
	ServeFile(file string, names ...string)
}

type response struct {
	http.ResponseWriter
	req *http.Request
	r   *render.Render
}

var _ ResponseWriter = &response{}

func NewResponseWriter(w http.ResponseWriter, req *http.Request, r *render.Render) ResponseWriter {
	return &response{
		ResponseWriter: w,
		req:            req,
		r:              r,
	}
}

func (w *response) R() *http.Request {
	return w.req
}

func (w *response) TemplateLookup(t string) *template.Template {
	return w.r.TemplateLookup(t)
}

func (w *response) Render(e render.Engine, data interface{}) {
	if err := w.r.Render(w, e, data); err != nil {
		http.Error(w, fmt.Sprintf("Render failed, reason: %v", err), http.StatusInternalServerError)
	}
}

func (w *response) Data(status int, v []byte) {
	if err := w.r.Data(w, status, v); err != nil {
		http.Error(w, fmt.Sprintf("Render failed, reason: %v", err), http.StatusInternalServerError)
	}
}

func (w *response) HTML(status int, name string, binding interface{}, htmlOpt ...render.HTMLOptions) {
	if err := w.r.HTML(w, status, name, binding, htmlOpt...); err != nil {
		http.Error(w, fmt.Sprintf("Render failed, reason: %v", err), http.StatusInternalServerError)
	}
}

func (w *response) JSON(status int, v interface{}) {
	if err := w.r.JSON(w, status, v); err != nil {
		http.Error(w, fmt.Sprintf("Render failed, reason: %v", err), http.StatusInternalServerError)
	}
}

func (w *response) JSONP(status int, callback string, v interface{}) {
	if err := w.r.JSONP(w, status, callback, v); err != nil {
		http.Error(w, fmt.Sprintf("Render failed, reason: %v", err), http.StatusInternalServerError)
	}
}

func (w *response) Text(status int, v string) {
	if err := w.r.Text(w, status, v); err != nil {
		http.Error(w, fmt.Sprintf("Render failed, reason: %v", err), http.StatusInternalServerError)
	}
}

func (w *response) XML(status int, v interface{}) {
	if err := w.r.XML(w, status, v); err != nil {
		http.Error(w, fmt.Sprintf("Render failed, reason: %v", err), http.StatusInternalServerError)
	}
}

func (w *response) Error(status int, contents ...string) {
	var v = http.StatusText(status)
	if len(contents) > 0 {
		v = contents[0]
	}
	http.Error(w, v, status)
}
