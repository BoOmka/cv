module github.com/boomka/cv

go 1.19

require (
	github.com/d0sbit/werr v0.0.0-20200314182725-e6d64c4e19e3
	github.com/gomarkdown/markdown v0.0.0-20221013030248-663e2500819c
	github.com/vugu/vgrouter v0.0.0-20200725205318-eeb478c42e5d
	github.com/vugu/vjson v0.0.0-20200505061711-f9cbed27d3d9
	github.com/vugu/vugu v0.3.4
)

require (
	github.com/vugu/html v0.0.0-20190914200101-c62dc20b8289 // indirect
	github.com/vugu/xxhash v0.0.0-20191111030615-ed24d0179019 // indirect
)

replace github.com/vugu/vugu v0.3.4 => github.com/boomka/vugu v0.0.0-20230115032035-bdecd1cdc4de
