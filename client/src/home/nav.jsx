import React from 'react';
import { Logo, NavWrapper } from '../components/nav';
import { Link } from 'react-router-dom';


class Nav extends React.Component {
    render() {
        return (
            <NavWrapper>
                <Link to="/" style={{textDecoration: 'none'}}>
                    <Logo>IOTA</Logo>
                </Link>
            </NavWrapper>
        )
    }
}

export default Nav;