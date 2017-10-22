# Palautus

Tunteja käytetty n. 28.


## Mitä olen tehnyt tällä viikolla?

  - documentaatioita
    - suorituskykytestaus ihmisen DNA:lla 
    - Go:n asennus
    - ohjelman asennus ja käyttö
    - testikattavuustulokset
    - testien toteutus
    - hieman godoc kommentteja (samankaltainen kuin javadoc)
    - ohjelman rakenne (korvaa godocsin)
  - ilmaisu `?`


## Miten ohjelma on edistynyt?

Dokumentaatiota on tullut paljon, mutta monimutkaisemmat ilmaisut on jätetty tekemättä


## Mitä opin tällä viikolla / tänään?

  - godoc
  - go:n tapa käistellä nimiavaruutta ja pakettien hallintaa


## Mikä jäi epäselväksi tai tuottanut vaikeuksia?

  - monimutkaisia ilmaisuja on vaikea tehdä nykyisellä toteutuksella


# Itsearviointi, ohjelman nykyinen tila ja retrospekti

Veikkaan, että tämä kohta on tarpeen, koska golang ja sen menetelmät eivät kaikille ole kovinkaan tuttuja.


## Ohjelma

Kun katsoo [Säännöllisten lausekkeiden rakenne](https://github.com/Gredu/TiraLabraRegexp/blob/master/docs/regExpRakenne.md) dokumentaatiota, se on suurimmaksi osaksi toteutettu. Kuitenkin ilman monimutkaisempia ilmaisuja ohjelma ei vielä tunnu oikealta säännöllisen lausekkeiden tulkilta. Kaikki ilmaisut kyseisestä dokumentista pystytään tekemään nykyisellä ohjelman rakenteella, mutta jos tästä vielä mennään eteenpäin, on ehkä tehtävä melko isojakin rakenteellisia muutoksia.

Idea oli aluksi tehdä ensin NFA ja siitä DFA. Projektin edetessä huomasin yrittäväni luoda DFA:ta suoraan säännöllisestä lausekkeesta. Siinä on kyllä onnistuttu, mutta vaikeammat ilmaisut tekevät menetelmästä todella monimutkaisen.

Koska selkeästi olin tekemässä jotain omaa, ja se vaikutti onnistuvan, niin päätin jatkaa. Yksi mahdollisuus olisi ollut tutkia automaatioita ja säännöllisten lauseiden algoritmeja tarkemmin ja seurata valmiita algoritmeja. 

Mitä kaikkea ohjelma osaa tehdä voi katsoa suoraan tiedostosta [match_test.go](https://github.com/Gredu/TiraLabraRegexp/blob/master/match_test.go). Testit on tehty ns. Table Driven Test menetelmällä. Tiedostossa näkee minkälaisia automaatteja voidaan luoda ja mitkä hahmot se tunnistaa.


### Kesken jääneitä koodipätkiä

Koodissa on muutamia osia, jotka saattavat vähän hämmentää. Todennäköisesti se johtuu siitä, että koodi on keskeneräinen ja ideaa ei päästy viemään ihan loppuun asti.


#### Token

Token Structin kentän token, joka sisältää toisen Tokenin, ei päästy kunnolla hyödyntämään. Kyseessä on siis ilmaisu, joka monistaa tai kertoo jonkun toisen Tokenin. Parseeraaja kyllä kuitenkin osaa jo asettaa kertoimille Tokenit.


### pipeline

Ohjelma tunnistaa milloin shellissä on johdettu dataa pipelinellä (eng. pipeline, en tiedä mikä se on suomeksi), mutta datan haltuunottaminen jäi kesken. Idea oli, että ohjelmaa voisi käyttää esimerkiksi kuin shellin `grep` ohjelmaa. Esim:

```sh
ls | grep *.go
```


## Tiedostorakenne

Jätän tiedostorakenteen sellaiseksi kuin mikä se on nyt. Tässä seurataan enemmän go:n tapaa. Näyttää siltä, että suositeltu tiedostorakenne on litteä, joka kyllä näyttää vähän sekavalta. Esimerkiksi [samaralla](https://github.com/Shopify/sarama), on tiedostorakenne litteä, ja samara on yksi seuratuimpia go-projekteja, jotka ovat githubissa.

Vaikuttaa siltä, että golang ohjelmoijat avaavat tiedostoja mielummin jollain hakukeinolla tai tageilla. Esimerkiksi kirjoittamalla tiedoston nimen, funktion nimen jne. editointi-ohjelma vie suoraan kyseiselle tiedostolle kyseiseen kohtaan. Kokonaiskuvaa on kieltämättä erittäin vaikea hahmottaa katsomalla vain tiedostorakennetta.

Haulla avaaminen on ollut minunkin tapa availla tiedostoja tai katsoa funktioiden lähteitä. Todella harvoin oikeastaan katselen tiedostorakennetta.


## godoc

Dokumentti [Ohjelman rakenne](https://github.com/Gredu/TiraLabraRegexp/blob/master/docs/ohjelmanRakenne.md) kertoo tarkasti kuinka ohjelma toimii.

Yritin saada godocsia toimimaan (godoc on melkein sama asia kuin javadoc), mutta huomasin siinä olevan esteen (tai ominaisuuden). Suoritettavan ohjelman pakettinimi on aina main. Tällöin godoc toimii vain `main.go` tiedostossa, mutta ei missään muualla.

Tässä go-suunnittelijat ovat yrittäneet (luultavasti) saada ohjelmoijat luomaan erillisiä kirjastoja siten, että vain main-ohjelma ajaa ne. Omalla kohdalla se alkaa jopa tuntua järkeenkäypältä. Tein suoritettavan ohjelman todella myöhään, vasta siinä vaiheessa kun oli aika demota. Tässä tuli tunne jo ennen godocsiin tutustumista, että minulla on oikeasti kaksi erillistä osaa, itse suoritettava ohjelma ja sen kirjasto. Näiden pakettinimet olisi pitänyt olla erilaisia.

En mennyt tätä asiaa tämän pidemmälle, vaan päätin kirjoittaa ohjelman rakenteesta erilliselle dokumentaatiolle godocsin sijaan. Koodissa on kyllä hiukan godocsiakin.

Olisi ollut kyllä mielenkiintoista käyttää godocsin yhtä erikoisuutta hyväksi. Nettisivu [godocs.org](https://godoc.org/) voi luoda html dokumentaation suoraan githubin reposta. godocs on luultavasti myös yksi syy miksi go-ohjelmoijat pärjäävät litteällä tiedostojärjestelmällä. Tärkeintä ovat funktiot ja niiden toiminta, ei se missä tiedostoissa tai kansioissa ne ovat. [Samara godocs.org sivuilla](https://godoc.org/github.com/Shopify/sarama), on dokumentaatio luotu githubissa olevan Samaran kommenttiosioista.


## Suorituskykymittauksen data?

Suorituskykymittauksen dna-tiedostot on zipattu, jotta tiedostot veisivät vähemmän tilaa. Tiedosto löytyy kansiosta `data/dna/`. Jos haluaa suorittaa suorituskykymittauksia, on zip tiedosto purettava samaan kansioon, kuin zip-tiedosto itse.


## Mitä olen oppinut koko kurssin aikana

Ohjelmaa tehdessä tuli opittua Go-kieli. Se on kieli, jota minun on pitänyt oppia, joten tietorakenteiden harjoitustyö tuli siihen sopivasti. On kyllä totta, että kielen osaamattomuus on hidastanut ohjelman kehitystä. Kyseessä ei kuitenkaan ole kurssi, jossa Go:ta opetellaan vaan tarkoitus on tehdä tietorakenteisiin liittyvää ohjelmointia. Toisaalta tietorakenteet on hyvin laaja käsite, ja näyttää siltä, että kurssilla harjoitellaan myös muita ohjelmointikäytäntöjä, kuten version hallintaa ja dokumentointia.

Pointterit tuli hyvin opeteltua. Jopa muutamiin uusiin tietorakenteisiin tuli tutustuttua Go:n ansiosta (slice). Go:n ohjelmointikehitys ja paradigma on ollut yllättävä ja hämmentävä. Kiinnostus Go:ta kohtaan on oikeastaan vain lisääntynyt.

Tietorakenteellisesti projekti on tuntunut tietyllä tavalla puun rakentamiselta linkitetyillä tietorakenteilla, eikä rekursiivisuus tunnu enää monimutkaiselta. Pointterit ovat vähän avartaneet tietokoneen sisäistä toimintaa.


## Mitä olisin tehnyt toisin

Ohjelmaa olisi kannattanut suunnitella huommattavasti kauemmin. Ohjelman suunnitteluvaiheessa otettiin huomioon vain yksinkertaiset ilmaukset, ja ohjelman ydin rakennettiin sitä ajatellen. Oletin myös, että monimutkaisempia ilmaisuja varten tarvittaisiin vain vähän lisäyksiä itse ohjelmaan, mutta näin ei asia ehkä ollutkaan.
