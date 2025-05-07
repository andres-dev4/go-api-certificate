# Inicializar módulo (crea go.mod):
go mod init github.com/andres-dev4/go-api-certificate

# Sincronizar dependencias (actualiza go.mod y go.sum):
go mod tidy

# Verificar integridad de dependencias:
go mod verify