import React from 'react'
import { Router } from '@reach/router'
import AppBar from './layout/AppBar'
import Publish from './pages/publish'

function App() {
  return (
    <>
      <AppBar />
      <Router>
        <Publish path="/" />
      </Router>
    </>
  )
}

export default App
