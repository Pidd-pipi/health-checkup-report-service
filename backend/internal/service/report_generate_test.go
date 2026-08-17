package service

import (
    "context"
    "errors"
    "testing"
    "time"

    "github.com/blueship581/gbcheckup/internal/model"
    "github.com/blueship581/gbcheckup/internal/repository"
    "github.com/blueship581/gbcheckup/internal/util"
)

func TestReportService_GenerateRequiresAllResultsEntered(t *testing.T) {
    db := newTestDB(t)
    ctx := context.Background()

    pkg := model.Package{Name: "年度体检", PackageType: "annual", Price: 599, Status: "active"}
    if err := db.Create(&pkg).Error; err != nil {
        t.Fatal(err)
    }
    item := model.PackageItem{PackageID: pkg.ID, ItemName: "空腹血糖", RefValueRange: "3.9-6.1", Department: "检验科"}
    if err := db.Create(&item).Error; err != nil {
        t.Fatal(err)
    }
    examinee := model.Examinee{Name: "李四", IDCardNo: "110101199202021234"}
    if err := db.Create(&examinee).Error; err != nil {
        t.Fatal(err)
    }
    reg := model.Registration{ExamineeID: examinee.ID, PackageID: pkg.ID, GuideNo: "GUIDE202608170001", Status: "registered", RegisteredAt: time.Now(), RegisterUserID: 2}
    if err := db.Create(&reg).Error; err != nil {
        t.Fatal(err)
    }
    if err := db.Create(&model.ExamResult{RegistrationID: reg.ID, ExamineeID: examinee.ID, PackageItemID: item.ID, Status: "entered", ResultValue: "5.2", DoctorID: 2}).Error; err != nil {
        t.Fatal(err)
    }
    if err := db.Create(&model.ExamResult{RegistrationID: reg.ID, ExamineeID: examinee.ID, PackageItemID: item.ID, Status: "pending"}).Error; err != nil {
        t.Fatal(err)
    }
    report := model.Report{RegistrationID: reg.ID, ExamineeID: examinee.ID, ReportNo: "GB202608170001", Status: "draft", DoctorID: 2}
    if err := db.Create(&report).Error; err != nil {
        t.Fatal(err)
    }

    svc := NewReportService(repository.NewReportRepository(db), repository.NewExamResultRepository(db), repository.NewRegistrationRepository(db), testLogger())
    _, err := svc.Generate(ctx, report.ID, 2, "", "", "")
    if err == nil {
        t.Fatal("expected error because not all results are entered")
    }
    var appErr *util.AppError
    if !errors.As(err, &appErr) || appErr.Code != 1404 {
        t.Fatalf("expected AppError code 1404, got %v", err)
    }
}
