package memory

import (
	"git.code.oa.com/ad/go-project-base/config/loader"
	"git.code.oa.com/ad/go-project-base/config/reader"
	"git.code.oa.com/ad/go-project-base/config/source"
)

// WithSource appends a source to list of sources
func WithSource(s source.Source) loader.Option {
	return func(o *loader.Options) {
		o.Source = append(o.Source, s)
	}
}

// WithReader sets the config reader
func WithReader(r reader.Reader) loader.Option {
	return func(o *loader.Options) {
		o.Reader = r
	}
}
