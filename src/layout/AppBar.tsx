import 'twin.macro'

import {
  AppBar,
  Drawer,
  IconButton,
  List,
  ListItem,
  ListItemIcon,
  ListItemText,
  Toolbar,
  Typography,
} from '@mui/material'
import { makeStyles } from '@mui/styles'
import { Inbox, Mail, Menu, GitHub } from '@mui/icons-material'

import React from 'react'
import { navigate } from '@reach/router'

const useStyles = makeStyles((theme) => ({
  root: {
    flexGrow: 1,
  },
  title: {
    flexGrow: 1,
  },
}))

function MenuAppBar() {
  const classes = useStyles()
  const [isDrawerVisible, setIsDrawVisible] = React.useState(false)

  return (
    <div className={classes.root}>
      <AppBar position="static">
        <Toolbar>
          <IconButton
            edge="start"
            color="inherit"
            aria-label="menu"
            onClick={() => setIsDrawVisible(true)}
          >
            <Menu />
          </IconButton>
          <Typography variant="h6" className={classes.title}>
            Pastebin/IO
          </Typography>
          <IconButton color="inherit" href="https://github.com/mayocream/pastebin-ipfs">
            <GitHub />
          </IconButton>
        </Toolbar>
        <Drawer anchor="left" open={isDrawerVisible} onClose={() => setIsDrawVisible(false)}>
          <List>
            {['Publish', 'Gallery'].map((text, index) => (
              <ListItem
                button
                key={text}
                tw="w-60"
                onClick={() => navigate(text !== 'Publish' ? text.replace(' ', '-').toLowerCase() : '/')}
              >
                <ListItemIcon>{index % 2 === 0 ? <Inbox /> : <Mail />}</ListItemIcon>
                <ListItemText primary={text} />
              </ListItem>
            ))}
          </List>
        </Drawer>
      </AppBar>
    </div>
  )
}

export { MenuAppBar }
