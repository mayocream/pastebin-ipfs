import React, { useEffect } from 'react'

function Gallery() {
  useEffect(() => {
    // 405
    fetch(import.meta.env.VITE_API_URL + '/gallery').then((res) => console.log(res.json))
  }, [])
  return <div>Gallery</div>
}

export { Gallery }
