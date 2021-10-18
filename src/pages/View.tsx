import 'twin.macro'
import { ReactNode, useEffect, useState } from 'react'
import { Avatar, Chip, Container, TextField, Paper, Button } from '@mui/material'
import { Match, navigate, RouteComponentProps } from '@reach/router'
import dayjs from 'dayjs'
import { CidResource } from '@/types'
import relativeTime from 'dayjs/plugin/relativeTime'
import { findFileLanguage } from '@/util/fileTypes'
import '@/css/prism-nord.css'
import { decryptData } from '@/util/crypt'
import { useSnackbar } from 'notistack'

dayjs.extend(relativeTime)

const cid404 = 'QmeFH2sPs1bvwbZ3fP7kD4Tza7suYxZYLBNnu4LDGMQeug'

interface ViewProps {
  cid: string
}

async function getMetadata(cid: string): Promise<CidResource> {
  return await fetch(import.meta.env.VITE_API_URL + `/ipfs/${cid}/__metadata.json`).then((res) => {
    console.log(res.json)
    return res.json() as Promise<CidResource>
  })
}

async function getFile(cid: string, filename: string): Promise<string> {
  return await fetch(import.meta.env.VITE_API_URL + `/ipfs/${cid}/${filename}`).then((res) => {
    console.log(res.text)
    return res.text()
  })
}

function renderContent(lang: string, text: string): ReactNode {
  if (lang !== 'md') {
    return (
      <Paper tw="mt-4 rounded-[10px]">
        <pre tw="rounded-[10px]">
          <code tw="text-[0.8125rem]" className={`language-${lang}`}>
            {text}
          </code>
        </pre>
      </Paper>
    )
  }

  // @ts-ignore
  const result: any = markdownit({
    html: true,
    langPrefix: 'language-',
    linkify: true,
  }).render(text)
  return (
    <div
      dangerouslySetInnerHTML={{
        __html: result,
      }}
    ></div>
  )
}

function View(props: RouteComponentProps<ViewProps>) {
  const cid = props.cid || cid404
  const [metadata, setMetadata] = useState<CidResource>()
  const [langCode, setLangCode] = useState('plain')
  const [text, setText] = useState('')
  const [password, setPassword] = useState('')
  const [revealText, setRevealText] = useState(localStorage.getItem(`reveal-${cid}`) || '')
  const { enqueueSnackbar, closeSnackbar } = useSnackbar()

  const highlight = () => {
    // @ts-ignore
    Prism.highlightAll()
  }

  useEffect(() => {
    getMetadata(cid).then((data) => {
      setMetadata(data)

      const lang = findFileLanguage(data.objects[0].name)
      setLangCode(lang)

      getFile(cid, data.objects[0].name)
        .then((txt) => setText(txt))
        .then(() => {
          highlight()
        })
    })
  }, [])

  const decrypt = async () => {
    if (password === '') return
    const data = await decryptData(text, password)
    if (data === '') {
      enqueueSnackbar('Password incorrect', { variant: 'error' })
      return
    }
    localStorage.setItem(`reveal-${cid}`, data)
    setRevealText(data)
    highlight()
  }

  return (
    <Container tw="pt-[60px]" maxWidth="md">
      <article tw="flex flex-col">
        <section tw="flex items-end mb-2">
          <h1 tw="text-4xl inline-block mr-5 mt-2 mb-0">{metadata?.objects[0].name}</h1>
          <Chip
            avatar={<Avatar src={`https://en.gravatar.com/avatar/${metadata?.author}?d=monsterid`} />}
            label={metadata?.author}
            variant="outlined"
          />
          {metadata?.encrypt_algorithm !== 'none' && (
            <Chip
              tw="ml-2"
              label={`${revealText === '' ? 'Encrypted' : 'Decrypted'} (${metadata?.encrypt_algorithm.toUpperCase()})`}
            />
          )}
          <Button
            tw="ml-[auto]"
            size="small"
            variant="outlined"
            color="secondary"
            download={metadata?.objects[0].name}
            href={
              revealText !== ''
                ? `data:text/plain;base64,${btoa(unescape(encodeURIComponent(revealText)))}`
                : import.meta.env.VITE_API_URL + `/ipfs/${cid}/${metadata?.objects[0].name}`
            }
          >
            Raw
          </Button>
        </section>
        <time tw="italic">Created {dayjs().to(dayjs(metadata?.created_at))}</time>
        <section>
          {metadata?.encrypt_algorithm !== 'none' && revealText === '' ? (
            <div tw="my-10 max-w-[300px]">
              <TextField
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                size="small"
                required
                id="password"
                type="password"
                label="Password"
                variant="outlined"
              />
              <Button onClick={decrypt} tw="mt-6" variant="contained" color="primary">
                Confirm
              </Button>
            </div>
          ) : (
            renderContent(langCode, revealText !== '' ? revealText : text)
          )}
        </section>
      </article>
    </Container>
  )
}

export { View }
