package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"myself/internal/profile"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("=== MySelf : Collecteur de Profil ===")

	p := profile.Profile{
		FullName: ask(reader, "Nom complet"),
		Age:      ask(reader, "Âge"),
		Email:    ask(reader, "Email"),
		City:     ask(reader, "Ville"),
		Notes:    ask(reader, "Notes personnelles"),
	}

	p.Clean()

	if err := p.Validate(); err != nil {
		fmt.Println("Erreur :", err)
		return
	}

	if err := saveProfile(p); err != nil {
		fmt.Println("Erreur lors de l'enregistrement :", err)
		return
	}

	fmt.Println("Profil enregistré dans MySelf/records/")
}

func ask(r *bufio.Reader, label string) string {
	fmt.Printf("%s : ", label)
	text, _ := r.ReadString('\n')
	return strings.TrimSpace(text)
}

func saveProfile(p profile.Profile) error {
	baseDir := "MySelf/records"
	if err := os.MkdirAll(baseDir, 0755); err != nil {
		return err
	}

	safeName := strings.ReplaceAll(strings.ToLower(p.FullName), " ", "_")
	dirName := fmt.Sprintf("%s_%d", safeName, time.Now().Unix())
	recordDir := filepath.Join(baseDir, dirName)

	if err := os.Mkdir(recordDir, 0755); err != nil {
		return err
	}

	filePath := filepath.Join(recordDir, "profile.txt")
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(p.ToText())
	return err
}
