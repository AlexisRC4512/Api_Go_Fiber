# Usa una imagen base de Go
FROM golang:1.22-bullseye AS builder

# Establece el directorio de trabajo
WORKDIR /app

# Copia los archivos del proyecto al contenedor
COPY go.mod go.sum ./
RUN go mod download

# Copia el código fuente al contenedor
COPY src ./src

# Compila la aplicación
RUN go build -o app ./src/app

# Usa la misma imagen base para ejecutar
FROM golang:1.22-bullseye

# Establece el directorio de trabajo
WORKDIR /root/

# Copia el ejecutable desde la etapa de construcción
COPY --from=builder /app/app .

# Copia el archivo de configuración
COPY src/config/config.yaml /root/config.yaml

# Expone el puerto 3000
EXPOSE 3000

# Comando para ejecutar la aplicación
CMD ["./app"]
