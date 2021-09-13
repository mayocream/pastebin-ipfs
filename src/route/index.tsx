import React from 'react'
import { Router } from '@reach/router'
import { MenuAppBar } from './layout/AppBar'
import { Publish } from './pages/Publish'

function App() {
  return (
    <>
      <MenuAppBar />
      <Router>
        <Publish path="/" />
      </Router>
    </>
  )
}

export default App
