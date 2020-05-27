import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './media/app';

class Game extends React.Component {
  render() {
    return (
      <div className="game">
        <App />
      </div>
    );
  }
}

// ========================================

ReactDOM.render(
  <Game />,
  document.getElementById('root')
);
