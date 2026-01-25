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
  └── game/             # UI widgets (Button, CheckBox, TextBox, VScrollBar)
```
