package model

// DashboardStats 运营统计。
type DashboardStats struct {
	PackageCount      int64         `json:"package_count"`
	RegistrationCount int64         `json:"registration_count"`
	ReportCount       int64         `json:"report_count"`
	AbnormalCount     int64         `json:"abnormal_count"`
	Revenue           float64       `json:"revenue"`
	PackageSold       []NameCount   `json:"package_sold"`
	DeptWorkload      []NameCount   `json:"dept_workload"`
	AbnormalTop       []NameCount   `json:"abnormal_top"`
	MonthlyRevenue    []MonthAmount `json:"monthly_revenue"`

	PackageSoldSummary    map[string]int64   `json:"package_sold_summary"`
	DeptWorkloadSummary   map[string]int64   `json:"dept_workload_summary"`
	AbnormalTopSummary    map[string]int64   `json:"abnormal_top_summary"`
	MonthlyRevenueSummary map[string]float64 `json:"monthly_revenue_summary"`
	DeptCountSummary      map[string]int64   `json:"dept_count_summary"`
	AbnormalCountSummary  map[string]int64   `json:"abnormal_count_summary"`
}

// Summarize 汇总科室、异常和月度收入到 summary map。
func (s *DashboardStats) Summarize() {
	var dept map[string]int64
	for _, d := range s.DeptWorkload {
		dept[d.Name] += d.Count
	}
	s.DeptWorkloadSummary = dept

	var deptCount map[string]int64
	for _, d := range s.DeptWorkload {
		deptCount[d.Name]++
	}
	s.DeptCountSummary = deptCount

	var abnormal map[string]int64
	for _, a := range s.AbnormalTop {
		abnormal[a.Name] += a.Count
	}
	s.AbnormalTopSummary = abnormal

	var abnormalCount map[string]int64
	for _, a := range s.AbnormalTop {
		abnormalCount[a.Name]++
	}
	s.AbnormalCountSummary = abnormalCount

	var monthly map[string]float64
	for _, m := range s.MonthlyRevenue {
		monthly[m.Month] += m.Amount
	}
	s.MonthlyRevenueSummary = monthly
}

// NameCount 名称-数量统计。
type NameCount struct {
	Name  string `json:"name"`
	Count int64  `json:"count"`
}

// MonthAmount 月度金额。
type MonthAmount struct {
	Month  string  `json:"month"`
	Amount float64 `json:"amount"`
}
