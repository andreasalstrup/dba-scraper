# dba-scraper

[Den Blå Avis](https://www.dba.dk/), the website we all know and love. You can now stress their servers by getting the complete list of items, fast and easily with a given search link eg. "`https://www.dba.dk/soeg/?soeg=gtx+1080+ti`".

The list of scraped items will be sorted by price and appear in [export.csv](export.csv)

## Setup
Install dependencies

```
go mod tidy
```

Run scraper

```
go run .
```

(Optional) Compile and run

```
go build .

./scraper
```
