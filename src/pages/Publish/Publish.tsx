import React, { useState } from 'react'
import { Button, Checkbox, Fab, Switch, TextField } from '@material-ui/core'
import { Container, Paper } from '@material-ui/core'
import { Controller, useForm } from 'react-hook-form'
import { highlight, languages } from 'prismjs'

import 'twin.macro'
import 'prismjs/components/prism-clike'
import 'prismjs/components/prism-javascript'
import 'prismjs/components/prism-markup'
import 'prismjs/components/prism-css'

import AddIcon from '@material-ui/icons/Add'
import Editor from 'react-simple-code-editor'

interface IPublicData {
  filename: string
  author: string
  public: boolean
  encrypt: boolean
  password?: string
}

function Publish() {
  const { handleSubmit, control, watch } = useForm()
  const [code, setCode] = useState('')

  const watchEncrypt = watch('encrypt', true)

  const onSubmit = async (data: IPublicData) => {
    console.log(data)
    const formData = new FormData()
    formData.append('author', data.author)
    // @ts-ignore
    formData.append('public', data.public)
    formData.append('filename', data.filename ?? 'plain.txt')

    let blob = new Blob([code], { type: 'text/plain' })
    // encrypt
    if (data?.password) {
      const enc = new TextEncoder()
      const iv = crypto.getRandomValues(new Uint8Array(16))
      const key = await crypto.subtle.importKey('raw', enc.encode(data.password), 'AES-GCM', false, [
        'encrypt',
        'decrypt',
      ])
      const encryptedCode: Uint8Array = await crypto.subtle.encrypt(
        {
          name: 'AES-GCM',
          iv: iv,
        },
        key,
        enc.encode(code)
      )
      blob = new Blob([encryptedCode], { type: 'text/plain' })
    }

    formData.append('file', blob)
    const resp = await fetch(import.meta.env.VITE_API_URL + '/api/v0/upload', {
      method: 'POST',
      body: formData,
      mode: 'cors',
    })
    const result = await resp.json()
    console.log(result)
  }

  return (
    <Container maxWidth="md">
      <div tw="fixed right-[60px] bottom-[60px]">
        <Fab color="primary" aria-label="add">
          <AddIcon />
        </Fab>
      </div>
      <Paper>
        <div tw="flex m-6 ml-16 pt-4 pb-6 w-full">
          <form onSubmit={handleSubmit(onSubmit)}>
            <div tw="w-[300px]">
              <Controller
                // TODO find code language by file suffix, eg: "a.md, b.js, c.yaml"
                name="filename"
                control={control}
                defaultValue=""
                rules={{ required: true }}
                render={({ field }) => <TextField fullWidth label="File name" {...field} />}
              />
              {/* {errors.filename && <span>This field is required</span>} */}
            </div>
            <div tw="my-4 min-w-[800px]">
              <Paper variant="outlined">
                <Editor
                  value={code}
                  onValueChange={(code) => setCode(code)}
                  // TODO support more languages
                  highlight={(code) => highlight(code, languages.markup, 'markup')}
                  padding={14}
                  tabSize={4}
                  insertSpaces={true}
                  name="code"
                  placeholder="Put something here.."
                  style={{
                    minHeight: '400px',
                    fontFamily: '"Fira code", "Fira Mono", monospace',
                    fontSize: 13,
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
            <div>
              <Controller
                name="public"
                control={control}
                defaultValue={true}
                render={({ field }) => <Switch {...field} />}
              />
              Show on gallery
            </div>
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
            {/* TODO add more files by repeat "<Paper>" component */}
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
