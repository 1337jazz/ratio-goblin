# 👺 ratiogoblin

`ratiogoblin` is a simple tool designed for use with status bars such as i3-bar, Waybar, and Polybar to get your current torrent ratio from IPTorrents.

---

## Quick Start

To get started with `ratiogoblin`, follow these steps:

1. Run `ratiogoblin init` to create a default configuration file - for most distros this will be at `~/.config/ratiogoblin/config.json`.

2. Edit the configuration file to add your IPTorrents credentials (from your cookie). Note, the `pass` field is not necessarily your account password, but the `pass` cookie value from IPTorrents.
    ```json
    {
      "uid": "your_uid",
      "pass": "your_pass"
    }
    ```

3. Run `ratiogoblin run` to see your current ratio.

4. Integrate `ratiogoblin` with your status bar of choice by adding the appropriate command to your bar's configuration.

---

## Installation

You can install `ratiogoblin` using one of the following methods:

### Pre-built Binaries

1. Download the latest binary from the [GitHub Releases](https://github.com/1337jazz/ratio-goblin/releases) page. Replace `VERSION` with the desired release version:
   ```bash
   curl -Lo ratiogoblin.tar.gz https://github.com/1337jazz/ratio-goblin/releases/download/VERSION/ratiogoblin_VERSION_linux_amd64.tar.gz
   ```
2. Extract the binary:
   ```bash
   tar -xvf ratiogoblin.tar.gz
   ```
3. Move the binary to a directory in your system's PATH (e.g., `/usr/local/bin/`):
   ```bash
   sudo mv ratiogoblin /usr/local/bin/
   ```

Now you can run the `ratiogoblin` command from anywhere in your terminal.

### Building from Source

### Prerequisites:
- [Go](https://golang.org/dl/) (version 1.23 or higher)

1. Clone the repository:
   ```bash
   git clone https://github.com/1337jazz/ratio-goblin.git
   cd ratio-goblin
   ```
2. Build the project:
   ```bash
   make
   ```
   This will create the binary at `bin/ratiogoblin`.

3. Optionally, move the binary to a global directory to make it accessible system-wide:
   ```bash
   sudo mv bin/ratiogoblin /usr/local/bin/
   ```

4. (Optional) Build with specific version information embedded:
   ```bash
   make build-version version=1.0.0-test
   ``
   Like `make build`, but also embeds the version into the binary for easier tracking.


---

## Automated Release Workflow

This project uses [Goreleaser](https://goreleaser.com/) to automate the release process. Here’s how it works:

### Prerequisites:
- You must have push access to the repository to make a release.

### Steps:

1. **Bump the Version**:
   Update the version number in `internal/version/version.go` using [semantic versioning](https://semver.org/).
   ```go
   package version

   const Version = "0.2.0" // Example of new version
   ```
2. **Commit the Version Update**:
   ```bash
   git add internal/version/version.go
   git commit -m "Bump version to 0.2.0"
   ```

3. **Tag the Release**:
   Create a new Git tag that starts with `v`, and push it to the repository.
   ```bash
   git tag v0.2.0
   git push origin v0.2.0
   ```

4. **Automated Workflow**:
   - Once the tag is pushed, the GitHub Actions workflow (`.github/workflows/release.yml`) will trigger automatically.
   - Goreleaser will:
     - Build the binaries for the project.
     - Package the application into release-ready archives (e.g., `.tar.gz`).
     - Create a release on GitHub with the versioned tag.
     - Generate a changelog from commit history.

5. **Release Completed**:
   - Upon successful completion, the release will be available on the [GitHub Releases](https://github.com/1337jazz/ratio-goblin/releases) page complete with downloadable binaries and changelogs.

---

## Contributing

Contributions are welcome! Please submit pull requests or open issues for any suggestions, bug reports, or improvements.
