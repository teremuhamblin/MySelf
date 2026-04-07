# ============================================
# MySelf — Makefile Professionnel
# Version : 1.2.0 / 1.3.0
# Auteur  : Teremu
# ============================================

APP_NAME = myself
API_DIR = cmd/api
COLLECTOR_DIR = cmd/collector
COLLECTOR_VB_DIR = cmd/collector_vB
COLLECTOR_VC_DIR = cmd/collector_vC
RECORDS_DIR = records

GO = go

# Couleurs
GREEN  = \033[0;32m
BLUE   = \033[0;34m
YELLOW = \033[1;33m
RED    = \033[0;31m
NC     = \033[0m

# ============================================
# Commandes principales
# ============================================

all: help

help:
	@echo ""
	@echo "$(BLUE)MySelf — Commandes disponibles$(NC)"
	@echo ""
	@echo "$(GREEN)make api$(NC)          → Lance l'API REST"
	@echo "$(GREEN)make collector$(NC)     → Lance le collecteur simple"
	@echo "$(GREEN)make collector_vB$(NC)  → Lance le collecteur 1.2.0"
	@echo "$(GREEN)make collector_vC$(NC)  → Lance le collecteur 1.3.0"
	@echo "$(GREEN)make fmt$(NC)           → Formatte le code Go"
	@echo "$(GREEN)make test$(NC)          → Lance les tests"
	@echo "$(GREEN)make build$(NC)         → Compile l'application"
	@echo "$(GREEN)make clean$(NC)         → Nettoie les binaires"
	@echo "$(GREEN)make tree$(NC)          → Affiche l'arborescence"
	@echo "$(GREEN)make init$(NC)          → Crée les dossiers nécessaires"
	@echo ""

# ============================================
# API
# ============================================

api:
	@echo "$(BLUE)Lancement de l'API MySelf...$(NC)"
	@$(GO) run $(API_DIR)/main.go

# ============================================
# Collecteurs
# ============================================

collector:
	@echo "$(BLUE>Lancement du collecteur simple...$(NC)"
	@$(GO) run $(COLLECTOR_DIR)/main.go

collector_vB:
	@echo "$(BLUE)Lancement du collecteur vB (1.2.0)...$(NC)"
	@$(GO) run $(COLLECTOR_VB_DIR)/main.go

collector_vC:
	@echo "$(BLUE)Lancement du collecteur vC (1.3.0)...$(NC)"
	@$(GO) run $(COLLECTOR_VC_DIR)/main.go

# ============================================
# Build
# ============================================

build:
	@echo "$(BLUE)Compilation de MySelf...$(NC)"
	@$(GO) build -o bin/$(APP_NAME) $(API_DIR)/main.go
	@echo "$(GREEN)Build terminé → bin/$(APP_NAME)$(NC)"

# ============================================
# Tests
# ============================================

test:
	@echo "$(BLUE)Lancement des tests...$(NC)"
	@$(GO) test ./... -v

# ============================================
# Formatage & lint
# ============================================

fmt:
	@echo "$(BLUE)Formatage du code...$(NC)"
	@$(GO) fmt ./...
	@echo "$(GREEN)Formatage terminé$(NC)"

vet:
	@echo "$(BLUE)Analyse statique (go vet)...$(NC)"
	@$(GO) vet ./...

# ============================================
# Initialisation
# ============================================

init:
	@echo "$(BLUE)Création des dossiers nécessaires...$(NC)"
	@mkdir -p $(RECORDS_DIR)
	@mkdir -p bin
	@echo "$(GREEN)Dossiers prêts$(NC)"

# ============================================
# Nettoyage
# ============================================

clean:
	@echo "$(YELLOW)Nettoyage des binaires...$(NC)"
	@rm -rf bin
	@echo "$(GREEN)Nettoyage terminé$(NC)"

# ============================================
# Arborescence
# ============================================

tree:
	@echo "$(BLUE)Arborescence du projet :$(NC)"
	@tree -L 4
