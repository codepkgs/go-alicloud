package dns

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
	"github.com/codepkgs/goalicloud/sdkconfig"
)

type DNS struct {
	client *alidns.Client
}

func NewClient(region string, credential auth.Credential, options ...sdkconfig.ConfigOption) (*DNS, error) {
	config := sdk.NewConfig()
	for _, option := range options {
		option(config)
	}
	client, err := alidns.NewClientWithOptions(region, config, credential)
	return &DNS{client: client}, err
}
