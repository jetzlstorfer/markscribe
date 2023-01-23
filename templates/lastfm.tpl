### Hi there 👋

## Favourite albums of all time 🎶

{{range favouriteAlbums 5}}
- {{.Artist.Name}} - {{.Name}}
{{- end}}


## Favourite artists of all time 👨‍🎤

{{range favouriteArtists 5}}
- {{.Name}} ({{.PlayCount}})
{{- end}}


## Favourite tracks of all time 💿

{{range favouriteTracks 5}}
- {{.Artist.Name}} - {{.Name}} ({{.PlayCount}})
{{- end}}


## Most recent tracks 🎺

{{range recentTracks 10}}
- {{.Artist.Name}} - {{.Name}}
{{- end}}
