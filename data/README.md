# Data

### Mod requirements
- The mod must be **stable** with as few open issues as possible
- The mod must **fully support** the current version
- The mod should provide **permission support**
- The mod should be compatible with other mods (TODO: Add a notice?)

### Usage
- Mod info is automatically downloaded from Modrinth, just writes down its `slug`
E.g: `https://modrinth.com/mod/fabric-api` has slug `fabric-api`
- If the mod is not listed in Modrinth, writes its info to `projects.custom.json`. The file is used for all versions.

### Versions

The latest version = the default file: `mods.default.json`
<br>

For older Minecraft versions:
```
mods.1.19.2.json
mods.1.19.json
...
mods.x.y.z.json
```
