package service

import (
	"context"
	"errors"
	"testing"

	"github.com/blueship581/gbcheckup/internal/model"
	"github.com/blueship581/gbcheckup/internal/repository"
)

func cancelledCtx() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	return ctx
}

func TestBatchImportStopsOnCancelledContext(t *testing.T) {
	db := newTestDB(t)
	svc := NewExamineeService(repository.NewExamineeRepository(db), testLogger())
	csv := "姓名,身份证号,手机号,性别,年龄\n张三,110101199001011234,13800000001,男,30\n李四,110101199001011235,13800000002,女,28"
	count, _, err := svc.BatchImport(cancelledCtx(), nil, csv)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("BatchImport err = %v, want context.Canceled", err)
	}
	if count != 0 {
		t.Fatalf("BatchImport imported %d rows on cancelled context", count)
	}
}

func TestCreateCancelledContextReturnsError(t *testing.T) {
	db := newTestDB(t)
	svc := NewExamineeService(repository.NewExamineeRepository(db), testLogger())
	_, err := svc.Create(cancelledCtx(), &model.Examinee{Name: "张三", IDCardNo: "110101199001011234"})
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("Create err = %v, want context.Canceled", err)
	}
}

func TestGetCancelledContextReturnsError(t *testing.T) {
	db := newTestDB(t)
	svc := NewExamineeService(repository.NewExamineeRepository(db), testLogger())
	_, err := svc.Get(cancelledCtx(), 1)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("Get err = %v, want context.Canceled", err)
	}
}

func TestListCancelledContextReturnsError(t *testing.T) {
	db := newTestDB(t)
	svc := NewExamineeService(repository.NewExamineeRepository(db), testLogger())
	_, _, err := svc.List(cancelledCtx(), "", 1, 10)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("List err = %v, want context.Canceled", err)
	}
}
