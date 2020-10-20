# markdown-codeblocks-processor

## Table of Contents

* [About markdown-codeblocks-processor](#about-markdown-codeblocks-processor)
  * [Specifications](#specifications)
* [Getting Started](#getting-started)
  * [Prerequisites](#prerequisites)
  * [Installation](#installation)
  * [Testing](#testing)
  * [Usage](#usage)
* [Contact Us](#contact-us)
* [Authors](#authors)
* [Licence](#licence)
* [Acknowledgments](#acknowledgments)
* [References](#references)

## About markdown-codeblocks-processor

markdown-codeblocks-processor is a golang software to execute codeblocks with the specified interpretor inside your markdown files

List elements represent a task
If your task is link to markdownfile, it will be included

### Specifications

You can find [here](SPECIFICATIONS.md), more information about what the final product should look like

## Getting Started

### Prerequisites

In order to build the program from sources, you need go to be installed on your computer
You can verify it by running:

```sh
go version
```

### Installation

```sh
git clone https://github.com/thdelmas/markdown-codeblocks-processor # Fetch source files

```

### Testing

How does someone test the code ?

```sh
cd markdown-codeblocks-processor # Go in this famous directory
go build &&
./markdown-codeblocks-processor --playbook ./tests_files --execute
```

### Usage

How does someone use the code ?

| Flag | Type | Usage |
|------|------|-------|
| `--playbook` | Path | The file you want to parse |
| `--execute` | Boolean | Whether you want to execute code bloc or not |

## Contact Us

<jerry@42.fr>

## Authors

* **Jerry** - *Initial design*

See also the list of [contributors](https://github.com/42paris/markdown-codeblocks-processor/graphs/contributors) who participated in this project.

## Licence

License - see the [LICENSE](LICENSE) file for details

## Acknowledgments

Thanks to all people spending their time to make a better world

## References

- [Readme boilerplate](https://github.com/thdelmas/better-README)
