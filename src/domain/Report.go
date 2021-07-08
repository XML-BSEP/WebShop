package domain

import (
	"encoding/xml"
	"time"
)

type Report struct {
	XMLName xml.Name `xml:"report"`
	ReportId string `xml:"report_id" json:"report_id"`
	Timestamp time.Time `xml:"timestamp"`
	StatisticReport []StatisticsReport `xml:"statistics" json:"statistic_report"`
}
