package objects

type (
	TextOptions struct {
		NoHeaders        bool
		ShowExec         bool
		UtilMap          map[string]UtilityPod
		UniqIdLength     int
		ShowHidden       bool
		Fields           []string
		AdditionalFields []string
	}
)
