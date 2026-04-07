package profile

// Types de records supportés par MySelf.
// Compatible avec le collecteur, l'analyse et l'API.
const (
	TypeEmotion = "emotion"
	TypeHabit   = "habit"
	TypeEvent   = "event"
	TypeNote    = "note"
)

// IsValidType vérifie si un type est reconnu.
func IsValidType(t string) bool {
	switch t {
	case TypeEmotion, TypeHabit, TypeEvent, TypeNote:
		return true
	default:
		return false
	}
}
