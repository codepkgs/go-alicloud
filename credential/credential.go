package credential

import (
	alicredentials "github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
)

func NewCredentialWithAKSK(accessKeyId, accessKeySecret string) *alicredentials.AccessKeyCredential {
	return alicredentials.NewAccessKeyCredential(accessKeyId, accessKeySecret)
}

func NewCredentialWithStsToken(accessKeyId, accessKeySecret, securityToken string) *alicredentials.StsTokenCredential {
	return alicredentials.NewStsTokenCredential(accessKeyId, accessKeySecret, securityToken)
}
