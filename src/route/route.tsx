import { Router, RouteComponentProps } from '@reach/router'
import { Publish } from '@/pages/Publish'
import { View } from '@/pages/View'
import { Gallery } from '@/pages/Gallery'
import { NotFound } from '@/pages/NotFound'
import 'twin.macro'

const RPublish = (props: RouteComponentProps) => <Publish />
const RView = (props: RouteComponentProps) => <View />
const RGallery = (props: RouteComponentProps) => <Gallery />
const RNotFound = (props: RouteComponentProps) => <NotFound />

function Route() {
  return (
    <div tw="h-full">
      <Router>
        <RNotFound default />
        <RPublish path="/" />
        <RView path="/cid" />
        <RGallery path="/gallery" />
      </Router>
    </div>
  )
}

export { Route }
