# Toteutus

Tarkoitus on mallintaa Laskennan mallin äärellisiä laitteita. Säännöllinen lauseke siis esitetään automaattina. Tällä toivottavasti on sellainen vaikutus, että eri säännöllisen lausekkeen ilmaisuja on helppo yhdistää samalla tavalla kuin automaatteja on helppo yhdistää.

Ohjelma sisältää automaatista tuttuja komponentteja. `state State` ilmaisee tilaa, eli Laskennan malleista Q:ta. Tämä kantaa kaksi ominaisuutta; tietoa siitä, onko hyväksytyssä tilassa `accept bool`, sekä mahdollisia tiloja joihin voi mennä `transition Transition`, joka vastaa Laskennan malleissa delta-funktiota. `Transition` vastaa yhtä deltaa. Sillä on siis tieto mihin tilaan voidaan päästä, jos `token Token` sen hyväksyy. `Token` sisältää tietueet `typeOperator` ja `value`. Nämä arvot kertovat minkälaista funktiota on käytettävä syötteen evaluoimiseen.

`typeOperator` vastaa siis sitä, onko kyse literaalista, metasta vai määrästä, ja näin osataan valita sopiva toiminto. Toiminnan, eli funktion parametrina käytetään sitten `value`-arvoa.
