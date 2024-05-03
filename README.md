# <p align=center>go-mp3</p>
> This is a slightly modified version of [hajimehoshi/go-mp3](https://github.com/hajimehoshi/go-mp3). <br/>An MP3 decoder in pure Go based on [PDMP3](https://github.com/technosaurus/PDMP3). 

[![Go Reference](https://pkg.go.dev/badge/github.com/imcarsen/go-mp3.svg)](https://pkg.go.dev/github.com/imcarsen/go-mp3) 

## Changes and reasoning
hajimehoshi's go-mp3 is amazing, but I had some issues with the seek functionality.

Changes made:
- Replace the Seek function with one that works within my application.
- Update to Go version 1.22

## Will I actively maintain this?
Kind of. If people have PRs or issues, I may be willing to look at/into them. But I can't promise anything.
