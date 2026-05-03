module github.com/kaua-nasc/gymtrack-go/apps/api/identity

go 1.26.2

replace github.com/kaua-nasc/gymtrack-go/libs/db => ../../../libs/db

replace github.com/kaua-nasc/gymtrack-go/libs/auth => ../../../libs/auth

require (
	github.com/gin-gonic/gin v1.12.0
	github.com/golang-jwt/jwt/v5 v5.3.1
	github.com/kaua-nasc/gymtrack-go/libs/auth v0.0.0-00010101000000-000000000000
	github.com/kaua-nasc/gymtrack-go/libs/db v0.0.0-00010101000000-000000000000
	go.uber.org/fx v1.24.0
	golang.org/x/crypto v0.50.0
)

require (
	github.com/gabriel-vasile/mimetype v1.4.12 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.30.1 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/lib/pq v1.12.3
	golang.org/x/sys v0.43.0 // indirect
	golang.org/x/text v0.36.0 // indirect
)
