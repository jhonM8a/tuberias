# Usa la imagen oficial de Go con Alpine
FROM golang:1.23-alpine AS builder

# Instala dependencias necesarias (opcional)
RUN apk add --no-cache git


# Establece el directorio de trabajo
WORKDIR /app

# Copia los archivos del proyecto al contenedor
COPY . .

# Compila el programa Go
RUN go build -o cmd/main .

# Crea una imagen final minimalista
FROM alpine:latest

# Copia el binario desde la etapa anterior
COPY --from=builder /app/main /main

# Configura el comando por defecto
CMD ["/main"]