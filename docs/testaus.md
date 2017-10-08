# Testaus

Muista asentaa Go ennen koodin testausta!

Testauksessa ja testikattavuudessa käytetään Go:n omaa testijärjestelmään. Kaikki komennot on annettava projektin juurella.

Testin suorittaminen:

```sh
go test
```

Testaus ja kattavuus saadaan suorittamalla:

```sh
go test -cover
```

Go:n oma testijärjestelmä antaa mahdollisuuden nähdä mitkä osat lähdekoodista on testattu ja mitkä ei. Terminaaliversiossa nähdään funktiokohtaiset kattavuudet:

```sh
go tool cover -func=coverage.out
```

HTML-versiossa nähdään tarkemmin ne kohdat lähdekoodista, joita ei ole testattu:

```sh
go tool cover -html=coverage.out
```

Kahdessa edellisessä esimerkissä tarvittiin tiedostoa `coverage.out`. Tämä tiedosto muodostetaan komennolla:

```sh
go test -coverprofile=coverage.out
```

Komento tuskin on käyttäjälle hyödyllinen, koska jokaisen github-päivityksen yhteydessä tämä tiedosto päivitetään ohjelmoijan toimesta kuitenkin. Jos omia testejä lisää tai poistaa, kannattaa se kuitenkin päivittää. Huomaa, että `coverage.out` on testitiedosto, joka siis sisältää testi-tulokset. Sen avaaminen `go tool cover` työkalulla ei itsestään käynnistä minkäänlaisia testejä.

On myös mahdollista saada testikattavuus lämpökarttana (eng. heat map), jolloin nähdään, kuinka hyvin tietyt osat lähdekoodista on testattu. Tiedosto `coverage.out` on kuitenkin päivitettävä lisäämällä `-covermode=count`:

```sh
go test -covermode=count -coverprofile=coverage.out
```

Tämän jälkeen tiedosto avataan HTML-versiona samalla tavalla kuin aikaisemmin jo nähtiin.


## Suorituskykytestaus

Suorituskykytestaus käynnistetään komennolla:

```sh
go test -bench=
```
