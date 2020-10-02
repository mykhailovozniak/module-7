package mongodb

import (
	"module-7/pkg/models"
	"testing"
)

func TestMaterialsModel_FindAll(t *testing.T) {
	db, clearTestData := newTestDB(t)
	defer clearTestData()

	m := MaterialsModel{Collection: db.Collection("materials-local")}
	materials, _ := m.FindAll()

	material1 := models.Material{Name: "Some name for testing #1"}

	if material1.Name != materials[0].Name {
		t.Errorf("Records are not the same")
	}
}
