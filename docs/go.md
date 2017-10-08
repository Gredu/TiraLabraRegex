# Go

Go-kielellä on muutamia erikoisuuksia, jotka saattavat hämmentää Java-kirjoittajaa. Tämän artikkelin tarkoitus on avittaa muita pääsemään helpommin ja nopeammin sinuiksi Go:hon, jotta itse arvointi ja koodinkatselmus olisi helpompaa.


## Tiedostorakenne

Muista kielistä on tuttu nimitys "workspace" (suom. työalue), jolla tarkoitetaan työaluetta, joka käytännössä sisältää kaiken, mitä yhdessä projektissa on. Työalue sisältää yleensä mm. kansion binääreille, paketeille ja lähdekoodeille. Erikseen saattaa olla myös projektiin liittyvät konfiguraatio tiedostot ja kansio testeille.

Yksi projekti vastaa yhtä työaluetta. Go:ssa näin ei ole. Go:ssa on ainoastaan yksi ainoa työalue, jonka alla kaikki koodi on.

Go ymmärtää projektit omiksi kokonaisuuksikseen kansiorakenteiden ja määreen `package <name>` avulla.

Käyttäjällä on oma Go-alueensa, joka on määritetty `GOPATH` ympäristömuuttujan avulla. Minun GOPATH on asetettu seuraavanlaiseksi:

```sh
❯ echo $GOPATH
/home/greatman/Code/go
```

Tässä kansiossa listaus antaa seuraavan:

```sh
❯ ls
bin  pkg  src
```

Jos asentaa Go:lla, binäärit menevät bin kansioon. Go voi käyttää githubbia myös lähteenä, eli github toimii Go:n yhtenä pakettivarastona. (Go-kielen asennuksen yhteydessä tulee go-kielisten ohjelmien asennukseen ja hallintaa liittyvää softaa)

Go ei ainakaan oletuksena jätä lähteitä lataamatta. Oikeastaan tietääkseni (lähes?) kaikki Go:n asennukset tehdään kääntämällä lähdekoodi ja asentamalla siitä tuleva binääri edellä mainittuun bin kansioon.

Mielenkiintoisinta on siis src kansio. Sen sisältö syvyyteen 2 ja vain kansiot näyttää seuraavalta:

```sh
❯ tree -L 2 -d src
src
├── github.com
│   ├── alecthomas
│   ├── bep
│   ├── BurntSushi
│   ├── cpuguy83
│   ├── davidrjenni
│   ├── dchest
│   ├── dominikh
│   ├── eknkc
│   ├── fatih
│   ├── fsnotify
│   ├── golang
│   ├── gorilla
│   ├── Gredu
│   ├── hashicorp
│   ├── josharian
│   ├── jstemmer
│   ├── kardianos
│   ├── kisielk
│   ├── klauspost
│   ├── kyokomi
│   ├── magiconair
│   ├── miekg
│   ├── mitchellh
│   ├── nicksnyder
│   ├── nsf
│   ├── pelletier
│   ├── PuerkitoBio
│   ├── rogpeppe
│   ├── russross
│   ├── shurcooL
│   ├── spf13
│   ├── yosssi
│   └── zmb3
├── golang.org
│   └── x
├── gopkg.in
│   ├── alecthomas
│   └── yaml.v2
└── honnef.co
    └── go

```

Tämä projekti sisältää kansiossa github.com/Gredu/TiraLabraRegex . Asennuksesta tulevan binäärin nimi on TiraLabraRegex.

Lähteet:
  - [go-tiedostorakenne](https://golang.org/doc/code.html)
  - [go-projektin pystyttäminen](https://dave.cheney.net/2014/12/01/five-suggestions-for-setting-up-a-go-project)



