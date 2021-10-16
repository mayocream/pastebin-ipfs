import React from 'react'
import { MenuAppBar } from './layout/AppBar'
import { Publish } from './pages/Publish'
import { View } from './pages/View'
import { Gallery } from './pages/Gallery'
import { Router, RouteComponentProps } from '@reach/router'
import { SnackbarProvider } from 'notistack'
import 'twin.macro'

function Sections() {
  return (
    <div tw="h-full">
      <Router>
        <Publish path="/" />
        <View path="/:cid" />
        <Gallery path="/gallery" />
      </Router>
    </div>
  )
}

function App() {
  return (
    <>
      <SnackbarProvider>
        <MenuAppBar />
        <div tw="min-h-[80vh]">
          <Sections />
        </div>
        <aside tw="my-5 text-center invisible">
          <span tw="italic">
            ©2021 Shoujo/IO, <a href="https://github.com/mayocream/pastebin-ipfs/issues">Feedback</a>
          </span>
        </aside>
      </SnackbarProvider>
    </>
  )
}

export default App
