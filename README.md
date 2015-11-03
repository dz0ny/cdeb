# cdeb

[![wercker status](https://app.wercker.com/status/1214b503d63d0bf2178dda7373d98983/m "wercker status")](https://app.wercker.com/project/bykey/1214b503d63d0bf2178dda7373d98983)

You can find precompiled binaries under *Releases*.

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
