package pgconnector

import "io"

type PgSinkReader struct {
	io.PipeReader
}

type PgSinkWriter struct {
	io.PipeWriter
}
