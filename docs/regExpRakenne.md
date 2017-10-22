# Säännöllisten lausekkeiden rakenne

Säännöllinen lause voidaan jakaa osiin erilaisten ilmausten perusteella. Tässä vaiheessa harjoitustyötä otetaan mukaan vain seuraavat: literaalit, metat ja määrän ilmaisut (kvanttorit?). Näiden määritelmä on myös määritelmä kuinka harjoitustyön säännöllisen lauseiden tulkki tulisi toimia. Ilmauksia tulisi pystyä yhdistelemään.


## Literaalit

Yksittäinen kirjain sellaisenaan. Esimerkiksi `a` hyväksyy sanat `aakkonen`, `paluu` ja `tuhatjalkainen`, mutta ei hyväksy sanoja `monni`, `mAlli` ja sanaa `meri`.

Literaaleja voi yhdistää, jolloin haku pitää sisältää kirjaimet kokonaan katkeamattomana ja samassa järjestyksessä. Esimerkiksi `abc` hyväksyy sanat `KissaabcKävelee` ja `Osaammeko abc-kirjaimet?`, mutta ei hyväksy sanoja `bca` ja `a b c`.


## Metat

Meta on ehto, mitä yhtä kirjainta ilmaus hyväksyy. Kenoviiva `\` ilmaisee alkavan metatyypin. Listaus eri metoista:
  - `\d` tarkoittaa mitä tahansa numeroa
  - `\w` tarkoittaa mitä tahansa kirjainta tai numeroa
  - `\s` tarkoittaa välilyöntiä
  - `.` mitä tahansa kirjain

Esimerkiksi `\d\d-\d\d-\d\d\d\d` hyväksyy sanat `20-09-1985` ja `01-01-2000`. Huomaa, että esimerkissä on käytetty literaaleja.


## Määrät

Ilmaisevat miten monta kertaa edellistä ilmaisua toistetaan.
  - `*` toistetaan nolla tai n-määrä (eli niin paljon kuin halutaan)
  - `+` toistetaan vähintään kerran tai n-määrä
  - `?` toistetaan kerran tai ei yhtään
  - `{k}` toistetaan k-verran

Esimerkiksi edellinen esimerkki olisi voitu kirjoittaa muotoon `\d{2}-\{2}-\d{4}` ja hyväksyy samat sanat.
