

# Watchdog 🐾

**Herrmienta liviana y de facil configuracion para el monitoreo de archivos y carpetas en tiempo real con notificaciones directas en telegram**

---

## ¿Qué es?

Watchdog es una herramienta CLI que monitorea archivos y directorios en tiempo real. Detecta cambios automáticamente, lo escribe en un .log y lo reporta  via notificacion a telegram. Sin dependencias externas — construido con Golang.

---

## Características

- 📁 Monitoreo de múltiples archivos y directorios simultáneamente
- 📬 Logs y alertas automáticas a Telegram al detectar cambios
- ⚡ Concurrencia con goroutines (hilos de golang) — una por archivo
- 🪶 Bajo consumo de recursos
- 🔧 Sin dependencias externas
- 🖥️ Compatible con Windows, Linux y Android (Termux)

---

## Instalación

### Opción 1 — Binario precompilado

Descargá el binario correspondiente a tu sistema desde [builds](./builds):

| Sistema | Archivo |
|---|---|
| Windows | `watchdog.exe` |
| Linux | `watchdog-linux` |
| Android (Termux) | `watchdog-termux` |

**Linux / Termux** — dar permisos de ejecución:
```bash
chmod +x watchdog-linux
./watchdog-linux
```

**Windows:**
```bash
watchdog.exe
```

---

### Opción 2 — Compilar desde el código fuente

Requiere tener [Go1.25.4](https://go.dev/dl/) instalado.

```bash
git clone https://github.com/mig-af/watchdog.git
cd watchdog
go build -o watchdog (si se quiere compilar)
go run main.go (si se quiere ejecutar directamente sin compilacion)
```

**Distribucion de Carpetas:**
```
watchdog/

├── builds/
│   ├── watchdog.exe
│   ├── watchdog-linux
│   └── watchdog-termux
├── info/
│   └── watchdog.log
    └── watchdog.json
├── src/
    └──config/
    └──dog/
    └──funcs/
    └──telegram/
    └──mystruct

├── go.mod
├── main.go
└── README.md

```

---

## Configuración



Editá el archivo `watchdog.json` dentro la carpeta `/info ` en el mismo directorio del binario :

Si se usa la opcion 1 de instalacion,  el binario precompilado, al iniciar este, se crea automáticamente los archivos necesarios

```json
{
   "timezone":"America/Bogota",

   "user": {
      "active": false,
      "telegram_id": 123456,
      "telegram_token": "telegram-token"
   },
   "file": {
      "paths": [
         "/rutaAbsoluta/hacia/el/archivo.txt"
      ]
   },
   "dir": {
      "paths": [
         "/rutaAbsoluta/hacia/la/carpeta/"
      ]
   }
}
```

| Campo | tipo | Descripción |
|---|---|----------|
| `timezone` | `string` | Zona horaria |
| `active` | `bool` | true para activar notificaciones en telegram
| `telegramId` | `int` | Id numerico del usuario de telegram  |
| `telegramToken` | `string` | Token del bot de telegram |
| `paths` | `array`| Lista de rutas de los archivos y carpetas a monitorear en formato string  |

---

## Uso

```bash
./watchdog-linux (linux)
./watchdog-termux (android)

```
Uso en segundo plano linux
```bash
./watchdog-linux  > /dev/null 2>&1 &
```
Uso en segundo plano windows
```powershell
Start-process .\watchdog.exe -WindowStyle hidden
```

El programa inicia el monitoreo, muestra en pantalla los archivos detectados y los almacena en el archivo /info/watchdog.log. Cualquier cambio genera una alerta en Telegram automáticamente con previa configuracion.

---

## Compilar para otros sistemas

```bash
# Windows
GOOS=windows GOARCH=amd64 go build -o builds/watchdog.exe

# Linux
GOOS=linux GOARCH=amd64 go build -o builds/watchdog-linux

# Android (Termux)
GOOS=android GOARCH=arm64 go build -o builds/watchdog-termux
```

---

