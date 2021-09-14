import { Container, Paper } from '@material-ui/core'
import React, { useState } from 'react'
import { useForm, Controller } from 'react-hook-form'
import { TextField, Checkbox, Switch, Button, Fab } from '@material-ui/core'
import AddIcon from '@material-ui/icons/Add'
import 'twin.macro'
import Editor from 'react-simple-code-editor'
import { highlight, languages } from 'prismjs'
import 'prismjs/components/prism-clike'
import 'prismjs/components/prism-javascript'
import 'prismjs/components/prism-markup'
import 'prismjs/components/prism-css'

function Publish() {
  const {
    handleSubmit,
    control,
    reset,
    formState: { errors },
  } = useForm()
  const [code, setCode] = useState('')
  const onSubmit = async (data: any) => {
    console.log(data)
    const formData = new FormData()
    formData.append('author', data.author)
    formData.append('public', data.public)
    formData.append('filename', data.filename ?? 'plain.txt')

    const blob = new Blob([code], { type: 'text/plain' })
    formData.append('file', blob)
    const resp = await fetch(import.meta.env.VITE_API_URL + '/api/v1/upload', {
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
