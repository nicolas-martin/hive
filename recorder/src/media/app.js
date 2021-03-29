import React, { Component } from 'react';
// import ReactDOM from 'react-dom';
import AudioExample from './AudioExample';
// import VideoExample from './VideoExample';

class App extends Component {
  constructor(props) {
    super(props);
    console.log('=====')
    console.log(this.props.userid)
    console.log(this.props.updateid)
    this.state = {
      // file: null
    };
  }
render() {
	return (
		<div>
			<h1>React Multimedia Capture Test</h1>
			<hr />
			<AudioExample userid={this.props.userid} updateid={this.props.updateid}/> 
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
