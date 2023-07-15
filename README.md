# GSEARCH

[![Go](https://github.com/cvxtct/gsearch/actions/workflows/go.yml/badge.svg)](https://github.com/cvxtct/gsearch/actions/workflows/go.yml)

Experimental **Full Text Search** in `markdown` files written in Golang.

Inspired by [Let's build a Full-Text Search engine](https://artem.krylysov.com/blog/2020/07/28/lets-build-a-full-text-search-engine/)

## Install

Run `make all`. Don't forget to set up `config.json` beforehand.

### Config

```json
{
    "path": "/path/to/root/folder/to/md/files/",
    "show_content": true, # show content of file
    "show_num_lines": 10 # limit lines of content
}
```