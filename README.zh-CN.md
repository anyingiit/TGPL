[English](README.md) · **简体中文**

> 英文版是规范版本。本页与 [README.md](README.md) 不一致时，以英文版为准。

<!-- translation-of: README.md sha256:7d397bd90bfe7238 -->

<!-- Source: Best-README-Template BLANK_README (Unlicense) — https://github.com/othneildrew/Best-README-Template -->
<a id="readme-top"></a>

# TGPL

一个精简、已归档的 Go 模块，其唯一的程序会向标准输出打印一行写死的 Hello World 问候语。

[![CI](https://github.com/anyingiit/TGPL/actions/workflows/ci.yml/badge.svg)](https://github.com/anyingiit/TGPL/actions/workflows/ci.yml)
[![License](https://img.shields.io/github/license/anyingiit/TGPL)](LICENSE)

[报告问题](https://github.com/anyingiit/TGPL/issues/new?template=bug_report.yml) · [提出需求](https://github.com/anyingiit/TGPL/issues/new?template=feature_request.yml)

<details>
  <summary>目录</summary>
  <ol>
    <li><a href="#about-the-project">关于本项目</a></li>
    <li><a href="#getting-started">开始使用</a></li>
    <li><a href="#usage">用法</a></li>
    <li><a href="#contributing">参与贡献</a></li>
    <li><a href="#license">许可证</a></li>
    <li><a href="#contact">联系方式</a></li>
  </ol>
</details>

## 关于本项目

这个模块（`github.com/anyingiit/TGPL`，见 `go.mod`）只包含一个程序：`8_1/main.go`，
它的全部内容就是一次 `fmt.Println` 调用，打印出 `Hello World!`。除此之外没有其他
源码，没有测试，也没有第二个包。

该仓库已被归档，因此它大概率不会再扩展到这一个文件之外。仓库中还提交了一个预先
编译好的 Windows 二进制文件 `8_1/8_1.exe`，它不是由 CI 生成的，不应被当作与
`main.go` 保持同步的产物。

计划中的功能与已知问题，见 [open issues](https://github.com/anyingiit/TGPL/issues)。

## 开始使用

### 环境要求

- Go 1.18 或更高版本，即 `go.mod` 声明的下限

### 安装

```sh
git clone https://github.com/anyingiit/TGPL.git
cd TGPL
go build -o bin/8_1 ./8_1/...
```

## 用法

运行编译出的二进制文件，它会打印这一行内容后退出：

```sh
./bin/8_1
# Hello World!
```

## 参与贡献

欢迎参与。[CONTRIBUTING.md](CONTRIBUTING.md) 说明如何提交 issue 或 pull request，[CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md) 说明对所有参与者的行为要求。

请不要在公开的 issue 或 pull request 中报告安全问题。[SECURITY.md](SECURITY.md) 说明了私下报告的方式。

## 许可证

以 MIT 许可证分发。详见 [LICENSE](LICENSE)。

## 联系方式

项目地址：[https://github.com/anyingiit/TGPL](https://github.com/anyingiit/TGPL)

<p align="right">(<a href="#readme-top">back to top</a>)</p>
