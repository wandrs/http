package http

import "time"

// Deadline is part of the interface for context.Context and we pass this to the request context
func (w *response) Deadline() (deadline time.Time, ok bool) {
	return w.req.Context().Deadline()
}

// Done is part of the interface for context.Context and we pass this to the request context
func (w *response) Done() <-chan struct{} {
	return w.req.Context().Done()
}

// Err is part of the interface for context.Context and we pass this to the request context
func (w *response) Err() error {
	return w.req.Context().Err()
}

// Value is part of the interface for context.Context and we pass this to the request context
func (w *response) Value(key interface{}) interface{} {
	return w.req.Context().Value(key)
}
