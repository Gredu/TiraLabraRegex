# Tietorakenteiden harjoitustyö: regexp tulkki

Säännölliset lausekkeet, eli regular expression, eli regexp, on syötteen tai hahmon tunnistamiseen tarkoitettu kieli. Säännölliset lauseet voidaan ajatella olevan hakuehtoja, joilla tunnistetaan tai löydetään tekstissä kyseisen haun täyttävän ehdon. Haku voi olla huomattavasti monimutkaisempi kuin perinteinen "etsi sivulta", koska säännöllisten lausekkeiden rakenteet voivat sisältää ehtoja ja muuttujia.

Erilaisia syntakseja säännöllisille lausekkeille on paljon. Eri ohjelmointikielet sisältävät yleensä jonkinlaisen säännöllisten kielten tulkin. Tässä harjoituksessa noudatetaan ainakin alkuun JavaScriptin säännöllisiä lausekkeita, mutta lopputyö tulee sisältämään luultavasti jotain muitakin ideoita muista kielistä.

Harjoitustyö tehdään Go kielellä. Tulkki pyritään rakentamaan Laskennan mallin kursista tutuilla äärellisillä laskentalaitteilla, eli niitä pyritään simuloimaan. Tällöin erilaisten säännöllisten kielten yksittäisiä ilmauksia voidaan yhdistää samalla tavalla kuin laskentalaitteita. Koodin O-vaatimusluokka on siten samaa tasoa kuin epädeterministisellä äärellisellä laskentalaitteella. Syöte luetaan vain yhden kerran vasemmalta oikealla.


## Säännöllisen lausekkeen rakenne

Säännöllinen lause voidaan jakaa osiin erilaisten ilmausten perusteella. Tässä vaiheessa harjoitustyötä otetaan mukaan vain seuraavat: literaalit, metat ja määrän ilmaisut (kvanttorit?). Näiden määritelmä on myös määritelmä kuinka harjoitustyön säännöllisen lauseiden tulkki tulisi toimia. Ilmauksia tulisi pystyä yhdistelemään.


### Literaalit

Yksittäinen kirjain sellaisenaan. Esimerkiksi `a` hyväksyy sanat `aakkonen`, `paluu` ja `tuhatjalkainen`, mutta ei hyväksy sanoja `monni`, `mAlli` ja sanaa `meri`.

Literaaleja voi yhdistää, jolloin haku pitää sisältää kirjaimet kokonaan katkeamattomana ja samassa järjestyksessä. Esimerkiksi `abc` hyväksyy sanat `KissaabcKävelee` ja `Osaammeko abc-kirjaimet?`, mutta ei hyväksy sanoja `bca` ja `a b c`.


### Metat

Meta on ehto, mitä yhtä kirjainta ilmaus hyväksyy. Kenoviiva `\\` ilmaisee alkavan metatyypin. Listaus eri metoista:
  - `\\d` tarkoittaa mitä tahansa numeroa
  - `\\w` tarkoittaa mitä tahansa kirjainta tai numeroa
  - `\\s` tarkoittaa välilyöntiä
  - `.` mitä tahansa kirjain

Esimerkiksi `\\d\\d-\\d\\d-\\d\\d\\d\\d` hyväksyy sanat `20-09-1985` ja `01-01-2000`. Huomaa, että esimerkissä on käytetty literaaleja.


### Määrät

Ilmaisevat miten monta kertaa edellistä ilmaisua toistetaan.
  - `*` toistetaan nolla tai n-määrä (eli niin paljon kuin halutaan)
  - `+` toistetaan vähintään kerran tai n-määrä
  - `?` toistetaan kerran tai ei yhtään
  - `{k}` toistetaan k-verran

Esimerkiksi edellinen esimerkki olisi voitu kirjoittaa muotoon `\\d{2}-\\{2}-\\d{4}` ja hyväksyy samat sanat.


## Ohjelmoinnin idea

Tarkoitus on mallintaa Laskennan mallin äärellisiä laitteita. Säännöllinen lauseke siis esitetään automaattina. Tällä toivottavasti on sellainen vaikutus, että eri säännöllisen lausekkeen ilmaisuja on helppo yhdistää samalla tavalla kuin automaatteja on helppo yhdistää.

Ohjelma sisältää automaatista tuttuja komponentteja. `state State` ilmaisee tilaa, eli Laskennan malleista Q:ta. Tämä kantaa kaksi ominaisuutta; tietoa siitä, onko hyväksytyssä tilassa `accept bool`, sekä mahdollisia tiloja joihin voi mennä `transition Transition`, joka vastaa Laskennan malleissa delta-funktiota. `Transition` vastaa yhtä deltaa. Sillä on siis tieto mihin tilaan voidaan päästä, jos `token Token` sen hyväksyy. `Token` sisältää tietueet `typeOperator` ja `value`. Nämä arvot kertovat minkälaista funktiota on käytettävä syötteen evaluoimiseen.

`typeOperator` vastaa siis sitä, onko kyse literaalista, metasta vai määrästä, ja näin osataan valita sopiva toiminto. Toiminnan, eli funktion parametrina käytetään sitten `value`-arvoa.
