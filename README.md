# tolog

vimでメモ/ログをとるためのプラグイン

- マークダウン
- 一日1ファイル
- TODO機能
- ログ機能
とか

## スクリーンショット

![](https://github.com/tortuepin/tolog/blob/images/images/tolog_image.png)


※下記の設定例の通りに設定した場合の動作
![](https://github.com/tortuepin/tolog/blob/images/images/tolog_mo.gif)

## インストール

### 手動

1. このリポジトリをクローン
2. vimのパスに追加
3. `make`を実行

### dein

```
[[plugins]]
repo = 'tortuepin/tolog'
build = 'make clean;make'
on_ft = 'markdown'
```
こんな感じに設定しておく


## 機能

### todo

#### 基本
`## todo`
という行があると、次の`## `ではじまる行か、文書の終わりまでがtodoを扱う部分として処理される。

例:
```
## todo

@home
- [ ] 水道代払う
- [ ] 家賃払う

@work
- [ ] 書類作成

```

`@`からはじまる行はタグとして処理される。
タグがあると、次の空白行までにあるtodoにそのタグが付与される。


#### Tolog_todo_set_active( [n] )

この関数を呼び出すと過去n日分のファイルから完了していないtodoを抜き出し、タグごとにまとめ現在のファイルに追記する。

nはデフォルトでは30。

### log

#### 基本

`## log`
という行があると、次の`## `で始まる行か、文書の終わりまでがlogを扱う部分として処理される。

例:
```
## log

[20:42]

夕飯をたべた

[20:43]

お風呂はいる

[20:53] @work @check

書類作成のために必要な情報まとめる

[20:46] @work

必要な情報まとめおわった
```

時間を表す`[??:??]`ではじまる行はから次の`[??:??]`までが1つのlog要素となる。

`[??:??]`から始まる行にタグを付けることができる。
タグは複数つけられる。

#### Tolog_add_log([tag1 tag2 ...])

この関数を呼び出すと引数に指定したtagがつけられたlog要素をファイルの末尾に追加する。

#### Tolog_log_search_bytag([tag1 tag2 ...])

この関数を呼び出すと引数に指定したtagがつけられたすべてのlog要素を過去のファイルから抜き出し、別windowで結果を表示する。

### その他

#### Tolog_tag_collect()

過去のファイルからタグを集めてリストにする。

#### Tolog_Complete_tag(...)

Tolog_tag_collectで集められたtagを返す。
タグの補完に使う。

#### Tolog_read_template()

`g:tolog_template_dir`に指定されたファイルを読み込んで現在開いているファイルに追記する。
テンプレートの読み込みに使う。

#### Tolog_get_[today, prev, next]_filename()

それぞれ、今日、前の日、次の日のファイル名を取得する。

#### Tolog_tumbling_dice() range

複数行を選択してこの関数を呼び出すとその中からランダムに1行抜き出して表示する。


## 設定例
```
let g:tolog_dir = "log/"
let g:tolog_template_dir = g:tolog_dir . "template.md"


command! -nargs=* -complete=custom,Tolog_Complete_tag T call Tolog_add_log(<f-args>) " :Tでlog追記
command! -nargs=* -complete=custom,Tolog_Complete_tag SearchByTag call Tolog_log_search_bytag(<f-args>) "SearchByTag @tag でtag検索
command! On call Tolog_open_log(Tolog_get_next_filename()) " 次の日のlog開く
command! Op call Tolog_open_log(Tolog_get_prev_filename()) " 前の日のlogを開く
command! Ot call Tolog_open_log(Tolog_get_today_filename()) " 今日のlogを開く
command! ReadTemp call Tolog_read_template() " テンプレートの読み込み
command! -range Dice <line1>,<line2>call Tolog_tumbling_dice() "サイコロをふる
```


## tolog_deoplete_source

https://github.com/tortuepin/tolog_deoplete_source

deopleteを使っている場合、上記のsourceを使うことでtagを入力する際に補完が効くようになる。
