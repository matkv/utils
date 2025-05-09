# Utils

<p align="center">
  <img src="https://img.shields.io/github/go-mod/go-version/matkv/utils" alt="Go Version">
  <img src="https://img.shields.io/github/license/matkv/utils" alt="License">
  <img src="https://img.shields.io/github/last-commit/matkv/utils" alt="Last Commit">
  <a href="https://pkg.go.dev/github.com/matkv/utils">
    <img src="https://pkg.go.dev/badge/github.com/matkv/utils.svg" alt="Go Reference">
  </a>
</p>

A CLI tool to make some common personal tasks easier & a practice project for learning Go.

## Installation

On Linux I can then install the latest version of `utils` by running:

```bash
go install github.com/matkv/utils@latest
```

If I want to install it on the machine I'm currently on this project - in the project directory, I can also just run:

```bash
go install
```

## Configuration

A config file will be created in ~/.config/utils/config.yaml. Or just use the one from the dotfiles repo.

Currently, the config file looks like this:

```yaml
configType: "" # "windows" or "linux"

windows:
  dotfiles:
    path: ""
  obsidian:
    vaultpath: ""

linux:
  dotfiles:
    path: ""
```

## Additional info

### Creating a tag

In order to install utils using the command below, I need to create a tag for every new version.

```bash
git tag v0.0.1
git push origin v0.0.1
```

Without the tag, I can't install it using @latest.

### Cobra-CLI commands

To create subcommands using `cobra-cli`: (for example creating a command that will be called as `utils hugo markdown-link-checker`) use `cobra-cli add markdown-link-checker -p hugoCmd` - here `hugoCmd` is the parent command.