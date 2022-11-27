# 1

Nawet nie próbujcie zrozumieć co tu się dzieje, tutaj opisze tylko
jak działą API i co trzeba docelowo osiągnąć

# CEL

chodzi nam o plik `./database.go` - zawiera on gigantyczną mapę
umieszczoną w zmiennej (tak, wiem, nie wolno tego robić, ale muszę :cry:)

## jak deklaruje się mapę w GO?

_krótkie wprowadzenie dla nieznających GO_

mapy w GO działają np jak słowniki w pythonie
`klucz:wartość,`

nasza mapa jest nieco bardziej skomplikowana, ponieważ jako wartość przyjmuje
`drawCommands`, ale o tym za chwilę

## Obsługa naszej mapy

### klucz

pierwsza część jest zdecydowanie prostsza do zrozumienia.
klucze w naszej mapie to `stringi` będące skrótowymi
nazwami aminokwasów (oznaczonymi jako `Sign` w JSONie)

### Wartość

wartość to (najczęściej) będzie wywołąnie "fabryki" schematu.

przyjżyjmy się następującemu łańcuchowi:

```go
	"[START]": draw(). // stworzenie fabryki (stały element)
		move(image.Pt(0, 60)). // przesunięcie "kursora" o 60 pikseli w dół (na przykład żeby zmieścić schemat)
		chemicalText("H_3_C", VAlignCenter, HAlignLeft). // dodanie tekst H_3 C
		connect(UpRight, standardLine). // narysowanie linii w prawo do góry
		connect(DownRight, standardLine). // narysowanie linii w prawo do dołu
		add( // tutaj dodajemy nowy "subłańcuch"
			draw(). // tworzymy nową "fabrykę" dla naszego łańcucha
				connect(Down, standardLine).
				chemicalText("NH_2_", VAlignTop, HAlignCenter).draw,
		).
		ignore(ignoreAll). // zignorowanie przesunięcia wytworzonego przez subłańcuch (kontynuujemy od miejsca) z linią w prawo-dół
		connect(UpRight, standardLine). // linia w prawo do góry
		chemicalText("OH", VAlignCenter, HAlignLeft). // tekst "OH"
		move(image.Point{}).draw, // zakończenie (może być bez move(...) tutaj, ale utrudniłoby pracę bo `.draw,` musi być)
```

także, pisanie tych faktorii polega jedynie na zabawie z poniższym API


# API drawer'a

wyobraźmy sobie że jesteśmy żółwiem z logomocji.
wszystkie komendy, które wydajemy (połączone `.`)
każą nam przemieszczać się wykonując określone czynności.
przykładowo, `....drawLine(Right, 15)` narysuje linię o długości
15 pikseli i przemieści nas te 15 pikseli w prawo.
generalnie, rysowanie zawsze przemieści nas od "lewego górnego" (początka)
do "końca" rysowanej figury z drobnymi wyjątkami (np, tekstu & alignment - patrz niżej)

_może ulec zmianie, ale jak coś zmienie to poprawie istniejący już kod_

- chemicalText: dodaje tekst z "formatowaniem chemicznym"
  na razie sprowadza się to do tego że wszystko co jest pomiędzy `_` jest w indeksie dolnym
  argumenty:
    1. tekst
    2. "wyróœnanie pionowe" (patrz poniżej)
    3. "wyrównanie poziome"
- drawLine: rysuje linie w wybranym kierunku
  argumenty:
    1. kierunek (Up/RightUp/Left e.t.c.)
    2. długość linii (w pikselach jeśli to coś komuś mówi)
- move: przesuwa naszego "żółwia" o XY pikseli.
  Należy zaznaczyć, że w ImGui lewy górny róg jest punktem 0,0
  a współżędne rosną w dół (Y) i w prawo (X)
  argumenty:
    1. punkt w formie `image.Pt(x, y)`
- ignore: ignoruje ostatni ruch.
  Na przykład jeśli chcemy kontynuować rysowanie po tym jak narysowaliśmy
  jakiś "sub-łańcuch". **UWAGA** wywołanie ignore() 2 razy nie zadziała tak
  jak możnaby oczekiwać - wtedy cofnięte zostanie cofnięcie!
- aromaticRing (jeszcze nie testowany): rysuje szceściokąt - pierścień
  aromatyczny
  argumenty:
    1. _wymiary zewnętrzne_ (jak kto woli 2*a)
- add: dodaje nowy łańcuch

## sub-łańcuchy

jeżeli chcemy narysować rozgałęzienie łańcucha, tj. narysować inny łańcuch z punktu
w którym jesteśmy, a następnie kontynuować rysowanie ze wcześniejszej pozycji
używamy funkcji `add(...)` i potem zaczynamy tak jakby od początku (patrz przykłąd w kodzie powyżej)
`draw().....`

po zamknięciu nawiasu od `add`, piszemy `ignore()` aby ignorować
całą poprzednią sekwencję

## alignment

alignment (VAlignment i HAlignment) pozwala określić który punkt
"pola tekstowego" jest tym od którego należy zacząć rysowanie.
