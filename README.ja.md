# json-filter

テキストストリームから JSON を抽出・検証・整形・修復するコマンドラインフィルター。LLM 出力・API レスポンス・ログデータの処理に最適です。

English documentation: [README.md](README.md)

## 特徴

- **JSON 抽出**: 正規表現を使ってテキストストリームに埋め込まれた JSON オブジェクト・配列を識別・抽出
- **自動整形**: 有効な JSON を適切なインデントでフォーマット
- **不完全 JSON の修復**: 欠損している閉じ括弧・角括弧を補完してマルフォームド／切り詰め JSON を修復
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

**curl と組み合わせて使用:**

```sh
curl -s https://api.github.com/users/octocat | json-filter
```

## ビルド

Go 1.16 以上が必要です。

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
