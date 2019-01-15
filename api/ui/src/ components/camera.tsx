import * as React from 'react'

interface State {
  display: boolean
  error: string
}

interface Props {}

class Camera extends React.Component<Props, State> {
  private player = React.createRef<HTMLVideoElement>()
  private canvas = React.createRef<HTMLCanvasElement>()

  constructor(props) {
    super(props)

    this.state = {
      display: true,
      error: ""
    }
  }

  handleCapture = () => {
    const canvas = this.canvas.current;
    const player = this.player.current;
    canvas.getContext('2d').drawImage(player, 0, 0, canvas.width, canvas.height);
  }

  componentDidMount() {
    const constraints = {
      video: true,
      audio: false
    }

    navigator.mediaDevices.getUserMedia(constraints).then((stream) => {
      this.player.current.srcObject = stream
    }).catch(() => {
      this.setState({
        display: false,
        error: "unable to locate a camera stream"
      })
    })
  }

  render() {
    if (this.state.display) {
      return (
        <section className="camera">
          <video id="player" ref={this.player} controls={true} autoPlay={true}></video>
          <button id="capture" onClick={this.handleCapture}>Capture</button>
          <canvas id="canvas" ref={this.canvas} width="320" height="240"></canvas>
        </section>
      )
    }

    return (
      <section className="camera">
        <p>Error: {this.state.error}</p>
      </section>
    )
  }
}

export default Camera
