import React from 'react'
import {
  IconButton,
  Typography,
  Toolbar,
  AppBar,
  Drawer,
  makeStyles,
  List,
  ListItem,
  ListItemIcon,
  ListItemText,
} from '@material-ui/core'
import { Menu, Inbox, Mail } from '@material-ui/icons'

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
  const [auth, setAuth] = React.useState(true)
  const [anchorEl, setAnchorEl] = React.useState(null)
  const [isDrawerVisible, setIsDrawVisible] = React.useState(false)

  const handleChange = (event) => {
    setAuth(event.target.checked)
  }

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
            {['Inbox', 'Starred', 'Send email', 'Drafts'].map((text, index) => (
              <ListItem button key={text}>
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
