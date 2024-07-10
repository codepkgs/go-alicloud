package dns

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
)

// DescribeDomainsRequestParams 查询域名列表时的可选参数
type DescribeDomainsRequestParams func(*alidns.DescribeDomainsRequest)

// DescribeDomainsWithKeyword 指定查询关键字
func (*DNS) DescribeDomainsWithKeyword(keyword string) DescribeDomainsRequestParams {
	return func(a *alidns.DescribeDomainsRequest) {
		a.KeyWord = keyword
	}
}

// DescribeDomainsWithSearchMode 指定搜索模式
func (*DNS) DescribeDomainsWithSearchMode(searchMode SearchMode) DescribeDomainsRequestParams {
	return func(a *alidns.DescribeDomainsRequest) {
		a.SearchMode = searchMode
	}
}

// DescribeDomains 查询域名列表
func (d *DNS) DescribeDomains(params ...DescribeDomainsRequestParams) ([]alidns.DomainInDescribeDomains, error) {
	request := alidns.CreateDescribeDomainsRequest()
	request.Scheme = "https"

	pageNumber := 1
	pageSize := 20

	domains := make([]alidns.DomainInDescribeDomains, 0)

	for {
		request.PageNumber = requests.NewInteger(pageNumber)
		request.PageSize = requests.NewInteger(pageSize)

		for _, param := range params {
			param(request)
		}

		resp, err := d.client.DescribeDomains(request)
		if err != nil {
			return domains, err
		}
		domains = append(domains, resp.Domains.Domain...)

		if resp.TotalCount <= int64(pageNumber*pageSize) {
			break
		}
		pageNumber++
	}

	return domains, nil
}

// DescribeDomainInfoRequestParams 查询域名信息的可选参数
type DescribeDomainInfoRequestParams func(*alidns.DescribeDomainInfoRequest)

// DescribeDomainInfoWithNeedDetailAttributes 查询域名信息时是否返回详细信息
func (*DNS) DescribeDomainInfoWithNeedDetailAttributes(needDetail bool) DescribeDomainInfoRequestParams {
	return func(a *alidns.DescribeDomainInfoRequest) {
		a.NeedDetailAttributes = requests.NewBoolean(needDetail)
	}
}

// DescribeDomainInfo 查询域名信息
func (d *DNS) DescribeDomainInfo(domainName string, params ...DescribeDomainInfoRequestParams) (response *alidns.DescribeDomainInfoResponse, err error) {
	request := alidns.CreateDescribeDomainInfoRequest()
	request.Scheme = "https"
	request.DomainName = domainName

	for _, param := range params {
		param(request)
	}

	return d.client.DescribeDomainInfo(request)
}

// DescribeDomainNS 获取域名的NS记录值
func (d *DNS) DescribeDomainNS(domainName string) (response *alidns.DescribeDomainNsResponse, err error) {
	request := alidns.CreateDescribeDomainNsRequest()
	request.Scheme = "https"
	request.DomainName = domainName

	return d.client.DescribeDomainNs(request)
}

// DescribeDomainGroups 查询所有域名分组
func (d *DNS) DescribeDomainGroups() ([]alidns.DomainGroup, error) {
	request := alidns.CreateDescribeDomainGroupsRequest()
	request.Scheme = "https"

	groups := make([]alidns.DomainGroup, 0)
	pageNumber := 1
	pageSize := 20
	for {
		request.PageNumber = requests.NewInteger(pageNumber)
		request.PageSize = requests.NewInteger(pageSize)
		resp, err := d.client.DescribeDomainGroups(request)
		if err != nil {
			return groups, err
		}
		groups = append(groups, resp.DomainGroups.DomainGroup...)
		if resp.TotalCount <= int64(pageNumber*pageSize) {
			break
		}
		pageNumber++
	}
	return groups, nil

}

// AddDomainGroup 添加域名分组
func (d *DNS) AddDomainGroup(groupName string) (response *alidns.AddDomainGroupResponse, err error) {
	request := alidns.CreateAddDomainGroupRequest()
	request.Scheme = "https"
	request.GroupName = groupName

	return d.client.AddDomainGroup(request)
}

// DeleteDomainGroup 删除域名分组，分组下的域名会被移动到默认分组
func (d *DNS) DeleteDomainGroup(groupId string) (response *alidns.DeleteDomainGroupResponse, err error) {
	request := alidns.CreateDeleteDomainGroupRequest()
	request.Scheme = "https"
	request.GroupId = groupId

	return d.client.DeleteDomainGroup(request)
}

// UpdateDomainGroup 修改域名分组
func (d *DNS) UpdateDomainGroup(groupName, groupId string) (response *alidns.UpdateDomainGroupResponse, err error) {
	request := alidns.CreateUpdateDomainGroupRequest()
	request.Scheme = "https"
	request.GroupId = groupId
	request.GroupName = groupName

	return d.client.UpdateDomainGroup(request)
}

// ChangeDomainGroup 将域名从原分组更换到新分组
func (d *DNS) ChangeDomainGroup(domainName, groupId string) (response *alidns.ChangeDomainGroupResponse, err error) {
	request := alidns.CreateChangeDomainGroupRequest()
	request.Scheme = "https"
	request.GroupId = groupId
	request.DomainName = domainName

	return d.client.ChangeDomainGroup(request)
}

// UpdateDomainRemark 修改域名备注
func (d *DNS) UpdateDomainRemark(domainName, remark string) (response *alidns.UpdateDomainRemarkResponse, err error) {
	request := alidns.CreateUpdateDomainRemarkRequest()
	request.Scheme = "https"
	request.DomainName = domainName
	request.Remark = remark

	return d.client.UpdateDomainRemark(request)
}
