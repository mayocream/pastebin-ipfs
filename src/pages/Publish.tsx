import React, { useState, useEffect } from 'react'
import { Button, Fab, Switch, TextField } from '@mui/material'
import { Container, Paper } from '@mui/material'
import { Controller, useForm } from 'react-hook-form'
import { highlight, languages } from 'prismjs'

import 'twin.macro'
import 'prismjs/components/prism-clike'
import 'prismjs/components/prism-javascript'
import 'prismjs/components/prism-markup'
import 'prismjs/components/prism-css'

import AddIcon from '@mui/icons-material/Add'
import Editor from 'react-simple-code-editor'
import { encryptData } from '@/util/crypt'
import { RouteComponentProps } from '@reach/router'

interface ISubmitData {
  author: string
  isPublic: boolean
  filename: string
  encrypt: boolean
  password: string
}

function Publish(props: RouteComponentProps) {
  const { handleSubmit, control, watch } = useForm()
  const [text, setText] = useState('')

  const watchEncrypt = watch('encrypt', false)

  const onSubmit = async (data: ISubmitData) => {
    console.debug(data)

    const formData = new FormData()

    if (data.author !== '') {
      formData.append('author', data.author)
    }
    // formData.append('public', data.isPublic)
    const filename = data.filename === '' ? 'plain.txt' : data.filename

    let content: Uint8Array | string
    if (data.encrypt) {
      content = await encryptData(text, data.password)
      formData.append('encrypt_algorithm', 'aes-gcm')
    } else {
      content = text
    }
    const blob = new Blob([content], { type: 'text/plain' })
    formData.append('file', blob, filename)

    const resp = await fetch(import.meta.env.VITE_API_URL + '/api/v0/upload', {
      method: 'POST',
      body: formData,
      mode: 'cors',
    })
    if (resp.status === 201) {
      const cid = await resp.json()
      window.location.href = `/` + cid
    }
  }

  return (
    <Container maxWidth="md">
      {/* <div tw="fixed right-[60px] bottom-[60px]">
        <Fab color="primary" aria-label="add">
          <AddIcon />
        </Fab>
      </div> */}
      <Paper>
        <div tw="flex m-6 ml-16 pt-4 pb-6 w-full">
          <form onSubmit={handleSubmit(onSubmit)}>
            <div tw="w-[300px]">
              <Controller
                name="filename"
                control={control}
                defaultValue=""
                render={({ field }) => <TextField fullWidth label="Filename" {...field} />}
              />
            </div>
            <div tw="my-4 min-w-[800px] max-w-[800px]">
              <Paper variant="outlined">
                <Editor
                  value={text}
                  onValueChange={(text) => setText(text)}
                  highlight={(text) => highlight(text, languages.markup, 'markup')}
                  padding={14}
                  tabSize={4}
                  insertSpaces={true}
                  name="code"
                  placeholder="Put something here.."
                  style={{
                    minHeight: '400px',
                    fontFamily: '"Fira code", "Fira Mono", monospace',
                    fontSize: 13,
                    overflow: 'scroll',
                  }}
                />
              </Paper>
            </div>
            <div tw="my-4 w-[200px]">
              <Controller
                name="author"
                control={control}
                defaultValue=""
                render={({ field }) => <TextField fullWidth label="Author" {...field} />}
              />
            </div>
            {/* <div>
              <Controller
                name="isPublic"
                control={control}
                defaultValue={true}
                render={({ field }) => <Switch {...field} />}
              />
              Show on gallery
            </div> */}
            <div>
              <Controller
                name="encrypt"
                control={control}
                defaultValue={false}
                render={({ field }) => <Switch {...field} />}
              />
              Encrypt
            </div>
            {watchEncrypt && (
              <div tw="my-4 w-[200px]">
                <Controller
                  name="password"
                  control={control}
                  defaultValue=""
                  render={({ field }) => <TextField fullWidth label="Password" {...field} />}
                />
              </div>
            )}
            <div tw="mt-6">
              <Button type="submit" color="primary" variant="contained">
                Publish
              </Button>
            </div>
          </form>
        </div>
      </Paper>
    </Container>
  )
}

export { Publish }
