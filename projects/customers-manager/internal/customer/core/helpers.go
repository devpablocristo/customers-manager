package core

import (
	"math"

	domain "github.com/devpablocristo/tech-house/projects/customers-manager/internal/customer/core/domain"
)

func calculateKPI(customers []domain.Customer) *domain.KPI {
	kpi := &domain.KPI{}
	if len(customers) == 0 {
		return kpi
	}

	var sumAge float64
	for _, c := range customers {
		sumAge += float64(c.Age)
	}
	kpi.AverageAge = sumAge / float64(len(customers))

	var sumSquaredDiff float64
	for _, c := range customers {
		diff := float64(c.Age) - kpi.AverageAge
		sumSquaredDiff += diff * diff
	}
	kpi.AgeStdDeviation = math.Sqrt(sumSquaredDiff / float64(len(customers)))

	return kpi
}
