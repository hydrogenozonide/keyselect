# keyselect
Pretty much a reimplementation of [wlr-which-key](https://github.com/MaxVerevkin/wlr-which-key), written as a quick go script to my taste.

The whole idea is pretty simple - I wanted a middleground between a million keybinds in Niri, and a full launcher such as [Fuzzel](https://codeberg.org/dnkl/fuzzel). My perference would've been modal keymaps, but to my knowlege Niri doesn't currently support them. As such, I've made keyselect as a quick runner for applications I often use which aren't important enough to warrant a proper keybind, such as Genshin Impact, Steam, shortcuts to a couple IRC networks, a power menu, and the like. Of course, you can configure it to whatever you want, I'm not your boss.

# Features
- A single script!
- As few dependecies as I can reasonably handle
- Configuration in YAML, in theory backwards compatible with most of wlr-which-key's config.
- Super fast!

# Installation
Just like any other program. ``git clone`` the whole repo and build it with go. If you really need them, the commands are below.
```bash
git clone github.com/hydrogenozonide/keyselect
cd keyselect
go build keyselect.go
```
You can then take the resulting binary and stick it in your $PATH or just anywhere you want to run it from.

# Usage
``` keyselect ``` opens the menu. As of right now, no arguments are supported, it just serves whatever's in your config. You can then navigate the menu via pressing the corresponding letter key for the submenu you'd like to expand or command you'd like to run. Submenus are indicated by the submenu indicator symbol set in your config file, defaulting to +. You also have keys for exiting the menu and going back one layer of submenu, defaulting to Q and H respectively. These binds **very much can** interfere with your actual menu option binds, so be careful as that can lead to unexpected behaviour. Just set your binds correctly, please. 

# Configuration
The config is located in ```$XDG_CONFIG_HOME/keyselect/config.yaml```. You will have to create one as it's not auto-generated. A demo config is probably the easiest way I could explain this, so I'll give one below with all the options set. It should all be self-explanatory, feel free to open an issue if anything's unclear. If you want a slightly more fleshed out example, you can take a look in my personal dotfiles.

```yaml
ui:
  show_path: true # default true
  show_lines: true # default true
  show_hints: false # default true
  submenu_symbol: ">" # defaults to +

binds:
  quit: "q"
  back: "h" # these will take priority over your menu options, which can lead to unexpected behaviour. please don't raise an issue about this, just set your binds properly please

# and now we have the main bit. in theory fully compatible with wlr-show-keys!
menu:
  - key: "r"
    desc: Comms
    submenu:
      - key: "v"
        desc: Vesktop
        cmd: vesktop
      - key: "i"
        desc: IRC
        submenu: # submenus in submenus!
          - key: "l"
            desc: Liberachat
            cmd: kitty -e catgirl liberachat # commands with args are supported
          - key: "t"
            desc: Tildechat
            cmd: kitty -e catgirl tildechat
          - key: "u"
            desc: Undernet
            cmd: kitty -e catgirl undernet
  - key: "g"
    desc: Games
    submenu:
      - key: "s"
        desc: Steam
        cmd: steam
      - key: "g"
        desc: Genshin Impact
        cmd: an-anime-game-launcher
      - key: "c"
        desc: Celeste
        cmd: steam -applaunch 504230
```

Just a quick warning, I programmed this in a single 45-minute session with only a little testing. It seems fine and I've been using it for a little bit now and I'll push fixes if I find anything, but the code is honestly pretty bad.
