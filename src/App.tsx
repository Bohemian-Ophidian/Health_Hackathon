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
          <Route path="/Hospital-Website/login" element={<Login />} />
          <Route path="/Hospital-Website/register" element={<Register />} />
          <Route path="/Hospital-Website" element={<HomePage />} />
          <Route path="/Hospital-Website/Services" element={<ServicePage/>} />
          <Route path="/Hospital-Website/About" element={<Landing />} />
          <Route path="/Hospital-Website/Doctors" element={<DoctorsPage />} />
        </Routes>
      </div>
    </Router>
  );
}

export default App;
