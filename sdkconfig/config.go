package sdkconfig

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"net/http"
	"time"
)

type Option func(*sdk.Config)

func WithScheme(scheme string) Option {
	return func(config *sdk.Config) {
		config.Scheme = scheme
	}
}

func WithAutoRetry(isAutoRetry bool) Option {
	return func(config *sdk.Config) {
		config.WithAutoRetry(isAutoRetry)
	}
}

func WithMaxRetryTime(maxRetryTime int) Option {
	return func(config *sdk.Config) {
		config.WithMaxRetryTime(maxRetryTime)
	}
}

func WithUserAgent(userAgent string) Option {
	return func(config *sdk.Config) {
		config.WithUserAgent(userAgent)
	}
}

func WithDebug(isDebug bool) Option {
	return func(config *sdk.Config) {
		config.WithDebug(isDebug)
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(config *sdk.Config) {
		config.WithTimeout(timeout)
	}
}

func WithHttpTransport(httpTransport *http.Transport) Option {
	return func(config *sdk.Config) {
		config.WithHttpTransport(httpTransport)
	}
}

func WithEnableAsync(isEnableAsync bool) Option {
	return func(config *sdk.Config) {
		config.WithEnableAsync(isEnableAsync)
	}
}

func WithMaxTaskQueueSize(maxTaskQueueSize int) Option {
	return func(config *sdk.Config) {
		config.WithMaxTaskQueueSize(maxTaskQueueSize)
	}
}

func WithGoRoutinePoolSize(goRoutinePoolSize int) Option {
	return func(config *sdk.Config) {
		config.WithGoRoutinePoolSize(goRoutinePoolSize)
	}
}
