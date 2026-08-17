package service

import (
    "context"
    "testing"

    "github.com/blueship581/gbcheckup/internal/repository"
)

func TestPackageService_UpdatePreservesOmittedFields(t *testing.T) {
    db := newTestDB(t)
    ctx := context.Background()
    svc := NewPackageService(repository.NewPackageRepository(db), repository.NewPackageItemRepository(db), testLogger())

    pkg, err := svc.Create(ctx, "高端体检", "premium", 1999, "active", "原说明")
    if err != nil {
        t.Fatal(err)
    }
    updated, err := svc.Update(ctx, pkg.ID, "", "", 0, "", "")
    if err != nil {
        t.Fatal(err)
    }
    if updated.Price != 1999 {
        t.Fatalf("price = %v, want 1999", updated.Price)
    }
    if updated.Description != "原说明" {
        t.Fatalf("description = %q, want 原说明", updated.Description)
    }
    if updated.Name != "高端体检" || updated.Status != "active" {
        t.Fatalf("name/status changed unexpectedly: name=%q status=%q", updated.Name, updated.Status)
    }
}
