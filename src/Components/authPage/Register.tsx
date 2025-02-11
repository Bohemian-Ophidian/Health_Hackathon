import axios from "axios";
import { useState } from "react";

const Register = () => {
  const [mobileNumber, setMobileNumber] = useState("");
  const [password, setPassword] = useState("");
  const [name, setName] = useState("");
  const [age, setAge] = useState("");
  const [sex, setSex] = useState("");
  const [medicalHistory, setMedicalHistory] = useState("");
  const [height, setHeight] = useState("");
  const [weight, setWeight] = useState("");

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    axios
      .post("http://localhost:8000/patient", {
        name,
        mobileNumber,
        age,
        sex,
        medicalHistory,
        height,
        weight,
        password,
      })
      .then((res) => console.log(res))
      .catch((err) => console.log(err));
  };

  return (
    <div className="flex justify-center items-center mt-12 min-h-screen bg-white p-6">
      <div className="bg-blue-600 p-10 rounded-2xl shadow-lg text-center w-96">
        <h2 className="text-white text-2xl font-semibold">Health Mentá</h2>
        <form onSubmit={handleSubmit} className="mt-6">
          {/* Name */}
          <div className="mb-4 text-left">
            <label className="block text-white font-bold">Name</label>
            <input
              type="text"
              name="name"
              required
              placeholder="Enter Full Name"
              className="w-full mt-2 p-2 border border-white bg-white rounded-md text-black"
              onChange={(e) => setName(e.target.value)}
              value={name}
            />
          </div>

          {/* Mobile Number */}
          <div className="mb-4 text-left">
            <label className="block text-white font-bold">Mobile Number</label>
            <input
              type="tel"
              name="mobileNumber"
              placeholder="Enter Mobile Number"
              required
              className="w-full mt-2 p-2 border border-white bg-white rounded-md text-black"
              onChange={(e) => setMobileNumber(e.target.value)}
              value={mobileNumber}
            />
          </div>

          {/* Password */}
          <div className="mb-4 text-left">
            <label className="block text-white font-bold">Password</label>
            <input
              type="password"
              name="password"
              placeholder="Enter Strong Password"
              required
              className="w-full mt-2 p-2 border border-white bg-white rounded-md text-black"
              onChange={(e) => setPassword(e.target.value)}
              value={password}
            />
          </div>

          {/* Age & Gender */}
          <div className="mb-4 flex justify-between gap-8 text-left">
            <div className="w-1/2">
              <label className="block text-white font-bold">Age</label>
              <input
                type="number"
                name="age"
                placeholder="Enter Age"
                min="1"
                required
                className="w-full mt-2 p-2 border border-white bg-white rounded-md text-black"
                onChange={(e) => setAge(e.target.value)}
                value={age}
              />
            </div>
            <div className="w-1/2">
              <label className="block text-white font-bold">Sex</label>
              <select
                name="sex"
                required
                className="w-full mt-2 p-2 border border-white bg-white rounded-md text-black"
                onChange={(e) => setSex(e.target.value)}
                value={sex}
              >
                <option value="">Select</option>
                <option value="male">Male</option>
                <option value="female">Female</option>
                <option value="other">Other</option>
              </select>
            </div>
          </div>

          {/* Height & Weight */}
          <div className="mb-4 flex justify-between gap-8 text-left">
            <div className="w-1/2">
              <label className="block text-white font-bold">Height (cm)</label>
              <input
                type="number"
                name="height"
                placeholder="Enter Height"
                min="0"
                required
                className="w-full mt-2 p-2 border border-white bg-white rounded-md text-black"
                onChange={(e) => setHeight(e.target.value)}
                value={height}
              />
            </div>
            <div className="w-1/2">
              <label className="block text-white font-bold">Weight (kg)</label>
              <input
                type="number"
                name="weight"
                placeholder="Enter Weight"
                min="0"
                required
                className="w-full mt-2 p-2 border border-white bg-white rounded-md text-black"
                onChange={(e) => setWeight(e.target.value)}
                value={weight}
              />
            </div>
          </div>

          {/* Medical History */}
          <div className="mb-4 text-left">
            <label className="block text-white font-bold">Medical History</label>
            <textarea
              name="medicalHistory"
              placeholder="Like allergies, chronic diseases, etc."
              rows={2}
              className="w-full mt-2 p-2 border border-white bg-white rounded-md text-black"
              onChange={(e) => setMedicalHistory(e.target.value)}
              value={medicalHistory}
            ></textarea>
          </div>

          {/* Submit Button */}
          <button
            type="submit"
            className="w-full bg-white text-black py-2 rounded-md font-semibold hover:bg-black hover:text-white transition duration-300"
          >
            Create Account
          </button>
        </form>
      </div>
    </div>
  );
};

export default Register;
