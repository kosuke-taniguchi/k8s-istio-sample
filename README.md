# kindを使ってk8s、istio周りをいじってみたのでメモ
## やったこと・手順
- kindでclusterを作成
  - `kind create cluster --config=multi-node.yaml`
- goのファイルを作成
  - 適当に環境変数を表示させるみたいな
- Dockerfile作成
- docker build
- `kind load docker-image`コマンドでDockerイメージをkindクラスターにロードする
  - GCRとか使ってイメージをPush, Pullしなくてよくなる
- istioのインストール
  - https://istio.io/latest/docs/setup/getting-started/#download
- k8sのリソース定義
  - Deployment
    - 今回は環境変数を変えるために2つのDeploymentを用意した
  - Service
- kubectl apply
- istioのマニフェスト用意
  - DestinationRule
  - VirtualService
  - Gateway
- istioのマニフェストapply
  - `kubectl apply -f {ファイル名}`
- `curl localhost:30070`で確認できる