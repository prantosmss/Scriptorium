# AGENTS.md

## Setup and Installation

### Install Dependencies
Run the following to install dependencies:
```bash
./scripts/install.sh
```

### Verify Environment
Run a pre-flight check to ensure the environment is ready:
```bash
./scripts/run-local.sh doctor
```

### Check Model Connectivity
Verify that the local environment can connect to the models:
```bash
./scripts/run-local.sh check
```

## Local Execution

### Run the Application
Execute the application directly from source (no pre-compiled binaries needed):
```bash
./scripts/run-local.sh
```

### Run Specific Commands
- **Pipeline**:
  ```bash
  ./scripts/run-local.sh pipeline --new-novel --prompt "Your prompt here"
  ```
- **Service**:
  ```bash
  ./scripts/run-local.sh service open
  ```
- **Review**:
  ```bash
  ./scripts/run-local.sh review [--from N --to M]
  ```
- **Rewrite**:
  ```bash
  ./scripts/run-local.sh rewrite [--from N --to M]
  ```

### Check Model Connectivity
Verify that the local environment can connect to the models:
```bash
./scripts/run-local.sh check
```

### Verify Environment
Run a pre-flight check to ensure the environment is ready:
```bash
./scripts/run-local.sh doctor
```

### Help
Display help for available commands:
```bash
./scripts/run-local.sh help
```

### Check Model Connectivity
Verify that the local environment can connect to the models:
```bash
./scripts/run-local.sh check
```

### Verify Environment
Run a pre-flight check to ensure the environment is ready:
```bash
./scripts/run-local.sh doctor
```

### Help
Display help for available commands:
```bash
./scripts/run-local.sh help
```

## Configuration

### Required Configuration
Ensure `config.example.jsonc` is updated with:
- **Providers**: Valid provider IDs and credentials.
- **Models**: Specify the model IDs and variants.
- **Roles**: Define roles and their configurations.

Example structure:
```json
{
  "providers": {
    "openai": {
      "apiKey": "your-api-key"
    }
  },
  "models": {
    "openai/gpt-4": {
      "variant": "latest"
    }
  },
  "roles": {
    "agent": {
      "model": "openai/gpt-4"
    }
  }
}
```

## Containerized Setup

### Start Services with Docker Compose
Ensure Docker and Docker Compose are installed, then start the services:
```bash
docker-compose up
```

### Stop Services
Stop the services when done:
```bash
docker-compose down
```

## Architecture and Workflow

### Key Directories
- **`cmd/novel-studio`**: Main application entrypoint.
- **`output/`**: Generated output files.
- **`./.novel-studio/`**: Configuration and state files.

### Workflow Notes
- **Go Run**: Always use `go run ./cmd/novel-studio` to ensure the latest source is used.
- **Relative Paths**: Ensure the working directory is set to the root of the project for correct relative paths.

### Operational Gotchas
- **Model Sync**: Ensure the local model binaries are in sync with the source code.
- **Dependencies**: Run `./scripts/install.sh` before any other commands.

## References
- [Configuration Example](config.example.jsonc)
- [Technical Documentation](README-TECHNICAL.md)
- [Docker Compose Setup](docker-compose.yml)

Today's date: Thu Oct 08 2026

Here is some useful information about the environment you are running in:
<env>
  Working directory: C:\Users\Sadik Pranto\Desktop\Scriptorium
  Workspace root folder: C:\Users\Sadik Pranto\Desktop\Scriptorium
  Is directory a git repo: yes
  Platform: win32
  Prefer C:\Users\Sadik Pranto\AppData\Local\Temp\opencode over generic system temporary directories such as /tmp; it is pre-created and approved for external access.
</env>

### Install Dependencies
Run the following to install dependencies:
```bash
./scripts/install.sh
```

### Verify Environment
Run a pre-flight check to ensure the environment is ready:
```bash
./scripts/run-local.sh doctor
```

### Check Model Connectivity
Verify that the local environment can connect to the models:
```bash
./scripts/run-local.sh check
```

## Local Execution

### Run the Application
Execute the application directly from source (no pre-compiled binaries needed):
```bash
./scripts/run-local.sh
```

### Run Specific Commands
- **Pipeline**:
  ```bash
  ./scripts/run-local.sh pipeline --new-novel --prompt "Your prompt here"
  ```
- **Service**:
  ```bash
  ./scripts/run-local.sh service open
  ```
- **Review**:
  ```bash
  ./scripts/run-local.sh review [--from N --to M]
  ```
- **Rewrite**:
  ```bash
  ./scripts/run-local.sh rewrite [--from N --to M]
  ```

### Check Model Connectivity
Verify that the local environment can connect to the models:
```bash
./scripts/run-local.sh check
```

### Verify Environment
Run a pre-flight check to ensure the environment is ready:
```bash
./scripts/run-local.sh doctor
```

### Help
Display help for available commands:
```bash
./scripts/run-local.sh help
```

### Check Model Connectivity
Verify that the local environment can connect to the models:
```bash
./scripts/run-local.sh check
```

### Verify Environment
Run a pre-flight check to ensure the environment is ready:
```bash
./scripts/run-local.sh doctor
```

### Help
Display help for available commands:
```bash
./scripts/run-local.sh help
```

## Configuration

### Required Configuration
Ensure `config.example.jsonc` is updated with:
- **Providers**: Valid provider IDs and credentials.
- **Models**: Specify the model IDs and variants.
- **Roles**: Define roles and their configurations.

Example structure:
```json
{
  "providers": {
    "openai": {
      "apiKey": "your-api-key"
    }
  },
  "models": {
    "openai/gpt-4": {
      "variant": "latest"
    }
  },
  "roles": {
    "agent": {
      "model": "openai/gpt-4"
    }
  }
}
```

## Containerized Setup

### Start Services with Docker Compose
Ensure Docker and Docker Compose are installed, then start the services:
```bash
docker-compose up
```

### Stop Services
Stop the services when done:
```bash
docker-compose down
```

## Architecture and Workflow

### Key Directories
- **`cmd/novel-studio`**: Main application entrypoint.
- **`output/`**: Generated output files.
- **`./.novel-studio/`**: Configuration and state files.

### Workflow Notes
- **Go Run**: Always use `go run ./cmd/novel-studio` to ensure the latest source is used.
- **Relative Paths**: Ensure the working directory is set to the root of the project for correct relative paths.

### Operational Gotchas
- **Model Sync**: Ensure the local model binaries are in sync with the source code.
- **Dependencies**: Run `./scripts/install.sh` before any other commands.

## References
- [Configuration Example](config.example.jsonc)
- [Technical Documentation](README-TECHNICAL.md)
- [Docker Compose Setup](docker-compose.yml)