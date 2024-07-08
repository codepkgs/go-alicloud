package domain

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/domain"
)

// GetDomainsRequestParams GetDomains参数
type GetDomainsRequestParams func(*domain.QueryDomainListRequest)

// GetDomainsWithDomainName 根据域名查询域名列表
func (d *Domain) GetDomainsWithDomainName(domainName string) GetDomainsRequestParams {
	return func(d *domain.QueryDomainListRequest) {
		d.DomainName = domainName
	}
}

// GetDomainsWithDomainState 根据域名状态查询域名列表
func (d *Domain) GetDomainsWithDomainState(state State) GetDomainsRequestParams {
	return func(d *domain.QueryDomainListRequest) {
		d.QueryType = state
	}
}

// GetDomains 根据GetDomainsRequestParams参数获取域名列表
func (d *Domain) GetDomains(params ...GetDomainsRequestParams) ([]domain.Domain, error) {
	request := domain.CreateQueryDomainListRequest()
	request.Scheme = "https"

	pageNumber := 1
	pageSize := 20

	domains := make([]domain.Domain, 0)

	for {
		request.PageNum = requests.NewInteger(pageNumber)
		request.PageSize = requests.NewInteger(pageSize)

		for _, param := range params {
			param(request)
		}

		resp, err := d.client.QueryDomainList(request)
		if err != nil {
			return domains, err
		}
		domains = append(domains, resp.Data.Domain...)

		if resp.TotalPageNum == pageNumber {
			break
		}
		pageNumber++
	}
	return domains, nil
}
