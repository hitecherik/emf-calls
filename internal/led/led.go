package led

type LedStatus *[3]byte
type LedMode byte

const (
	ModeNone LedMode = iota
	ModeNormal
	ModeRave
	ModePulse
	ModeSpectrum
	ModeRainbow
)

var (
	StatusOff      LedStatus = &[3]byte{0, 0, 0}
	StatusRed      LedStatus = &[3]byte{255, 0, 0}
	StatusGreen    LedStatus = &[3]byte{0, 255, 0}
	StatusBlue     LedStatus = &[3]byte{0, 0, 255}
	StatusWhite    LedStatus = &[3]byte{255, 255, 255}
	StatusMerlot   LedStatus = &[3]byte{110, 51, 20}
	StatusOrange   LedStatus = &[3]byte{255, 165, 0}
	StatusNotFound LedStatus = &[3]byte{60, 33, 68}
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
	}
	ModeTranslation = map[string]LedMode{
		"normal":   ModeNormal,
		"rave":     ModeRave,
		"pulse":    ModePulse,
		"spectrum": ModeSpectrum,
		"rainbow":  ModeRainbow,
	}

	currentLedStatus       = StatusOrange
	currentMode            = ModeNormal
	currentBrightness byte = 255
)

func SetStatus(status LedStatus) {
	currentLedStatus = status
}

func SetMode(mode LedMode) {
	currentMode = mode
}

func SetBrightness(brightness byte) {
	currentBrightness = brightness
}

func GetStatus() LedStatus {
	return currentLedStatus
}

func GetMode() LedMode {
	return currentMode
}

func GetBrightness() byte {
	return currentBrightness
}
