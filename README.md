# gomatchup
**gomatchup** is a command-line interface (CLI) tool written in Go that assists in generating team compositions according to customizable constraints. It features a REPL (Read–Eval–Print Loop) interface that enables users to interactively manage players and apply constraints before generating balanced teams.

---

## Features

- **Interactive REPL**<br>
Add players and define constraints through an interactive interface.

- **Customizable Constraints**<br>
Add constraints, to prevent certain players from being in the same team.

- **Team Generation**<br>
Split added players into a specified number of teams, while respecting the constraints and keeping team sizes balanced.

## Installation
You can install `gomatchup` using `go install`
```bash
go install github.com/MrBhop/gomatchup@latest
```
After installation, ensure your `GOPATH/bin` or Go's default bin directory is in your `PATH`. Then run via
```bash
gomatchup
```

## Usage
When you start `gomatchup`, you'll enter an interactive prompt where you can add players, define constraints, and generate teams. The interface provides guidance as you go, so most functionality should be self-explanatory. Use the built-in `help` command at any time to see available actions, and `exit` to leave the program.

## License
No License.

## Acknowledgements
This project was created for a course on boot.dev (https://www.boot.dev/courses/build-personal-project-1).
