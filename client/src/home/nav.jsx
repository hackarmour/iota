import React from 'react';
import { Logo, NavWrapper } from '../components/nav';


class Nav extends React.Component {
    render() {
        return (
            <NavWrapper>
                <Logo>IOTA</Logo>
            </NavWrapper>
        )
    }
}

export default Nav;