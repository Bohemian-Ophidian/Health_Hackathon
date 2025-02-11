import axios from "axios";
import { useState } from "react";
import React from "react";
import { useNavigate } from "react-router-dom";

const Login = () => {
  const [mobile, setMobileNumber] = useState("");
  const [password, setPassword] = useState("");
  const navigate = useNavigate();

  const handleSubmit = (e:React.FormEvent) => {
    e.preventDefault();
    axios
      .post("http://localhost:8000/login", { mobile, password }) 
      .then((res) => {
        console.log(res);
        if (res.data === "success") {
          navigate("/Hospital-Website/home");
        }
      })
      .catch((error) => {
        console.error("Error during login:", error);
      });
  };

  return (
    <div className="flex justify-center items-center h-screen bg-white">
      <div className="bg-blue-600 p-10 rounded-2xl shadow-lg text-center w-96">
        <h2 className="text-white text-2xl font-semibold">Health Mentá</h2>
        <form onSubmit={handleSubmit} className="mt-6"> {/* Added onSubmit */}
          <div className="mb-4 text-left">
            <label className="block text-white font-bold">Mobile Number</label>
            <input
              type="tel"
              id="mobile"
              placeholder="Enter Mobile Number"
              name="mobile"
              required
              onChange={(e) => setMobileNumber(e.target.value)}
              className="w-full mt-2 p-2 border border-white bg-white rounded-md text-black"
            />
          </div>
          <div className="mb-4 text-left">
            <label htmlFor="password" className="block text-white font-bold">Password</label>
            <input
              type="password"
              id="password"
              placeholder="Enter Last Password"
              name="password"
              required
              className="w-full mt-2 p-2 border border-white bg-white rounded-md text-black"
              onChange={(e) => setPassword(e.target.value)}
            />
          </div>
          <button
            type="submit"
            className="w-1/2 bg-white text-black py-2 rounded-md font-semibold hover:bg-black hover:text-white transition duration-300"
          >
            Login
          </button>
          <div className="mt-4 text-sm">
            <span className="text-white">
              Don't have an account? <a href="/Hospital-Website/register" className="underline">Sign up</a>
            </span>
          </div>
          <div className="mt-2 text-sm">
            <a href="#" className="text-white underline">Forgot Password?</a>
          </div>
        </form>
      </div>
    </div>
  );
};

export default Login;
