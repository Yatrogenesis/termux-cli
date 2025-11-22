# Termux CLI

Una herramienta CLI poderosa y eficiente para Android construida específicamente para Termux.

## Descripción

Termux CLI es un framework de línea de comandos diseñado para aprovechar al máximo las capacidades de Termux en Android. Proporciona acceso simplificado a las APIs de Android, gestión de almacenamiento, networking y automatización de tareas.

## Características

- **Sistema**: Información del sistema, CPU, memoria, kernel
- **Batería**: Monitoreo de estado y salud de la batería
- **Almacenamiento**: Análisis y gestión de espacio en disco
- **Red**: Diagnóstico de conexión WiFi y herramientas de networking
- **Notificaciones**: Integración con el sistema de notificaciones de Android
- **Optimizado**: Binario pequeño (<10MB), mínimo consumo de recursos

## Requisitos

- Android 7.0+
- Termux instalado
- Termux-API (para funciones avanzadas)
- Arquitectura: ARM64 (aarch64)

## Instalación

### Método 1: Descarga Directa (Recomendado)

```bash
# Descargar la última versión
wget https://github.com/Yatrogenesis/termux-cli/releases/latest/download/termux-cli-arm64

# Dar permisos de ejecución
chmod +x termux-cli-arm64

# Mover a bin
mv termux-cli-arm64 $PREFIX/bin/termux-cli
```

### Método 2: Install Script

```bash
curl -sSL https://raw.githubusercontent.com/Yatrogenesis/termux-cli/main/install.sh | bash
```

### Método 3: Compilar desde Fuente

```bash
# Clonar el repositorio
git clone https://github.com/Yatrogenesis/termux-cli.git
cd termux-cli

# Compilar
go build -o termux-cli main.go

# Instalar
cp termux-cli $PREFIX/bin/
```

## Uso

```bash
# Ver comandos disponibles
termux-cli --help

# Info del sistema
termux-cli system info

# Estado de batería
termux-cli battery status

# Análisis de almacenamiento
termux-cli storage analyze

# Info de red
termux-cli network info

# Enviar notificación
termux-cli notify "Título" "Mensaje"
```

## Comandos Disponibles

### System
```bash
termux-cli system info        # Información del sistema
termux-cli system cpu         # Info de CPU
termux-cli system memory      # Estado de memoria
```

### Battery
```bash
termux-cli battery status     # Estado de batería
termux-cli battery health     # Salud de batería
```

### Storage
```bash
termux-cli storage info       # Info de almacenamiento
termux-cli storage analyze    # Analizar uso de espacio
termux-cli storage cleanup    # Limpiar archivos temporales
termux-cli storage large      # Encontrar archivos grandes
```

### Network
```bash
termux-cli network info       # Info de red
termux-cli network wifi       # Estado WiFi
termux-cli network ping       # Ping a host
termux-cli network speed      # Test de velocidad
```

### Notifications
```bash
termux-cli notify <title> <message>  # Enviar notificación
```

## Arquitectura

```
termux-cli/
├── cmd/                    # Comandos CLI
│   ├── root.go
│   ├── system.go
│   ├── battery.go
│   ├── storage.go
│   └── network.go
├── internal/
│   ├── termux/           # Integración Termux-API
│   ├── storage/          # Gestión de almacenamiento
│   └── network/          # Utilidades de red
├── pkg/                  # Paquetes públicos
├── docs/                 # Documentación
└── main.go
```

## Desarrollo

### Requisitos de Desarrollo

- Go 1.25+
- Termux con acceso a git
- GitHub CLI (gh)

### Setup

```bash
# Clonar repositorio
git clone https://github.com/Yatrogenesis/termux-cli.git
cd termux-cli

# Instalar dependencias
go mod download

# Ejecutar tests
go test ./...

# Compilar
go build -o termux-cli main.go
```

### Contribuir

1. Fork el proyecto
2. Crea una rama (`git checkout -b feature/AmazingFeature`)
3. Commit tus cambios (`git commit -m 'Add: amazing feature'`)
4. Push a la rama (`git push origin feature/AmazingFeature`)
5. Abre un Pull Request

## Roadmap

### v0.1.0 (MVP)
- [x] Estructura básica del proyecto
- [ ] Comandos de sistema
- [ ] Info de batería
- [ ] Análisis de almacenamiento
- [ ] Info de red

### v0.2.0
- [ ] Limpieza automática de espacio
- [ ] Gestión avanzada de archivos
- [ ] Notificaciones ricas
- [ ] Tests completos

### v0.3.0
- [ ] Terminal UI interactiva
- [ ] Automatización de tareas
- [ ] Integración cloud
- [ ] Sincronización de archivos

## Documentación

Para documentación completa, consulta:
- [Technical Design Document](./TDD.md)
- [API Reference](./docs/api.md)
- [Contributing Guide](./CONTRIBUTING.md)

## Stack Tecnológico

- **Lenguaje**: Go 1.25
- **CLI Framework**: Cobra + Viper
- **Testing**: testify
- **CI/CD**: GitHub Actions

## Licencia

MIT License - ver archivo [LICENSE](LICENSE) para detalles.

## Autor

**Yatrogenesis**

- GitHub: [@Yatrogenesis](https://github.com/Yatrogenesis)

## Agradecimientos

- Comunidad de Termux
- Proyecto Termux-API
- Cobra CLI Framework

## Soporte

Si encuentras un bug o tienes una sugerencia:
- Abre un [Issue](https://github.com/Yatrogenesis/termux-cli/issues)
- Únete a las [Discussions](https://github.com/Yatrogenesis/termux-cli/discussions)

---

**Nota**: Este proyecto está optimizado para dispositivos Android con Termux. El binario compilado es específico para arquitectura ARM64 (aarch64).

Hecho con ❤️ para la comunidad de Termux
