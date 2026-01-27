# ratio-goblin

The `ratio-goblin` is a simple tool designed for use with status bars such as i3-bar, Waybar, and Polybar to get your current torrent ratio from IPTorrents.

---

## Usage:

Run the following commands in your terminal:

**Help**
```
$ ratiogoblin help
```
Displays usage information and available commands.

**Initialize Configuration**
```
$ ratiogoblin init
```
Creates a default configuration file in your system's user config directory under the application name.

**Run Scraper**
```
$ ratiogoblin run
```
Executes the ratio scraper and displays the output.

**Check Version**
```
$ ratiogoblin version
```
Displays the current version of the application.

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
