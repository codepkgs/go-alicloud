package domain

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/domain"
)

// GetDomainInfo 查询域名信息
func (d *Domain) GetDomainInfo(domainName string) (response *domain.QueryDomainByDomainNameResponse, err error) {
	request := domain.CreateQueryDomainByDomainNameRequest()
	request.Scheme = "https"
	request.DomainName = domainName

	return d.client.QueryDomainByDomainName(request)
}
