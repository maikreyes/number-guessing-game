# Number Guessing Game 🎲

Este repositorio contiene una solución para el proyecto [Number Guessing Game](https://roadmap.sh/projects/number-guessing-game) de [roadmap.sh](https://roadmap.sh).

Es un juego de consola (CLI) escrito en **Go** donde el usuario debe adivinar un número generado aleatoriamente por la computadora, con un número limitado de intentos basado en la dificultad seleccionada.

## 📋 Características

El juego cumple con los siguientes requisitos del proyecto:

- **Selección de Dificultad**:
  - 🟢 **Easy (Fácil)**: 10 oportunidades.
  - 🟡 **Medium (Medio)**: 5 oportunidades.
  - 🔴 **Hard (Difícil)**: 3 oportunidades.
- **Generación Aleatoria**: La computadora selecciona un número entre 1 y 100.
- **Feedback Interactivo**: Indica si el número ingresado es mayor o menor que el número objetivo.
- **Mensajes de Estado**: Bienvenida, reglas, victoria (con conteo de intentos) y derrota.
- **Interfaz de Línea de Comandos (CLI)**: Interacción limpia y directa desde la terminal.

## 🚀 Tecnologías

- **Lenguaje**: Go (Golang)
- **Arquitectura**: Estructura estándar de proyectos Go (`cmd/` para el punto de entrada, `internal/` para la lógica del juego).

## 📂 Estructura del Proyecto

```text
number-guessing-game/
├── cmd/        # Punto de entrada de la aplicación (main.go)
├── internal/   # Lógica del juego y funcionalidades internas
├── .gitignore  # Archivos ignorados por git
├── go.mod      # Definición del módulo y dependencias
├── LICENSE     # Licencia MIT
└── README.md   # Documentación del proyecto
```

🛠️ Instalación y Uso
Prerrequisitos

Tener instalado Go (versión 1.18 o superior recomendada).

Pasos para ejecutar

Clonar el repositorio:

```Bash
git clone [https://github.com/maikreyes/number-guessing-game.git](https://github.com/maikreyes/number-guessing-game.git)
cd number-guessing-game
Ejecutar el juego: Dependiendo de cómo hayas nombrado tu archivo principal dentro de cmd, el comando será uno de los siguientes (usualmente es el nombre del proyecto o main.go):
```

```Bash
go run ./cmd/main.go
# O alternativamente:
# go run ./cmd/main.go
Compilar (Opcional): Si deseas generar un ejecutable binario:
```

```Bash
go build -o number-guessing-game ./cmd/number-guessing-game
./number-guessing-game
```

🎮 Cómo Jugar
Al iniciar, verás un mensaje de bienvenida.

Selecciona un nivel de dificultad ingresando el número correspondiente:

1 para Fácil.

2 para Medio.

3 para Difícil.

Ingresa tu adivinanza (un número entre 1 y 100).

El juego te dirá si el número secreto es mayor o menor que tu intento.

¡Gana si adivinas el número antes de que se acaben tus intentos!

📄 Licencia
Este proyecto está bajo la Licencia MIT - mira el archivo LICENSE para más detalles.

👤 Autor
Michael Estiven Reyes Escobar

GitHub: @maikreyes

Este proyecto es parte de mi camino de aprendizaje en desarrollo backend.