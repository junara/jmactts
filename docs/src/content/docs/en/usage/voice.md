---
title: Voice Selection
description: Pick a voice by name (-v) or by language/country code (-L)
---

`jmactts` lets you pick a voice two ways.

## By voice name (`-v`)

Use any exact name from `say -v ?`.

```bash
jmactts -v Kyoko こんにちは
jmactts -v Samantha "Hello, world"
```

List available voices with `-l`:

```bash
jmactts -l                # all voices (with sample text)
jmactts -l -L ja          # Japanese voices only
```

## By language/country code (`-L`)

When `-L` is given without `-v`, jmactts auto-picks the primary voice for that language or country.

```bash
jmactts -L ja_JP こんにちは     # full locale
jmactts -L ja こんにちは         # language code
jmactts -L JP こんにちは         # country code
jmactts -L en Hello              # any en_*
```

### Matching rules

`-L` values are matched in this order:

1. `ja_JP` format (contains `_`) → exact locale match
2. Language code (`ja`, `en`, …) → all `xx_*` matches
3. If no match, treat as country code (`JP`, `US`, …) and search `*_YY`

### Primary voice heuristic

When multiple voices match, jmactts prefers **names that do NOT contain a parenthesis `(`**.

macOS conventionally names each language's primary voice without parentheses (e.g. `Kyoko`, `Samantha`, `Daniel`) and parenthesizes secondary/novelty voices (e.g. `Eddy (日本語（日本）)`).

## Speech rate (`-r`)

`-r` adjusts words per minute.

```bash
jmactts -L en -r 250 "Speak faster"
jmactts -L en -r 120 "Speak slower"
```
