<!-- Source: Best-README-Template BLANK_README (Unlicense) — https://github.com/othneildrew/Best-README-Template -->
<a id="readme-top"></a>

# TGPL

A minimal, archived Go module whose single program in 8_1/main.go prints a hard-coded Hello World greeting to standard output.

**English** · [简体中文](README.zh-CN.md)

[![CI](https://github.com/anyingiit/TGPL/actions/workflows/ci.yml/badge.svg)](https://github.com/anyingiit/TGPL/actions/workflows/ci.yml)
[![License](https://img.shields.io/github/license/anyingiit/TGPL)](LICENSE)

[Report a bug](https://github.com/anyingiit/TGPL/issues/new?template=bug_report.yml) · [Request a feature](https://github.com/anyingiit/TGPL/issues/new?template=feature_request.yml)

<details>
  <summary>Table of Contents</summary>
  <ol>
    <li><a href="#about-the-project">About The Project</a></li>
    <li><a href="#getting-started">Getting Started</a></li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>

## About The Project

This module (`github.com/anyingiit/TGPL`, `go.mod`) holds one program: `8_1/main.go`,
whose entire body is a call to `fmt.Println` that prints `Hello World!`. There is no
other source, no tests, and no second package.

The repository is archived, so this is unlikely to grow past that single file. A
pre-built Windows binary, `8_1/8_1.exe`, is committed next to the source; it is not
produced by CI and should not be treated as up to date with `main.go`.

See the [open issues](https://github.com/anyingiit/TGPL/issues) for planned features and known issues.

## Getting Started

### Prerequisites

- Go 1.18 or newer, the floor `go.mod` declares

### Installation

```sh
git clone https://github.com/anyingiit/TGPL.git
cd TGPL
go build -o bin/8_1 ./8_1/...
```

## Usage

Run the built binary; it prints its one line and exits:

```sh
./bin/8_1
# Hello World!
```

## Contributing

Contributions are welcome. Read [CONTRIBUTING.md](CONTRIBUTING.md) for how to open an issue or a pull request, and [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md) for the standards expected of everyone taking part.

Please do not report security issues in public issues or pull requests. [SECURITY.md](SECURITY.md) explains how to report them privately.

## License

Distributed under the MIT License. See [LICENSE](LICENSE) for details.

## Contact

Project link: [https://github.com/anyingiit/TGPL](https://github.com/anyingiit/TGPL)

<p align="right">(<a href="#readme-top">back to top</a>)</p>
