package service

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/blueship581/gbcheckup/internal/model"
	"github.com/blueship581/gbcheckup/internal/repository"
	"github.com/blueship581/gbcheckup/internal/util"
)

func newExamResultService(t *testing.T) *ExamResultService {
	db := newTestDB(t)
	return NewExamResultService(repository.NewExamResultRepository(db), repository.NewRegistrationRepository(db), repository.NewAbnormalMetricRepository(db), testLogger())
}

func TestEnterMissingResult404(t *testing.T) {
	svc := newExamResultService(t)
	_, err := svc.Enter(context.Background(), 999, 1, EnterInput{ResultValue: "5.5"})
	var appErr *util.AppError
	if !errors.As(err, &appErr) || appErr.HTTPStatus != http.StatusNotFound {
		t.Fatalf("err = %v, want NotFound(404)", err)
	}
}

func TestReviewMissingResult404(t *testing.T) {
	svc := newExamResultService(t)
	err := svc.Review(context.Background(), 999)
	var appErr *util.AppError
	if !errors.As(err, &appErr) || appErr.HTTPStatus != http.StatusNotFound {
		t.Fatalf("err = %v, want NotFound(404)", err)
	}
}

func TestEnterAlreadyEntered409(t *testing.T) {
	db := newTestDB(t)
	item := model.PackageItem{ItemName: "血糖", RefValueRange: "3.9-6.1", Department: "检验科"}
	if err := db.Create(&item).Error; err != nil {
		t.Fatal(err)
	}
	examinee := model.Examinee{Name: "张三", IDCardNo: "110101199001011234"}
	if err := db.Create(&examinee).Error; err != nil {
		t.Fatal(err)
	}
	res := model.ExamResult{ExamineeID: examinee.ID, PackageItemID: item.ID, Status: "entered", DoctorID: 1}
	if err := db.Create(&res).Error; err != nil {
		t.Fatal(err)
	}
	svc := NewExamResultService(repository.NewExamResultRepository(db), repository.NewRegistrationRepository(db), repository.NewAbnormalMetricRepository(db), testLogger())
	_, err := svc.Enter(context.Background(), res.ID, 1, EnterInput{ResultValue: "5.5"})
	var appErr *util.AppError
	if !errors.As(err, &appErr) || appErr.HTTPStatus != http.StatusConflict {
		t.Fatalf("err = %v, want Conflict(409)", err)
	}
}
