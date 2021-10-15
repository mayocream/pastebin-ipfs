import 'twin.macro'
import { useEffect, useState } from 'react'
import { Avatar, Chip, Container, TextField, Paper, Button } from '@material-ui/core'
import { Match, navigate, RouteComponentProps } from '@reach/router'
import dayjs from 'dayjs'
import { CidResource } from '@/types'
import relativeTime from 'dayjs/plugin/relativeTime'
import { findFileLanguage } from '@/util/fileTypes'
import '@/css/prism-nord.css'

dayjs.extend(relativeTime)

const cid404 = 'QmeFH2sPs1bvwbZ3fP7kD4Tza7suYxZYLBNnu4LDGMQeug'

interface ViewProps {
  cid: string
}

async function getMetadata(cid: string): Promise<CidResource> {
  return await fetch(import.meta.env.VITE_API_URL + `/api/v0/${cid}/__metadata.json`).then((res) => {
    console.log(res.json)
    return res.json() as Promise<CidResource>
  })
}

async function getFile(cid: string, filename: string): Promise<string> {
  return await fetch(import.meta.env.VITE_API_URL + `/api/v0/${cid}/${filename}`).then((res) => {
    console.log(res.text)
    return res.text()
  })
}

function View(props: RouteComponentProps<ViewProps>) {
  const cid = props.cid || cid404
  const [metadata, setMetadata] = useState<CidResource>()
  const [langCode, setLangCode] = useState('plain')
  const [text, setText] = useState('')

  useEffect(() => {
    getMetadata(cid).then((data) => {
      setMetadata(data)

      const lang = findFileLanguage(data.objects[0].name)
      setLangCode(lang)

      getFile(cid, data.objects[0].name)
        .then((txt) => setText(txt))
        .then(() => {
          // @ts-ignore
          Prism.highlightAll()
        })
    })
  }, [])

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
          <Button tw='ml-[auto]' size="small" variant="outlined" color="secondary" href={import.meta.env.VITE_API_URL + `/api/v0/${cid}/${metadata?.objects[0].name}`}>
            Raw
          </Button>
        </section>
        <time tw="italic">Created {dayjs().to(dayjs(metadata?.created_at))}</time>
        <Paper tw="mt-2 rounded-[10px]">
          <section>
            <pre tw="rounded-[10px]">
              <code tw="text-[0.8125rem]" className={`language-${langCode}`}>
                {text}
              </code>
            </pre>
          </section>
        </Paper>
      </article>
    </Container>
  )
}

export { View }
