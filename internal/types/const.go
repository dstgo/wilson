package types

const (
	LogIpKey               = "ip"
	LogHttpMethodKey       = "method"
	LogRequestPathKey      = "path"
	LogRequestUrlKey       = "url"
	LogRequestQuery        = "query"
	LogRequestHeader       = "header"
	LogRequestBody         = "body"
	LogHttpStatusKey       = "status"
	LogRequestIdKey        = "requestId"
	LogRequestContentType  = "content-type"
	LogResponseContentType = "response-type"
	LogHttpContentLength   = "content-length"
	LogHttpResponseLength  = "response-length"
	LogRequestCostKey      = "cost"

	LogRecoverRequestKey = "request"
	LogRecoverStackKey   = "stack"
	LogRecoverErrorKey   = "error"
)

const (
	DateTimeFormat = "2006-01-02 15:04:05.999 -07:00"
	DateFormat     = "2006-01-02"
	TimeFormat     = "15:04:05.999"
)

const (
	QueryOk    = "op.query.ok"
	QueryFail  = "op.query.fail"
	CreateOk   = "op.create.ok"
	CreateFail = "op.create.fail"
	UpdateOk   = "op.update.ok"
	UpdateFail = "op.update.fail"
	DeleteOk   = "op.delete.ok"
	DeleteFail = "op.delete.fail"
)
