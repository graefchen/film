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
- storage
  - https://pkg.go.dev/encoding/json
  - https://github.com/jameycribbs/hare
  - https://github.com/ostafen/clover
  - https://github.com/akyoto/cache

more help: https://github.com/avelino/awesome-go

## commands

### version

print version

### init

initialize a foto project

### add

add image or directory to database

`-t | --title` to give a title (else the image name is choosen)
`-d | --date` add an custome date

### remove

removes an image

### list

list currently used images

### build

build the blog

<!-- ### serve

build and serve the blog on localhost -->
