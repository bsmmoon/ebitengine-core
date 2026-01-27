# BACKLOG

<!-- This file is for AI's reference to track potential improvements and future work -->

## Architecture Review
- [ ] **Evaluate internal/ui package scope** - Currently contains mix of truly reusable (InteractionManager, DebugOverlay, TileMap) and single-use components (Button, CheckBox, TextBox, VScrollBar only used by example_ui)
  - Option A: Keep as-is (batteries included approach)
  - Option B: Move single-use widgets to internal/example_ui, keep only reusable components in internal/ui
  - Option C: Remove internal/ui entirely, each game uses Ebitengine directly
  - Recommendation: Option B - minimal UI toolkit with genuinely reusable components

## Testing
- [ ] Add tests for UI components (Button, CheckBox, TextBox, VScrollBar)
- [ ] Add integration tests for example_ui

## Documentation
- [ ] Add usage examples in code comments for each widget
- [ ] Document widget API and event callbacks

## Features
- [ ] Scene/state management system (menu → gameplay → pause)
- [ ] Asset management structure (fonts, images, audio)
- [ ] Additional UI widgets (Slider, Dropdown, Modal)
- [ ] Centralized configuration system

## Code Quality
- [ ] Consistent error handling across components
- [ ] Performance profiling for UI rendering
