import { Router, RouteComponentProps } from '@reach/router'
import { Publish } from '@/pages/Publish'
import { View } from '@/pages/View'
import { Gallery } from '@/pages/Gallery'

const RPublish = (props: RouteComponentProps) => <Publish />
const RView = (props: RouteComponentProps) => <View />
const RGallery = (props: RouteComponentProps) => <Gallery />

function Route() {
  return (
    <Router>
      <RPublish path="/" />
      <RView path="/:cid" />
      <RGallery path="/gallery" />
    </Router>
  )
}

export { Route }
