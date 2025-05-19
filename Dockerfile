FROM golang:1.24.3-alpine AS builder

WORKDIR /app

# Installer les dépendances nécessaires
RUN apk add --no-cache gcc musl-dev

# Copier les fichiers de dépendances
COPY go.mod go.sum ./
RUN go mod download

# Copier le code source
COPY . .

# Générer le code SQLC
RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
RUN sqlc generate

# Compiler l'application
RUN CGO_ENABLED=0 GOOS=linux go build -o bot ./cmd/bot

# Image finale
FROM alpine:latest

WORKDIR /app

# Copier le binaire compilé
COPY --from=builder /app/bot .

# Exposer le port si nécessaire
EXPOSE 8080

# Commande par défaut
CMD ["./bot"] 