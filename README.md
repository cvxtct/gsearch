# GSEARCH
![Coverage](https://img.shields.io/badge/Coverage-26.0%25-red)
[![Go](https://github.com/cvxtct/gsearch/actions/workflows/go.yml/badge.svg)](https://github.com/cvxtct/gsearch/actions/workflows/go.yml) [![Golangci-lint](https://github.com/cvxtct/gsearch/actions/workflows/main.yml/badge.svg)](https://github.com/cvxtct/gsearch/actions/workflows/main.yml) [![CodeQL](https://github.com/cvxtct/gsearch/actions/workflows/codeql.yml/badge.svg)](https://github.com/cvxtct/gsearch/actions/workflows/codeql.yml)

---

Experimental **Full Text Search** in `markdown` files written in Golang.

Inspired by [Let's build a Full-Text Search engine](https://artem.krylysov.com/blog/2020/07/28/lets-build-a-full-text-search-engine/)

## Install / Uninstall

Create config.json in `configs` folder first, fill accordingly.
Run `make all`.

To remove package run `make clean`.



```Json
{
    "path": "/path/to/root/",
    "show_content": true,
    "show_num_lines": 10,
    "file_type": ".md"
}
```