package led

type LedStatus *[3]byte

var (
	StatusOff   LedStatus = &[3]byte{0, 0, 0}
	StatusRed   LedStatus = &[3]byte{255, 0, 0}
	StatusGreen LedStatus = &[3]byte{0, 255, 0}
	StatusBlue  LedStatus = &[3]byte{0, 0, 255}
	StatusWhite LedStatus = &[3]byte{255, 255, 255}
)

var (
	StatusTranslation = map[string]LedStatus{
		"off":   StatusOff,
		"red":   StatusRed,
		"green": StatusGreen,
		"blue":  StatusBlue,
		"white": StatusWhite,
	}

	currentLedStatus = StatusWhite
)

func SetStatus(status LedStatus) {
	currentLedStatus = status
}

func GetStatus() LedStatus {
	return currentLedStatus
}
