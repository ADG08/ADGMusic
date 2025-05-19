# ADGMusic Bot

Bot Discord pour la lecture de musique avec une architecture hexagonale.

## Structure du Projet

```
.
├── cmd/                    # Point d'entrée de l'application
│   └── bot/               # Main du bot
├── internal/              # Code interne de l'application
│   ├── domain/           # Entités et règles métier
│   ├── application/      # Cas d'utilisation
│   ├── infrastructure/   # Implémentations techniques
│   │   ├── discord/     # Implémentation Discord
│   │   ├── database/    # Implémentation PostgreSQL
│   │   └── music/       # Implémentation de la lecture audio
│   └── ports/           # Ports (interfaces)
│       ├── input/       # Ports d'entrée
│       └── output/      # Ports de sortie
├── pkg/                  # Packages réutilisables
├── sqlc/                # Configuration et requêtes SQLC
├── Dockerfile           # Configuration Docker
└── docker-compose.yml   # Configuration Docker Compose
```

## Technologies Utilisées

- Go 1.21
- Discord Go
- PostgreSQL avec pgx
- SQLC pour la génération de code SQL
- Docker pour la conteneurisation

## Installation

1. Cloner le repository
2. Copier `.env.example` vers `.env` et configurer les variables d'environnement
3. Lancer avec Docker Compose : `docker-compose up -d`

## Configuration

Les variables d'environnement nécessaires sont :
- `DISCORD_TOKEN` : Token du bot Discord
- `DATABASE_URL` : URL de connexion PostgreSQL 