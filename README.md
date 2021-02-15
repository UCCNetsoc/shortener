# Shortener

[![Build Status](https://ci.netsoc.co/api/badges/UCCNetsoc/shortener/status.svg)](https://ci.netsoc.co/UCCNetsoc/shortener)

A URL shortener written in go

|METHOD|PATH|Description|
|--|--|--|
|Get|/{slug}|HTTP Redirect to resolved URL from db|
|POST|/api|Create slug/url pair on db {slug: string, url: string}|
|DELETE|/api/{slug}|Delete slug/url pair on db|