package service

import (
    "context"
    "testing"

    "github.com/blueship581/gbcheckup/internal/constants"
    "github.com/blueship581/gbcheckup/internal/model"
    "github.com/blueship581/gbcheckup/internal/repository"
)

func TestAbnormalMetricService_ListWithoutExamineeReturnsAll(t *testing.T) {
    db := newTestDB(t)
    ctx := context.Background()

    item1 := model.PackageItem{ItemName: "血糖", RefValueRange: "3.9-6.1", Department: "检验科"}
    item2 := model.PackageItem{ItemName: "肝功能", RefValueRange: "0-40", Department: "检验科"}
    if err := db.Create(&item1).Error; err != nil {
        t.Fatal(err)
    }
    if err := db.Create(&item2).Error; err != nil {
        t.Fatal(err)
    }
    if err := db.Create(&model.AbnormalMetric{ExamineeID: 1, PackageItemID: item1.ID, AbnormalLevel: constants.AbnormalMild, Value: "8.0", RefValueRange: "3.9-6.1", FollowUpStatus: constants.FollowUpPending}).Error; err != nil {
        t.Fatal(err)
    }
    if err := db.Create(&model.AbnormalMetric{ExamineeID: 2, PackageItemID: item2.ID, AbnormalLevel: constants.AbnormalModerate, Value: "120", RefValueRange: "0-40", FollowUpStatus: constants.FollowUpPending}).Error; err != nil {
        t.Fatal(err)
    }

    svc := NewAbnormalMetricService(repository.NewAbnormalMetricRepository(db), testLogger())
    items, total, err := svc.List(ctx, 0, 1, 20)
    if err != nil {
        t.Fatal(err)
    }
    if total != 2 || len(items) != 2 {
        t.Fatalf("items = %d, total = %d, want both 2", len(items), total)
    }
}
