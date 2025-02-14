// import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import HomePage from './Components/Main/HomePage';
import ServicePage from './Components/Services/ServicePage';
import Landing from './Components/AboutPage/Landing';
import DoctorsPage from './Components/Services/DoctorsPage';
import Header from './Components/Header/Header';
import Login from './Components/authPage/Login';
import Register from './Components/authPage/Register';


function App() {
  return (
    <Router>
      <div className="text-[#1d4d85] app min-w-[280px] min-h-screen bg-background">
         <Header/>
        <Routes>
          <Route path="/Health-Mentá/login" element={<Login />} />
          <Route path="/Health-Mentá/register" element={<Register />} />
          <Route path="/Health-Mentá" element={<HomePage />} />
          <Route path="/Health-Mentá/Services" element={<ServicePage/>} />
          <Route path="/Health-Mentá/About" element={<Landing />} />
          <Route path="/Health-Mentá/Doctors" element={<DoctorsPage />} />
        </Routes>
      </div>
    </Router>
  );
}

export default App;
