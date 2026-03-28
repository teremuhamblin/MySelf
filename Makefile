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
