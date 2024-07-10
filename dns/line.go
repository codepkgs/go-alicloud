package dns

import "github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"

func (d *DNS) DescribeSupportLines(domainName string) (response *alidns.DescribeSupportLinesResponse, err error) {
	request := alidns.CreateDescribeSupportLinesRequest()
	request.Scheme = "https"
	request.DomainName = domainName

	return d.Client.DescribeSupportLines(request)
}
