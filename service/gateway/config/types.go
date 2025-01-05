package config

import "time"

type Endpoint struct {
	Path           string
	Method         string
	Description    string
	Protocol       string
	ResponseFormat bool
	Timeout        time.Duration
	Metadata       map[string]string
	Host           string
	Middlewares    []Middleware
	Backends       []Backend
	Retry          *Retry
}

type Middleware struct {
	Name     string
	Options  map[string]interface{}
	Required bool
}

type Backend struct {
	Target string
	Weight *int64
}

type Header struct {
	Name  string
	Value string
}

type Condition struct {
	Header     *Header
	StatusCode string
}

type Retry struct {
	Count      int
	Timeout    time.Duration
	Conditions []Condition
}

type Tracing struct {
	Endpoint    string
	SampleRatio float64
	Timeout     time.Duration
	Insecure    bool
}

type CircuitBreaker struct {
	Trigger *struct {
		SuccessRatio *struct {
			Success float64
			Request int
			Bucket  int
			Window  time.Duration
		}
		Ratio int64
	}
	Action struct {
		ResponseData *struct {
			StatusCode int
			Header     []struct {
				Key   string
				Value []string
			}
			Body []byte
		}
		BackupService *struct {
			Endpoint Endpoint
		}
	}
	Conditions []Condition
}

type Cors struct {
	AllowCredentials    bool
	AllowOrigins        []string
	AllowMethods        []string
	AllowHeaders        []string
	ExposeHeaders       []string
	MaxAge              time.Duration
	AllowPrivateNetwork bool
}

type RewriteHeadersPolicy struct {
	Set    map[string]string
	Add    map[string]string
	Remove []string
}

type Rewrite struct {
	PathRewrite            string
	RequestHeadersRewrite  *RewriteHeadersPolicy
	ResponseHeadersRewrite *RewriteHeadersPolicy
	StripPrefix            string
	HostRewrite            string
}
