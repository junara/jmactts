---
title: Quick Start
description: Start using jmactts in five minutes
---

After installing, you can dive in immediately.

## Speak something

```bash
jmactts "Hello, world"
```

## Speak in a specific language

`-L` selects the primary voice for a language or country code.

```bash
jmactts -L ja Hello reads with a Japanese voice
jmactts -L en Hello will be read in English
```

## Save to a file

The format is detected from the extension.

```bash
jmactts -L en -o hello.m4a Hello
jmactts -L en -o hello.mp3 MP3 also works
```

## Read from the clipboard

Play back text you just copied with `pbcopy`:

```bash
pbpaste | jmactts        # via pipe
jmactts -c               # via flag
```

## Next steps

- [Input Sources](/jmactts/en/usage/input/) — arguments / files / stdin / clipboard
- [Voice Selection](/jmactts/en/usage/voice/) — `-v` vs. `-L`
- [Output Formats](/jmactts/en/usage/output/) — AIFF / WAV / M4A / MP3
- [Flags](/jmactts/en/reference/flags/) — every flag in one place
