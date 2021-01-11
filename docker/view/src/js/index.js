// アクセスキー、シークレットキーを格納したjsonを作成する関数
function makeConfJson(accessKey,secretKey){
    var conf = new Object();
    conf.Access_Key = accessKey;
    conf.Secret_Key = secretKey;
    var json = JSON.stringify(conf);
    return json;
}

// リクエストの送信先urlを作成する関数
function setUrl(){
    var host = document.getElementById('hostName').value;
    var path ='/api/v1/conf';
    var url = 'http://' +  host + ':8000'  + path;
    return url;
}

// PUTメソッドを実行するための関数
// urlとjsonのリクエストボディを引数とする
function doPUT(url,data){
    console.log(url);
    var xhr = new XMLHttpRequest();
    xhr.open('PUT', url);
    xhr.setRequestHeader("Content-Type", "application/json");
    xhr.send(data);
    xhr.onreadystatechange = function() {
        if(xhr.readyState === 4 && xhr.status === 200) {
            //データを取得後の処理を書く
            console.log("200 OK");
            // HACK:別の関数で要素の書き換えは今後行う
            var status = document.getElementById('status');
            status.textContent = '更新成功';
        }
    }
}

// GETメソッドを実行するための関数
// urlを引数とする
// function doGET(url){
//     var xhr = new XMLHttpRequest();
//     xhr.open('GET', url);
//     xhr.onreadystatechange = function() {
//         if(xhr.readyState === 4 && xhr.status === 200) {
//             //データを取得後の処理を書く
//             console.log("200 OK");
//             var data = xhr.responseText;
//             console.log(data);
//             return data;
//         }
//     }
//     xhr.send(null);
// }

// アクセスキー、シークレットキーを更新するための関数
function updateConf(){
    var accessKey = document.getElementById('accessKey').value;
    var secretKey = document.getElementById('secretKey').value;
    var data = makeConfJson(accessKey,secretKey);
    var url = setUrl();
    doPUT(url,data);
}

// ページ読み込み時に、現在設定されているアクセスキー、シークレットキーを取得するための関数
// function readConf(){
//     var accessKey = document.getElementById('accessKey');
//     var secretKey = document.getElementById('secretKey');
//     var url = setUrl();
//     var json = doGET(url);
//     var obj = JSON.parse(json);
//     console.log(json);
//     console.log(obj.Access_Key);
//     console.log(obj.Secret_Key);
// }