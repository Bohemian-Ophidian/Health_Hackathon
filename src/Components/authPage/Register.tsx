import { useState } from "react";

const Register = () => {
  const [useEmail, setUseEmail] = useState(true);

  const handleToggle = () => {
    setUseEmail(!useEmail);
  };

  return (
    <div className="flex justify-center items-center mt-12 min-h-screen bg-white p-6">
      <div className="bg-blue-600 p-10 rounded-2xl shadow-lg text-center w-96">
        <h2 className="text-white text-2xl font-semibold">Health Mentá</h2>
        <form action="#" method="POST" className="mt-6">
          <div className="mb-4 text-left">
            <label htmlFor="full-name" className="block text-white font-bold">Name</label>
            <input
              type="text"
              id="full-name"
              name="full-name"
              required
              placeholder="Enter Full Name"
              className="w-full mt-2 p-2 border border-white bg-white rounded-md text-black"
            />
          </div>
          
          {/* Conditionally render either Email or Mobile Number */}
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
            <label htmlFor="password" className="block text-white font-bold">Enter a strong password</label>
            <input
              type="password"
              placeholder="Enter Strong Password"
              id="password"
              name="password"
              required
              className="w-full mt-2 p-2 border border-white bg-white rounded-md text-black"
            />
          </div>
          <div className="mb-4 flex justify-between gap-8 text-left">
            <div className="w-1/2">
              <label htmlFor="age" className="block text-white font-bold">Age</label>
              <input
                type="number"
                id="age"
                name="age"
                placeholder="Enter Age"
                min="1"
                required
                className="w-full mt-2 p-2 border border-white bg-white rounded-md text-black"
              />
            </div>
            <div className="w-1/2">
              <label htmlFor="sex" className="block text-white font-bold">Sex</label>
              <select
                id="sex"
                name="sex"
                placeholder="Enter Sex"
                required
                className="w-full mt-2 p-2 border border-white bg-white rounded-md text-black"
              >
                <option value="">Select</option>
                <option value="male">Male</option>
                <option value="female">Female</option>
                <option value="other">Other</option>
              </select>
            </div>
          </div>
          <div className="mb-4 text-left">
            <label htmlFor="medical-history" className="block text-white font-bold">Medical History</label>
            <textarea
              id="medical-history"
              name="medical-history"
              placeholder="like allergies, chronic diseases, etc."
              rows={2}
              className="w-full mt-2 p-2 border border-white bg-white rounded-md text-black"
            ></textarea>
          </div>
          <button
            type="submit"
            className="w-full bg-white text-black py-2 rounded-md font-semibold hover:bg-black hover:text-white transition duration-300"
          >
            Create Account
          </button>
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

export default Register;
