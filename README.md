# ebitengine-core

The goal of this repository is to provide a "core" minimal ebitengine project ready for use.

## Getting Started

1. Open in devcontainer (Docker required)
2. Build and run an example:
   ```bash
   cd cmd/helloworld
   make build
   ./helloworld
   ```

## Examples

- **helloworld** - Minimal example displaying "Hello, World!"
- **example_ui** - Interactive UI with buttons, checkboxes, and text boxes

## Project Structure

```
cmd/                    # Executable entry points
  ├── helloworld/       # Basic demo
  └── example_ui/       # UI widgets demo
internal/               # Reusable components
  ├── helloworld/       # Hello world logic
  ├── ui/               # UI widgets (Button, CheckBox, TextBox, VScrollBar)
  └── example_ui/       # Example UI game implementation
```

## AI Assistant Reference

> This section helps AI assistants efficiently navigate the codebase by providing direct file paths and their purposes, minimizing token usage from exploratory file reads.

**Entry Points:**
- `cmd/helloworld/main.go` - Minimal game example
- `cmd/example_ui/main.go` - UI demo with widgets

**Core Game Logic:**
- `internal/example_ui/game.go` - Example UI game loop (Update/Draw/Layout)
- `internal/example_ui/game_config.go` - Game configuration

**UI Components:**
- `internal/ui/context.go` - UI state management
- `internal/ui/button.go` - Clickable button widget
- `internal/ui/check_box.go` - Checkbox with toggle
- `internal/ui/textbox.go` - Scrollable text display
- `internal/ui/v_scroll_bar.go` - Vertical scrollbar
- `internal/ui/input.go` - Input handling

**Configuration:**
- `go.mod` - Dependencies (Ebitengine v2.8.8)
- `cmd/example_ui/config.go` - Screen/UI settings
