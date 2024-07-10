package dns

type SearchMode = string

const (
	LIKE  SearchMode = "LIKE"
	EXACT            = "EXACT"
)

type RecordTypes = string

const (
	A           RecordTypes = "A"
	CNAME                   = "CNAME"
	NS                      = "NS"
	MX                      = "MX"
	TXT                     = "TXT"
	SRV                     = "SRV"
	AAAA                    = "AAAA"
	REDIRECTURL             = "REDIRECT_URL"
	FORWARDURL              = "FORWARD_URL"
)

type RecordStatus = string

const (
	ENABLE  RecordStatus = "Enable"
	DISABLE              = "Disable"
)
