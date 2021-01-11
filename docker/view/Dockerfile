FROM  nginx:latest

#設定ファイル(conf)のみビルド時にイメージに追加する
#その他の静的ファイルはdocker-composeでvolumeにマウントする
ADD ./conf/default.conf /etc/nginx/conf.d/