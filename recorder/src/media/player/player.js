import React, { Component } from 'react'
import ReactPlayer from 'react-player'

// NOTE: Local videos needs to be in "public" folder
const VIDEOS = [
  'ok.webm',
  'https://www.youtube.com/watch?v=C0DPdy98e4c',
  'https://www.youtube.com/watch?v=Mxesac55Puo'
]

class Player extends Component {
	state = {
    playIndex: 0
  }
  nextVideo = () => {
  	this.setState({ playIndex: this.state.playIndex + 1 })
  }
  render () {
    return (
    	<div>
        <ReactPlayer
          url={VIDEOS[this.state.playIndex]}
          playing
          controls
          onEnded={this.nextVideo}
        />
        <button onClick={this.nextVideo}>Next</button>
      </div>
    )
  }
}

export default Player;
