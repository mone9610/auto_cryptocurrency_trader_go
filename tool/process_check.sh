#goのヘルスチェック用のパスを指定する
_URL="localhost:8000/health"
#無限ループにてヘルスチェックを実施する
while true
do
    sleep 10
    HTTP_RESPONSE=$(curl --silent --write-out "HTTPSTATUS:%{http_code}" -X GET $_URL)
    HTTP_STATUS=$(echo $HTTP_RESPONSE | tr -d '\n' | sed -e 's/.*HTTPSTATUS://')
    echo "$HTTP_STATUS"
    if [ "$HTTP_STATUS" -eq '200' ]; then
    echo "ok"
    else
    echo "goのヘルスチェックに失敗しました。goのプロセスを再起動します。"
    docker-compose exec go bash -c "go run main"
    fi
done
