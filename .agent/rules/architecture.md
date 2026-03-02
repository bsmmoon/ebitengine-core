---
description: Architectural guidelines and best practices for the ebitengine-core project
---
# Ebitengine-Core Project Rules

When working on this project, adhere to the following architectural guidelines and best practices:

## Project Structure
- **`cmd/{game_name}/`**: Thin entry points. Each game has a `main.go` that configures and runs the game, and a `config.go` for game-specific screen and configuration settings.
- **`internal/{game_name}/`**: Game-specific logic and widgets (one package per game). This package handles the game loop implementation and owns its state.
- **`internal/shared/`**: Truly reusable cross-game utilities (e.g., Camera2D, IsometricProjection, InteractionManager, TileMap).

## Interface Pattern
Each game must implement the `ebiten.Game` interface implicitly via Go's duck typing:
- `Update() error` - Game logic per frame
- `Draw(screen *ebiten.Image)` - Rendering per frame
- `Layout(w, h int) (int, int)` - Screen dimensions setup

## Separation of Concerns
- Shared utilities stay in `internal/shared/` and should be reusable across multiple games.
- Game-specific widgets (like a specialized Button or HUD) stay with their respective game package (e.g., `internal/example_ui/`).
- Main functions (`cmd/`) only handle window setup and game initialization, not game logic.

## AI Assistant Guidelines
- **Prioritize shared components**: Always check `internal/shared/` for existing utilities before implementing new functionality.
- **Propose extensions**: If shared components could be extended to support new use cases, ask the user for their opinion before mutating `internal/shared/` files.
- **Focus on Clarity**: Evaluate whether abstractions provide syntactic sugar that improves code clarity, not just raw functionality.
