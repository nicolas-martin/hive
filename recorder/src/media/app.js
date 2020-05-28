import React, { Component } from 'react';
import ReactDOM from 'react-dom';
import AudioExample from './AudioExample';
import VideoExample from './VideoExample';
import * as api from '../functions/api';

class App extends Component {
  constructor(props) {
    super(props);
    console.log('=====')
    console.log(this.props.id)
    this.state = {
      // file: null
    };
  }
render() {
	return (
		<div>
			<h1>React Multimedia Capture Test</h1>
			<hr />
			<AudioExample id={this.props.id}/> 
			<hr />
		</div>
	);
}
};

// ReactDOM.render(
// 	<App />,
// 	document.getElementById('root')
// );
//
export default App
