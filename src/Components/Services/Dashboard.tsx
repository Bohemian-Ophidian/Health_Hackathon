import React, { useState, useEffect } from "react";
import axios from "axios";

const Dashboard: React.FC = () => {
  const [medicines, setMedicines] = useState<any[]>([]);
  const [appointments, setAppointments] = useState<any[]>([]);
  const [profile, setProfile] = useState<any | null>(null);
  const [newMedicine, setNewMedicine] = useState({
    name: "",
    description: "",
    dosage: "",
    time: "",
  });
  const [isEditing, setIsEditing] = useState<boolean>(false);
  const [editableProfile, setEditableProfile] = useState<any>({
    name: "",
    height: "",
    weight: "",
    medical_history: [],
  });
  const [tip, setTip] = useState<string>('');

  const token = localStorage.getItem("token");

  // List of small tips
  const tips = [
    'Eat one serving of fruit!',
    'Walk for 15 minutes!',
    'Do 10 squats!',
    'Drink 8 glasses of water!',
    'Spend 5 minutes practicing deep breathing!',
    'Prepare a healthy meal!',
    'Get 7-8 hours of sleep!',
    'Limit screen time before bed!',
    'Take your vitamins!',
    'Stand up and stretch every hour!',
    'Eat a balanced breakfast!',
    'Reflect on something positive!',
    'Call a friend or family member!',
    'Try a new healthy recipe!',
    'Pack a healthy snack for work/school!',
    'Swap sugary drinks for water or unsweetened tea!',
    'Take the stairs instead of the elevator!',
    'Listen to a calming podcast or music!',
    'Set a bedtime reminder!',
    'Plan your meals for the week!',
    'Write down 3 things you are grateful for!',
    'Do some light stretching while watching TV!',
    'Have a tech-free hour before bed!',
    'Make a smoothie with fruits and vegetables!',
    'Learn a new healthy recipe!',
    'Try a new type of exercise (e.g., yoga, pilates)!',
    'Go for a bike ride!',
    'Have a conversation with someone new!',
    'Learn about nutrition and healthy eating!',
    'Set a small, achievable fitness goal for the week!',
    'Track your food intake for one day!',
    'Try a new type of cuisine!',
    'Cook a meal with fresh, whole ingredients!',
  ];

  // Fetch data on login or token change
  useEffect(() => {
    const fetchData = async () => {
      try {
        if (token) {
          // Fetch Medicines
          const medicinesResponse = await axios.get("http://localhost:3001/medicines", {
            headers: { Authorization: `Bearer ${token}` },
          });
          setMedicines(medicinesResponse.data);

          // Fetch Appointments
          const appointmentsResponse = await axios.get("http://localhost:3001/api/appointments", {
            headers: { Authorization: `Bearer ${token}` },
          });
          setAppointments(appointmentsResponse.data);

          // Fetch Profile
          const profileResponse = await axios.get("http://localhost:3001/api/getPatientId", {
            headers: { Authorization: `Bearer ${token}` },
          });
          setProfile(profileResponse.data);
        }
      } catch (error) {
        console.error("Error fetching data:", error);
      }
    };

    fetchData();
  }, [token]);

  // Random tip generator on component mount
  useEffect(() => {
    const randomTip = tips[Math.floor(Math.random() * tips.length)];
    setTip(randomTip);
  }, []);

  // Function to add a new medicine
  const addMedicine = async () => {
    if (!token) {
      alert("Please log in to add medicines.");
      return;
    }

    try {
      const response = await axios.post(
        "http://localhost:3001/add-medicine",
        newMedicine,
        {
          headers: { Authorization: `Bearer ${token}` },
        }
      );

      setMedicines(response.data.medicines);
      setNewMedicine({ name: "", description: "", dosage: "", time: "" });
    } catch (error) {
      console.error("Error adding medicine:", error);
    }
  };

  // Function to remove a medicine
  const removeMedicine = async (medicineId: string) => {
    if (!token) {
      alert("Please log in to remove medicines.");
      return;
    }

    try {
      await axios.delete(`http://localhost:3001/remove-medicine/${medicineId}`, {
        headers: { Authorization: `Bearer ${token}` },
      });

      setMedicines(medicines.filter((med) => med._id !== medicineId));
    } catch (error) {
      console.error("Error removing medicine:", error);
    }
  };

  const timeSlots = [
    "08:00 AM", "10:00 AM", "12:00 PM", "02:00 PM", "04:00 PM", "06:00 PM", "08:00 PM", "10:00 PM"
  ];
  const timeOptions = timeSlots.map((time, index) => (
    <option key={index} value={time}>{time}</option>
  ));

  const toggleEditMode = () => {
    setIsEditing(!isEditing);
  };

  // Handle input changes
  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>, field: string) => {
    setEditableProfile({
      ...editableProfile,
      [field]: e.target.value,
    });
  };

  // Handle changes in medical history
// Handle change for medical history
const handleMedicalHistoryChange = (e: React.ChangeEvent<HTMLInputElement>) => {
  setEditableProfile({
    ...editableProfile,
    medical_history: e.target.value,  // Directly update the medical history field
  });
};


  // Submit the edited profile data
  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      const response = await axios.put(
        "http://localhost:3001/api/updateProfile", // Ensure the URL is correct
        editableProfile, // Send the editableProfile data
        {
          headers: { Authorization: `Bearer ${token}` }, // Ensure the token is correctly passed
        }
      );
      setProfile(response.data); // Update profile with new data
      setIsEditing(false); // Exit edit mode
    } catch (error) {
      console.error("Error updating profile:", error); // Log any error to the console
    }
  };

  return (
    <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 max-w-7xl mx-auto p-6">
      {/* Medicines Section */}
      <div className="bg-white shadow-lg rounded-lg p-6">
        <h2 className="text-xl font-semibold">Medicines</h2>
        <div className="space-y-2 mt-4">
          <input
            type="text"
            placeholder="Medicine Name"
            value={newMedicine.name}
            onChange={(e) => setNewMedicine({ ...newMedicine, name: e.target.value })}
            className="border p-2 w-full rounded"
          />
          <input
            type="text"
            placeholder="Description"
            value={newMedicine.description}
            onChange={(e) => setNewMedicine({ ...newMedicine, description: e.target.value })}
            className="border p-2 w-full rounded"
          />
          <input
            type="text"
            placeholder="Dosage (e.g., 100mg)"
            value={newMedicine.dosage}
            onChange={(e) => setNewMedicine({ ...newMedicine, dosage: e.target.value })}
            className="border p-2 w-full rounded"
          />
          <select
            value={newMedicine.time}
            onChange={(e) => setNewMedicine({ ...newMedicine, time: e.target.value })}
            className="border p-2 w-full rounded"
          >
            <option value="">Select Time</option>
            {timeOptions}
          </select>

          <button onClick={addMedicine} className="bg-blue-500 text-white px-4 py-2 rounded-lg w-full">
            Add Medicine
          </button>
        </div>

        <h3 className="mt-6 text-lg font-semibold">Your Medicines</h3>
        <ul className="mt-2 text-gray-600">
          {medicines.length > 0 ? (
            medicines.map((med) => (
              <li key={med._id} className="border p-2 rounded mt-2">
                <strong>{med.name}</strong> - {med.description}
                {med.dosage && ` (Dosage: ${med.dosage})`}
                {med.time && ` (Time: ${med.time})`}
                <button
                  onClick={() => removeMedicine(med._id)}
                  className="bg-red-500 text-white px-2 py-1 rounded-lg ml-2"
                >
                  Remove
                </button>
              </li>
            ))
          ) : (
            <p>No medicines added.</p>
          )}
        </ul>
      </div>

      {/* Tip of the Day Section */}
      <div className="bg-white shadow-lg rounded-lg p-6">
        <h2 className="text-xl font-semibold">Tip of the Day</h2>
        <p className="text-gray-600 mt-2 text-xl">{tip}</p>
      </div>

      {/* Schedule Section */}
      <div className="bg-white shadow-lg rounded-lg p-6">
        <h2 className="text-xl font-semibold">Diet/Medicine Reminder</h2>
        <ul className="mt-2 text-gray-600">
          {[
            { name: "Breakfast", time: "08:00 AM" },
            { name: "Lunch", time: "01:00 PM" },
            { name: "Dinner", time: "07:00 PM" },
            ...medicines.map((med) => ({ name: med.name, time: med.time }))
          ]
            .sort((a, b) =>
              new Date(`1970/01/01 ${a.time}`).getTime() - new Date(`1970/01/01 ${b.time}`).getTime()
            )
            .map((task, index) => (
              <li key={index} className="border p-2 rounded mt-2">
                <strong>{task.name}</strong> - {task.time}
              </li>
            ))}
        </ul>
      </div>

      {/* Appointment Section */}
      <div className="bg-white shadow-lg rounded-lg p-6">
        <h2 className="text-xl font-semibold">Appointments</h2>
        <ul>
          {appointments.map((appointment, index) => (
            <li key={index} className="mb-2 flex justify-between items-center">
              <span>
                {`${appointment.doctorName}: ${new Date(appointment.date).toLocaleDateString()} at ${appointment.time}`}
              </span>
            </li>
          ))}
        </ul>
      </div>

      {/* Profile Section */}
<div className="bg-white shadow-lg rounded-lg p-6">
  <h2 className="text-xl font-semibold">Profile</h2>
  {profile ? (
    <>
      {isEditing ? (
        <form onSubmit={handleSubmit}>
          <div className="space-y-2">
            <label className="block text-gray-600">Name:</label>
            <input
              type="text"
              value={editableProfile.name}
              onChange={(e) => handleInputChange(e, "name")}
              className="border p-2 w-full rounded"
            />
            <label className="block text-gray-600 mt-2">Height (in cm):</label>
            <input
              type="number"
              value={editableProfile.height}
              onChange={(e) => handleInputChange(e, "height")}
              className="border p-2 w-full rounded"
            />
            <label className="block text-gray-600 mt-2">Weight (in kg):</label>
            <input
              type="number"
              value={editableProfile.weight}
              onChange={(e) => handleInputChange(e, "weight")}
              className="border p-2 w-full rounded"
            />
            <label className="block text-gray-600 mt-2">Medical History:</label>
            <input
              type="text"
              value={editableProfile.medical_history}
              onChange={handleMedicalHistoryChange}  // Directly handle medical history change
              className="border p-2 w-full rounded"
              placeholder="Enter medical history"
            />
            <button
              type="submit"
              className="bg-blue-500 text-white px-4 py-2 rounded-lg mt-4 w-full"
            >
              Save Changes
            </button>
          </div>
        </form>
      ) : (
        <>
          <p className="text-gray-600 mt-2">
            <strong>Name:</strong> {profile.name}
          </p>
          <p className="text-gray-600 mt-2">
            <strong>Height:</strong> {profile.height} cm
          </p>
          <p className="text-gray-600 mt-2">
            <strong>Weight:</strong> {profile.weight} kg
          </p>
          <p className="text-gray-600 mt-2">
            <strong>Medical History:</strong>
          </p>
          <p className="text-gray-600">{profile.medical_history}</p>  {/* Display the medical history */}
          <button
            onClick={toggleEditMode}
            className="bg-blue-500 text-white px-4 py-2 rounded-lg mt-4"
          >
            Edit Profile
          </button>
        </>
      )}
    </>
  ) : (
    <p>Loading profile...</p>
  )}
</div>

      {/* Graph Section */}
      <div className="bg-white shadow-lg rounded-lg p-6">
        <h2 className="text-xl font-semibold">Graph</h2>
        <svg viewBox="0 0 400 220" width="100%" height="auto">
          <rect width="100%" height="220" fill="white" />
          <line x1="50" y1="0" x2="50" y2="200" stroke="black" strokeWidth="2" />
          <line x1="50" y1="200" x2="400" y2="200" stroke="black" strokeWidth="2" />
          <polyline
            fill="none"
            stroke="rgba(75, 192, 192, 1)"
            strokeWidth="3"
            points={[10, 20, 15, 30, 25, 10]
              .map((value, index) => `${index * 50 + 50},${200 - value * 5}`)
              .join(" ")}
          />
          <g className="axis x-axis">
            {["Jan", "Feb", "Mar", "Apr", "May", "Jun"].map((label, index) => (
              <text key={index} x={index * 50 + 50} y="215" textAnchor="middle" fontSize="12">
                {label}
              </text>
            ))}
          </g>
          <g className="axis y-axis">
            {[0, 10, 20, 30, 40].map((value, index) => (
              <text key={index} x="40" y={200 - value * 5} textAnchor="end" fontSize="12" dominantBaseline="middle">
                {value}
              </text>
            ))}
          </g>
        </svg>
      </div>
    </div>
  );
};

export default Dashboard;
