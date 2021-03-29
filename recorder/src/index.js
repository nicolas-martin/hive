import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './media/app';
import Player from './media/player/player';
import { BrowserRouter, Route } from 'react-router-dom'

class Game extends React.Component {
  render() {
    return (
      <div className="game">
        <Route path="/" exact component={App}/>
        <Route path="/play" exact component={Player}/>
        <Route path="/record/:userid/:updateid" 
          render={(props) => <App userid={props.match.params.userid} updateid={props.match.params.updateid} {...props} />}
        />
      </div>
    );
  }
}

ReactDOM.render(
	<BrowserRouter>
		<Game />
	</BrowserRouter>,
  document.getElementById('root')
);
