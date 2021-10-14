# pastebin-ipfs
![go](https://github.com/mayocream/pastebin-ipfs/actions/workflows/go.yml/badge.svg)
![node](https://github.com/mayocream/pastebin-ipfs/actions/workflows/node.yml/badge.svg)
![docker](https://github.com/mayocream/pastebin-ipfs/actions/workflows/docker.yml/badge.svg)

Pastebin built on IPFS, securely served by Distributed Web and Edge Network.

It's like [gist](https://gist.github.com/) but for anonymous.

[Preview](https://paste.shoujo.io)

## Features

<!-- - [Gallery](https://paste.shoujo.io/gallery) shows *Public* snippets -->
- Stored in [IPFS](https://ipfs.io/) distributed network
- *AES-GCM* Encryption
- Open API (CORS Origin `*`)
- File upload (API Only)
- Cache by CDN (or [IPFS Gateway](https://cloudflare-ipfs.com))
- Syntax highlight by [Prismjs](https://github.com/PrismJS/prism)
- No Expiration

## Usage

### Web

Webpage serves [Gallery](https://paste.shoujo.io/gallery) and provide GUI to paste your snippets.

Vist [Webpage](https://paste.shoujo.io).

### Terminal

Create snippets:

```bash
$ curl -T doc.md https://paste.shoujo.io/ # remember to have a slash '/' at the end
# or
$ curl -X POST https://paste.shoujo.io -d 'いつか君に伝えたいと思っていた気持ちは'
# or
$ curl -X PUT https://paste.shoujo.io/lyrics.txt -d 'Stars fall, birds sleep'
```

Cat snippets:

```bash
curl https://paste.shoujo.io/QmTnhJH8azDsudkxgp8wNLEN5Zq86NAE6DAkzwGBDpaQ6Z/raw/plain.txt
```

## Self-Hosted

### Docker

You must have ipfs-daemon running on your host first.

```bash
docker run -p 8080:3939 pastebin-ipfs:latest
```

### Docker Compose

Edit [docker-compose.yml](https://github.com/mayocream/pastebin-ipfs/blob/main/deploy/docker/docker-compose.yml) file.

```bash
docker-compose up -d
```

## Develop

```bash
make run # start ipfs daemon at http://127.0.0.1:5001
         # run pastebin API at http://127.0.0.1:3939
make web-live # run Webpage
```

## QA

You can provide suggestion or ask question by open a Github Issue in any languages. 🧐

## Todo

- [ ] replace ipfs daemon with [ipfs-lite](github.com/hsanjuan/ipfs-lite).

## LICENSE

MIT
