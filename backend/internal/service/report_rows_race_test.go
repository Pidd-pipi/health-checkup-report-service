package service

import (
	"testing"

	"github.com/blueship581/gbcheckup/internal/model"
	"github.com/blueship581/gbcheckup/internal/util"
)

func makeReportResults() []model.ExamResult {
	return []model.ExamResult{
		{ResultValue: "5.5", IsAbnormal: false, PackageItem: model.PackageItem{ItemName: "血压", RefValueRange: "3.5-9.5"}},
		{ResultValue: "99", IsAbnormal: true, PackageItem: model.PackageItem{ItemName: "血糖", RefValueRange: "3.9-6.1"}},
		{ResultValue: "7.0", IsAbnormal: false, PackageItem: model.PackageItem{ItemName: "视力", RefValueRange: "4.0-5.5"}},
	}
}

func assertRowNames(t *testing.T, got []util.PDFRow, want ...string) {
	t.Helper()
	if len(got) != len(want) {
		t.Fatalf("rows = %d, want %d", len(got), len(want))
	}
	for i, w := range want {
		if len(got[i].Columns) == 0 || got[i].Columns[0] != w {
			t.Fatalf("row[%d] item = %q, want %q", i, got[i].Columns[0], w)
		}
	}
}

func TestReportGenerateKeepsAllResultRows(t *testing.T) {
	assertRowNames(t, reportRows(makeReportResults()), "血压", "血糖", "视力")
	if got := summarizeAbnormal(makeReportResults()); got != 1 {
		t.Fatalf("summarizeAbnormal = %d, want 1", got)
	}
	assertRowNames(t, util.BuildResultRows(makeReportResults()), "血压", "血糖", "视力")
	if got := util.CountAbnormalResults(makeReportResults()); got != 1 {
		t.Fatalf("CountAbnormalResults = %d, want 1", got)
	}
}
