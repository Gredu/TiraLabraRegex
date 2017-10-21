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


## Testien toteutus

Ohjelma koostuu kolmesta funktiosta, `parseRegexp()`, `generateMachine()` ja `match()`. Funktiota `generateMachine()` on tukemassa funktio `generateTransition()`. Koska funktiot toimivat erittäin lähekkäin, on ne päätetty sijoittaa samaan tiedostoon. Samalla tavalla myös `isDigit()` on funktion `match()` tukena ja samassa tiedostossa.

Projektista löytyy kolme testitiedostoa, jotka on nimetty pääfunktioiden mukaan. Esimerkiksi `generateMachine_test.go`. Tiedosto sisältää myös testin sitä tukevalle apufunktiolle `generateTransition()`.


### generateMachine_test.go

Testaa, että ohjelma rakentaa automaatin oikein. Tähän tiedostoon on käytetty hyvin vähän automaatiota verrattuna muihin testitiedostoihin. Automaatti siis luodaan ja sitä testaan "käsin", että tilat ja niiden siirtymät ovat varmasti oikeita. Koska kaikki on tehty manuaalisesti, ei testejä ole paljoa, vaikka koodin määrä näyttää suurelta.

Tähän ratkaisuun ollaan tultu, koska testausohjelman rakentaminen on jo sen verran monimutkaista, että siinä samalla oikeastaan tulisi rakennettua jo itse ohjelman...


### parseRegexp_test.go

Testataan, että funktio parseeraa säännöllisen lausekkeen oikein palautettavaksi Tokeneiksi array/slice muodossa.

Testejä on kumpaakin, sekä automaattisesti, että manuaalisesti. Osa säännöllisistä lauseista tekee jotain hieman epäsäännöllisempää, joita on helpompi testata manuaalisesti.


### match_test.go

Kahta edelistä on testattu hyvin, mutta tiedostoa `match.go` on testattu todella paljon. Syy tähän on, että testaus on helppo tehdä Table Driven Test periaatteella. On vain testattava, mitä syötteitä säännölliset lauseet hyväksyvät, ja tätä on helppo automatisoida. Käytännössä luodaan struct, joka sisältää sisään tulon (eng. inputs) ja halutun tuloksen. Nämä sitten iteroidaan läpi.

Tiedostosta näkee hyvin selkeästi minkälaisia ilmaisuja ohjelma tällä hetkellä hyväksyy.

Funktion `match()` testaaminen on ollut helppo tehdä TDD-menetelmällä, koska testitiedostoja on ollut helppo kirjoittaa. Sekin on ollut yksi syy, miksi tiedosto on paisunut testeistä.

Lähde: https://github.com/golang/go/wiki/TableDrivenTests
