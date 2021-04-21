package mockdata

import (
	"github.com/sageflow/sageapi/internal/mockql/model"
	"github.com/sageflow/sageflow/pkg/strs"
)

func getWorkspaceProjectDocuments() []*model.Document {
	return []*model.Document{
		&model.Document{
			ID:            "35326b33-65a2-460a-aa02-55422bcd569d",
			Name:          "Customers",
			Specification: strs.GetAddress(``),
		},
		&model.Document{
			ID:            "2fee0da1-4a95-47b1-a774-9a26c920bd99",
			Name:          "Employees",
			Specification: strs.GetAddress(``),
		},
	}
}
