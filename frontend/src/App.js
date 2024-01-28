import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Footer from './Components/Footer';
import Header from './Components/Header';
import StartPage from './Pages/StartPage';
import Furnitures from './Pages/Furnitures';

export default class App extends React.Component {
  render() {
    return (
      <Router>
        <div className='wrapper'>
          <Header />

          <main>
            <Routes>
              <Route path='/' element={<StartPage />} />
              <Route path='/furnitures/:tag' element={<Furnitures />} />
              <Route path='*' element={<StartPage />} />
            </Routes>
          </main>

          <Footer />
        </div>
      </Router>
    );
  }
}