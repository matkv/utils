# Utils

A CLI tool to make some common personal tasks easier. For now, this is primarly a practice project to learn Go using some real-life problems. Ideally this tool will become some sort of personal "pocket-knife" that helps me automate some day to day tasks - especially since Go allows to build any program as a single static binary so I should be able to run it anywhere I need it without any set-up process.

The first problem I want to tackle is that I want to automate some things I have to do manually when editing content for my hugo website.

**Hint**: To create subcommands using `cobra-cli`: (for example creating a command that will be called as `utils hugo markdown-link-checker`) use `cobra-cli add markdown-link-checker -p hugoCmd` - here `hugoCmd` is the parent command.

## Installation

### Creating a tag

In order to install utils using the command below, I need to create a tag for every new version.

```bash
git tag v0.0.1
git push origin v0.0.1
```

Without the tag, I can't install it using @latest.

### Installing

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
configType: "" # "windows" or "archlinux"

windows:
  dotfiles:
    path: ""
  obsidian:
    vaultpath: ""

archlinux:
  dotfiles:
    path: ""
```
