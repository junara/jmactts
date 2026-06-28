package cli

import (
	"fmt"
	"io"
)

func PrintUsage(w io.Writer) {
	fmt.Fprint(w, `jmactts - macOS の say コマンドを使った多言語テキスト読み上げ CLI

Usage:
  jmactts [flags] [text...]

Examples:
  jmactts こんにちは 世界
  echo "Hello, world" | jmactts -v Samantha
  jmactts -f speech.txt
  jmactts -c                                # クリップボードの内容を読み上げ
  jmactts -L ja_JP こんにちは              # 国/言語コードで自動ボイス選択
  jmactts -L JP こんにちは                  # 国コード (JP) でも可
  jmactts -L en Hello world                # 言語コード (en) でも可
  jmactts -v Kyoko -r 200 ゆっくり読ませたいテキスト
  jmactts -o hello.m4a こんにちは
  jmactts -o hello.mp3 こんにちは          # ffmpeg が必要
  jmactts -l -L ja                         # 日本語ボイス一覧のみ表示

長文 (1500文字超) は文区切り (。 .! ? ! ? 改行) で自動分割して順次再生。
再生中の Ctrl-C で即座に停止します。

Flags:
  -v, --voice <name>    ボイス名 (例: Kyoko / Otoya / Samantha)
  -L, --lang <code>     国/言語コード (ja_JP / ja / JP / en / US 等)。-v 未指定時に自動選択
  -r, --rate <wpm>      話速 (words per minute)
  -f, --file <path>     入力テキストファイル ('-' で標準入力)
  -c, --clipboard       クリップボード (pbpaste) からテキスト読み込み
  -o, --output <path>   音声ファイル出力 (拡張子で判定: .aiff/.wav/.m4a/.aac/.mp3)
  -l, --list-voices     利用可能なボイス一覧 (-L で絞り込み可)
  -V, --version         バージョン表示
  -h, --help            このヘルプ
`)
}
