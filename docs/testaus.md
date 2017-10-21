# Testaaminen

Muista asentaa Go ennen koodin testausta!

Testauksessa, testikattavuudessa ja suorituskykytestauksessa käytetään Go:n omaa testijärjestelmään. Kaikki komennot on annettava projektin juurella. [Suorituskykymittaus löytyy muualta](docs/suorituskyky.md).


## Testaus

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


## Testikattavuustulokset

Kuten edellä on jo todettu, testikattavuus saadaan komennolla:

```sh
go tool cover --func=coverage.out
```

Tämä antaa testikattavuuden funktioille:

```sh
github.com/Gredu/TiraLabraRegex/generateMachine.go:3:   generateMachine         100.0%
github.com/Gredu/TiraLabraRegex/generateMachine.go:23:  generateTransition      100.0%
github.com/Gredu/TiraLabraRegex/main.go:9:              main                    0.0%
github.com/Gredu/TiraLabraRegex/main.go:30:             matchLine               0.0%
github.com/Gredu/TiraLabraRegex/match.go:3:             match                   100.0%
github.com/Gredu/TiraLabraRegex/match.go:34:            isDigit                 100.0%
github.com/Gredu/TiraLabraRegex/parseRegexp.go:7:       parseRegexp             100.0%
total:                                                  (statements)            79.3%
```

Funktiot main ja matchLine eivät ole testattuja sen takia, että ne sisältävät paljon sellaista koodia, joka on todennäköisesti jo testattu Googlen toimesta. Testaamaton koodi liittyy suurimmaksi osaksi tiedoston avaamiseen, lukemiseen ja sulkemiseen, eikä siis mitään säännöllisten lausekkeiden tulkkiin liittyvää koodia suoranaisesti.
