# jmactts docs

Astro + Starlight で構築した jmactts のドキュメントサイト。

公開先: <https://junara.github.io/jmactts/>

## 開発

```sh
npm install
npm run dev      # http://localhost:4321/jmactts/
npm run build    # ./dist へ書き出し
npm run preview  # ビルド済みサイトをローカルプレビュー
```

ドキュメント本文は `src/content/docs/{ja,en}/` 以下の Markdown/MDX を編集してください。

## デプロイ

`main` ブランチに `docs/**` の変更を push すると [`.github/workflows/docs.yml`](../.github/workflows/docs.yml) が起動し、GitHub Pages へ自動デプロイされます。
