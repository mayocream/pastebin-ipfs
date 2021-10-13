import React from 'react'
import { MenuAppBar } from './layout/AppBar'
import { Publish } from './pages/Publish'
import { View } from './pages/View'
import { Gallery } from './pages/Gallery'
import { Router, RouteComponentProps } from '@reach/router'

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
      <MenuAppBar />
      <Sections />
    </>
  )
}

export default App
