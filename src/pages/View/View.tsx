import React from 'react'
import { TextField, Container } from '@material-ui/core'
import 'twin.macro'

function View() {
  return (
    <Container tw="pt-[35vh]" maxWidth="md">
      <TextField
        id="outlined-cid"
        label="Please enter cids to search"
        variant="outlined"
        // color="secondary"
        fullWidth
        tw="shadow-lg"
      />
    </Container>
  )
}

export { View }
