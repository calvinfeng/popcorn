import * as React from 'react'
import './camera.scss'

interface State {
  display: boolean
  error: string
}

interface Props {}

class Camera extends React.Component<Props, State> {
  private player = React.createRef<HTMLVideoElement>()
  private canvas = React.createRef<HTMLCanvasElement>()

  constructor(props: any) {
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
        error: "unable to locate a camera stream, here's a test footage"
      })
    })
  }

  render() {
    if (this.state.display) {
      return (
        <section className="camera">
          <div className="video-container">
            <video id="player" ref={this.player} controls={false} autoPlay={true}></video>
          </div>
          <button id="capture" onClick={this.handleCapture}>Capture</button>
          <canvas id="canvas" ref={this.canvas}></canvas>
        </section>
      )
    }

    return (
      <section className="camera">
        <p>Error: {this.state.error}</p>
        <div className="video-container">
          <video id="test-player" src="http://clips.vorwaerts-gmbh.de/big_buck_bunny.mp4" 
            autoPlay={true}></video>
        </div>
      </section>
    )
  }
}

export default Camera
