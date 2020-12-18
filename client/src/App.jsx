import React from 'react';
import { Route } from 'react-router-dom';
import Nav from './home/nav';

import './fonts.css';
import './App.sass';

import Home from './home/home';
import CreateProject from './home/create';
import CreateEntity from './main/createEntity';
import Main from './main/main';

class App extends React.Component {
  render() {
    return (
      <div>
          <Nav />
          <Route path="/" exact component={Home} />
          <Route path="/create" component={CreateProject} />

          <Route path="/project/:id" component={Main} exact />
          <Route path="/project/:id/new" component={CreateEntity} />
      </div>
    )
  }
}

export default App;
