# json-filter

テキストストリームから JSON を抽出・検証・整形・修復するコマンドラインフィルター。LLM 出力・API レスポンス・ログデータの処理に最適です。

English documentation: [README.md](README.md)

## 特徴

- **JSON 抽出**: 再帰降下パーサー（[nlk/jsonfix](https://github.com/nlink-jp/nlk)）を使ってテキストストリームに埋め込まれた JSON オブジェクト・配列を識別・抽出
- **自動整形**: 有効な JSON を 2 スペースインデントでフォーマット
- **強力な JSON 修復**: LLM 出力・API レスポンス・ログデータに見られる 20 種以上の問題を修復:
  - Markdown コードフェンス（`` ```json ... ``` ``）
  - シングルクォートのキー・値
  - 末尾カンマ
  - クォートなしキー
  - 閉じ括弧・角括弧の欠損（深いネストにも対応）
  - コメント（`//`、`/* */`、`#`）
  - Python 形式リテラル（`True`、`False`、`None`）
  - 二重エスケープ JSON（`{\"key\": \"value\"}`）
  - その他 — 全一覧は [nlk/jsonfix ドキュメント](https://github.com/nlink-jp/nlk) を参照
- **バイパスモード**: `--bypass` フラグで抽出失敗時に元の入力をそのまま stdout に流し、パイプラインの中断を防止

## インストール

[リリースページ](https://github.com/nlink-jp/json-filter/releases) からプラットフォームに合ったバイナリをダウンロードしてください。

```sh
unzip json-filter-<version>-<os>-<arch>.zip
mv json-filter /usr/local/bin/
```

## 使用方法

`json-filter` は stdin から読み込み、処理した JSON を stdout に書き出します。

```sh
<コマンド> | json-filter [flags]
```

### フラグ

| フラグ | 説明 |
|--------|------|
| `--bypass` | JSON 抽出が失敗した場合、元の入力をそのまま出力 |
| `--version` | バージョン情報を表示して終了 |

### 例

**ログ出力から JSON を抽出・整形:**

```sh
echo 'INFO: data: {"id": 1, "name": "Alice"}' | json-filter
# {
#   "id": 1,
#   "name": "Alice"
# }
```

**不完全な JSON を修復:**

```sh
echo '{"data": {"item": "value"' | json-filter
# {
#   "data": {
#     "item": "value"
#   }
# }
```

**Markdown コードフェンスから JSON を抽出:**

```sh
printf '```json\n{"key": "value"}\n```\n' | json-filter
# {
#   "key": "value"
# }
```

**シングルクォート・末尾カンマ・クォートなしキーを修復:**

```sh
echo "{'name': 'Alice', 'age': 30,}" | json-filter
# {
#   "name": "Alice",
#   "age": 30
# }
```

**curl と組み合わせて使用:**

```sh
curl -s https://api.github.com/users/octocat | json-filter
```

## ビルド

Go 1.26 以上が必要です。

```sh
git clone https://github.com/nlink-jp/json-filter.git
cd json-filter
make build        # 現在のプラットフォーム向けにビルド → dist/json-filter
make build-all    # 全プラットフォーム向けにクロスコンパイル
make package      # ビルドして .zip アーカイブを作成
make test         # テストを実行
make clean        # dist/ を削除
```

対応プラットフォーム: `linux/amd64`, `linux/arm64`, `darwin/amd64`, `darwin/arm64`, `windows/amd64`

## 関連リンク

- [CHANGELOG.md](CHANGELOG.md)
