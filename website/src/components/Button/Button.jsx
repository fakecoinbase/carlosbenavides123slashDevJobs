import React from 'react';
import styled from 'styled-components';

const StyledButton = styled.div`
    border-radius: .2rem;
    font-size: 1.2rem;
    width: fit-content;

    padding-left: 1rem;
    padding-right: 1rem;

    ${({btype}) => btype === "joblink" && `
        outline: .1rem solid white;
        color: #CEF074;
        float:right;
        &:hover{
            outline: .1rem solid #CEF074;
            background-color: #CEF074;
            cursor:pointer;
            color: #484CF6;
        }
    `}
`
function Button({btype, onClick, onMouseEnter, text}) {
    return (
        <StyledButton btype={btype} onClick={onClick} onMouseEnter={onMouseEnter}>
            <span>{text}</span>
        </StyledButton>
    )
}

export default Button;