import { NextApiRequest } from 'next'

export interface UserResponse {
  accessToken: string
  idToken: string
  userId: string
}

export interface ResponseError {
  error: boolean
  message: string
}

export interface UserAccount {
  id: string
  slug: string
  userId: string
}

export interface IronSessionRequest extends NextApiRequest {
  session: {
    //TODO:remove any
    get: (arg0: string) => any
    set: (arg0: string) => any
  }
}
