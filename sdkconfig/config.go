package sdkconfig

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"net/http"
	"time"
)

type ConfigOption func(*sdk.Config)

func WithScheme(scheme string) ConfigOption {
	return func(config *sdk.Config) {
		config.Scheme = scheme
	}
}

func WithAutoRetry(isAutoRetry bool) ConfigOption {
	return func(config *sdk.Config) {
		config.WithAutoRetry(isAutoRetry)
	}
}

func WithMaxRetryTime(maxRetryTime int) ConfigOption {
	return func(config *sdk.Config) {
		config.WithMaxRetryTime(maxRetryTime)
	}
}

func WithUserAgent(userAgent string) ConfigOption {
	return func(config *sdk.Config) {
		config.WithUserAgent(userAgent)
	}
}

func WithDebug(isDebug bool) ConfigOption {
	return func(config *sdk.Config) {
		config.WithDebug(isDebug)
	}
}

func WithTimeout(timeout time.Duration) ConfigOption {
	return func(config *sdk.Config) {
		config.WithTimeout(timeout)
	}
}

func WithHttpTransport(httpTransport *http.Transport) ConfigOption {
	return func(config *sdk.Config) {
		config.WithHttpTransport(httpTransport)
	}
}

func WithEnableAsync(isEnableAsync bool) ConfigOption {
	return func(config *sdk.Config) {
		config.WithEnableAsync(isEnableAsync)
	}
}

func WithMaxTaskQueueSize(maxTaskQueueSize int) ConfigOption {
	return func(config *sdk.Config) {
		config.WithMaxTaskQueueSize(maxTaskQueueSize)
	}
}

func WithGoRoutinePoolSize(goRoutinePoolSize int) ConfigOption {
	return func(config *sdk.Config) {
		config.WithGoRoutinePoolSize(goRoutinePoolSize)
	}
}
