---
title: アーキテクチャ
description: jmactts のクリーンアーキテクチャ層構成
---

`jmactts` は **Clean Architecture** の 3 層構成です。依存方向は内向き一方向で、`usecase` 層がポート (インターフェース) を所有します (Dependency Inversion Principle)。

## 層構成

```
main.go                              ← Composition Root
└─ internal/
   ├─ domain/                        ← Entities (純粋ドメイン, 外部依存なし)
   │  ├─ voice.go                    Voice / Locale / VoiceList
   │  ├─ speech.go                   OutputFormat / DetectFormat
   │  └─ chunk.go                    ChunkText (純粋関数)
   ├─ usecase/                       ← Application Logic
   │  ├─ ports.go                    Synthesizer / MP3Encoder / VoiceCatalog / TextSource
   │  ├─ speak.go                    SpeakUseCase
   │  └─ voices.go                   ListVoices / PickVoice
   └─ adapter/                       ← Interface Adapters
      ├─ say/                        macOS say を Synthesizer + VoiceCatalog として実装
      ├─ ffmpeg/                     ffmpeg を MP3Encoder として実装
      ├─ textsource/                 Clipboard / File / Stdin / Args (TextSource 実装)
      └─ cli/                        フラグ解析、Run(ctx, args, deps, stdout, stderr) int
```

## 依存方向のルール

- `domain` ← `usecase` ← `adapter/*` ← `main.go`
- `domain` はどのパッケージも import しない
- `usecase` は `domain` のみ import
- `adapter/*` は `usecase` と `domain` を import
- `main.go` (Composition Root) のみすべての層を import し、依存関係を組み立てる

## ポートとアダプタ

`usecase/ports.go` で定義された 4 つのインターフェースが、ユースケース層から外部資源を抽象化します。

| ポート | 役割 | 実装アダプタ |
|---|---|---|
| `Synthesizer` | テキストを音声化 (再生 or 保存) | `adapter/say` |
| `MP3Encoder` | AIFF を MP3 に変換 | `adapter/ffmpeg` |
| `VoiceCatalog` | 利用可能ボイス一覧の取得 | `adapter/say` |
| `TextSource` | 入力テキストの取得 | `adapter/textsource` (4 種) |

新しい入力ソースや出力形式を追加したい場合は、対応するアダプタを実装するだけで `usecase` / `domain` を変更する必要はありません。

## テスト戦略

- **`domain` 層**: 純粋関数なので単体テストが書きやすい (`internal/domain/chunk_test.go` 参照)
- **`usecase` 層**: モックアダプタを注入してテスト可能
- **`adapter` 層**: macOS 環境での統合テスト (主にスモークテスト)
