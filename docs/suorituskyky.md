# Suorituskykymittaus

Suorituskykymittaukseen käytetty data on zipattuna kansiossa `data/dna/`. Jotta suorituskykymittausta voidaan tehdä, on tämä paketti purettava samaan kansioon kuin missä zip-tiedosto on. Data on zipattu tilan säästämiseksi.


## Kuinka suorituskykymittaus on tehty

Suorituskykytestaus tapahtuu tiedostossa `benchmark_test.go`. Suorituskykytestauksen käynnistäminen tapahtuu projektin juurella komennolla:

```sh
go test -bench=.
```

Saadaan seuraavanlaista tietoa ulos:

```sh
BenchmarkMatchLine63M-6         2000000000               0.25 ns/op
BenchmarkMatchLineDots63M-6     2000000000               0.22 ns/op
BenchmarkMatchLine81M-6         2000000000               0.36 ns/op
BenchmarkMatchLineDots81M-6     2000000000               0.34 ns/op
BenchmarkMatchLine104M-6        2000000000               0.46 ns/op
BenchmarkMatchLineDots104M-6    1000000000               0.41 ns/op
```

Suorituskykytestauksessa katsotaan kuinka nopeasti ohjelma käy läpi palan ihmisen DNA:ta, joka on FASTA-formaatissa. Tiedostoissa on siis vain osa ihmisen genomia. Tiedostoilla on koko, joka näkyy ensimmäisessä sarakkeessa nimen lopussa juuri ennen M kirjainta. Eli tiedostoja on 63Mt, 81Mt ja 104Mt kokoisina. Jos funktion nimessä on "Dot", tarkoittaa se sitä, että automaatissa on käytetty myös ilmaisua `.`. Loput ovat vain literaaleja. Funktionimen lopussa on numero 6, joka tarkoittaa, että testi on suoritettu x86-64 arkkitehtuurille.

Rivejä tiedostoissa on seuraavasti:
  - 63Mt, rivejä 1 073 071
  - 81Mt, rivejä 1 387 626
  - 104Mt, rivejä 1 784 063

Toinen sarake kertoo kuinka monta kertaa koe suoritettiin. Go:n oma suorituskykymittaus pyrkii toistamaan koetta niin monta kertaa, että tulos on luotettava. Kertojen määrää saatetaan vähentää automaattisesti jos suoriutumiseen menee jo muutenkin paljon aikaa. Kaksi miljardia kertaa on aina maksimi.

Viimeisin sarake on tulos. ns/op tarkoittaa nanosekuntia per operaatio, eli kertoo sen, kuinka kauan yhden mittauksen suorttamiseen meni aikaa keskimäärin (vai mediaani?).

Testissä on otettu huomioon, että tiedoston lataamista ja tiedoston sulkemista ei lasketa mukaan suorituskykymittauksessa. Nämä tapahtuvat funktioilla `ResetTimer()` ja `StopTimer()`.

Kokonaisuudessa kaikkien suorituskykymittauksen läpiajo kestää n. 2 minuuttia.


## Tulosten tulkinta

Ohjelma luo automaatin, joka on deterministinen. Itse syötteiden käynti tapahtuu deterministisessä automaatissa O(n) nopeudella. Tulokset tukevat tätä havaintoa, sillä suorituskykymittauksessa kestävä aika vaativuus kasvaa lineaarisesti tiedoston kokoon nähden.

Oli myös odotettavaa, että suorituskykytestaus ilmaisun `.` kanssa on vähän nopeampi. Ilmaisu `.` tarkoitti, että mikä tahansa seuraava syöte kelpaa. Automaatti on rakennettu siten, että `.` ilmaisu tarkoittaa oikeastaan vaan sitä, että hypätään syötteen ohi ja mennään suoraan seuraavaan tilaan. Käytännössä tämä tarkoittaa sitä, että vertailua ei tehdä.

Nopeus on yllättänyt jopa sen verran, että tulokset vaikuttavat epäilyttäviltä. Empiirinen tutkimus on kuitenkin osoittaa, että tuloksissa on järkeä.

Ihmetyttää myös Go:n kyky suorittaa funktio kaksi miljardia kertaa hyvin nopeasti.


## Mahdolliset lisätoimet

On selkeää, että tiedostot voivat olla suuruudestaan huolimatta edelleenkin liian pieniä. Suorituskykytestaus ei antanut aina samankaltaisia tuloksia. Välillä isommat tiedostot suoritettiin nopeammin kuin pienemmät, mutta pääosin tulokset olivat kuten tässä artikkelissa on näytetty. Pienet tiedostot siis suoritettiin nopeammin kuin isot tiedostot useimmissa tapauksissa. Virhe todennäköisesti johtuu vain siitä, että mittausta suorittava tietokone teki juuri silloin työtä myös jossain muualla.

Isommat tiedostot saattaisivat poistaa tämän mahdollisen virhemarginaalin, mutta jo näillä tiedostoilla pystytään tehdä melko varmoja päättelyjä.


## Mittaukseen käytetty data

Mittaukseen käytetty data on zipattuna kansiossa `data/dna/`. Pura zipattu tiedosto samaan paikkaan, jos haluat ajaa testejä itse.


## Lähteet

https://golang.org/pkg/testing/
https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go
