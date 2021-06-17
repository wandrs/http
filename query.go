package http

// FIXME: We should differ Query and Form, currently we just use form as query
// Currently to be compatible with macaron, we keep it.

// Query returns request form as string with default
func (w *response) Query(key string, defaults ...string) string {
	return (*Forms)(w.req).MustString(key, defaults...)
}

// QueryTrim returns request form as string with default and trimmed spaces
func (w *response) QueryTrim(key string, defaults ...string) string {
	return (*Forms)(w.req).MustTrimmed(key, defaults...)
}

// QueryStrings returns request form as strings with default
func (w *response) QueryStrings(key string, defaults ...[]string) []string {
	return (*Forms)(w.req).MustStrings(key, defaults...)
}

// QueryInt returns request form as int with default
func (w *response) QueryInt(key string, defaults ...int) int {
	return (*Forms)(w.req).MustInt(key, defaults...)
}

// QueryInt64 returns request form as int64 with default
func (w *response) QueryInt64(key string, defaults ...int64) int64 {
	return (*Forms)(w.req).MustInt64(key, defaults...)
}

// QueryBool returns request form as bool with default
func (w *response) QueryBool(key string, defaults ...bool) bool {
	return (*Forms)(w.req).MustBool(key, defaults...)
}
