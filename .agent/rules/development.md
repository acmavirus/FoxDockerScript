# Development Rules for FoxDocker Panel

To ensure stability and prevent build failures in GitHub Actions, every agent must follow these rules:

## 1. Vue & TypeScript Standards
- **Strict Typing**: Always use `lang="ts"` and `<script setup>`.
- **No Unused Variables**: Ensure all variables defined in `<script>` are either used in logic or referenced in the `<template>`. Unused variables will FAIL the build.
- **No Duplicate Declarations**: Avoid redeclaring `ref` or constants. Always check if a similar variable already exists before adding a new one.
- **Template Integrity**: When editing large `<template>` blocks, double-check that every tag is correctly closed (`</div>`, `</button>`, etc.). Avoid "corrupted" diffs that leave orphan tags.
- **Variable Scoping**: Ensure variables used in `v-for` (e.g., `item`, `subTab`) do not conflict with outer scopes and are correctly typed.

## 2. Deployment & Build Safety
- **Local Validation**: Before finishing a task involving Vue files, RUN the build command locally to verify there are no TypeScript errors:
  ```powershell
  cd web; npm run build
  ```
- **Dockerfile Sync**: If you change the Go version in `go.mod`, ensure the `Dockerfile` stage for the backend builder matches the same version.

## 3. Automation Scripts
- **Command Detection**: Scripts in `scripts/` (like `update.sh` or `install.sh`) must detect if the system uses `docker compose` or `docker-compose` to ensure compatibility across different VPS environments.
- **Directory Checks**: Always ensure necessary directories (like `/opt/foxdocker/scripts`) are created and populated.

## 4. UI/UX Excellence
- **Premium Aesthetics**: Maintain the glassmorphism and modern dark mode design system.
- **Responsive**: All new sections in `App.vue` must be responsive (mobile-friendly).
