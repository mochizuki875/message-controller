# message-controller
`message`リソースの`spec`で設定した文字列を`status`に反映するコントローラー。

```
$ kubectl get message
NAME               WORD               NUMBER
message-sample-1   Hello World        5
message-sample-2   Hello Controller   8
```

## Environment

```
$ kubebuilder version
Version: main.version{KubeBuilderVersion:"3.7.0", KubernetesVendor:"1.24.1", GitCommit:"3bfc84ec8767fa760d1771ce7a0cb05a9a8f6286", BuildDate:"2022-09-20T17:21:57Z", GoOs:"darwin", GoArch:"amd64"}
$ go version
go version go1.19 darwin/amd64
$ kind version
kind v0.15.0 go1.19 darwin/amd64
```