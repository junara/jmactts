---
title: Quick Start
description: Start using jmactts in five minutes
---

After installing, you can dive in immediately.

## Speak something

```bash
jmactts "Hello, world"
```

Spoken with `say`'s default voice (the one configured in System Settings).

## Speak in a specific language

`-L` takes a language or country code and auto-picks that language's primary voice.

```bash
jmactts -L en "Hello, world"        # Samantha
jmactts -L en_GB "Good evening"     # Daniel (UK English)
jmactts -L ja こんにちは            # Kyoko
```

## Save to a file

The format is detected from the `-o` extension.

```bash
jmactts -L en -o hello.m4a "Hello"
jmactts -L en -o hello.mp3 "MP3 also works"      # requires ffmpeg
```

## Read from the clipboard

Play back text you just copied with `pbcopy`:

```bash
jmactts -c             # via flag
pbpaste | jmactts      # via pipe
```

## Next steps

- [Input Sources](/jmactts/en/usage/input/) — when to use arguments / files / stdin / clipboard
- [Voice Selection](/jmactts/en/usage/voice/) — `-v` vs. `-L` and the matching rules
- [Output Formats](/jmactts/en/usage/output/) — when to pick AIFF / WAV / M4A / MP3
- [Flags](/jmactts/en/reference/flags/) — every flag and exit code
