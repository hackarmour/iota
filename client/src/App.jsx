import React from 'react';
import { Route } from 'react-router-dom';

import Home from './home/home';
import CreateEntity from './home/create';
import Main from './main/main';

class App extends React.Component {
  render() {
    return (
      <div>
          <Route path="/" exact component={Home} />
          <Route path="/create" component={CreateEntity} />

          <Route path="/entity/:id" component={Main} />
      </div>
    )
  }
}

export default App;
