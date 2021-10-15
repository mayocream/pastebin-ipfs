import { RouteComponentProps } from '@reach/router'
import React, { useEffect } from 'react'

function Gallery(props: RouteComponentProps) {
  useEffect(() => {
    // 405
    fetch(import.meta.env.VITE_API_URL + '/api/v0/gallery').then((res) => console.log(res.json))
  }, [])
  return <div>Gallery</div>
}

export { Gallery }
