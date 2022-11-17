# dba-scraper

[Den Blå Avis](https://www.dba.dk/), the website we all know and love. You can now stress their servers by getting the info of a listing fast and easy with a given search link etc. "`https://www.dba.dk/soeg/?soeg=gtx+1080+ti`".

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
go run .

./scraper
```
