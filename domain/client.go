package domain

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth"
	alidomain "github.com/aliyun/alibaba-cloud-sdk-go/services/domain"
	"github.com/codepkgs/goalicloud/sdkconfig"
)

type Domain struct {
	client *alidomain.Client
}

func NewClient(region string, credential auth.Credential, options ...sdkconfig.ConfigOption) (*Domain, error) {
	config := sdk.NewConfig()
	for _, option := range options {
		option(config)
	}
	client, err := alidomain.NewClientWithOptions(region, config, credential)
	return &Domain{client: client}, err
}
