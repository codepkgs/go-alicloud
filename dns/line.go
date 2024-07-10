package dns

import "github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"

// DescribeSupportLines 查询支持的线路
func (d *DNS) DescribeSupportLines(domainName string) (*alidns.DescribeSupportLinesResponse, error) {
	request := alidns.CreateDescribeSupportLinesRequest()
	request.Scheme = "https"
	request.DomainName = domainName

	return d.client.DescribeSupportLines(request)
}
