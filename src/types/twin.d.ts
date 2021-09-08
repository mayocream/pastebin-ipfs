import 'twin.macro'
import { css as cssImport } from '@emotion/react'

declare module 'twin.macro' {
    // The styled and css imports
    const css: typeof cssImport
}

