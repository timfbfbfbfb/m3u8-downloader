module github.com/http-live-streaming/m3u8-downloader

go 1.15

require (
	github.com/http-live-streaming/m3u8-downloader/dl v0.0.0-20210218104036-d0d4d45ee0bc
	github.com/http-live-streaming/m3u8-downloader/parse v0.0.0-20210218104036-d0d4d45ee0bc // indirect
	github.com/http-live-streaming/m3u8-downloader/tool v0.0.0-20210218104036-d0d4d45ee0bc // indirect
)

replace github.com/http-live-streaming/m3u8-downloader/dl => ./dl

replace github.com/http-live-streaming/m3u8-downloader/tool => ./tool

replace github.com/http-live-streaming/m3u8-downloader/parse => ./parse
