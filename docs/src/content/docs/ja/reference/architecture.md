---
title: アーキテクチャ
description: jmactts のクリーンアーキテクチャ層構成
---

`jmactts` は **Clean Architecture** の 3 層構成です。依存方向は内向き一方向で、ポート (インターフェース) は `usecase` 層が所有します (Dependency Inversion Principle)。

## 層構成

```
main.go                              ← Composition Root (配線のみ)
└─ internal/
   ├─ domain/                        ← Entities (純粋ドメイン、外部依存なし)
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
      └─ cli/                        フラグ解析と Run(ctx, args, deps, stdout, stderr) int
```

## 依存方向

`domain` ← `usecase` ← `adapter/*` ← `main.go`

- `domain` は他のどのパッケージも import しない
- `usecase` は `domain` のみ import
- `adapter/*` は `usecase` と `domain` のみ import
- `main.go` (Composition Root) だけがすべての層を import し、依存関係を組み立てる

## ポートとアダプタ

`usecase/ports.go` で定義された 4 つのインターフェースが、ユースケース層から外部資源を抽象化します。

| ポート | 役割 | 実装アダプタ |
|---|---|---|
| `Synthesizer` | テキストを音声化 (再生または保存) | `adapter/say` |
| `MP3Encoder` | AIFF を MP3 に変換 | `adapter/ffmpeg` |
| `VoiceCatalog` | 利用可能ボイス一覧の取得 | `adapter/say` |
| `TextSource` | 入力テキストの取得 | `adapter/textsource` (4 実装) |

新しい入力ソースや出力フォーマットを追加する場合は、対応するアダプタを実装するだけで `usecase` / `domain` を変更する必要はありません。

## テスト戦略

| 層 | テストしやすさ | 方針 |
|---|---|---|
| `domain` | 高 (純粋関数) | 単体テスト (例: `internal/domain/chunk_test.go`) |
| `usecase` | 中 | モックアダプタを注入したテスト |
| `adapter` | 低 (外部依存) | macOS 上のスモークテスト |
