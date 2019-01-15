import * as React from 'react'
import * as ReactDOM from 'react-dom'
import { GoogleLogin, GoogleLoginResponse, GoogleLogout } from 'react-google-login'
import axios, { AxiosResponse, AxiosError, AxiosRequestConfig } from 'axios'
import * as Cookies from 'es-cookie'
import "./app.scss"
import { configureInterceptor, setUserStateToCookie, getUserStateFromCookie } from './util'
import Camera from './ components/camera'

// OAuth Web Client ID
const WEB_CLIENT_ID = '1098793859190-nrralnk204e67g7seeok729qrt0jhd32.apps.googleusercontent.com';

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
    let img: JSX.Element;
    if (this.state.authenticated) {
      button = <GoogleLogout buttonText="Logout" onLogoutSuccess={this.handleLogoutSuccess} />
      img = <img src={this.state.imageUrl} height="43" width="43" />

    } else {
      button = <GoogleLogin clientId={WEB_CLIENT_ID}
        buttonText="Login" onSuccess={this.handleLoginSuccess} onFailure={this.handleLoginError} /> 
      img = <div />
    }

    return (
      <section className="header">
        <div className="left-box">
          {img}
        </div>
        <div className="right-box">
          {button}
        </div>
      </section>
    )
  }

  get content() {
    if (this.state.authenticated) {
      const greeting = (
        <section className="greeting">
          <h2>Welcome, {this.state.name}</h2>
          <p>{this.state.email}</p>
        </section>
      )

      return (
        <section className="real-content">
          {greeting}
          <Camera />
        </section>
      )
    }

    return (
      <section className="landing-content">
        <img src="/android-chrome-512x512.png"></img>
        <p>Login to find out what movie we recommend!</p>
      </section>
    )
  }

  render() {
    return (
      <section className="popcorn">
        {this.header}
        {this.content}
      </section>
    )
  }
}

document.addEventListener("DOMContentLoaded", () => {
  ReactDOM.render(<App />, document.getElementById("app"))
})