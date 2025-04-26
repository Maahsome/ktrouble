package defaults

import (
	"ktrouble/objects"
)

func EphemeralSleepList() []objects.EphemeralSleep {

	return []objects.EphemeralSleep{
		{
			Name:    "30 Minutes",
			Seconds: "1800",
		},
		{
			Name:    "4 Hours",
			Seconds: "14400",
		},
		{
			Name:    "8 Hours",
			Seconds: "28800",
		},
		{
			Name:    "12 Hours",
			Seconds: "43200",
		},
		{
			Name:    "1 Day",
			Seconds: "86400",
		},
	}
}
