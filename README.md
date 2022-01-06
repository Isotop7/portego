# portego

## Description

Many people use arch because of its awesome ecosystem of community-curated `PKGBUILD` files in AUR.
_portego_ uses these files to create Crux-compatible `Pkgfile`s.

It is written in `Go` and mainly a fun project.

## Functions

Main functions are/should be:

- fetch `PKGBUILD` files from AUR
- load all values from `PKGBUILD` to object-like representation
- transform these values to a compatible `Pkgfile`
- store the created file in a user given path

## Changelog

Right now we're on `0.0.0-prealpha-testing`. Version `1.0.0` might never come. 
