# MemoIndex

指定フォルダ内の `.txt`, `.md`, `.html` ファイルを自動インデックス化し、全文検索・新規作成を可能にするローカルアプリです。  
CLI（コマンドライン）とGUI（専用ウィンドウ）の両方から利用できます。

---

## 機能概要

- 指定フォルダ配下の `.txt`, `.md`, `.html` ファイルのみ自動インデックス化
- ファイルの追加・変更・削除をリアルタイムで監視し、インデックスを自動更新
- 高速全文検索とヒット箇所のスニペット（前後30文字）表示
- 新規メモ帳作成（エディタ自動起動＆即インデックス化）
- CLI／GUI（専用ウィンドウ）両対応

---

## 画面イメージ

### 専用ウィンドウ（GUI）

- 検索ボックス＋検索ボタン
- 新規メモ帳ボタン
- 検索結果一覧（ファイル名＋スニペット、クリックでファイルを開く／検索結果は閉じない）

```

---

## | 🔍 \[検索ワード                 ]\[検索]   \[＋新規メモ]

\| 検索結果:
\|  - notes/bleve\_tips.txt
\|     ...Goで全文検索をするならBleve。基本的な使い方としては...
\|  - docs/search\_engine.md
\|     ...全文検索（Full-text search）は、Go言語のBleveが...
\|  - web/index.html
\|     ...<h2>Go言語による全文検索の実装例</h2>...
-------------------------------------

````

### CLI

#### ファイル検索（上位3件＋スニペット表示）

```bash
memoindex search "検索ワード"
````

出力例:

```
1. notes/bleve_tips.txt
   ...Goで全文検索をするならBleve。基本的な使い方としては...

2. docs/search_engine.md
   ...全文検索（Full-text search）は、Go言語のBleveが...

3. web/index.html
   ...<h2>Go言語による全文検索の実装例</h2>...
```

#### 新規メモ帳作成

```bash
memoindex new [ファイル名]
```

指定したファイル名（省略時は日時など自動付与）で新規メモ帳を作成し、既定のエディタで開きます。

---

## システム仕様

* **インデックス対象**:

  * 指定ディレクトリ以下の `.txt`, `.md`, `.html` ファイルのみ
* **自動監視・更新**:

  * フォルダを監視し、ファイルの追加・編集・削除に応じて自動でBleveインデックスを更新
* **検索機能**:

  * 全文検索（AND/OR/部分一致）
  * ヒットしたファイルのパス＋該当箇所のスニペット（前後30文字）表示
* **新規メモ帳**:

  * ワンクリックまたはコマンド一発で空ファイル作成＋既定エディタ自動起動＋自動インデックス登録
* **複数窓同時起動可**（CLI・GUI混在可）
* **設定ファイル**:

  * 監視フォルダやエディタパス、除外ディレクトリなどを設定可能

---

## 技術構成

* **言語**: Go
* **全文検索エンジン**: [Bleve](https://github.com/blevesearch/bleve)
* **ファイル監視**: [fsnotify](https://github.com/fsnotify/fsnotify)
* **GUI**: fyne, go-astilectron等
* **CLI**: Cobra, urfave/cli等

---

## 使い方

### 1. インストール・初期設定

1. Goインストール
2. 本リポジトリをクローン
3. `go build -o memoindex` で実行ファイルを作成
4. `config.yaml.sample` を `config.yaml` にコピーして監視フォルダ等を設定
   - `memo_dirs` や `index_path` には絶対パスも指定可能です

### 2. 起動

* CLI:

  ```bash
  memoindex search "Go全文検索"
  memoindex new mymemo.txt
  memoindex reindex
  ```
* GUI:

  ```
  memoindex gui
  ```

### 3. ビルドとリリース

#### クロスコンパイル

`GOOS` と `GOARCH` を指定すると他 OS/アーキテクチャ向けにビルドできます。

```bash
# 例: Windows 64bit 向けの実行ファイルを生成
GOOS=windows GOARCH=amd64 go build -o memoindex.exe
```

#### 配布例

生成したバイナリと `config.yaml.sample` を zip 等にまとめて GitHub Releases へアップロードします。  
CI (`.github/workflows/go.yml`) を拡張すれば、タグ作成時に自動ビルド・アーカイブ生成も可能です。

---

## ライセンス

MIT

---

## TODO / 今後の予定

* 検索結果のフィルタ・ソート
* 高度な検索構文（正規表現・タグ付け等）
* インデックス対象の拡張（PDF等）

---

## 開発・コントリビュート

PR・Issue歓迎です！
