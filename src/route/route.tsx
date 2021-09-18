import { Router, RouteComponentProps } from '@reach/router'
import { Publish } from '@/pages/Publish'
import { View } from '@/pages/View'
import { Gallery } from '@/pages/Gallery'
import 'twin.macro'

const RPublish = (props: RouteComponentProps) => <Publish />
const RView = (props: RouteComponentProps) => <View />
const RGallery = (props: RouteComponentProps) => <Gallery />

function Route() {
  return (
    <div tw="h-full">
      <Router>
        <RPublish path="/" />
        <RView path="/:cid" />
        <RGallery path="/gallery" />
      </Router>
    </div>
  )
}

export { Route }
