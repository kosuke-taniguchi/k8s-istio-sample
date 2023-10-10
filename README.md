# kindを使ってk8s、istio, ArgoRollouts周りをいじってみたのでメモ
## やったこと・手順
- [kind](https://kind.sigs.k8s.io/)でclusterを作成
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

## ArgoRollouts手順
- rollout.yaml作成
- `kubectl apply -f ./deployments/rollout.yaml`
### 新しいDockerイメージをデプロイする際
- docker build
- deployments/deployment.yamlのimageバージョン更新
- `kubectl apply -f ./deployments/deployment.yaml`
