# Alfred workflows
  <div>
    <a href="/LICENSE.md"><img alt="GitHub" src="https://img.shields.io/github/license/kgnugur/tureng-alfred"></a>
  </div>

---

## fd workflow

<img alt="Gif"
src="https://user-images.githubusercontent.com/28161197/65831745-d61a6700-e2c5-11e9-9f49-05b8b95ad9e9.gif"
width="60%" />

fd Your 🏠 and get all of your files respecting gitignore.

### Features
- Blazing fast
- Uses fd for fetching files.
- Fuzzy filtering is supported.
- Native alfred file view (Not recommended as it diminishes fuzzy filtering
functionality)
- Full path view. (Recommended)
- Respecting ignore file when searching.

### Usage
- If you want to change full path to files view set `fileView` workflow
  environment to `true`
- 'Shift' action modifier opens file in `nvim` by default change
  `terminalEditor` environment value if you want different terminal editor.


### Requirements

Comes with single binary. [fd](https://github.com/sharkdp/fd) is needed.

---

## Tureng workflow

<img alt="Gif" src="https://user-images.githubusercontent.com/28161197/65831745-d61a6700-e2c5-11e9-9f49-05b8b95ad9e9.gif" width="60%" />

Search in tureng with search suggestions.


### Requirements

Comes with single binary only requirement is [official Tureng
App](https://apps.apple.com/tr/app/id854063979?mt=12) for native experience.
