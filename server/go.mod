module ppamo/server

go 1.17

require (
	controllers v1.0.1 // indirect
	customerrors v1.0.1 // indirect
	entities v1.0.1 // indirect
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/mattn/go-colorable v0.1.11 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.1 // indirect
	golang.org/x/crypto v0.0.0-20210817164053-32db794688a5 // indirect
	golang.org/x/net v0.0.0-20211015210444-4f30a5c0130f // indirect
	golang.org/x/sys v0.0.0-20211103235746-7861aae1554b // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/time v0.0.0-20201208040808-7e3f01d25324 // indirect
	repositories v1.0.1 // indirect
	routers v1.0.1 // indirect
	services v1.0.1 // indirect
)

replace (
	controllers v1.0.1 => ./controllers
	customerrors v1.0.1 => ./customerrors
	entities v1.0.1 => ./entities
	repositories v1.0.1 => ./repositories
	routers v1.0.1 => ./routers
	services v1.0.1 => ./services
)
