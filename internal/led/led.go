package led

type LedStatus *[5]byte
type LedMode byte

const (
	ModeNone LedMode = iota
	ModeNormal
	ModeRave
	ModeSpectrum
	ModeRainbow
)

// Byte 0 is the Mode. If the Mode is ModeNormal, bytes 1-4 represent RGBA.
// Otherwise, bytes 1-4 are 0.
var (
	StatusOff      LedStatus = &[5]byte{byte(ModeNormal), 0, 0, 0, 255}
	StatusRed      LedStatus = &[5]byte{byte(ModeNormal), 255, 0, 0, 255}
	StatusGreen    LedStatus = &[5]byte{byte(ModeNormal), 0, 255, 0, 255}
	StatusBlue     LedStatus = &[5]byte{byte(ModeNormal), 0, 0, 255, 255}
	StatusWhite    LedStatus = &[5]byte{byte(ModeNormal), 255, 255, 255, 255}
	StatusMerlot   LedStatus = &[5]byte{byte(ModeNormal), 110, 51, 20, 255}
	StatusOrange   LedStatus = &[5]byte{byte(ModeNormal), 255, 165, 0, 255}
	StatusNotFound LedStatus = &[5]byte{byte(ModeNormal), 60, 33, 68, 255}
	StatusRave     LedStatus = &[5]byte{byte(ModeRave), 0, 0, 0, 0}
	StatusSpectrum LedStatus = &[5]byte{byte(ModeSpectrum), 0, 0, 0, 0}
	StatusRainbow  LedStatus = &[5]byte{byte(ModeRainbow), 0, 0, 0, 0}
)

var (
	StatusTranslation = map[string]LedStatus{
		"off":       StatusOff,
		"red":       StatusRed,
		"green":     StatusGreen,
		"blue":      StatusBlue,
		"white":     StatusWhite,
		"merlot":    StatusMerlot,
		"orange":    StatusOrange,
		"not found": StatusNotFound,
		"rave":      StatusRave,
		"spectrum":  StatusSpectrum,
		"rainbow":   StatusRainbow,
	}

	currentLedStatus = StatusOrange
	currentLedMode   = ModeNormal
)

func SetStatus(status LedStatus) {
	currentLedStatus = status
}

func SetMode(mode LedMode) {
	currentLedMode = mode
}

func GetStatus() LedStatus {
	return currentLedStatus
}

func GetMode() LedMode {
	return currentLedMode
}
