export interface CidResource {
    author: string
    created_at: string
    encrypt_algorithm: string
    objects: CidObject[]
}

export interface CidObject {
    name: string
    mime_type: string
    size: number
}
