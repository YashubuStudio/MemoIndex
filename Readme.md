# MemoIndex

MemoIndex は、指定したフォルダ内のテキストファイルを自動でインデックス化し、素早く全文検索できるローカルアプリです。GUI ウィンドウまたは CLI から利用できます。現在はビルドの都合上、CLI 用と GUI 用で個別に実行ファイルを作成する構成になっています。

## 特長

- `.txt` / `.md` / `.html` の内容を自動でインデックス化
- ファイルを追加・編集するとインデックスを自動更新
- 高速全文検索とスニペット表示
- ワンクリックで新規メモ帳を作成

## インストール

1. GitHub Releases からお使いの OS 用バイナリ（CLI 版 `memoindex` / GUI 版 `memoindex_gui`）をダウンロードし、任意の場所に展開します。
2. ソースからビルドする場合は Go をインストールした上で、このリポジトリをクローンします。CLI 用は `go build -o memoindex`、GUI 用は `go build -ldflags="-H windowsgui" -o memoindex_gui` でバイナリを作成します（GUI は fyne ツールでのパッケージングも可能）。CLI のみ試す場合は `go run .` でも実行できます。
3. `config.yaml.sample` を `config.yaml` にコピーし、監視したいフォルダやエディタパスを設定します。

## 使い方

### CLI 版

```bash
./memoindex search "キーワード"   # 検索
./memoindex new [ファイル名]       # 新規メモ作成
./memoindex reindex               # インデックスを再構築
```

設定変更は `memoindex config` サブコマンドで行えます。例:

```bash
./memoindex config lang ja            # 言語を日本語に設定
./memoindex config add-dir ./notes    # フォルダを追加（自動で絶対パスに変換）
./memoindex config index-path ./idx.bleve  # インデックスの保存先変更
./memoindex config editor vim         # 使用エディター変更
```

### GUI 版

`./memoindex_gui` を実行すると検索ボックス付きのウィンドウが開きます。検索や新規メモ作成をボタン操作で行えます。

## 設定

`config.yaml` で以下を設定できます。

- `memo_dirs` : 監視するフォルダのパス (複数指定可)
- `index_path` : インデックスファイルの保存先
- `editor` : 新規メモ作成時に開くエディタ
- `language` : 使用する言語 (検索結果表示など)

### 対応言語

- 日本語 (`ja`)
- 英語 (`en`)

### 代表的なエディター

- notepad
- vim
- nano
- gedit
- code (Visual Studio Code)

## ライセンス

MIT

開発に関する詳細な情報は `development_Readme.md` を参照してください。
