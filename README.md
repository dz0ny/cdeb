# cdeb

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
