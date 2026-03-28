package profile

import (
	"errors"
	"strings"
)

// Profile représente un ensemble d'informations personnelles
// collectées et stockées dans MySelf.
type Profile struct {
	FullName string
	Age      string
	Email    string
	City     string
	Notes    string
}

// Clean normalise les champs (trim, espaces, etc.)
func (p *Profile) Clean() {
	p.FullName = strings.TrimSpace(p.FullName)
	p.Age = strings.TrimSpace(p.Age)
	p.Email = strings.TrimSpace(p.Email)
	p.City = strings.TrimSpace(p.City)
	p.Notes = strings.TrimSpace(p.Notes)
}

// Validate vérifie les champs essentiels.
// Tu peux renforcer cette logique selon tes besoins.
func (p *Profile) Validate() error {
	if p.FullName == "" {
		return errors.New("le nom complet est obligatoire")
	}
	if p.Email == "" {
		return errors.New("l'email est obligatoire")
	}
	return nil
}

// ToText génère une représentation textuelle propre,
// utilisée pour l’écriture dans un fichier.
func (p *Profile) ToText() string {
	var b strings.Builder

	b.WriteString("Nom complet : " + p.FullName + "\
