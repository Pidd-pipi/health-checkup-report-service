package service

import (
    "context"
    "testing"

    "github.com/blueship581/gbcheckup/internal/model"
    "github.com/blueship581/gbcheckup/internal/repository"
)

func TestPackageService_DeleteItemInUseFails(t *testing.T) {
    db := newTestDB(t)
    ctx := context.Background()
    pkgRepo := repository.NewPackageRepository(db)
    itemRepo := repository.NewPackageItemRepository(db)
    svc := NewPackageService(pkgRepo, itemRepo, testLogger())

    pkg, err := svc.Create(ctx, "入职体检", "entry", 300, "active", "基础套餐")
    if err != nil {
        t.Fatal(err)
    }
    item, err := svc.AddItem(ctx, pkg.ID, &model.PackageItem{ItemName: "血常规", RefValueRange: "3.5-9.5", Department: "检验科"})
    if err != nil {
        t.Fatal(err)
    }
    examinee := model.Examinee{Name: "张三", IDCardNo: "110101199001011234"}
    if err := db.Create(&examinee).Error; err != nil {
        t.Fatal(err)
    }
    result := model.ExamResult{ExamineeID: examinee.ID, PackageItemID: item.ID, Status: "entered", DoctorID: 1}
    if err := db.Create(&result).Error; err != nil {
        t.Fatal(err)
    }

    if err := svc.DeleteItem(ctx, item.ID); err == nil {
        t.Fatal("expected error when deleting item in use")
    }
    var got model.PackageItem
    if err := db.First(&got, item.ID).Error; err != nil {
        t.Fatalf("item should still exist: %v", err)
    }
}
