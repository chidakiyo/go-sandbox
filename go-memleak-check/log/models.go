package log

import (
	"encoding/json"
	"time"
)

var _ json.Marshaler = LogSeverity(0)

type LogSeverity int

const (
	SeverityDefault   LogSeverity = 0
	SeverityDebug                 = 100
	SeverityInfo                  = 200
	SeverityNotice                = 300
	SeverityWarning               = 400
	SeverityError                 = 500
	SeverityCritical              = 600
	SeverityAlert                 = 700
	SeverityEmergency             = 800
)

func (severity LogSeverity) MarshalJSON() ([]byte, error) {
	return json.Marshal(severity.String())
}

func (severity LogSeverity) String() string {
	switch severity {
	case 0:
		return "DEFAULT"
	case 100:
		return "DEBUG"
	case 200:
		return "INFO"
	case 300:
		return "NOTICE"
	case 400:
		return "WARNING"
	case 500:
		return "ERROR"
	case 600:
		return "CRITICAL"
	case 700:
		return "ALERT"
	case 800:
		return "EMERGENCY"
	default:
		return "ERROR"
	}
}

type LogEntry struct {
	Severity       LogSeverity             `json:"severity"`
	HttpRequest    *HttpRequest            `json:"httpRequest,omitempty"`
	Time           Time                    `json:"time,omitempty"`
	Trace          string                  `json:"logging.googleapis.com/trace,omitempty"`
	SpanID         string                  `json:"logging.googleapis.com/spanId,omitempty"`
	Operation      *LogEntryOperation      `json:"logging.googleapis.com/operation,omitempty"`
	SourceLocation *LogEntrySourceLocation `json:"logging.googleapis.com/sourceLocation,omitempty"`
	Message        string                  `json:"message,omitempty"`
}

type HttpRequest struct {
	RequestMethod                  string   `json:"requestMethod"`
	RequestURL                     string   `json:"requestUrl"`
	RequestSize                    int64    `json:"requestSize,string,omitempty"`
	Status                         int      `json:"status"`
	ResponseSize                   int64    `json:"responseSize,string,omitempty"`
	UserAgent                      string   `json:"userAgent,omitempty"`
	RemoteIP                       string   `json:"remoteIp,omitempty"`
	Referer                        string   `json:"referer,omitempty"`
	Latency                        Duration `json:"latency,omitempty"`
	CacheLookup                    *bool    `json:"cacheLookup,omitempty"`
	CacheHit                       *bool    `json:"cacheHit,omitempty"`
	CacheValidatedWithOriginServer *bool    `json:"cacheValidatedWithOriginServer,omitempty"`
	CacheFillBytes                 *int64   `json:"cacheFillBytes,string,omitempty"`
	Protocol                       string   `json:"protocol"`
}

type LogEntryOperation struct {
	ID       string `json:"id,omitempty"`
	Producer string `json:"producer,omitempty"`
	First    *bool  `json:"first,omitempty"`
	Last     *bool  `json:"last,omitempty"`
}

type LogEntrySourceLocation struct {
	File     string `json:"file,omitempty"`
	Line     int64  `json:"line,string,omitempty"`
	Function string `json:"function,omitempty"`
}

var _ json.Marshaler = Time(time.Time{})

type Time time.Time

func (t Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(t).Format(time.RFC3339Nano))
}

var _ json.Marshaler = Duration(0)

type Duration time.Duration

func (d Duration) MarshalJSON() ([]byte, error) {
	nanos := time.Duration(d).Nanoseconds()
	secs := nanos / 1e9
	nanos -= secs * 1e9

	v := make(map[string]interface{})
	v["seconds"] = int64(secs)
	v["nanos"] = int32(nanos)

	return json.Marshal(v)
}
