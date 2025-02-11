import { useState } from "react";

const Login = () => {
  const [useEmail, setUseEmail] = useState(true);

  const handleToggle = () => {
    setUseEmail(!useEmail);
  };

  return (
    <div className="flex justify-center items-center h-screen bg-white">
      <div className="bg-blue-600 p-10 rounded-2xl shadow-lg text-center w-96">
        <h2 className="text-white text-2xl font-semibold">Health Mentá</h2>
        <form action="#" method="POST" className="mt-6">
          <div className="mb-4 text-left">
            <label className="block text-white font-bold">{useEmail ? "Email" : "Mobile Number"}</label>
            {useEmail ? (
              <input
                type="email"
                id="email"
                name="email"
                placeholder="Enter Email"
                required
                className="w-full mt-2 p-2 border border-white bg-white rounded-md text-black"
              />
            ) : (
              <input
                type="tel"
                id="mobile"
                placeholder="Enter Mobile Number"
                name="mobile"
                required
                className="w-full mt-2 p-2 border border-white bg-white rounded-md text-black"
              />
            )}
          </div>
          <div className="mb-4 text-left">
            <label htmlFor="password" className="block text-white font-bold">Password</label>
            <input
              type="password"
              id="password"
              placeholder="Enter Strong Password"
              name="password"
              required
              className="w-full mt-2 p-2 border border-white bg-white rounded-md text-black"
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
        <div className="mt-4">
          <label className="text-white">
            <input
              type="checkbox"
              checked={!useEmail}
              onChange={handleToggle}
              className="mr-2"
            />
            Use {useEmail ? "Mobile Number" : "Email"} instead?
          </label>
        </div>
      </div>
    </div>
  );
};

export default Login;
