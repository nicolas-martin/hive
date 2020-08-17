import React, { Component } from 'react'
import ReactPlayer from 'react-player'

// NOTE: Local videos needs to be in "public" folder
const VIDEOS = [
  'https://www.youtube.com/watch?v=C0DPdy98e4c',
]

class Player extends Component {
  constructor(props) {
    super(props);
    console.log('+++++')
    console.log(this.props.id)
    if (!this.props.id){
      VIDEOS.push(this.props.id)
    }
  }

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
