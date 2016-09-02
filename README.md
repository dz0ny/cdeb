# cdeb

[![wercker status](https://app.wercker.com/status/1214b503d63d0bf2178dda7373d98983/m "wercker status")](https://app.wercker.com/project/bykey/1214b503d63d0bf2178dda7373d98983)
[![Go Report Card](https://goreportcard.com/badge/github.com/dz0ny/cdeb)](https://goreportcard.com/report/github.com/dz0ny/cdeb)


```
usage: cdeb [<flags>] <command> [<args> ...]

Build packages for debian with great ease or declaritive template.

Flags:
  --help         Show context-sensitive help (also try --help-long and --help-man).
  -v, --verbose  Verbose mode.
  --version      Show application version.

Args:
  <root>     Path to deb root folder
  <control>  Path to debian control folder
  <deb>      Path to output deb file

Commands:
  help [<command>...]
    Show help.


  dumb <root> <control> <deb>
    Create debian file from root and control folder.


  gen <root> <deb>
    Create debian file from root folder and .cdeb template.


  init
    Create .cdeb template.

```

You can find compiled binaries for your platform under "Releases" or if you prefer quick install:

```
$ curl -L https://github.com/dz0ny/cdeb/releases/download/v1.0.0/cdeb-`uname -s`-`uname -m` > /usr/local/bin/cdeb
$ chmod +x /usr/local/bin/cdeb
```


cdeb is sponsored by [NiteoWeb Ltd.](http://www.niteoweb.com/)
