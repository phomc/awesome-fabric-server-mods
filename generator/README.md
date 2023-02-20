# Generator

The generator automatically parses the `data` folder, fetches information from Modrinth and writes the `README.md` file.

## CLI
```
generator \
-template "template.md" \
-custom "../data/mods.custom.json" \
-data "../data/mods.default.json" \
-cache "cache.json" \
-out "../README.md"
```

All arguments are optional.
