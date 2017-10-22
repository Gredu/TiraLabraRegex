# Ohjelman rakenne

Ohjelma koostuu kolmesta funktiosta, `parseRegexp()`, `generateMachine()` ja `match()`. Funktiota `generateMachine()` on tukemassa funktio `generateTransition()`. Koska funktiot toimivat erittäin lähekkäin, on ne päätetty sijoittaa samaan tiedostoon. Samalla tavalla myös `isDigit()` on funktion `match()` tukena ja samassa tiedostossa.

Funktiot ovat siis tiedostoissa siten, että itse tiedosto kuvaa mitä se tekee. Ennen tiedostojen toiminnan kuvausta, kuvataan kuitenkin ensin structuurit.


## structs.go

Sisältää ohjelman kaikki struktuurit.


### Token

Yksittäinen operaatio. Tarkoitettu funktiolle `generateMachine()` arraynä. Tokeneiden avulla luodaan automaatti. Token kertoo kuinka tilat ja siirtymät tulisi luoda.

**typeOperator** kenttä on operaatio-tyyppi. Säännöllisen lauseen yksittäiset ilmaisut voidaan jakaa kategoreihin. Tämä kenttä kertoo mihin kategoriaan ilmaisu kuuluu. Funktio `generateTransition()`, joka on samassa tiedostossa kuin `generateMachine()`, luo transitiot **typeOperator** kentän mukaan.

**value** kenttä on arvo, johon **typeOperator** tehdään.

**token** kenttä voi pitää sisällään toisen token. Idea on siinä, että kvanttori tai kerto-ilmaisut oikeastaan monistavat Tokenia. Ominaisuutta ei ole saatu valmiiksi, vaikka sen rakentaminen on aloitettu.


### Transition

Laskennan malleista tutulla termillä tätä structia voisi sanoa deltaksi. Sisältää tilan, eli **state** ja **token**. Jos **token** hyväksyy syötteen, siirrytään kyseiseen tilaan.


### State

Tila. Jokaisessa tilasta on mahdollisuus päästä toiseen tilaan. Tätä kuvaa kenttä **transitions**, joka on array/slice, ja sisältää siirtymiä. Toinen kenttä on **accept**, joka kertoo onko tila hyväksytty. Funktio `match()` vertaa accept tietoa ja jäljellä olevan syötteen pituutta, ja päättää sen avulla onko syöte hyväksytty.


## parseRegexp.go

Funktio `parseRegexip()` ottaa argumentikseen säännöllisen lauseen ja palauttaa arrayn Tokeneita. Tätä palautusta voi `generateMachine()` käyttää.

Säännöllinen lause annetaan tyyppinä `string`. Tämä lause pyritään pilkkomaan yksittäisiin kokonaisuuksiin, eli Tokeneihin. Esimerkiksi lause "abc" parseeraja antaisi takaisin kolme **tokenia**, joiden **typeOperator** olisi "literal" ja arvot "a", "b" ja "c". Monimutkaisempi "ab*c" parseeraaja antaisi takaisin myös kolme **tokenia**, mutta nyt keskimmäisen tokenin **typeOperator** olisi "star" ja sen arvo "b".

Funktiossa on paljon muuttujan `i` ja `j` muutoksia. Nämä kuvaavat missä kohtaa säännöllistä lausetta ollaan ja kuinka pitkä array ollaan palauttamassa. Oletuksena luodaan array, joka on yhtä pitkä kuin säännöllinen lause, joka sitten palautetaan lyhennettynä. Kun säännöllinen lause muutetaan **tokeneiksi** on se varmasti lyhyempi tai yhtäsuuri kuin säännöllisen lauseen pituus. Palautetaan siis aina "täynnä" oleva array.


## generateMachine.go

Funktio `generateMachine()` ottaa vastaan arrayn **tokeneita** ja palauttaa tilan **state**. Tämä tila nimetään usein heti "machine", joka voi hämätä, koska oikeasti tarkoittaa vain ensimmäiseen tilaan (automaatissa).

Muuttuja `currentState` pitää sisällään tämän hetkistä tilaa. Kun **tokeneita** käydään läpi luodaan uusi tila, ja transitio kuinka siihen päästään. Transitio annetaan sitten `currentState` tilalle kenttään **transitions**. Tämän jälkeen uusi tila asetetaan `currentState` muuttujaan.

Tämä toistetaan kunnes **Tokenit** loppuvat.

Jos kyseessä on viimein **Token**, on kyseessä myös viimenen tila, jolloin tämän tila asetetaan hyväksytyksi (True).

Siirtymät asetetaan sen mukaan, mikä **typeOperator** **tokenilla** on.


## match.go

Funktio `match()` ottaa argumenteikseen syötteen ja tilan. Tila on automaation ensimmäinen tila (automaatio on vain tilojen luoma verkko). Syöte laitetaan automaatin läpi ja katsotaan päädytäänkö viimeiseen tilaan silloin kun syöte on loppu. Palauttaa totuusarvon.


## main.go

Tekee ohjelmasta käytettävän. Käy läpi tiedoston ja syöttää yksitellen rivit automaatioon. Tulostaa syötteen jos se kelpaa automaatiolle.

