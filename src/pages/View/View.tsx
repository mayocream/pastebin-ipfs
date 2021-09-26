import 'twin.macro'

import { Container, TextField } from '@material-ui/core'
import { Match, navigate } from '@reach/router'

function View() {
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
