# 📦 MemoIndex v1.0.0 リリースノート

**MemoIndex** は、指定フォルダ内の `.txt` / `.md` / `.html` ファイルを自動でインデックス化し、全文検索できるローカルアプリです。  
GUI と CLI の両方に対応しており、用途に応じて使い分けることができます。

# 簡単な使い方
## 初回起動
### Windowsの場合
1. 任意のフォルダにてZIPファイルを解凍する。
2. MemoIndex_GUI.exe をダブルクリックして起動
3. 入力欄にテキストを追加し「新規メモ帳」を押して内容を保存

### ターミナルの場合
1. ターミナルで解凍した階層へ移動
2. ./memoindex or memoindex を実行（GUIが起動します。）


## 既にあるテキストファイルをインデックス化する
※過去に1度以上起動済みであることが前提です。
1. memoindex config add-dir [ディレクトリパス]
2. memoindex reindex
Linuxでは、memoindex → ./memoindex で動作します。


---

## 🆕 主な機能

- 複数フォルダにまたがるメモを自動インデックス化
- 高速全文検索とスニペット付き結果表示
- ファイルの更新を自動検知し、インデックスを更新
- ワンクリックで新規メモ作成（GUI）
- 日本語 / 英語 切り替え対応

---

## 💻 ダウンロードファイル

| ファイル名 | 内容 |
|------------|------|
| `MemoIndex-windows-amd64.zip` | Windows 64bit 向け（GUI & CLI 同梱） |
| `MemoIndex-linux-amd64.zip`   | Linux 64bit 向け（CLI、GUIはFyne環境必須） |

`.zip` を展開後、`config.yaml.sample` を `config.yaml` にリネームして設定を行ってください。

---

## 🚀 使い方（CLI）

```bash
./memoindex search "キーワード"     # キーワード検索
./memoindex new mynote.txt          # 新規メモ作成
./memoindex reindex                 # インデックス再構築
````

設定の変更例：

```bash
./memoindex config lang ja                 # 日本語に変更
./memoindex config add-dir ./notes        # フォルダ追加
./memoindex config index-path ./idx.bleve # 保存先変更
./memoindex config editor code            # 使用エディタ変更
```

---

## 🪟 GUI の使い方

`memoindex_gui` を実行すると、検索ウィンドウが起動します。
キーワード入力、検索、メモ作成がボタンで行えます。

---

## ⚙ 設定ファイル（config.yaml）

```yaml
memo_dirs:
  - "./notes"
index_path: "./myindex.bleve"
editor: "notepad"
language: "ja"
```

---

## 🌐 対応言語

* 日本語 (`ja`)
* 英語 (`en`)

---

## 📄 ライセンス

[[MIT License](https://chatgpt.com/c/LICENSE)](LICENSE)

## 制作・著作
本ソフトウェア MemoIndex の制作著作は [YashubuStudio](https://ykvario.com) に帰属します。


開発向け情報は `development_Readme.md` を参照してください。
