# Génère un profil de test dans MySelf/records/
test-profile:
	@echo "Génération d'un profil de test..."
	@mkdir -p MySelf/records/test_profile
	@echo "Nom complet : Test User" > MySelf/records/test_profile/profile.txt
	@echo "Âge : 30" >> MySelf/records/test_profile/profile.txt
	@echo "Email : test@example.com" >> MySelf/records/test_profile/profile.txt
	@echo "Ville : Toulouse" >> MySelf/records/test_profile/profile.txt
	@echo "Notes : Profil généré automatiquement pour les tests." >> MySelf/records/test_profile/profile.txt
	@echo "Profil de test créé dans MySelf/records/test_profile/"

# Nom du binaire généré
BINARY = collector

# Dossier du code source
CMD_DIR = cmd/collector

# Commande par défaut : build + run
all: build run

# Compilation du collector
build:
	go build -o $(BINARY) $(CMD_DIR)

# Exécution du collector
run:
	./$(BINARY)

# Nettoyage des artefacts
clean:
	rm -f $(BINARY)

# Formatage du code
fmt:
	go fmt ./...

# Vérification statique
vet:
	go vet ./...

# Build + vérifications
check: fmt vet build


