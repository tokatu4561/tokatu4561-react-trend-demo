import { NextApiHandler } from 'next'
import { fetchJson } from '@/lib/api'

const API_URL = process.env.APP_API_URL

const handler: NextApiHandler = async (req, res) => {
  const { jwt } = req.cookies

  if (!jwt) {
    res.status(401).end()
    return
  }

  try {
    const user = await fetchJson(`${API_URL}/auth/user`, {
      headers: { Authorization: `Bearer ${jwt}` },
    })
    console.log(user)

    res.status(200).json({})
  } catch (error) {
    res.status(401).end()
  }
}

export default handler
