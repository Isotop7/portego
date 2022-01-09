# portego

## Description

Many people use _Archlinux_ because of its awesome ecosystem of community-curated `PKGBUILD` files in AUR.
_portego_ uses these files to create Crux-compatible port build files: `Pkgfile`.

It is written in `Go` and ***mainly a fun project***.

## Functions

Main functions are/should be:

- [x] fetch `PKGBUILD` files from AUR
- [ ] load all values from `PKGBUILD` to object-like representation
- [ ] transform these values to a compatible `Pkgfile`
- [ ] store the created file in a user given path
- [ ] output `Pkgfile` to console

## Usage

```bash
usage: portego [-h|--help] -q|--query "<value>" -o|--output "<value>"
               [-c|--console] [-d|--debug]

               Portego uses PKGBUILD files form AUR to create Pkgfiles for Crux
               

Arguments:

  -h  --help     Print help information
  -q  --query    Name of queried package from AUR
  -o  --output   Output path of generated Pkgfile
  -c  --console  Print Pkgfile to console. Default: false
  -d  --debug    Debug output. Default: false
```

## Changelog

Right now we're on `0.0.0-prealpha-testing`. Version `1.0.0` might never come. Version `0.1` should contain basic functionality.

## Sources / Links

- [AURweb RPC Interface](https://wiki.archlinux.org/title/Aurweb_RPC_interface)
