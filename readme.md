# film

🎞 _a photoblog generator_

inspired by photoblog generators:

1. [moul](https://github.com/moul-co/moul)
2. [foto](https://github.com/waynezhang/foto)
3. [1600pr.sh](https://github.com/andersju/1600pr.sh)
4. [Exposé](https://github.com/Jack000/Expose)
5. [Sigal](https://github.com/saimn/sigal/)

Dead simple go code. Photo Gallery like Polaroid(?)
Use go with following deps:

- command line arguments
  - https://pkg.go.dev/flag
  - https://github.com/alecthomas/kong
  - https://github.com/spf13/cobra
- image
  - https://pkg.go.dev/image
  - https://github.com/anthonynsimon/bild
- html
  - https://pkg.go.dev/html/template
  - https://github.com/a-h/templ
  - https://github.com/osteele/liquid
  - https://github.com/valyala/fasttemplate
- storage
  - https://pkg.go.dev/encoding/json
  - https://github.com/jameycribbs/hare
  - https://github.com/ostafen/clover
  - https://github.com/akyoto/cache
- config
  - https://github.com/spf13/viper
  - https://github.com/tucnak/store

more help: https://github.com/avelino/awesome-go

## goals

- simple/easy
- fast
- small

## commands

### create

create a foto project

```bash
$ film create --help
```

Ask for specifics and so on like: "Name of your new project: `user input`",
wait for user input and then execute these.

Ask for:

- project name

### preview

open server and show the pictures

### export

build the blog into a directory

```bash
$ film export --help
```

## generated site

```bash
$ tree
.
├── _site/
└── film

```

### home

The most recent picture

### _&lt;number&gt;_

The site that generates the picture with _number_

### archive

An archive that displays all the pictures in a masonry-like grid

![example](example.jpg)
