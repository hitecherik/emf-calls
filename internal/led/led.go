package led

type LedStatus *[4]byte
type LedMode byte

const (
	ModeNone LedMode = iota
	ModeNormal
	ModeRave
	ModePulse
	ModeSpectrum
	ModeRainbow
)

// Byte 0 is the Mode. If the Mode is ModeNormal, bytes 1-3 represent RGB.
// Otherwise, bytes 1-3 are 0.
var (
	StatusOff      LedStatus = &[4]byte{byte(ModeNormal), 0, 0, 0}
	StatusRed      LedStatus = &[4]byte{byte(ModeNormal), 255, 0, 0}
	StatusGreen    LedStatus = &[4]byte{byte(ModeNormal), 0, 255, 0}
	StatusBlue     LedStatus = &[4]byte{byte(ModeNormal), 0, 0, 255}
	StatusWhite    LedStatus = &[4]byte{byte(ModeNormal), 255, 255, 255}
	StatusMerlot   LedStatus = &[4]byte{byte(ModeNormal), 110, 51, 20}
	StatusOrange   LedStatus = &[4]byte{byte(ModeNormal), 255, 165, 0}
	StatusNotFound LedStatus = &[4]byte{byte(ModeNormal), 60, 33, 68}
	StatusPulse    LedStatus = &[4]byte{byte(ModePulse), 0, 0, 0}
	StatusRave     LedStatus = &[4]byte{byte(ModeRave), 0, 0, 0}
	StatusSpectrum LedStatus = &[4]byte{byte(ModeSpectrum), 0, 0, 0}
	StatusRainbow  LedStatus = &[4]byte{byte(ModeRainbow), 0, 0, 0}
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
		"pulse":     StatusPulse,
		"spectrum":  StatusSpectrum,
		"rainbow":   StatusRainbow,
	}

	currentLedStatus       = StatusOrange
	currentBrightness byte = 255
)

func SetStatus(status LedStatus) {
	currentLedStatus = status
}

func SetBrightness(brightness byte) {
	currentBrightness = brightness
}

func GetStatus() LedStatus {
	return currentLedStatus
}

func GetBrightness() byte {
	return currentBrightness
}
