import React from 'react';
import styled from 'styled-components';

const StyledButton = styled.div`
    border-radius: .2rem;
    font-size: 1.2rem;
    width: fit-content;

    padding-left: 1rem;
    padding-right: 1rem;

    ${({btype}) => btype === "joblink" && `
        border: 2px solid #1cce7cab;
        color: #7d846c;
        float:right;
        &:hover{
            border: .1rem solid #7d846c;
            background-color: #7d846c;
            color: white;
        }
    `}

    ${({btype}) => btype === "dropdown" && `
        border: 1px solid gray;
        cursor: pointer;
        padding-left: 0rem;
        padding-right: 0rem;
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