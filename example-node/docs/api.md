# API

DMCのバックエンドとして最低限実装する必要がある必須APIについて説明します。  
APIの定義については `swagger.yaml` を参照してください。  

## /swagger.json

swagger.jsonを配信するAPIです。  
クライアント側でエンドポイントとして指定するパスになるため、  
example-nodeのように `/` を `/swagger.json` にリダイレクトさせると良いと思います。  
swagger.jsonを配信するAPIであればパスは `/swagger.json` である必要はありません。  

## /dmc_authtype

サーバがサポートしている認証方式をクライアント側に伝えるためのAPIです。  
認証に利用するパスなどもこのレスポンスに含めます。  
非認証状態でコールできる必要があり、パスは固定です。  

## /dmc

アプリケーションの情報と画面構成を返すAPIです。  
権限に応じて画面構成を変更させるため認証必須で、パスは固定です。  