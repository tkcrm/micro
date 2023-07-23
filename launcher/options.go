package launcher

import (
	"context"

	"github.com/tkcrm/mx/logger"
)

type Option func(*Options)

type Options struct {
	logger logger.Logger

	Name    string
	Version string

	// Before and After funcs
	BeforeStart []func() error
	BeforeStop  []func() error
	AfterStart  []func() error
	AfterStop   []func() error

	Signal bool

	Context context.Context
}

func newOptions(opts ...Option) Options {
	opt := Options{
		logger: logger.New(),

		BeforeStart: make([]func() error, 0),
		BeforeStop:  make([]func() error, 0),
		AfterStart:  make([]func() error, 0),
		AfterStop:   make([]func() error, 0),

		Signal: true,

		Context: context.Background(),
	}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}

// Name of the launcher
func WithName(n string) Option {
	return func(o *Options) {
		o.Name = n
	}
}

// Version of the launcher
func WithVersion(v string) Option {
	return func(o *Options) {
		o.Version = v
	}
}

func WithContext(ctx context.Context) Option {
	return func(o *Options) {
		o.Context = ctx
	}
}

func WithSignal(b bool) Option {
	return func(o *Options) {
		o.Signal = b
	}
}

func WithLogger(l logger.Logger) Option {
	return func(o *Options) {
		o.logger = l
	}
}

// Before and Afters

// WithBeforeStart run funcs before service starts
func WithBeforeStart(fn func() error) Option {
	return func(o *Options) {
		o.BeforeStart = append(o.BeforeStart, fn)
	}
}

// WithBeforeStop run funcs before service stops
func WithBeforeStop(fn func() error) Option {
	return func(o *Options) {
		o.BeforeStop = append(o.BeforeStop, fn)
	}
}

// WithAfterStart run funcs after service starts
func WithAfterStart(fn func() error) Option {
	return func(o *Options) {
		o.AfterStart = append(o.AfterStart, fn)
	}
}

// WithAfterStop run funcs after service stops
func WithAfterStop(fn func() error) Option {
	return func(o *Options) {
		o.AfterStop = append(o.AfterStop, fn)
	}
}
