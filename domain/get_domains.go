package domain

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	alidomain "github.com/aliyun/alibaba-cloud-sdk-go/services/domain"
)

type GetDomainsRequestParams func(*alidomain.QueryDomainListRequest)

func (d *Domain) GetDomainsWithDomainName(domainName string) GetDomainsRequestParams {
	return func(d *alidomain.QueryDomainListRequest) {
		d.DomainName = domainName
	}
}

func (d *Domain) GetDomainsWithDomainState(state State) GetDomainsRequestParams {
	return func(d *alidomain.QueryDomainListRequest) {
		d.QueryType = state
	}
}

func (d *Domain) GetDomains(params ...GetDomainsRequestParams) ([]alidomain.Domain, error) {
	request := alidomain.CreateQueryDomainListRequest()
	request.Scheme = "https"

	pageNumber := 1
	pageSize := 20

	domains := make([]alidomain.Domain, 0)

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
