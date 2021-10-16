# pastebin-ipfs
![go](https://github.com/mayocream/pastebin-ipfs/actions/workflows/go.yml/badge.svg)
![node](https://github.com/mayocream/pastebin-ipfs/actions/workflows/node.yml/badge.svg)
![docker](https://github.com/mayocream/pastebin-ipfs/actions/workflows/docker.yml/badge.svg)
![deploy](https://github.com/mayocream/pastebin-ipfs/actions/workflows/deploy.yml/badge.svg)

<img width="200px" src="./docs/images/ipfs-archivists.svg" />

<small>_(IPFS Archivists)_</small>    

*仍处于开发阶段，欢迎提交 Pull Request。*

基于 IPFS 的 Pastebin，由 _去中心化网络_ 和 _边缘网络_ 安全驱动。

类似于 [gist](https://gist.github.com/)，但不需要登陆账号。
[Ubuntu Pastebin](https://paste.ubuntu.com/) 的替代品。

[paste.shoujo.io](https://paste.shoujo.io)

<table>
  <td><img width="500px" src="./docs/images/index.png" /></td>
  <td><img width="500px" src="./docs/images/view.png" /></td>
</table>
         
## 特性

<!-- - [Gallery](https://paste.shoujo.io/gallery) shows *Public* snippets -->
- 数据由 [IPFS](https://ipfs.io/) 去中心化储存
- *AES-GCM* 加密
- 开发 API（OpenAPI v3，跨域 CORS `*`）
- 文件上传（仅限 API）
- CDN 缓存（或 [IPFS 网关](https://cloudflare-ipfs.com)）
- [Prismjs](https://github.com/PrismJS/prism) 语法高亮
- 无过期时间（受限于 IPFS）

## 使用

### 网页（Web）

Web 版提供近期发布[画册看板（未实现）](https://paste.shoujo.io/gallery)、操作文件的可视化面板。

访问 [Web 页面](https://paste.shoujo.io)。

### API

**API 文档**: [*OpenAPI v3 - Swagger UI*](https://mayocream.github.io/pastebin-ipfs/api/)    

Pastebin 限制每个用户的请求速率为 20 QPS。

### 终端（Terminal）

创建 Snippet：

```bash
$ curl -T doc.md https://paste.shoujo.io/api/v0/ # remember to have a slash '/' at the end
# or
$ curl -X POST https://paste.shoujo.io/api/v0/ -d 'いつか君に伝えたいと思っていた気持ちは'
# or
$ curl -X PUT https://paste.shoujo.io/api/v0//lyrics.txt -d 'Stars fall, birds sleep'
```

获取 Snippet：

```bash
curl https://paste.shoujo.io/api/v0/QmTnhJH8azDsudkxgp8wNLEN5Zq86NAE6DAkzwGBDpaQ6Z/plain.txt
```

## 私有化部署（Self-Hosted）

### Kubernetes

使用 [Helm](https://helm.sh/) 部署 pastebin-ipfs.

```bash
git clone https://github.com/mayocream/pastebin-ipfs
cd pastebin-ipfs/helm
helm install pastebin-ipfs .
```

参阅 [values.yaml](./helm/values.yaml) 了解详细参数。

### Docker Compose

编辑 [deploy/docker/docker-compose.yml](https://github.com/mayocream/pastebin-ipfs/blob/main/deploy/docker/docker-compose.yml) 文件.

```bash
docker-compose up -d
```

### Docker

你必须先在主机上运行 ipfs-daemon。

IPFS 运行示例： [docker-compose.yml](https://github.com/mayocream/pastebin-ipfs/blob/main/docker-compose.yml).

```bash
docker run -p 8080:3939 pastebin-ipfs:latest
```

## 开发

```bash
make run # start ipfs daemon at http://127.0.0.1:5001
         # run pastebin API at http://127.0.0.1:3939
make web-live # run Webpage
```

## 反馈

欢迎通过 Github Issue 提交建议和反馈，不限制语言。🧐

## Todo

- [ ] replace ipfs daemon with [ipfs-lite](github.com/hsanjuan/ipfs-lite).

## 致谢

- [Web Crypto Encryption and Decryption Example](https://github.com/bradyjoslin/webcrypto-example)

## LICENSE

MIT
