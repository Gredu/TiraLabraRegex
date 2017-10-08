# Testaus

Muista asentaa Go ennen koodin testausta!

Testauksessa ja testikattavuudessa käytetään Go:n omaa testijärjestelmään. Kaikki komennot on annettava projektin juurella.

Testin suorittaminen:

```go
go test
```

Testaus ja kattavuus saadaan suorittamalla:

```go
go test -cover
```

Go:n oma testijärjestelmä antaa mahdollisuuden nähdä mitkä osat lähdekoodista on testattu ja mitkä ei. Terminaaliversiossa nähdään funktiokohtaiset kattavuudet:

```go
go tool cover -func=coverage.out
```

HTML-versiossa nähdään tarkemmin ne kohdat lähdekoodista, joita ei ole testattu:

```go
go tool cover -html=coverage.out
```

Kahdessa edellisessä esimerkissä tarvittiin tiedostoa `coverage.out`. Tämä tiedosto muodostetaan komennolla:

```go
go test -coverprofile=coverage.out
```

Komento tuskin on käyttäjälle hyödyllinen, koska jokaisen github-päivityksen yhteydessä tämä tiedosto päivitetään kuitenkin. Jos omia testejä lisää tai poistaa, kannattaa tämä kuitenkin päivittää.

On myös mahdollista saada testikattavuus lämpökarttana (eng. heat map), jolloin nähdään, kuinka hyvin tietyt osat lähdekoodista on testattu. Tiedosto `coverage.out` on kuitenkin päivitettävä lisäämällä `-covermode=count`:

```go
go test -covermode=count -coverprofile=coverage.out
```

Tämän jälkeen tiedosto avataan HTML-versiona samalla tavalla kuin aikaisemmin jo nähtiin.
