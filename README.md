# cryptkey

`cryptkey` is a minimalist command-line utility written in Go for generating
cryptographically secure random keys. It offers options for generating
8-byte, 16-byte, or 32-byte hexadecimal keys, making it suitable for quick
key generation for various applications.

## Features

- **Secure Key Generation**: Utilizes `crypto/rand` for robust,
  cryptographically secure random key generation.
- **Flexible Key Lengths**: Supports 8, 16, and 32-byte key outputs.
- **Clipboard Integration**: Designed to be easily piped with OS-specific
  clipboard utilities for instant use.

## Usage

To generate a key, simply run the `cryptkey` executable. By default, it will
output a 32-byte hexadecimal key.

```bash
cryptkey
```

### Options

- `-8`: Generate an 8-byte key.
- `-16`: Generate a 16-byte key.
- `-32`: Generate a 32-byte key (default).

If multiple length flags are provided, the first one encountered will take
precedence.

### Examples

Generate a 16-byte key:

```bash
cryptkey -16
```

Integrate with your system's clipboard for immediate use:

- **Windows**:
  ```bash
  cryptkey -16 | clip
  ```
- **macOS**:
  ```bash
  cryptkey | pbcopy
  ```
- **Linux** (requires `xclip`):
  ```bash
  cryptkey | xclip -selection clipboard
  ```

## Installation

### From Source

1.  **Clone the repository** (or download the `main.go` file):
    ```bash
    git clone https://github.com/your-username/cryptkey.git # Replace with your repo
    cd cryptkey
    ```
2.  **Build the executable**:
    ```bash
    go build -o cryptkey main.go
    ```
3.  **Optional: Move to PATH**: For system-wide access, move the `cryptkey`
    executable to a directory included in your system's `PATH` (e.g.,
    `/usr/local/bin` on Linux/macOS, or a directory added to `Path`
    environment variable on Windows).

    ```bash
    # On Linux/macOS
    sudo mv cryptkey /usr/local/bin/
    ```
### Via _go install_
```bash
go install github.com/theaaronn/keycrypt
```

## License

This project is open source and available under the [MIT License](LICENSE).
