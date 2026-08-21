package service

import (
	"context"
	"testing"

	"github.com/blueship581/gbcheckup/internal/model"
	"github.com/blueship581/gbcheckup/internal/repository"
)

func TestBatchUpdateItemsReleasesAndKeepsError(t *testing.T) {
	db := newTestDB(t)
	svc := NewPackageService(repository.NewPackageRepository(db), repository.NewPackageItemRepository(db), testLogger())
	err := svc.BatchUpdateItems(context.Background(), []model.PackageItem{{ItemName: ""}})
	if err == nil {
		t.Fatal("BatchUpdateItems should return error for empty item name")
	}
}

func TestBatchDeleteItemsReleasesAndKeepsError(t *testing.T) {
	db := newTestDB(t)
	svc := NewPackageService(repository.NewPackageRepository(db), repository.NewPackageItemRepository(db), testLogger())
	err := svc.BatchDeleteItems(context.Background(), []uint{0})
	if err == nil {
		t.Fatal("BatchDeleteItems should return error for empty id")
	}
}
