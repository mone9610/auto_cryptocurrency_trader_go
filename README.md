auto_cryptocurrency_trader_go
====

## 概要  
bitFlyerLightningAPIを利用して、仮想通貨を自動取引してくれるツールです。  

## 詳細  
any%ルールに基づいて、ETH/JPYをbitFlyer上で自動取引できます。  
デフォルトの設定では、tickerから取得した最終最高価格の-5%で指値で買い注文を発注します。買い注文が約定した場合には、買い注文の最終約定価格から+5%で指値で売り注文を発注し、それと同時に逆指値で売り注文を-10%の価格で発注します。売り注文が約定した場合には、再度同様のフローで買い注文を発注します。

## 機能  
* any%ルールに基づく指値買い注文・売り注文・逆指値注文の自動発注機能
* ダッシュボード(開発中)
* 買い注文履歴表示機能(開発中)
* 売り注文履歴表示機能(開発中)
* Ticker履歴表示機能(開発中)
* 取引高表示機能(開発中)

## 実行環境
* View(開発中)
  * Nginx
  * BootStrap4(予定)
  * TypeScript(予定)
* Model, Contorller
  * go 1.15
  * MySQL 5.6
* その他
  * phpMyAdmin
  * bash
  * docker
  * docker-compose 

## アーキテクチャ
<img width="814" alt="スクリーンショット 2021-01-17 15 43 26" src="https://user-images.githubusercontent.com/64692797/104833281-c1a71700-58da-11eb-94cf-812b38a46c6e.png">

## 使い方
1. bitFlyerに新規アカウントを作成した上で、口座開設、入金、アクセスキーおよびシークレットキーの取得を実施してください。  
[bitFlyer新規アカウント作成はこちらから](https://bitflyer.com/ja-jp/account-create)  

2. 実行するホスト上にこのリポジトリをcloneしてください。
```
$ git clone https://github.com/mone9610/auto_cryptocurrency_trader_go.git
```

3. リポジトリをcloneしたら、コンテナを起動してください。
```
$ docker-compose up -d
```
4. 下記のファイルを実行し、MySQLとgoのプロセスを起動してください。
```
$ bash auto-cryptcurrency-trader/tool/init.sh & auto-cryptcurrency-trader/tool/process_check.sh.sh &
```

5. 80番ポートにブラウザからアクセスし、簡易設定画面からアクセスキーとシークレットキーを入力してください。
```
localhost:80
```
※パブリックIPが付与されている場合は、localhostをパブリックIPに読み替えてください。
## 注意事項
* 本ツールはMITライセンスですが、bitflyerLightningAPIに関する権利はbitFlyer社様へ帰属します。当ツールを転用して商用利用等を検討される場合には、bitFlyer社様へお問い合わせください。
* 取得したアクセスキー、シークレットキーに関しては、ご自身の責任において十分に注意してお取り扱いをお願いいたします。
* 本ツールをパブリッククラウド等にデプロイして運用する際には、ご自身の責任においてSSL化等を実施し、セキュリティの確保に努めてください。

## 免責事項
* 当方は当ツールの継続性や機能等を何ら保証するものではなく、これらの欠陥、瑕疵等について、これらを使用したこと、又は使用できなかったことから生じる一切の損害に関して、いかなる責任も負いません。
* 当ツールの使用により、birFlyer社または第三者が損害を被った場合は、使用者が一切の責任を負うものとします。

## ライセンス
[MIT](https://github.com/mone9610/auto_cryptocurrency_trader_go/blob/main/LICENSE)

## 制作者
[mone9610](https://github.com/mone9610)
