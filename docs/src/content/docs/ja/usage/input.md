---
title: 入力ソース
description: 引数 / ファイル / 標準入力 / クリップボードからテキストを読み込む
---

`jmactts` は 4 つの入力経路をサポートします。複数指定された場合の優先順位は次のとおりです。

```
-c (クリップボード)  >  -f (ファイル)  >  位置引数  >  パイプされた標準入力
```

## 位置引数

引数として渡したテキストを読み上げます。複数の引数はスペースで連結されます。

```bash
jmactts こんにちは 世界
jmactts -L ja "句読点を含む、文章。"     # スペースを保ちたいときはクォート
```

## ファイル (`-f`)

`-f` / `--file` でテキストファイルを指定します。

```bash
jmactts -f speech.txt
jmactts -L ja -f long_article.txt -o article.mp3
```

`-f -` を指定すると標準入力から読み込みます。これは「明示的に stdin を入力ソースにする」指定で、位置引数があっても無視されます。

```bash
echo "Hello" | jmactts -f -
```

## 標準入力 (パイプ)

引数も `-f` も `-c` も指定されず stdin がパイプされている場合は、自動的に標準入力から読み込みます。

```bash
echo "標準入力からのテスト" | jmactts
cat article.txt | jmactts -L ja
git log -1 --pretty=%s | jmactts -L ja      # 直近コミットメッセージを読み上げ
```

## クリップボード (`-c`)

`-c` / `--clipboard` で macOS のクリップボード (`pbpaste`) から読み込みます。

```bash
jmactts -c
jmactts -c -L ja      # 日本語ボイスで
```

ブラウザでコピーした記事をそのまま音声化したいときに便利です。
