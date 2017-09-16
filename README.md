# Tietorakenteiden harjoitustyö: regex tulkki

Säännölliset lausekkeet, eli regular expression, eli regex, on syötteen tai hahmon tunnistamiseen käytettävä kieli. Säännölliset lauseet voi ajatella olevan hakuehtoja, joilla tunnistetaan tai löydetään tekstissä kyseisen haun täyttävän ehdon. Haku voi olla huomattavasti monimutkaisempi kuin perinteinen "etsi sivulta", koska säännöllisten lausekkeiden rakenteet voivat sisältää ehtoja ja muuttujia.

Erilaisia syntakseja säännöllisille lausekkeille on paljon. Eri ohjelmointikielet sisältävät yleensä jonkinlaisen säännöllisten kielten tulkin. Tässä harjoituksessa noudatetaan ainakin alkuun JavaScriptin säännöllisiä lausekkeita, mutta lopputyö tulee sisältämään luultavasti jotain muitakin ideoita muista kielistä.

Harjoitustyö tehdään Go kielellä. Tulkki pyritään rakentamaan Laskennan mallin kursista tutuilla äärellisillä laskentalaitteilla, eli pyritään simuloimaan niitä. Tällöin erilaisten säännöllisten kielten yksittäisiä ilmauksia voidaan yhdistää samalla tavalla kuin laskentalaitteet. Koodin O-vaatimusluokka on siten samaa tasoa kuin epädeterministisellä äärellisellä laskentalaitteella. Syöte luetaan vain yhden kerran vasemmalta oikealla.


## Säännöllisen lausekkeen rakenne

Säännöllinen lause voidaan jakaa osiin erilaisten ilmausten perusteella. Tässä vaiheessa harjoitustyöstä otetaan mukaan vain seuraavat: literaalit, metat ja määrän ilmaisu (kvanttorit?). Näiden määritelmä on myös määritelmä kuinka harjoitustyön säännöllisen lauseiden tulkki tulisi toimia. Ilmauksia tulisi pystyä yhdistelemään.


### Literaalit

Yksittäinen kirjain sellaisenaan. Esimerkiksi `a` hyväksyy sanat `aakkonen`, `paluu` ja `tuhatjalkainen`, mutta ei hyväksy sanoja `monni`, `mAlli` ja sanaa `meri`.

Literaaleja voi yhdistää, jolloin haku pitää sisältää kirjaimet kokonaan katkeamattomana ja samassa järjestyksessä. Esimerkiksi `abc` hyväksyy sanat `KissaabcKävelee` ja `Osaammeko abc-kirjaimet?`, mutta ei hyväksy sanoja `bca` ja `a b c`.


### Metat

Meta on ehto, mitä yhtä kirjainta ilmaus hyväksyy. Kenoviiva `\\` ilmaisee alkavan metatyypin. Listaus eri metoista:
  - `\\d` tarkoittaa mitä tahansa numeroa
  - `\\w` tarkoittaa mitä tahansa kirjainta tai numeroa
  - `\\s` tarkoittaa välilyöntiä
  - `.` mikä tahansa kirjain

Esimerkiksi `\\d\\d-\\d\\d-\\d\\d\\d\\d` hyväksyy sanat `20-09-1985` ja `01-01-2000`. Huomaa, että esimerkissä on käytetty literaaleja.


### Määrät

Ilmaisevat miten monta kertaa on edellistä ilmaisua toistettava.
  - `*` toistetaan nolla tai n-määrä (eli niin paljon kuin halutaan)
  - `+` toistetaan vähintään kerran tai n-määrä
  - `?` toistetaan nolla tai kerran
  - `{k}` toistetaan k-verran

Esimerkiksi edellinen esimerkki olisi voitu kirjoittaa muotoon `\\d{2}-\\{2}-\\d{4}` ja hyväksyy samat sanat.
