---
title: Voice Selection
description: Pick a voice by name (-v) or by language/country code (-L)
---

`jmactts` lets you pick a voice two ways: `-v` (voice name) and `-L` (language/country code). When both are given, **`-v` wins**.

## By language/country code (`-L`)

Pass `-L` alone and jmactts auto-picks that language's primary voice — you don't need to remember any voice names to switch languages.

```bash
jmactts -L ja_JP こんにちは     # full locale → Kyoko
jmactts -L ja こんにちは         # language code → Kyoko
jmactts -L JP こんにちは         # country code → Kyoko
jmactts -L en "Hello"           # English → Samantha
jmactts -L en_GB "Cheerio"      # UK English → Daniel
```

### Matching rules

`-L` is resolved in this order:

1. **Full locale** (contains `_`, e.g. `ja_JP`) → exact match
2. **Language code** (`ja`, `en`, …) → all `xx_*` voices
3. **Country code** (`JP`, `US`, …) → all `*_YY` voices

### Primary voice heuristic

When multiple voices match, jmactts prefers **names without a parenthesis `(`**.

By macOS convention each language's primary voice (`Kyoko`, `Samantha`, `Daniel`, …) is named without parentheses, while secondary/novelty voices (`Eddy (日本語（日本）)`, …) are parenthesized.

## By voice name (`-v`)

Pass any name listed by `say -v ?`.

```bash
jmactts -v Kyoko こんにちは
jmactts -v Samantha "Hello world"
jmactts -v Otoya こんばんは
```

## Listing voices

`-l` prints the available voices. Combine with `-L` to filter.

```bash
jmactts -l                # all voices (with sample text)
jmactts -l -L ja          # Japanese voices only
jmactts -l -L en_GB       # UK English voices only
```

## Speech rate (`-r`)

`-r` adjusts the speaking rate in words per minute. Default values vary per voice.

```bash
jmactts -L en -r 250 "Speak faster"
jmactts -L en -r 120 "Speak slower"
```
