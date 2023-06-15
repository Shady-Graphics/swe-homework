import React from 'react';
import { BrowserRouter as Router, Routes, Route, NavLink } from 'react-router-dom';
import Gallery from './components/Gallery';
import Favorites from './components/Favorites';

function App() {
  return (
    <Router>
      <div className="container-fluid">
        <nav className="navbar navbar-expand-lg navbar-light bg-light">
          <NavLink className="navbar-brand" to="/">
            Happiness
          </NavLink>
          <button
            className="navbar-toggler"
            type="button"
            data-toggle="collapse"
            data-target="#navbarNav"
            aria-controls="navbarNav"
            aria-expanded="false"
            aria-label="Toggle navigation"
          >
            <span className="navbar-toggler-icon"></span>
          </button>
          <div className="collapse navbar-collapse" id="navbarNav">
            <ul className="navbar-nav">
              <li className="nav-item">
                <NavLink className="nav-link" activeClassName="active" to="/" end>
                  Gallery
                </NavLink>
              </li>
              <li className="nav-item">
                <NavLink className="nav-link" activeClassName="active" to="/favorites">
                  Favorites
                </NavLink>
              </li>
            </ul>
          </div>
        </nav>

        <Routes>
          <Route path="/" element={<Gallery />} />
          <Route path="/favorites" element={<Favorites />} />
        </Routes>
      </div>
    </Router>
  );
}

export default App;
