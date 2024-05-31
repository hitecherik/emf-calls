package led

type LedStatus byte

const (
	StatusOff LedStatus = iota
	StatusRed
	StatusGreen
	StatusBlue
	StatusWhite
)

var (
	StatusTranslation = map[string]LedStatus{
		"off":   StatusOff,
		"red":   StatusRed,
		"green": StatusGreen,
		"blue":  StatusBlue,
		"white": StatusWhite,
	}

	currentLedStatus = StatusOff
)

func SetStatus(status LedStatus) {
	currentLedStatus = status
}

func GetStatus() LedStatus {
	return currentLedStatus
}
