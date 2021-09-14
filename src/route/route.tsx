import { Router, RouteComponentProps } from '@reach/router'
import { Publish } from '@/pages/Publish'
import { Cid } from '@/pages/Cid'
import { Gallery } from '@/pages/Gallery'
import { ApiTests } from '@/pages/ApiTests'
import { ApiDocs } from '@/pages/ApiDocs'

const RPublish = (props: RouteComponentProps) => <Publish />
const RCid = (props: RouteComponentProps) => <Cid />
const RGallery = (props: RouteComponentProps) => <Gallery />
const RApiTests = (props: RouteComponentProps) => <ApiTests />
const RApiDocs = (props: RouteComponentProps) => <ApiDocs />

function Route() {
  return (
    <Router>
      <RPublish path="/" />
      <RCid path="/cid" />
      <RGallery path="/gallery" />
      <RApiTests path="/api-tests" />
      <RApiDocs path="/api-docs" />
    </Router>
  )
}

export { Route }
