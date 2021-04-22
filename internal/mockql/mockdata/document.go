package mockdata

import (
	"github.com/sageflow/sageapi/internal/mockql/model"
)

func getSessionDocuments() []*model.SessionDocument {
	return []*model.SessionDocument{
		&model.SessionDocument{
			ID:   "35326b33-65a2-460a-aa02-55422bcd569d",
			Name: "Customers",
		},
		&model.SessionDocument{
			ID:   "2fee0da1-4a95-47b1-a774-9a26c920bd99",
			Name: "Employees",
		},
	}
}
