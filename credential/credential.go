package credential

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
)

func NewCredentialWithAKSK(accessKeyId, accessKeySecret string) *credentials.AccessKeyCredential {
	return credentials.NewAccessKeyCredential(accessKeyId, accessKeySecret)
}

func NewCredentialWithStsToken(accessKeyId, accessKeySecret, securityToken string) *credentials.StsTokenCredential {
	return credentials.NewStsTokenCredential(accessKeyId, accessKeySecret, securityToken)
}
