import 'twin.macro'
import { useEffect } from 'react'
import { Container, TextField } from '@material-ui/core'
import { Match, navigate } from '@reach/router'

function View() {
  useEffect(() => {
    fetch(import.meta.env.VITE_API_URL + '/QmPNKURHDipHmBvND1SkcHQBXbHZYXBrDnroiu1dvLZFVM/raw/nihao').then((res) =>
      console.log(res.json)
    )
  }, [])
  let cid = ''

  function submit(code: string) {
    if (code === 'Enter') navigate(`/view/${cid}`)
  }

  return (
    <Container tw="pt-[35vh]" maxWidth="md">
      <Match path="/view/:cid">
        {(props) =>
          props.match ? (
            <div>fdsfds</div>
          ) : (
            <TextField
              id="outlined-cid"
              label="Please enter cids to search"
              variant="outlined"
              // color="secondary"
              fullWidth
              tw="shadow-lg"
              onKeyDown={(e) => submit(e.code)}
              onChange={(e) => (cid = e.target.value)}
            />
          )
        }
      </Match>
    </Container>
  )
}

export { View }
