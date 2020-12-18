import React, {useEffect} from 'react';

function Main(props){

    useEffect(() => {
        console.log(props.match.params.id)
    })

    return(
        <div>Hey</div>
    )
}

export default Main;