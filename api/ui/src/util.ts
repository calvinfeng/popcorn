import axios, { AxiosRequestConfig } from 'axios'
import * as Cookies from 'es-cookie'


interface UserState {
  authenticated: boolean
  name: string
  email: string
  imageUrl: string
}

export function configureInterceptor(token: string) {
  axios.interceptors.request.use((config: AxiosRequestConfig) => {    
    config.headers.token = `${token}`
  
    return config
  }, (err: any) => {
    return Promise.reject(err)
  });
}

export function setUserStateToCookie(state: UserState) {
  Cookies.set("name", state.name)
  Cookies.set("email", state.email)
  Cookies.set("image_url", state.imageUrl)
}

export function getUserStateFromCookie(): UserState {
  if (Cookies.get("name") === "") {
    return {
      authenticated: false,
      name: "",
      email: "",
      imageUrl: ""
    }
  }

  return {
    authenticated: true,
    name: Cookies.get("name"),
    email: Cookies.get("email"),
    imageUrl: Cookies.get("image_url")
  }
}

