# Game Dev Command Line

Custom command line tooling to help with game development.

## Installation

```sh
go install github.com/Tiden-Dev/game-dev-cmd@latest
```

## Usage

```
gdev [command] [flags]
```

### Global Flags

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--log` | `-l` | `info` | Log level: `debug`, `info`, `warn`, `error`, `fatal` |

---

## Commands

### `gdev ue`

Reads the `.uproject` file in the current directory and prints the project name and engine version.

Run from the root of an Unreal Engine project:

```sh
gdev ue
```

**Example output:**

```
2026-06-10T12:00:00Z INF unreal project found! name=MyGame engine=5.7
```

**With debug logging:**

```sh
gdev ue --log debug
```
