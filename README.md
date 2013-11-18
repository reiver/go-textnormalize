Go TextNormalize
================

TextNormalize is a Golang library that provides ways of *normalizing text*.

Normalizing text includes either normalizing *line separators* or normalizing *paragraph separators* (or both).

For normalizing *line separators*, this means making all the following conversions:
* "\r\n"   → "\u2028",
* "\r"     → "\u2028",
* "\n"     → "\u2028",
* "\u0085" → "\u2028".

'\u2028' is the UNICODE *line separator* character.
'\u2029' is the UNICODE *paragraph separator* character.
They have similar functionality to the HTML <br> and <p> elements, respectively.
 

Usage
-----

A simlpe example of performing normalization of *line separators* is:
```
s := "This is the 1st line\r\nThis is the 2nd line\r\n"

normalized := textnormalize.NormalizeLineSeparatorsString(s)
```

