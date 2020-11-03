package mock

import (
	"context"
	"module-7/pkg/models"
)

var mockMaterial = &models.Material{
	Name: "This is a mocked Material",
}

type MaterialModel struct {}

func (m *MaterialModel) FindAll(ctx context.Context) (mat []*models.Material, err error) {
	var list []*models.Material

	list = append(list, mockMaterial)

	return list, nil
}
