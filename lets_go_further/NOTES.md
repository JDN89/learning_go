# NOTES

## Added live reload with air

[Live reload for go apps](https://github.com/air-verse/air)
go get ...
go tool air init
go tool air

## Go tool info

[go tool](https://blog.howardjohn.info/posts/go-tools-command/)

## API verioning - snippet from book

API versioning

APIs which support real-world businesses and users often need to change their functionality and endpoints over time — sometimes in a backward-incompatible way. So, to avoid problems and confusion for clients, it’s a good idea to always implement some form of API versioning.

There are [two common approaches](https://stackoverflow.com/questions/389169/best-practices-for-api-versioning) to doing this:

    1. By prefixing all URLs with your API version, like `/v1/healthcheck` or `/v2/healthcheck`.
    2. By using custom Accept and Content-Type headers on requests and responses to convey the API version, like Accept: application/vnd.greenlight-v1.
