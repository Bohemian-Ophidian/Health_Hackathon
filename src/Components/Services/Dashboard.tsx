import React, { useState, useEffect } from "react";

const Dashboard: React.FC = () => {
  const [medicines, setMedicines] = useState<any[]>([]);
  const [newMedicineName, setNewMedicineName] = useState("");
  const [newMedicineDescription, setNewMedicineDescription] = useState("");
  const [newMedicineDosage, setNewMedicineDosage] = useState("");
  const [newMedicineTime, setNewMedicineTime] = useState("");

  // Retrieve token from localStorage (assumes you save it upon login)
  const token = localStorage.getItem("token");

  // Fetch medicines from the backend
  useEffect(() => {
    const fetchMedicines = async () => {
      try {
        const response = await fetch("http://localhost:3001/medicines", {
          headers: { Authorization: token ? `Bearer ${token}` : "" },
        });

        if (!response.ok) {
          const errorData = await response.json();
          console.error("Error fetching medicines:", errorData.message);
          return;
        }

        const data = await response.json();
        if (Array.isArray(data)) {
          setMedicines(data);
        } else {
          console.error("Unexpected data format", data);
          setMedicines([]);
        }
      } catch (error) {
        console.error("Error fetching medicines data:", error);
      }
    };

    if (token) {
      fetchMedicines();
    }
  }, [token]);

  // Function to add a new medicine
  const addMedicine = async () => {
    if (!token) {
      console.error("No token found. Please log in again.");
      return;
    }
    
    try {
      console.log("Adding medicine with token:", token);
      const response = await fetch("http://localhost:3001/add-medicine", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify({
          name: newMedicineName,
          description: newMedicineDescription,
          dosage: newMedicineDosage,
          time: newMedicineTime,
        }),
      });

      if (!response.ok) {
        const errorData = await response.json();
        console.error("Error adding medicine:", errorData.message);
        return;
      }

      const data = await response.json();
      if (data.medicines && Array.isArray(data.medicines)) {
        setMedicines(data.medicines);
      } else {
        console.error("Unexpected data format when adding medicine", data);
        setMedicines([]);
      }

      // Reset the form
      setNewMedicineName("");
      setNewMedicineDescription("");
      setNewMedicineDosage("");
      setNewMedicineTime("");
    } catch (error) {
      console.error("Error adding medicine:", error);
    }
  };

  return (
    <div className="flex flex-wrap gap-4 p-4">
      {/* Medicines Section */}
      <div className="bg-white shadow-md rounded-2xl p-6 flex-1 min-w-[300px]">
        <h2 className="text-xl font-bold">Medicines</h2>
        {/* Add Medicine Form */}
        <div className="flex flex-col gap-2 my-4">
          <input
            type="text"
            placeholder="Medicine Name"
            value={newMedicineName}
            onChange={(e) => setNewMedicineName(e.target.value)}
            className="border p-2 rounded"
          />
          <input
            type="text"
            placeholder="Description"
            value={newMedicineDescription}
            onChange={(e) => setNewMedicineDescription(e.target.value)}
            className="border p-2 rounded"
          />
          <input
            type="text"
            placeholder="Dosage (optional)"
            value={newMedicineDosage}
            onChange={(e) => setNewMedicineDosage(e.target.value)}
            className="border p-2 rounded"
          />
          <input
            type="text"
            placeholder="Time (optional)"
            value={newMedicineTime}
            onChange={(e) => setNewMedicineTime(e.target.value)}
            className="border p-2 rounded"
          />
          <button onClick={addMedicine} className="bg-blue-500 text-white p-2 rounded">
            Add Medicine
          </button>
        </div>

        {/* Display Medicines */}
        {Array.isArray(medicines) && medicines.length > 0 ? (
          <ul className="text-gray-600">
            {medicines.map((medicine, index) => (
              <li key={index} className="mb-2">
                <strong>{medicine.name}</strong> - {medicine.description}
                {medicine.dosage && ` - Dosage: ${medicine.dosage}`}
                {medicine.time && ` - Time: ${medicine.time}`}
              </li>
            ))}
          </ul>
        ) : (
          <p>No medicines...</p>
        )}
      </div>

      {/* Graph Section */}
      <div className="bg-white shadow-md rounded-2xl p-6 flex-1 min-w-[300px]">
        <h2 className="text-xl font-bold">Graph</h2>
        <svg viewBox="0 0 400 220" width="100%" height="auto">
          {/* Background */}
          <rect width="100%" height="220" fill="white" />
          {/* Y-Axis */}
          <line x1="50" y1="0" x2="50" y2="200" stroke="black" strokeWidth="2" />
          {/* X-Axis */}
          <line x1="50" y1="200" x2="400" y2="200" stroke="black" strokeWidth="2" />
          {/* Data Polyline */}
          <polyline
            fill="none"
            stroke="rgba(75, 192, 192, 1)"
            strokeWidth="3"
            points={[12, 19, 3, 5, 2, 3]
              .map((value, index) => `${index * 50 + 50},${200 - value * 10}`)
              .join(" ")}
          />
          {/* X-Axis Labels */}
          <g className="axis x-axis">
            {["Jan", "Feb", "Mar", "Apr", "May", "Jun"].map((label, index) => (
              <text key={index} x={index * 50 + 50} y="215" textAnchor="middle" fontSize="12">
                {label}
              </text>
            ))}
          </g>
          {/* Y-Axis Labels */}
          <g className="axis y-axis">
            {[0, 5, 10, 15, 20].map((value, index) => (
              <text
                key={index}
                x="40"
                y={200 - value * 10}
                textAnchor="end"
                fontSize="12"
                dominantBaseline="middle"
              >
                {value}
              </text>
            ))}
          </g>
        </svg>
      </div>

      {/* Streak Section */}
      <div className="bg-white shadow-md rounded-2xl p-6 flex-1 min-w-[300px]">
        <h2 className="text-xl font-bold">Streak</h2>
        <p className="text-gray-600">Track your streak...</p>
      </div>

      {/* Tasks Section */}
      <div className="bg-white shadow-md rounded-2xl p-6 flex-1 min-w-[300px]">
        <h2 className="text-xl font-bold">Tasks</h2>
        <p className="text-gray-600">Your pending tasks...</p>
      </div>

      {/* Notes Section */}
      <div className="bg-white shadow-md rounded-2xl p-6 flex-1 min-w-[300px]">
        <h2 className="text-xl font-bold">Notes</h2>
        <p className="text-gray-600">Your notes...</p>
      </div>

      {/* Profile Section */}
      <div className="bg-white shadow-md rounded-2xl p-6 flex-1 min-w-[300px]">
        <h2 className="text-xl font-bold">Profile</h2>
        <p className="text-gray-600">Your profile information...</p>
      </div>
    </div>
  );
};

export default Dashboard;
