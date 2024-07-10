package dns

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
)

type DescribeDomainRecordsRequestParams func(request *alidns.DescribeDomainRecordsRequest)

// DescribeDomainRecordsWithKeyword 根据关键字搜索
func (*DNS) DescribeDomainRecordsWithKeyword(keyword string) DescribeDomainRecordsRequestParams {
	return func(a *alidns.DescribeDomainRecordsRequest) {
		a.KeyWord = keyword
	}
}

// DescribeDomainRecordsWithSearchMode 搜索模式
func (*DNS) DescribeDomainRecordsWithSearchMode(searchMode SearchMode) DescribeDomainRecordsRequestParams {
	return func(a *alidns.DescribeDomainRecordsRequest) {
		a.SearchMode = searchMode
	}
}

// DescribeDomainRecordsWithRRKeyword 根据主机记录的关键字进行搜索
func (*DNS) DescribeDomainRecordsWithRRKeyword(rrKeyword string) DescribeDomainRecordsRequestParams {
	return func(a *alidns.DescribeDomainRecordsRequest) {
		a.RRKeyWord = rrKeyword
	}
}

// DescribeDomainRecordsWithValueKeyword 根据记录值的关键字搜索
func (*DNS) DescribeDomainRecordsWithValueKeyword(valueKeyword string) DescribeDomainRecordsRequestParams {
	return func(a *alidns.DescribeDomainRecordsRequest) {
		a.ValueKeyWord = valueKeyword
	}
}

// DescribeDomainRecordsWithRecordType 根据记录类型搜索
func (*DNS) DescribeDomainRecordsWithRecordType(recordType RecordTypes) DescribeDomainRecordsRequestParams {
	return func(a *alidns.DescribeDomainRecordsRequest) {
		a.Type = recordType
	}
}

// DescribeDomainRecordsWithRecordLine 根据线路搜索
// 可以用 DescribeSupportLines 获取支持的线路列表
func (*DNS) DescribeDomainRecordsWithRecordLine(line string) DescribeDomainRecordsRequestParams {
	return func(a *alidns.DescribeDomainRecordsRequest) {
		a.Line = line
	}
}

// DescribeDomainRecordsWithRecordStatus 根据记录状态搜索
func (*DNS) DescribeDomainRecordsWithRecordStatus(status RecordStatus) DescribeDomainRecordsRequestParams {
	return func(a *alidns.DescribeDomainRecordsRequest) {
		a.Status = status
	}
}

// DescribeDomainRecords 根据传入参数获取指定主域名的所有解析记录列表
func (d *DNS) DescribeDomainRecords(domainName string, params ...DescribeDomainRecordsRequestParams) ([]alidns.Record, error) {
	request := alidns.CreateDescribeDomainRecordsRequest()
	request.Scheme = "https"
	request.DomainName = domainName

	pageNumber := 1
	pageSize := 100

	records := make([]alidns.Record, 0)

	for {
		request.PageNumber = requests.NewInteger(pageNumber)
		request.PageSize = requests.NewInteger(pageSize)

		for _, param := range params {
			param(request)
		}

		resp, err := d.Client.DescribeDomainRecords(request)
		if err != nil {
			return records, nil
		}

		records = append(records, resp.DomainRecords.Record...)
		if resp.TotalCount <= int64(pageNumber*pageSize) {
			break
		}
		pageNumber++
	}
	return records, nil
}

// SetDomainRecordStatus 设置解析记录状态
func (d *DNS) SetDomainRecordStatus(recordId string, status RecordStatus) (response *alidns.SetDomainRecordStatusResponse, err error) {
	request := alidns.CreateSetDomainRecordStatusRequest()
	request.Scheme = "https"
	request.RecordId = recordId
	request.Status = status

	return d.Client.SetDomainRecordStatus(request)
}

// AddDomainRecordRequestParams 添加主机记录相关的可选参数
type AddDomainRecordRequestParams func(*alidns.AddDomainRecordRequest)

func (*DNS) AddDomainRecordWithTTL(ttl int) AddDomainRecordRequestParams {
	return func(a *alidns.AddDomainRecordRequest) {
		a.TTL = requests.NewInteger(ttl)
	}
}

func (*DNS) AddDomainRecordWithMxPriority(priority int) AddDomainRecordRequestParams {
	return func(a *alidns.AddDomainRecordRequest) {
		a.Priority = requests.NewInteger(priority)
	}
}

func (*DNS) AddDomainRecordWithLine(line string) AddDomainRecordRequestParams {
	return func(a *alidns.AddDomainRecordRequest) {
		a.Line = line
	}
}

// AddDomainRecord 添加主机记录
func (d *DNS) AddDomainRecord(domainName string, rr string, recordType RecordTypes, value string, params ...AddDomainRecordRequestParams) (response *alidns.AddDomainRecordResponse, err error) {
	request := alidns.CreateAddDomainRecordRequest()
	request.Scheme = "https"
	request.DomainName = domainName
	request.RR = rr
	request.Value = value
	request.Type = recordType

	for _, param := range params {
		param(request)
	}

	return d.Client.AddDomainRecord(request)
}

// DeleteDomainRecord 删除主机记录
func (d *DNS) DeleteDomainRecord(recordId string) (response *alidns.DeleteDomainRecordResponse, err error) {
	request := alidns.CreateDeleteDomainRecordRequest()
	request.Scheme = "https"
	request.RecordId = recordId
	return d.Client.DeleteDomainRecord(request)
}

// UpdateDomainRecordRequestParams 更新主机记录相关的可选参数
type UpdateDomainRecordRequestParams func(*alidns.UpdateDomainRecordRequest)

func (*DNS) UpdateDomainRecordWithTTL(ttl int) UpdateDomainRecordRequestParams {
	return func(a *alidns.UpdateDomainRecordRequest) {
		a.TTL = requests.NewInteger(ttl)
	}
}

func (*DNS) UpdateDomainRecordWithMxPriority(priority int) UpdateDomainRecordRequestParams {
	return func(a *alidns.UpdateDomainRecordRequest) {
		a.Priority = requests.NewInteger(priority)
	}
}

func (*DNS) UpdateDomainRecordWithLine(line string) AddDomainRecordRequestParams {
	return func(a *alidns.AddDomainRecordRequest) {
		a.Line = line
	}
}

// UpdateDomainRecord 修改解析记录
func (d *DNS) UpdateDomainRecord(recordId string, rr string, recordType RecordTypes, value string, params ...UpdateDomainRecordRequestParams) (response *alidns.UpdateDomainRecordResponse, err error) {
	request := alidns.CreateUpdateDomainRecordRequest()
	request.Scheme = "https"
	request.RecordId = recordId
	request.RR = rr
	request.Value = value
	request.Type = recordType

	for _, param := range params {
		param(request)
	}

	return d.Client.UpdateDomainRecord(request)
}

// UpdateDomainRecordRemark 修改解析记录的备注
func (d *DNS) UpdateDomainRecordRemark(recordId, remark string) (response *alidns.UpdateDomainRecordRemarkResponse, err error) {
	request := alidns.CreateUpdateDomainRecordRemarkRequest()
	request.Scheme = "https"
	request.RecordId = recordId
	request.Remark = remark
	return d.Client.UpdateDomainRecordRemark(request)
}
