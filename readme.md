# film

_a photoblog generator_

inspired by photoblog generators:

- [1600pr.sh](https://github.com/andersju/1600pr.sh)
- [Exposé](https://github.com/Jack000/Expose)
- [Sigal](https://github.com/saimn/sigal/)
- [foto](https://github.com/waynezhang/foto)

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

more help: https://github.com/avelino/awesome-go

## commands

### version

print version

```bash
$ film
```

### help

The help message for the command

```bash
$ film help
Create a fotoblog
```

### create

create a foto project

```bash
$ film create --help
```

Ask for specifics and so on like: "Name of your new project: `user input`", wait for user input and then execute these.

### add

add image or directory to database

```bash
$ film add --help
```

`-t | --title` to give a title (else the image name is choosen)
`-d | --date` add an custome date

### remove

removes an image

```bash
$ film remove --help
```

### list

list currently used images

```bash
$ film list --help
```

### build

build the blog into a directory

```bash
$ film build --help
```

## generated site

### home

The most recent picture

### &lt;id&gt;

The site that generates the picture of id

### archive

An archive that displays all the pictures in a masonry-like grid
