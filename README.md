# pastebin-ipfs

Pastebin built on IPFS, securely served by Distributed Web and Edge Network.

It's like [gist](https://gist.github.com/) but for anonymous.

[Preview](https://pasteb.shoujo.io)

## Features

- [Gallery](https://paste.shoujo.io/gallery) shows *Public* snippets
- Stored in [IPFS](https://ipfs.io/) distributed network
- *AES-GCM* Encryption
- Open API (CORS Origin `*`)
- File upload
- Cache by CDN (or [IPFS Gateway](https://cloudflare-ipfs.com))
- Syntax highlight by [Prismjs](https://github.com/PrismJS/prism)
- No Expiration

## Usage

### Web

Webpage serves [Gallery](https://paste.shoujo.io/gallery) and provide GUI to paste your snippets.

Vist [Webpage](https://paste.shoujo.io).

### Terminal

```bash
curl -T doc.md https://paste.shoujo.io/ # remember to have a slash '/' at the end
```

## Develop

```bash
make compose # start ipfs daemon at http://127.0.0.1:5001
make run # run pastebin API at http://127.0.0.1:3939
npm run dev # run Webpage
```

## QA

You can provide suggestion or ask question by open a Github Issue in any languages. 🧐

## Todo

- [ ] replace ipfs daemon with [ipfs-lite](github.com/hsanjuan/ipfs-lite).

## LICENSE

MIT
