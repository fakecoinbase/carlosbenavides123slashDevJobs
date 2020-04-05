import React, { Component } from "react";
import "./Card.scss";

function Card() {
  return (
    <div class="cards">
      <div class="card">
        <span>
          <img
            src={
              "https://res.cloudinary.com/dhxwdb3jl/image/upload/v1586121171/unnamed_wqeqel.png"
            }
            className="icon"
          />
          Los Angeles, CA
        </span>
        <br />
        Honey
        <br />
        Software Engineer
        <br />
        <button>Apply</button>
      </div>

      <div class="card">
        <span>
          <img
            src={
              "https://res.cloudinary.com/dhxwdb3jl/image/upload/v1586121171/unnamed_wqeqel.png"
            }
            className="icon"
          />
          Los Angeles, CA
        </span>
        <br />
        Honey
        <br />
        Software Engineer
        <br />
        <button>Apply</button>
      </div>

      <div class="card">
        <span>
          <img
            src={
              "https://res.cloudinary.com/dhxwdb3jl/image/upload/v1586121171/unnamed_wqeqel.png"
            }
            className="icon"
          />
          Los Angeles, CA
        </span>
        <br />
        Honey
        <br />
        Software Engineer
        <br />
        <button>Apply</button>
      </div>

      <div class="card">
        <span>
          <img
            src={
              "https://res.cloudinary.com/dhxwdb3jl/image/upload/v1586121171/unnamed_wqeqel.png"
            }
            className="icon"
          />
          Los Angeles, CA
        </span>
        <br />
        Honey
        <br />
        Software Engineer
        <br />
        <button>Link to Apply</button>
      </div>
    </div>
  );
}

export default Card;
