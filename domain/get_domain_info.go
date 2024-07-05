package domain

import (
	alidomain "github.com/aliyun/alibaba-cloud-sdk-go/services/domain"
)

func (d *Domain) GetDomainInfo(domainName string) (response *alidomain.QueryDomainByDomainNameResponse, err error) {
	request := alidomain.CreateQueryDomainByDomainNameRequest()
	request.Scheme = "https"
	request.DomainName = domainName

	return d.client.QueryDomainByDomainName(request)
}
