// AES Encrypt
export async function aesEncrypt(password: string, data: string): Promise<Uint8Array> {
  const enc = new TextEncoder()
  const iv = crypto.getRandomValues(new Uint8Array(16))
  const key = await crypto.subtle.importKey('raw', enc.encode(password), 'AES-GCM', false, ['encrypt', 'decrypt'])
  const encryptedData: Uint8Array = await crypto.subtle.encrypt(
    {
      name: 'AES-GCM',
      iv: iv,
    },
    key,
    enc.encode(data)
  )
  return encryptedData
}
