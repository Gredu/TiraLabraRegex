# Viikkoraportti 5

  - tunteja käytetty n. 14


## Mitä olen tehnyt tällä viikolla?

  - koodikatselmointi: https://github.com/jesushilden/Ristinolla
  - testejä
  - tutkinut testikattavuusraportteja
  - lisännyt ilmaisuja, ainakin seuraavat: \d, + ja .
  - dokumenttikansiorakenne on järjestetty uudelleen
  - dokumentaatiota on pilkottu otsikkoihin
  - uusi dokumentti testaamisesta
  - uusi dokumentti Go-kielestä helpottamaan arvioimista ja koodikatselmointia
  - yksi suorituskykytestaus
  - tutkinut Go:n tiedostorakenteita


## Miten ohjelma on edistynyt?

Huomattavasti nopeammin nyt, kun Go alkaa sujua ja koodin runko on valmis. Ilmauksien lisääminen on ollut helppoa, koska monet ovat olleet vain yhdistelmiä aikaisemmista ilmaisuista.


## Mitä opin tällä viikolla / tänään?

Go:n työkalupakki on laaja. Testikattavuusrapotti on esimerkiksi jo sisäänrakennettu Go:hon, joten codecov ja travista ei tarvita.

Tuli tehtyä yksi suorituskykytestaus Go-tyyliin, mutta sillä ei paljoa ollut merkitystä. Testi suoriutui liian nopeasti, että tuloksella olisi ollut mitään väliä. Käytännössä pitäisi saada isompi määrä syötteitä ja/tai huomatavan monimutkainen säännöllinen lauseke.


## Mikä jäi epäselväksi tai tuottanut vaikeuksia?

Alunperin tarkoitus oli tehdä ensin NFA ja sitten sen DFA-muuntaja, mutta olen nyt huomannut, että teenkin suoraan DFA:ta. Oikeastaan kyseessä taitaa olla jonkinlainen hybridi. Automaatti käydään läpi syvyyshakuisesti, eli käytännössä syötettä ei oikeasti käydä läpi vasemmalta oikealle, vaan jokaisessa etenemisessä annetaan syöte kopiona, josta on poistettu ensimmäinen kirjain. Laite on hybridi siinä mielessä, että haaraumia ei oikeastaan ole paljoa, mutta en ole ihan varma johtuuko se siitä, että olen itse niitä vähentänyt tai siitä, että ilmaukset ovat vielä todella yksinkertaisia.

Esimerkiksi meta \d, joka siis kuvasi, että syötteen on oltava numero, on toteutettu funktiolla `isDigit(input string)`, joka palauttaa totuusarvon. Funktion sijainti on `match.go` tiedostossa. Funktio siis tarkistaa onko kyseessä numero ja päästää seuraavaan tilaan, jos on. Tämä olisi voitu toteuttaa myös siten, että samaan tilaan pääsee kymmenellä eri tavalla (eli eri numeroilla). Tavallaan tuntuu huijaukselta, mutta näin päästään myös huomattavasti vähemmällä koodilla.

Käytännössä tämänkaltainen ratkaisu ei ole haitannut. Ilmaukset ovat olleet yksinkertaisia ja loppujen ilmauksien toteuttaminen on todella helppoa. Siis niiden ilmausten, jotka on dokumentoitu toteutettavaksi. Kuitenkin ilmaukset () ja [], jotka ajattelin toteuttaa, jos aikaa riittää, ovat huomattavasti vaikeampia. Nykyinen runko ei oikein taivu tuollaisiin ilmauksiin. Olen miettinyt kolmea vaihtoehtoa ongelman ratkaisemiseksi:

  1. uudelleen kirjoitetaan NFA ja tehdään aito DFA-muuntaja
  2. jatketaan nykyisellä tavalla ja tehdään ilmaukset () ja [] toimimaan apinan raivolla
  3. unohdetaan monimutkaisemmat ilmaukset

Ratkaisu 3 antaa aikaa hiota ohjelmaa demotilaisuutta varten. Ratkaisu 1 saattaa viedä liikaa aikaa. Ratkaisu 2 voi onnistua melko nopeastikin, mutta rakenteisiin täytyy kuitenkin tehdä muutoksia. Esimerkiksi laitteen luojafunktio joutuu ymmärtämään "ali-tokenit (eli ali-ilmaisut)". Esimerkiksi säännöllinen lause `(abc|dfg)*`. Kyseessä on "star" operaattori, jonka arvoina on kaksi ilmaisua. Kyse on oikeastaan melko helposta asiasta, mutta ohjelma ei saisi rajoittaa vain yhteen ali-ilmaisuun. Ali-ilmaisujen määrä ei saisi olla rajoitettu. Esimerkiksi `((10|100)*abc|dfg)*` pitäisi myös kyetä muuntamaan laitteeksi.

Tällä hetkellä mennään vaihtoehdolla 3 ja sitten 2 jos jää aikaa.


## Mitä teen seuraavaksi?

  - suurempi määrä syötteitä suorituskykytestien tekemiseen
  - parempi ja selkeämpi tiedostorakenne
  - käyttöliittymä
    - toimii shellissä putkituksen kautta
    - toimii siten, että tiedosto voidaan antaa yhtenä argumenttina ja säännöllinen lause toisella
  - loput ilmaisut, jotka on määritelty dokumentaatioon
  - ylimääräiset ilmaisut () ja []
