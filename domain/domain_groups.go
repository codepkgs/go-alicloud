package domain

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/domain"
)

type GetDomainGroupsParams func(*domain.QueryDomainGroupListRequest)

// GetDomainGroupsWithDomainGroupName 根据域名分组名称查询
func GetDomainGroupsWithDomainGroupName(domainGroupName string) GetDomainGroupsParams {
	return func(request *domain.QueryDomainGroupListRequest) {
		request.DomainGroupName = domainGroupName
	}
}

// GetDomainGroupsWithShowDeletingGroup 是否展示正在删除中的域名分组
func GetDomainGroupsWithShowDeletingGroup(showDeleting bool) GetDomainGroupsParams {
	return func(request *domain.QueryDomainGroupListRequest) {
		request.ShowDeletingGroup = requests.NewBoolean(showDeleting)
	}
}

// GetDomainGroups 查询域名分组列表
func (d *Domain) GetDomainGroups(params ...GetDomainGroupsParams) (*domain.QueryDomainGroupListResponse, error) {
	request := domain.CreateQueryDomainGroupListRequest()
	request.Scheme = "https"

	for _, param := range params {
		param(request)
	}

	return d.client.QueryDomainGroupList(request)
}

// SaveDomainGroupParams 保存域名分组的参数
type SaveDomainGroupParams func(*domain.SaveDomainGroupRequest)

// SaveDomainGroupWithDomainGroupId 保存域名分组。如果有该参数则修改该分组，如果没有该参数则新增该分组
func SaveDomainGroupWithDomainGroupId(domainGroupId int) SaveDomainGroupParams {
	return func(request *domain.SaveDomainGroupRequest) {
		request.DomainGroupId = requests.NewInteger(domainGroupId)
	}
}

// SaveDomainGroup 保存域名分组。如果有 SaveDomainGroupWithDomainGroupId 则修改分组，否则创建分组
func (d *Domain) SaveDomainGroup(domainGroupName string, params ...SaveDomainGroupParams) (*domain.SaveDomainGroupResponse, error) {
	request := domain.CreateSaveDomainGroupRequest()
	request.Scheme = "https"
	request.DomainGroupName = domainGroupName

	for _, param := range params {
		param(request)
	}

	return d.client.SaveDomainGroup(request)
}

// DeleteDomainGroup 删除域名分组
func (d *Domain) DeleteDomainGroup(domainGroupId int) (*domain.DeleteDomainGroupResponse, error) {
	request := domain.CreateDeleteDomainGroupRequest()
	request.Scheme = "https"
	request.DomainGroupId = requests.NewInteger(domainGroupId)

	return d.client.DeleteDomainGroup(request)
}

// UpdateDomainToDomainGroup 向分组中设置域名
// replace 否替换当前分组中域名
func (d *Domain) UpdateDomainToDomainGroup(domainGroupId int, replace bool, domainNames ...string) (*domain.UpdateDomainToDomainGroupResponse, error) {
	request := domain.CreateUpdateDomainToDomainGroupRequest()
	request.Scheme = "https"
	request.DomainGroupId = requests.NewInteger(domainGroupId)
	request.Replace = requests.NewBoolean(replace)
	request.DataSource = requests.NewInteger(1)
	request.DomainName = &domainNames

	return d.client.UpdateDomainToDomainGroup(request)
}
