# Viikkoraportti 3

  - tunteja käytetty n. 8


## Mitä olen tehnyt tällä viikolla?

Regexp parseeraaja, joka luo `tokens Token` arrayn annestusta säännöllisestä lauseesta. Array sisältää regexpin, joka on nyt hieman helpommin käsiteltävässä muodossa. Ilmaisu on yhdessä solussa `Token` rakenteessa. Esimerkiksi tähtioperaatio `*` on voinut olla vaikea lukea suoraan säännöllisestä lauseesta, koska on aina pitänyt tarkistaa taaksepäin mitä kirjainta tähtioperaatiolla tarkoitetaan. Toisaalta taas `\\` kohdalla on pitänyt katsoa eteenpäin selvitettäessä mistä metasta on kysymys.

Automaatin luova funktio käyttää nyt säännöllisen lauseen sijasta tätä juuri luotua `tokens` arraytä, josta se luo automaatin.

Olen lukenut seuraavat artikkelit:
  - http://www.cs.princeton.edu/courses/archive/spr09/cos333/beautiful.html
  - https://blog.golang.org/go-slices-usage-and-internals

Ja tietenkin lisää go-kielestä.


## Miten ohjelma on edistynyt?

Hitaasti. Toisaalta taas on sellainen tunne, että kun vaikeimmat asiat on ohjelmoitu, niin triviaalia koodia syntyy sitten todella paljon ja nopeasti.


## Mitä opin tällä viikolla / tänään?

Slice, Go:n suosittelema tietorakenne. Yllätyin siitä, että listaa ei juurikaan suositella. Keskustelupalstoilla puhutaan siitä, että go:n `list` tietorakenne voisi ihan hyvin julistaa deprekoiduksi. Syy miksi se on vielä olemassa, johtuu siitä, että muilta kieliltä go-kieleen siirtyvillä voi olla kotoisampi olo tuttujen listojen kanssa.

Slice on arrayhin perustuva tietorakenne. Käytännössä slice on filtteri tai näkymä, jonka avulla näytetään vain haluttuja osia arraystä pointtereiden avulla. Uusia arraytä voidaan tehdä tarvittaessa, mutta mitään ei poisteta tai kopioida. Slicellä pystyy käytännössä tehdä kaiken, mitä listoilla voi tehdä. Oletettavasti se on ominaisuuksiltaan parempi kuin listat, koska sliceä suositaan enemmän.


## Mikä jäi epäselväksi tai tuottanut vaikeuksia?

Välillä hidastaa todella omituiset go-kielen omituisuudet, mutta mikään ongelma ei ole vielä osoittanut täysin murskaavaksi.


## Mitä teen seuraavaksi?

  - ensimmäiset testit
  - projektin refrakturointia: koodia voisi järjestää erilaisiin tiedostoihin
  - lisää ilmaisujen toteutuksia 
  - ohjeet go-ohjelmien ajoon, asennukseen ja testeihin
