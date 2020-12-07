import React from 'react';

class Main extends React.Component{
   constructor(props) {
       super(props);
       this.state = {}
   }

   componentDidMount() {
       console.log(this.props.match.params.id)
   }

   render() {
       return (
           <div>
               
           </div>
       )
   }
}

export default Main;