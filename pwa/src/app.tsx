import * as React from 'react'
import * as ReactDOM from 'react-dom'
import { GoogleLogin, GoogleLoginResponse } from 'react-google-login'
import axios from 'axios'

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
      email: "",
      name: "",
      imageUrl: "",
      authenticated: false
    }
  }

  handleLoginSuccess = (res: GoogleLoginResponse) => {
    const profile = res.getBasicProfile();
    this.setState({
      email: profile.getEmail(),
      name: profile.getName(),
      imageUrl: profile.getImageUrl(),
      authenticated: true
    })

    const auth = res.getAuthResponse()
    console.log(auth.id_token)
  }

  handleLoginError = (err: any) => {
    console.log(err)
  }

  render() {
    if (this.state.authenticated) {
      return <section>
        <h2>Welcome, {this.state.name}</h2>
        <h3>{this.state.email}</h3>
        <img src={this.state.imageUrl} />
      </section>
    }

    return <GoogleLogin clientId={CLIENT_ID} buttonText="Login" 
      onSuccess={this.handleLoginSuccess} 
      onFailure={this.handleLoginError} />
  }
}

document.addEventListener("DOMContentLoaded", () => {
  ReactDOM.render(<App />, document.getElementById("app"))
})