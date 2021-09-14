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
  makeStyles,
} from '@material-ui/core'
import { Inbox, Mail, Menu } from '@material-ui/icons'

import React from 'react'
import { navigate } from '@reach/router'

const useStyles = makeStyles((theme) => ({
  root: {
    flexGrow: 1,
  },
  menuButton: {
    marginRight: theme.spacing(2),
  },
  title: {
    flexGrow: 1,
  },
}))

function MenuAppBar() {
  const classes = useStyles()
  // const [auth, setAuth] = React.useState(true)
  // const [anchorEl, setAnchorEl] = React.useState(null)
  const [isDrawerVisible, setIsDrawVisible] = React.useState(false)

  // const handleChange = (event) => {
  //   setAuth(event.target.checked)
  // }

  return (
    <div className={classes.root}>
      <AppBar position="static">
        <Toolbar>
          <IconButton
            edge="start"
            className={classes.menuButton}
            color="inherit"
            aria-label="menu"
            onClick={() => setIsDrawVisible(true)}
          >
            <Menu />
          </IconButton>
          <Typography variant="h6" className={classes.title}>
            Pastebin/IO
          </Typography>
        </Toolbar>
        <Drawer anchor="left" open={isDrawerVisible} onClose={() => setIsDrawVisible(false)}>
          <List>
            {['Publish', 'Gallery', 'Cid', 'API Tests', 'API Docs'].map((text, index) => (
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
