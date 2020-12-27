---

## fd workflow

fd Your 🏠 and get all files, respecting gitignore.

### Demo

<img
src="https://user-images.githubusercontent.com/28161197/103171270-fbe14380-485b-11eb-9aa2-c076a858e3b5.mov"
/>

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
