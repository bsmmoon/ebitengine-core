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
- **tiles** - Tilemap rendering demo with layered tiles and debug overlay

## Project Structure

```
cmd/                    # Executable entry points
  ├── helloworld/       # Basic demo
  ├── example_ui/       # UI widgets demo
  └── tiles/            # Tilemap demo
internal/               # Reusable components
  ├── helloworld/       # Hello world logic
  ├── ui/               # UI widgets and utilities
  ├── example_ui/       # Example UI game implementation
  └── tiles/            # Tiles game implementation
```

## Architecture

**Package Structure:**
- `cmd/` - Thin entry points that configure and run games
- `internal/ui/` - Reusable UI widgets and utilities (shared across games)
- `internal/{game_name}/` - Game-specific logic (one package per game)

**Interface Pattern:**
Each game implements `ebiten.Game` interface implicitly (Go's duck typing):
- `Update() error` - Game logic per frame
- `Draw(screen *ebiten.Image)` - Rendering per frame
- `Layout(w, h int) (int, int)` - Screen dimensions

**Separation of Concerns:**
- UI widgets are stateless and reusable
- Game structs own their state and widget instances
- Main functions handle window setup and game initialization

## AI Assistant Reference

> This section helps AI assistants efficiently navigate the codebase by providing direct file paths and their purposes, minimizing token usage from exploratory file reads.

**Entry Points:**
- `cmd/helloworld/main.go` - Minimal game example
- `cmd/example_ui/main.go` - UI demo with widgets
- `cmd/tiles/main.go` - Tilemap rendering demo

**Core Game Logic:**
- `internal/example_ui/game.go` - Example UI game loop (Update/Draw/Layout)
- `internal/example_ui/game_config.go` - Game configuration
- `internal/tiles/game.go` - Tiles game loop with TileMap and DebugOverlay
- `internal/tiles/game_config.go` - Tiles game configuration

**UI Components:**
- `internal/ui/context.go` - UI state management
- `internal/ui/button.go` - Clickable button widget
- `internal/ui/check_box.go` - Checkbox with toggle
- `internal/ui/textbox.go` - Scrollable text display
- `internal/ui/v_scroll_bar.go` - Vertical scrollbar
- `internal/ui/input.go` - Input handling
- `internal/ui/tilemap.go` - TileMap component for layered tile rendering
- `internal/ui/debug_overlay.go` - TPS/FPS debug display widget
- `internal/ui/image_utils.go` - Image loading and Spritesheet utilities

**Configuration:**
- `go.mod` - Dependencies (Ebitengine v2.8.8)
- `cmd/example_ui/config.go` - Screen/UI settings
- `cmd/tiles/config.go` - Tiles screen settings
