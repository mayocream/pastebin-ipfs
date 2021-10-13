const fileAlias = new Map([
  ['rs', 'rust'],
  ['sh', 'bash'],
])

// find file language by file extension
export function findFileLanguage(filename: string): string {
  const ext = filename.split('.').pop()
  if (ext === undefined || ext === '') {
    return 'plain'
  }
  const alias = fileAlias.get(ext)
  if (alias) {
    return alias
  }

  return ext
}
