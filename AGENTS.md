# Project Context for AI Agents

## Overview
This project wraps an existing Vite + TypeScript web application (`fm-mindmap`) into a desktop application using **Wails v2**.
The primary goal is packaging; logic changes should be minimal and restricted to the Wails backend configuration.

## Core Principles
1.  **Immutable Frontend**: Do NOT modify code within the `frontend/` directory. It is a git submodule containing the target web application.
2.  **Mise for Management**: All development tasks (running, building, tool management) must be performed via `mise`.
3.  **Wails Wrapper**: The Go code (`main.go`, `app.go`) serves purely as a container/bridge for the frontend.

## Tech Stack
- **Wrapper**: Wails v2 (Go)
- **Frontend**: React, TypeScript, Vite, Tailwind CSS (inside `frontend/`)
- **Tool Manager**: mise

## Environment Setup
- **Tools**: Node.js (lts), Go (latest), Wails (managed via mise/go) are defined in `mise.toml`.
- **Initialization**: `mise install` setups the environment.

## Commands
| Action | Command | Description |
| :--- | :--- | :--- |
| **Development** | `mise run dev` | Starts the Wails development server and frontend watcher. |
| **Build** | `mise run build` | Builds the production desktop binary. |
| **Init** | `mise run init-project` | (One-off) Initializes Wails config. |

## Directory Structure
- `frontend/`: The existing web application (Git Submodule). **Read-only context.**
- `build/`: Wails build artifacts and assets.
- `wails.json`: Wails configuration.
- `main.go`: Entry point for the Wails application.
- `mise.toml`: Task runner and tool version definitions.
- `docs/`: Project documentation.

## Workflow
1.  **Setup**: Ensure `mise` is installed and run `mise install`.
2.  **Dev**: Run `mise run dev` to test the application locally.
3.  **Build**: Run `mise run build` to generate the executable.

## Important Notes
- The `wails.json` is configured to run `npm install` and `npm run build` inside the `frontend` directory during the build process.
- If `frontend` dependencies change, `wails.json` scripts might need adjustment, but the `frontend` source code itself should remain untouched by this project's agents.
