---
title: Architecture
description: jmactts follows a three-layer clean architecture
---

`jmactts` is structured in three layers per **Clean Architecture**. Dependencies point inward only, and the `usecase` layer owns the port interfaces (Dependency Inversion Principle).

## Layers

```
main.go                              ← Composition Root
└─ internal/
   ├─ domain/                        ← Entities (pure domain, no external deps)
   │  ├─ voice.go                    Voice / Locale / VoiceList
   │  ├─ speech.go                   OutputFormat / DetectFormat
   │  └─ chunk.go                    ChunkText (pure function)
   ├─ usecase/                       ← Application Logic
   │  ├─ ports.go                    Synthesizer / MP3Encoder / VoiceCatalog / TextSource
   │  ├─ speak.go                    SpeakUseCase
   │  └─ voices.go                   ListVoices / PickVoice
   └─ adapter/                       ← Interface Adapters
      ├─ say/                        macOS say as Synthesizer + VoiceCatalog
      ├─ ffmpeg/                     ffmpeg as MP3Encoder
      ├─ textsource/                 Clipboard / File / Stdin / Args (TextSource impls)
      └─ cli/                        Flag parsing; Run(ctx, args, deps, stdout, stderr) int
```

## Dependency direction

- `domain` ← `usecase` ← `adapter/*` ← `main.go`
- `domain` imports nothing else
- `usecase` imports `domain` only
- `adapter/*` imports `usecase` and `domain`
- Only `main.go` (the Composition Root) wires every layer together

## Ports and adapters

The four interfaces in `usecase/ports.go` abstract external resources from application logic:

| Port | Role | Adapter |
|---|---|---|
| `Synthesizer` | Render text to speech (play or save) | `adapter/say` |
| `MP3Encoder` | Transcode AIFF to MP3 | `adapter/ffmpeg` |
| `VoiceCatalog` | List available voices | `adapter/say` |
| `TextSource` | Provide input text | `adapter/textsource` (4 variants) |

Adding a new input source or output format only requires writing an adapter; `usecase` and `domain` stay untouched.

## Testing strategy

- **`domain` layer** — pure functions, easy to unit test (see `internal/domain/chunk_test.go`)
- **`usecase` layer** — injectable mock adapters
- **`adapter` layer** — integration smoke tests on macOS
