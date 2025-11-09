
# 🌊 Tide — Terminal IDE for Go Developers

Tide is a lightweight **Terminal IDE** built in Go with [Bubble Tea](https://github.com/charmbracelet/bubbletea).  
It brings an IDE-like experience directly into your terminal — fast, keyboard-driven, and built for productivity.

---

## ✨ Features

- 📁 File Explorer — Navigate projects without leaving your terminal  
- 📝 Text Editor — Open, edit, and save files with syntax-aware features  
- 🔍 Command Palette — Quick access to tools and actions  
- ⚡  Multi-Screen Layout — File explorer, editor, and terminal running side-by-side  
- 🎨 Nerd Font Icons — Rich visual feedback using Nerd Fonts  
- 🧩 Modular Architecture — Each screen (e.g., Editor, Explorer) has its own model and update loop  

---

## 🧱 Project Structure

```bash
tide/
├── cmd/                # Main app entry point
├── internal/
│   ├── core/           # App model and main program flow
│   ├── components/     # Reusable TUI components
│   ├── styles/         # Colors, borders, themes
├── assets/             # Icons, fonts, sample configs
└── README.md
````

---

## 🚀 Installation

### Clone and build manually

```bash
git clone https://github.com/GoKells/tide.git
cd tide
make build-<your-os>
```

### Run directly

```bash
go run ./cmd
```

---

## 🧰 Requirements

* Go 1.22+
* Nerd Fonts installed (for icons)
* Terminal with truecolor support

---

## ⚙️ Configuration

Tide stores configuration in a local `.tide/` folder within your project directory.

Example structure:

```
.tide/
├── config.yaml
├── keybindings.yaml
└── theme.yaml
```

You can customize keymaps, theme colors, and layout options here.

---

## 💡 Keybindings (Default)

| Action       | Shortcut   |
| ------------ | ---------- |
| Open File    | `Enter`    |
| Go Back      | `Esc`      |
| Switch Focus | `Tab`      |
| Save File    | `Ctrl + S` |
| Quit Tide    | `Ctrl + Q` |

---

## 🧪 Development

Run Tide in development mode with live reload (using [Air](https://github.com/cosmtrek/air)): (Not yet implemented)

```bash
air
```

Format and tidy:

```bash
go fmt ./...
go mod tidy
```

---

## 🤝 Contributing

Contributions are welcome!
Fork the repo, create a feature branch, and submit a pull request.

Before committing:

```bash
go fmt ./...
golangci-lint run
```

---

## 📜 License

MIT License © 2025 [Your Name or Organization]

---

## 💬 Roadmap

* [ ] Text editor with inteliscense and code highlighting 
* [ ] Working explorer window
* [ ] Plugin System for language support
* [ ] Git integration (status, commits, diffs)
* [ ] Built-in terminal panel
* [ ] Syntax highlighting
* [ ] Fuzzy file search

---

## 🧠 Inspiration

Tide is inspired by tools like:

* [Helix](https://helix-editor.com)
* [Neovim](https://neovim.io)
* [Charmbracelet Bubble Tea](https://github.com/charmbracelet/bubbletea)

---

> “The tide doesn’t ask permission to rise — neither should your code.” 🌊

