# Intro

niniejszy dokument ma na celu przybliżenie (oraz swego rodzaju przybliżenie)
technologii ImPlot.

## Po co nam ImPlot?

planujemy wykonać część zadania [dotyczącą prezentacji białka](design-doc.md#wizualizacja-kandydatów-na-białka)
przy jego użyciu.

# strona techniczna w GO

## Overview

- framework [giu](https://github.com/AllenDang/giu) posiada "wbudowaną" obsługę ImPlot
- API jest podobne do api samego GIU
- wszystkie funkcje odpowiedzialne za obsługę wykresów zgromadzone są w pliku [Plot.go](https://github.com/AllenDang/giu/blob/master/Plot.go)
  (przy czym może to ulec zmianie, bo jak patrzę na 467 linijek w 1 pliku to mogę "nie wytrzymać" i to podzielić :grinning:)

## Zasady ogólne korzystania z pakietu

1. tworzymy "PlotCanvasWidget", czyli po prostu wołamy `giu.Plot("tytół")`
2. dostajemy (jak zwykle w GIU) factory
3. Możemy poustawiać rozmaite właściwości (limity na osiach, jednostkę, rozmiar e.t.c.)
4. jak już skończymy się bawić settingsami do Canvasa, wołamy `.Plots(...)`
5. no i tu jako argumenty podajemy listę naszych wykresów

API moim zdaniem nie powinno sprawiać problemów, teraz idziemy do:

<!--
## High-Math
-->

RIP: nie ma możliwości wstawiania tekstu na wykresach :-O