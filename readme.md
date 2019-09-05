# go-nrutil

[![release](https://img.shields.io/badge/Experimental-v0.1.0-yellow.svg)](https://github.com/brettski/go-nrutil) [![Go Report Card](https://goreportcard.com/badge/github.com/brettski/go-nrutil)](https://goreportcard.com/report/github.com/brettski/go-nrutil)

## New Relic utility

Cli to read and write Synthetics scripts to local files.  It's main purpose is to make it easier to store Synthetics scripts in source control.

## Shout-out and Thank You

A thank you to the great Go community and the individuals and groups who provide and maintain the libraries we use! We really appreciate all your efforts.

Modules used in this project:

* Mow.Cli - A versatile library for building CLI applications in Go
  * https://github.com/jawher/mow.cli
* go-homedir - Go library for detecting and expanding the user's home directory without cgo.
  * github.com/mitchellh/go-homedir
* go-yaml - YAML support for the Go language.
  * https://github.com/go-yaml/yaml
* Termtables - Go ASCII Table Generator, ported from the Ruby terminal-tables library
  * ~~github.com/apcera/termtables~~
  * github.com/brettski/go-termtables
    * Apcera deleted termtables from GitHub so I restored the repo for others to use

## Usage

nrutil is a single executable cli which synchronizes New Relic Synthetics scripted-typed checks with a local file system. This allows you to manage these scripts in some type of source control (e.g. Git). New Relic Synthetics itself has no observance of version history, etc.

nrutil has inline help. Add -h to any command to get details about it. E.g. `nrutil getmonitorlist -h`

**nrutil requires a configuration file in your $HOME directory (C:\Users\me, ~/, etc.). You can create this file by running:**

`nrutil config create -nradminkey <Your New Relic Admin Key>`

The configuration file manages the following data:

```yaml
---
nradminkey: <your-admin-key>
basepath: ~/nrsynthetics
syntheticmonitors
- guid-of-monitor-1-23456
- guid-of-monitor-2-34567
- guid-of-monitor-n-opqrs
```

* **nradminkey**        - The admin API key from New Relic. Per NR: "You must use your Admin User's API key to make Synthetics REST API calls. The account's REST API key will not work."
* **basepath**          - The base path to store the synthetic scripts.
* **syntheticmonitors** - (not yet implemented) A list of monitors the cli is managing.

In this early stage of nrutil, the list of managed Synthetics *isn't implemented*. You can however manage Synthetic monitors individually:

`nrutil getscript --id <Synthetics guid id>`  
`nrutil setscript --id <Synthetics guid id>`

These commands will store the file in the basepath set in the configuration file allowing you to manage them in any source control.

Add any questions or bugs discovered in the [GitHub Issue Tracker](https://github.com/brettski/go-nrutil/issues)
