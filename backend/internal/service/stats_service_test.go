package service

import (
	"context"
	"reflect"
	"testing"

	"github.com/blueship581/gbcheckup/internal/model"
	"github.com/blueship581/gbcheckup/internal/repository"
)

func TestDashboardMonthlySummaryNoPanic(t *testing.T) {
	db := newTestDB(t)
	pkg := model.Package{Name: "入职体检", PackageType: "entry", Price: 300, Status: "active"}
	if err := db.Create(&pkg).Error; err != nil {
		t.Fatal(err)
	}
	item := model.PackageItem{PackageID: pkg.ID, ItemName: "血常规", Department: "检验科"}
	if err := db.Create(&item).Error; err != nil {
		t.Fatal(err)
	}
	examinee := model.Examinee{Name: "张三", IDCardNo: "110101199001011234"}
	if err := db.Create(&examinee).Error; err != nil {
		t.Fatal(err)
	}
	reg := model.Registration{ExamineeID: examinee.ID, PackageID: pkg.ID, GuideNo: "GUIDE001", Status: "registered"}
	if err := db.Create(&reg).Error; err != nil {
		t.Fatal(err)
	}

	svc := NewStatsService(
		repository.NewPackageRepository(db), repository.NewRegistrationRepository(db),
		repository.NewReportRepository(db), repository.NewExamResultRepository(db),
		repository.NewAbnormalMetricRepository(db), repository.NewPackageItemRepository(db),
		testLogger(),
	)
	stats, err := svc.Dashboard(context.Background())
	if err != nil {
		t.Fatalf("Dashboard error = %v", err)
	}
	v := reflect.ValueOf(stats).Elem()
	for _, name := range []string{"PackageSoldSummary", "DeptWorkloadSummary", "AbnormalTopSummary", "MonthlyRevenueSummary"} {
		f := v.FieldByName(name)
		if !f.IsValid() || f.IsNil() {
			t.Fatalf("DashboardStats.%s missing or nil", name)
		}
	}
}
