package models

import "testing"

func TestMaterialModel(t *testing.T)  {
	material := &Material{Name: "Test material"}

	if material.Name == "" {
		t.Errorf("Name property should not be empty")
	}
}
