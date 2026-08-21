package service

import (
	"context"
	"errors"
	"testing"

	"github.com/blueship581/gbcheckup/internal/repository"
)

func regCancelledCtx() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	return ctx
}

func TestRegisterCtxCancel(t *testing.T) {
	db := newTestDB(t)
	svc := NewRegistrationService(repository.NewRegistrationRepository(db), repository.NewExamineeRepository(db), repository.NewPackageRepository(db), repository.NewPackageItemRepository(db), repository.NewExamResultRepository(db), testLogger())
	_, err := svc.Register(regCancelledCtx(), 1, 1, 1)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("Register err = %v, want context.Canceled", err)
	}
}

func TestGetCtxCancel(t *testing.T) {
	db := newTestDB(t)
	svc := NewRegistrationService(repository.NewRegistrationRepository(db), repository.NewExamineeRepository(db), repository.NewPackageRepository(db), repository.NewPackageItemRepository(db), repository.NewExamResultRepository(db), testLogger())
	_, err := svc.Get(regCancelledCtx(), 1)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("Get err = %v, want context.Canceled", err)
	}
}

func TestListCtxCancel(t *testing.T) {
	db := newTestDB(t)
	svc := NewRegistrationService(repository.NewRegistrationRepository(db), repository.NewExamineeRepository(db), repository.NewPackageRepository(db), repository.NewPackageItemRepository(db), repository.NewExamResultRepository(db), testLogger())
	_, _, err := svc.List(regCancelledCtx(), "", 1, 10)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("List err = %v, want context.Canceled", err)
	}
}

func TestUpdateStatusCtxCancel(t *testing.T) {
	db := newTestDB(t)
	svc := NewRegistrationService(repository.NewRegistrationRepository(db), repository.NewExamineeRepository(db), repository.NewPackageRepository(db), repository.NewPackageItemRepository(db), repository.NewExamResultRepository(db), testLogger())
	err := svc.UpdateStatus(regCancelledCtx(), 1, "completed")
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("UpdateStatus err = %v, want context.Canceled", err)
	}
}
