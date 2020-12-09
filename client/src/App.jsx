import React from 'react';
import { Route } from 'react-router-dom';
import Nav from './home/nav';

import './fonts.css';
import './App.sass';

import Home from './home/home';
import CreateEntity from './home/create';
import Main from './main/main';

class App extends React.Component {
  render() {
    return (
      <div>
          <Nav />
          <Route path="/" exact component={Home} />
          <Route path="/create" component={CreateEntity} />

          <Route path="/project/:id" component={Main} />
      </div>
    )
  }
}

export default App;
