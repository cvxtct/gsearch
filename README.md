# GSEARCH
![Coverage](https://img.shields.io/badge/Coverage-30.1%25-yellow)

[![Go](https://github.com/cvxtct/gsearch/actions/workflows/go.yml/badge.svg)](https://github.com/cvxtct/gsearch/actions/workflows/go.yml) [![Golangci-lint](https://github.com/cvxtct/gsearch/actions/workflows/main.yml/badge.svg)](https://github.com/cvxtct/gsearch/actions/workflows/main.yml) [![CodeQL](https://github.com/cvxtct/gsearch/actions/workflows/codeql.yml/badge.svg)](https://github.com/cvxtct/gsearch/actions/workflows/codeql.yml)

---

Experimental **Full Text Search** in `markdown` files written in Golang.

Inspired by [Let's build a Full-Text Search engine](https://artem.krylysov.com/blog/2020/07/28/lets-build-a-full-text-search-engine/)

## Install

Run `make all`. Don't forget to set up `config.json` beforehand.

### Config

```json
{
    "path": "/path/to/root/folder/to/md/files/",
    "show_content": true,   # show content of file
    "show_num_lines": 10    # limit lines of content
    "file_type": ".md"      # file type to index 
}
```

## Plans

- Save index to db
- Reindex if file changed
- More sophisticated document storing i.e. url recognition 
- Search result by relevancy (word frequency)
- REST API interface
- Github Gist interaface -> gist content search
