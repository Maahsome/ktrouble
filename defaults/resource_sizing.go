package defaults

import (
	"ktrouble/objects"
)

func ResourceSizingList() []objects.ResourceSize {

	return []objects.ResourceSize{
		{
			Name:       "Small",
			LimitsCPU:  "250m",
			LimitsMEM:  "2Gi",
			RequestCPU: "100m",
			RequestMEM: "512Mi",
		},
		{
			Name:       "Medium",
			LimitsCPU:  "500m",
			LimitsMEM:  "4Gi",
			RequestCPU: "200m",
			RequestMEM: "1Gi",
		},
		{
			Name:       "Large",
			LimitsCPU:  "1000m",
			LimitsMEM:  "8Gi",
			RequestCPU: "500m",
			RequestMEM: "2Gi",
		},
		{
			Name:       "XLarge",
			LimitsCPU:  "10000m",
			LimitsMEM:  "30Gi",
			RequestCPU: "6000m",
			RequestMEM: "2Gi",
		},
	}
}
