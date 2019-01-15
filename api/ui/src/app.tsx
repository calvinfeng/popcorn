import * as React from 'react'
import * as ReactDOM from 'react-dom'
import { GoogleLogin, GoogleLoginResponse, GoogleLogout } from 'react-google-login'
import axios, { AxiosResponse, AxiosError, AxiosRequestConfig } from 'axios'
import * as Cookies from 'es-cookie'
import "./app.scss"
import { configureInterceptor, setUserStateToCookie, getUserStateFromCookie } from './util'
import Camera from './ components/camera'

// OAuth Client ID
const CLIENT_ID = '159011759683-01llm5cirgtboge88g73342bl9nn1ihb.apps.googleusercontent.com'

interface State {
  email: string
  name: string
  imageUrl: string
  authenticated: boolean
}

interface Props {}

class App extends React.Component<Props, State> {
  constructor(props) {
    super(props)
    this.state = {
      authenticated: false,
      name: "",
      email: "",
      imageUrl: ""
    }
  }

  handleLoginSuccess = (res: GoogleLoginResponse) => {
    const profile = res.getBasicProfile();

    const newState = {
      authenticated: true,
      name: profile.getName(),
      email: profile.getEmail(),
      imageUrl: profile.getImageUrl()
    }

    setUserStateToCookie(newState)
    this.setState(newState)

    const auth = res.getAuthResponse()
    Cookies.set("google_id_token", auth.id_token)
    configureInterceptor(auth.id_token)
  }

  handleLoginError = (err: any) => {
    console.error(err)
  }

  handleLogoutSuccess = () => {
    Cookies.remove("google_id_token");

    const newState = {
      authenticated: false,
      email: "",
      name: "",
      imageUrl: ""
    }

    setUserStateToCookie(newState)
    this.setState(newState)
  }
  
  componentDidMount() {
    if (Cookies.get("google_id_token")) {
      configureInterceptor(Cookies.get("google_id_token"))
      const newState = getUserStateFromCookie()
      this.setState(newState)
    }
  }

  get header() {
    let button: JSX.Element;
    if (this.state.authenticated) {
      button = <GoogleLogout buttonText="Logout" onLogoutSuccess={this.handleLogoutSuccess} />
    } else {
      button = <GoogleLogin clientId={CLIENT_ID}
        buttonText="Login" onSuccess={this.handleLoginSuccess} onFailure={this.handleLoginError} /> 
    }

    return (
      <section className="header">
        {button}
      </section>
    )
  }

  get content() {
    if (this.state.authenticated) {
      return (
        <section className="content">
          <h2>Welcome, {this.state.name}</h2>
          <h4>{this.state.email}</h4>
          <img src={this.state.imageUrl} />
          <Camera />
        </section>
      )
    }

    return (
      <section className="content">
        <p>Login to find out what movie we recommend!</p>
      </section>
    )
  }

  render() {
    return (
      <section id="popcorn">
        {this.header}
        {this.content}
      </section>
    )
  }
}

document.addEventListener("DOMContentLoaded", () => {
  ReactDOM.render(<App />, document.getElementById("app"))
})