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
`DrawCommands`, ale o tym za chwilę

## Obsługa naszej mapy

### klucz

pierwsza część jest zdecydowanie prostsza do zrozumienia.
klucze w naszej mapie to `stringi` będące skrótowymi
nazwami aminokwasów (oznaczonymi jako `Sign` w JSONie)

### Wartość

wartość to (najczęściej) będzie wywołanie "fabryki" schematu.

przyjżyjmy się następującemu łańcuchowi:

```go
	"[START]": Draw(). // stworzenie fabryki (stały element)
		ChemicalText("H_3_C", VAlignCenter, HAlignLeft). // dodanie tekst H_3 C
		DrawLine(UpRight, standardLine). // narysowanie linii w prawo do góry
		DrawLine(DownRight, standardLine). // narysowanie linii w prawo do dołu
		AddSubcommand( // tutaj dodajemy nowy "subłańcuch"
			Draw(). // tworzymy nową "fabrykę" dla naszego łańcucha
				DrawLine(Down, standardLine).
				DrawLine("NH_2_", VAlignTop, HAlignCenter),
		).
		Ignore(ignoreAll). // zignorowanie przesunięcia wytworzonego przez subłańcuch (kontynuujemy od miejsca) z linią w prawo-dół
		DrawLine(UpRight, standardLine). // linia w prawo do góry
		ChemicalText("OH", VAlignCenter, HAlignLeft), // tekst "OH"
```

także, pisanie tych faktorii polega jedynie na zabawie z poniższym API


# API drawer'a

wyobraźmy sobie że jesteśmy żółwiem z logomocji.
wszystkie komendy, które wydajemy (połączone `.`)
każą nam przemieszczać się wykonując określone czynności.
przykładowo, `....drawLine(Right, 15)` narysuje linię o długości 15 pikseli i przemieści nas te 15 pikseli w prawo. generalnie, rysowanie zawsze przemieści nas od "lewego górnego" (początka)
do "końca" rysowanej figury z drobnymi wyjątkami (np, tekstu & alignment - patrz niżej)

_może ulec zmianie, ale jak coś zmienie to poprawie istniejący już kod_

- ChemicalText: dodaje tekst z "formatowaniem chemicznym"
  na razie sprowadza się to do tego że wszystko co jest pomiędzy `_` jest w indeksie dolnym
  argumenty:
    1. tekst
    2. "wyróœnanie pionowe" (patrz poniżej)
    3. "wyrównanie poziome"
- DrawLine: rysuje linie w wybranym kierunku
  argumenty:
    1. kierunek (Up/RightUp/Left e.t.c.)
    2. długość linii (w pikselach jeśli to coś komuś mówi)
- Move: przesuwa naszego "żółwia" o XY pikseli.
  Należy zaznaczyć, że w ImGui lewy górny róg jest punktem 0,0
  a współżędne rosną w dół (Y) i w prawo (X)
  argumenty:
    1. punkt w formie `image.Pt(x, y)`
- Ignore: ignoruje ostatni ruch.
  Na przykład jeśli chcemy kontynuować rysowanie po tym jak narysowaliśmy
  jakiś "sub-łańcuch". **UWAGA** wywołanie ignore() 2 razy nie zadziała tak
  jak możnaby oczekiwać - wtedy cofnięte zostanie cofnięcie!
- AromaticRing (jeszcze nie testowany): rysuje szceściokąt - pierścień
  aromatyczny
  argumenty:
    1. _wymiary zewnętrzne_ (jak kto woli 2*a)
- AddSubcommand: dodaje nowy łańcuch

## sub-łańcuchy

jeżeli chcemy narysować rozgałęzienie łańcucha, tj. narysować inny łańcuch z punktu
w którym jesteśmy, a następnie kontynuować rysowanie ze wcześniejszej pozycji
używamy funkcji `AddSubcommand(...)` i potem zaczynamy tak jakby od początku (patrz przykłąd w kodzie powyżej)
`Draw().....`

po zamknięciu nawiasu od `AddSubcommand`, piszemy `Ignore()` aby ignorować
całą poprzednią sekwencję

## alignment

alignment (VAlignment i HAlignment) pozwala określić który punkt
"pola tekstowego" jest tym od którego należy zacząć rysowanie.
