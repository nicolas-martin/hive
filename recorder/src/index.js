import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './media/app';
import { BrowserRouter, Route } from 'react-router-dom'

class Game extends React.Component {
  render() {
    return (
      <div className="game">
        <Route path="/" exact component={App}/>
        <Route path="/record/:id" 
          render={(props) => <App id={props.match.params.id} {...props} />}
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
