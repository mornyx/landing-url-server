package handlers

import "net/http"

var _ http.ResponseWriter = &mockResponseWriter{}

type mockResponseWriter struct {
	statusCode int
	header     map[string][]string
	body       []byte
}

func newMockResponseWriter() *mockResponseWriter {
	return &mockResponseWriter{
		header: map[string][]string{},
	}
}

func (m *mockResponseWriter) Header() http.Header {
	return m.header
}

func (m *mockResponseWriter) Write(buf []byte) (int, error) {
	m.body = buf
	return len(buf), nil
}

func (m *mockResponseWriter) WriteHeader(statusCode int) {
	m.statusCode = statusCode
}
