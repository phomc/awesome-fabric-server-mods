# Awesome Fabric server-side mods

A curated list of awesome Fabric server-side mods.

![](https://img.shields.io/badge/Fabric-{{ .DataSet.Version }}-blue)

{{ range $key, $value := $.DataSet.Data }}
## {{ $key }}
| Name  | Description | Links |
|-------|-------------|-------|
{{- range $mod := $value}}
{{- $project := (index $.Projects.Data $mod)}}
| {{ if ne $project.Slug "" }} [{{ $project.Title }}](https://modrinth.com/mod/{{ $project.Slug }}) {{ else if eq $project.Slug "" }} {{ $project.Title }} {{ end }} | {{ $project.Description }} | {{ if ne $project.SourceURL "" }} [\[GitHub\]]({{ $project.SourceURL }}) {{ end }} {{ if ne $project.WikiURL "" }} [\[Wiki\]]({{ $project.WikiURL }}) {{ end }} {{ if ne $project.DiscordURL "" }} [\[Discord\]]({{ $project.DiscordURL }}) {{ end }} {{ if ne $project.IssuesURL "" }} [\[Issues\]]({{ $project.IssuesURL }}) {{ end }} |
{{- end}}
{{ end }}
