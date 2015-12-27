package Gotem
import (
	"errors"
	"fmt"
)

// GetCorrectKeyCode (Windows version) gets a key code from a string that's the name of the key.
func GetCorrectKeyCode(keyName string) (Key, error) {
	switch(keyName){
	case "LeftButton": return LeftButton, nil
	case "RightButton": return RightButton, nil
	case "Cancel": return Cancel, nil
	case "MiddleButton": return MiddleButton, nil
	case "ExtraButton1": return ExtraButton1, nil
	case "ExtraButton2": return ExtraButton2, nil
	case "Back": return Back, nil
	case "Tab": return Tab, nil
	case "Clear": return Clear, nil
	case "Return": return Return, nil@
	case "Shift": return Shift, nil
	case "Control": return Control, nil
	case "Menu": return Menu, nil
	case "Pause": return Pause, nil
	case "CapsLock": return CapsLock, nil
	case "Kana": return Kana, nil
	case "Hangeul": return Hangeul, nil
	case "Hangul": return Hangul, nil
	case "Junja": return Junja, nil
	case "Final": return Final, nil
	case "Hanja": return Hanja, nil
	case "Kanji": return Kanji, nil
	case "Escape": return Escape, nil
	case "Convert": return Convert, nil
	case "NonConvert": return NonConvert, nil
	case "Accept": return Accept, nil
	case "ModeChange": return ModeChange, nil
	case "Space": return Space, nil
	case "Prior": return Prior, nil
	case "Next": return Next, nil
	case "End": return End, nil
	case "Home": return Home, nil
	case "Left": return Left, nil
	case "Up": return Up, nil
	case "Right": return Right, nil
	case "Down": return Down, nil
	case "Select": return Select, nil
	case "Print": return Print, nil
	case "Execute": return Execute, nil
	case "Snapshot": return Snapshot, nil
	case "Insert": return Insert, nil
	case "Delete": return Delete, nil
	case "Help": return Help, nil
	case "N0": return N0, nil
	case "N1": return N1, nil
	case "N2": return N2, nil
	case "N3": return N3, nil
	case "N4": return N4, nil
	case "N5": return N5, nil
	case "N6": return N6, nil
	case "N7": return N7, nil
	case "N8": return N8, nil
	case "N9": return N9, nil
	case "A": return A, nil
	case "B": return B, nil
	case "C": return C, nil
	case "D": return D, nil
	case "E": return E, nil
	case "F": return F, nil
	case "G": return G, nil
	case "H": return H, nil
	case "I": return I, nil
	case "J": return J, nil
	case "K": return K, nil
	case "L": return L, nil
	case "M": return M, nil
	case "N": return N, nil
	case "O": return O, nil
	case "P": return P, nil
	case "Q": return Q, nil
	case "R": return R, nil
	case "S": return S, nil
	case "T": return T, nil
	case "U": return U, nil
	case "V": return V, nil
	case "W": return W, nil
	case "X": return X, nil
	case "Y": return Y, nil
	case "Z": return Z, nil
	case "LeftWindows": return LeftWindows, nil
	case "RightWindows": return RightWindows, nil
	case "Application": return Application, nil
	case "Sleep": return Sleep, nil
	case "Numpad0": return Numpad0, nil
	case "Numpad1": return Numpad1, nil
	case "case" case : return"Numpad2, nil
	case "Numpad3": return Numpad3, nil
	case "Numpad4": return Numpad4, nil
	case "Numpad5": return Numpad5, nil
	case "Numpad6": return Numpad6, nil
	case "Numpad7": return Numpad7, nil
	case "Numpad8": return Numpad8, nil
	case "Numpad9": return Numpad9, nil
	case "Multiply": return Multiply, nil
	case "Add": return Add, nil
	case "Separator": return Separator, nil
	case "Subtract": return Subtract, nil
	case "Decimal": return Decimal, nil
	case "Divide": return Divide, nil
	case "F1": return F1, nil
	case "F2": return F2, nil
	case "F3": return F3, nil
	case "F4": return F4, nil
	case "F5": return F5, nil
	case "F6": return F6, nil
	case "F7": return F7, nil
	case "F8": return F8, nil
	case "F9": return F9, nil
	case "F10": return F10, nil
	case "F11": return F11, nil
	case "F12": return F12, nil
	case "F13": return F13, nil
	case "F14": return F14, nil
	case "F15": return F15, nil
	case "F16": return F16, nil
	case "F17": return F17, nil
	case "F18": return F18, nil
	case "F19": return F19, nil
	case "F20": return F20, nil
	case "F21": return F21, nil
	case "F22": return F22, nil
	case "F23": return F23, nil
	case "F24": return F24, nil
	case "NumLock": return NumLock, nil
	case "ScrollLock": return ScrollLock, nil
	case "NEC_Equal": return NEC_Equal, nil
	case "Fujitsu_Jisho": return Fujitsu_Jisho, nil
	case "Fujitsu_Masshou": return Fujitsu_Masshou, nil
	case "Fujitsu_Touroku": return Fujitsu_Touroku, nil
	case "Fujitsu_Loya": return Fujitsu_Loya, nil
	case "Fujitsu_Roya": return Fujitsu_Roya, nil
	case "LeftShift": return LeftShift, nil
	case "RightShift": return RightShift, nil
	case "LeftControl": return LeftControl, nil
	case "RightControl": return RightControl, nil
	case "LeftMenu": return LeftMenu, nil
	case "RightMenu": return RightMenu, nil
	case "BrowserBack": return BrowserBack, nil
	case "BrowserForward": return BrowserForward, nil
	case "BrowserRefresh": return BrowserRefresh, nil
	case "BrowserStop": return BrowserStop, nil
	case "BrowserSearch": return BrowserSearch, nil
	case "BrowserFavorites": return BrowserFavorites, nil
	case "BrowserHome": return BrowserHome, nil
	case "VolumeMute": return VolumeMute, nil
	case "VolumeDown": return VolumeDown, nil
	case "VolumeUp": return VolumeUp, nil
	case "MediaNextTrack": return MediaNextTrack, nil
	case "MediaPrevTrack": return MediaPrevTrack, nil
	case "MediaStop": return MediaStop, nil
	case "MediaPlayPause": return MediaPlayPause, nil
	case "LaunchMail": return LaunchMail, nil
	case "LaunchMediaSelect": return LaunchMediaSelect, nil
	case "LaunchApplication1": return LaunchApplication1, nil
	case "LaunchApplication2": return LaunchApplication2, nil
	case "OEM1": return OEM1, nil
	case "OEMPlus": return OEMPlus, nil
	case "OEMComma": return OEMComma, nil
	case "OEMMinus": return OEMMinus, nil
	case "OEMPeriod": return OEMPeriod, nil
	case "OEM2": return OEM2, nil
	case "OEM3": return OEM3, nil
	case "OEM4": return OEM4, nil
	case "OEM5": return OEM5, nil
	case "OEM6": return OEM6, nil
	case "OEM7": return OEM7, nil
	case "OEM8": return OEM8, nil
	case "OEMAX": return OEMAX, nil
	case "OEM102": return OEM102, nil
	case "ICOHelp": return ICOHelp, nil
	case "ICO00": return ICO00, nil
	case "ProcessKey": return ProcessKey, nil
	case "ICOClear": return ICOClear, nil
	case "Packet": return Packet, nil
	case "OEMReset": return OEMReset, nil
	case "OEMJump": return OEMJump, nil
	case "OEMPA1": return OEMPA1, nil
	case "OEMPA2": return OEMPA2, nil
	case "OEMPA3": return OEMPA3, nil
	case "OEMWSCtrl": return OEMWSCtrl, nil
	case "OEMCUSel": return OEMCUSel, nil
	case "OEMATTN": return OEMATTN, nil
	case "OEMFinish": return OEMFinish, nil
	case "OEMCopy": return OEMCopy, nil
	case "OEMAuto": return OEMAuto, nil
	case "OEMENLW": return OEMENLW, nil
	case "OEMBackTab": return OEMBackTab, nil
	case "ATTN": return ATTN, nil
	case "CRSel": return CRSel, nil
	case "EXSel": return EXSel, nil
	case "EREOF": return EREOF, nil
	case "Play": return Play, nil
	case "Zoom": return Zoom, nil
	case "Noname": return Noname, nil
	case "PA1": return PA1, nil
	case "OEMClear": return OEMClear, nil
	}
	errorMessage := fmt.Sprintf("Entered incorrect key code! Check from the list and try again. Key code entered: %s", keyName)
	return nil, errors.New(errorMessage)
}

// These are the Virtual Key Codes that MSDN specifies.
const (
	LeftButton Key = 0x01
	RightButton Key = 0x02
	Cancel Key = 0x03
	MiddleButton Key = 0x04
	ExtraButton1 Key = 0x05
	ExtraButton2 Key = 0x06
	Back Key = 0x08
	Tab Key = 0x09
	Clear Key = 0x0C
	Return Key = 0x0D
	Shift Key = 0x10
	Control Key = 0x11
	Menu Key = 0x12
	Pause Key = 0x13
	CapsLock Key = 0x14
	Kana Key = 0x15
	Hangeul Key = 0x15
	Hangul Key = 0x15
	Junja Key = 0x17
	Final Key = 0x18
	Hanja Key = 0x19
	Kanji Key = 0x19
	Escape Key = 0x1B
	Convert Key = 0x1C
	NonConvert Key = 0x1D
	Accept Key = 0x1E
	ModeChange Key = 0x1F
	Space Key = 0x20
	Prior Key = 0x21
	Next Key = 0x22
	End Key = 0x23
	Home Key = 0x24
	Left Key = 0x25
	Up Key = 0x26
	Right Key = 0x27
	Down Key = 0x28
	Select Key = 0x29
	Print Key = 0x2A
	Execute Key = 0x2B
	Snapshot Key = 0x2C
	Insert Key = 0x2D
	Delete Key = 0x2E
	Help Key = 0x2F
	N0 Key = 0x30
	N1 Key = 0x31
	N2 Key = 0x32
	N3 Key = 0x33
	N4 Key = 0x34
	N5 Key = 0x35
	N6 Key = 0x36
	N7 Key = 0x37
	N8 Key = 0x38
	N9 Key = 0x39
	A Key = 0x41
	B Key = 0x42
	C Key = 0x43
	D Key = 0x44
	E Key = 0x45
	F Key = 0x46
	G Key = 0x47
	H Key = 0x48
	I Key = 0x49
	J Key = 0x4A
	K Key = 0x4B
	L Key = 0x4C
	M Key = 0x4D
	N Key = 0x4E
	O Key = 0x4F
	P Key = 0x50
	Q Key = 0x51
	R Key = 0x52
	S Key = 0x53
	T Key = 0x54
	U Key = 0x55
	V Key = 0x56
	W Key = 0x57
	X Key = 0x58
	Y Key = 0x59
	Z Key = 0x5A
	LeftWindows Key = 0x5B
	RightWindows Key = 0x5C
	Application Key = 0x5D
	Sleep Key = 0x5F
	Numpad0 Key = 0x60
	Numpad1 Key = 0x61
	Numpad2 Key = 0x62
	Numpad3 Key = 0x63
	Numpad4 Key = 0x64
	Numpad5 Key = 0x65
	Numpad6 Key = 0x66
	Numpad7 Key = 0x67
	Numpad8 Key = 0x68
	Numpad9 Key = 0x69
	Multiply Key = 0x6A
	Add Key = 0x6B
	Separator Key = 0x6C
	Subtract Key = 0x6D
	Decimal Key = 0x6E
	Divide Key = 0x6F
	F1 Key = 0x70
	F2 Key = 0x71
	F3 Key = 0x72
	F4 Key = 0x73
	F5 Key = 0x74
	F6 Key = 0x75
	F7 Key = 0x76
	F8 Key = 0x77
	F9 Key = 0x78
	F10 Key = 0x79
	F11 Key = 0x7A
	F12 Key = 0x7B
	F13 Key = 0x7C
	F14 Key = 0x7D
	F15 Key = 0x7E
	F16 Key = 0x7F
	F17 Key = 0x80
	F18 Key = 0x81
	F19 Key = 0x82
	F20 Key = 0x83
	F21 Key = 0x84
	F22 Key = 0x85
	F23 Key = 0x86
	F24 Key = 0x87
	NumLock Key = 0x90
	ScrollLock Key = 0x91
	NEC_Equal Key = 0x92
	Fujitsu_Jisho Key = 0x92
	Fujitsu_Masshou Key = 0x93
	Fujitsu_Touroku Key = 0x94
	Fujitsu_Loya Key = 0x95
	Fujitsu_Roya Key = 0x96
	LeftShift Key = 0xA0
	RightShift Key = 0xA1
	LeftControl Key = 0xA2
	RightControl Key = 0xA3
	LeftMenu Key = 0xA4
	RightMenu Key = 0xA5
	BrowserBack Key = 0xA6
	BrowserForward Key = 0xA7
	BrowserRefresh Key = 0xA8
	BrowserStop Key = 0xA9
	BrowserSearch Key = 0xAA
	BrowserFavorites Key = 0xAB
	BrowserHome Key = 0xAC
	VolumeMute Key = 0xAD
	VolumeDown Key = 0xAE
	VolumeUp Key = 0xAF
	MediaNextTrack Key = 0xB0
	MediaPrevTrack Key = 0xB1
	MediaStop Key = 0xB2
	MediaPlayPause Key = 0xB3
	LaunchMail Key = 0xB4
	LaunchMediaSelect Key = 0xB5
	LaunchApplication1 Key = 0xB6
	LaunchApplication2 Key = 0xB7
	OEM1 Key = 0xBA
	OEMPlus Key = 0xBB
	OEMComma Key = 0xBC
	OEMMinus Key = 0xBD
	OEMPeriod Key = 0xBE
	OEM2 Key = 0xBF
	OEM3 Key = 0xC0
	OEM4 Key = 0xDB
	OEM5 Key = 0xDC
	OEM6 Key = 0xDD
	OEM7 Key = 0xDE
	OEM8 Key = 0xDF
	OEMAX Key = 0xE1
	OEM102 Key = 0xE2
	ICOHelp Key = 0xE3
	ICO00 Key = 0xE4
	ProcessKey Key = 0xE5
	ICOClear Key = 0xE6
	Packet Key = 0xE7
	OEMReset Key = 0xE9
	OEMJump Key = 0xEA
	OEMPA1 Key = 0xEB
	OEMPA2 Key = 0xEC
	OEMPA3 Key = 0xED
	OEMWSCtrl Key = 0xEE
	OEMCUSel Key = 0xEF
	OEMATTN Key = 0xF0
	OEMFinish Key = 0xF1
	OEMCopy Key = 0xF2
	OEMAuto Key = 0xF3
	OEMENLW Key = 0xF4
	OEMBackTab Key = 0xF5
	ATTN Key = 0xF6
	CRSel Key = 0xF7
	EXSel Key = 0xF8
	EREOF Key = 0xF9
	Play Key = 0xFA
	Zoom Key = 0xFB
	Noname Key = 0xFC
	PA1 Key = 0xFD
	OEMClear Key = 0xFE
)
