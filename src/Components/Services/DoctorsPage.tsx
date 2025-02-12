import React, { useState } from 'react';
import DatePicker from 'react-datepicker';
import "react-datepicker/dist/react-datepicker.css";

interface Doctor {
  name: string;
  title: string;
  speciality: string;
  experience: string;
  gender: string;
  profileLink: string;
  imageUrl: string;
}

const doctors: Doctor[] = [
  {
    name: "Dr. Sandeep Budhiraja",
    title: "Group Medical Director - Max Healthcare & Senior Director - Institute of Internal Medicine",
    speciality: "Internal Medicine",
    experience: "29+ Years",
    gender: "Male",
    profileLink: "#",
    imageUrl: "https://www.shutterstock.com/image-photo/profile-picture-smiling-young-caucasian-600nw-1954278664.jpg",
  },
  {
    name: "Dr. Balbir Singh",
    title: "Group Chairman - Cardiac Sciences, Pan Max & Chief of Interventional Cardiology and Electrophysiology, Max Saket",
    speciality: "Cardiac Sciences, Cardiology, Cardiac Electrophysiology-Pacemaker, Interventional Cardiology",
    experience: "34+ Years",
    gender: "Male",
    profileLink: "#",
    imageUrl: "https://www.shutterstock.com/image-photo/profile-picture-smiling-young-caucasian-600nw-1954278664.jpg",
  },
  {
    name: "Dr. Rajesh Ahlawat",
    title: "Group Chairman - Urology & Chairman - Centre of Excellence of Prostate & Urological Cancers and Male Health",
    speciality: "Urology, Uro-Oncology, Kidney Transplant",
    experience: "40+ Years",
    gender: "Male",
    profileLink: "#",
    imageUrl: "https://www.shutterstock.com/image-photo/profile-picture-smiling-young-caucasian-600nw-1954278664.jpg",
  },
  {
    name: "Dr. T.S. Kler",
    title: "Chairman & HOD - BLK-Max Heart & Vascular Institute, Chairman Pan Max - Electrophysiology - Heart & Vascular Institute",
    speciality: "Cardiac Electrophysiology-Pacemaker, Cardiology, Interventional Cardiology, Cardiac Arrhythmia, Cardiac Sciences",
    experience: "30+ Years",
    gender: "Male",
    profileLink: "#",
    imageUrl: "https://www.shutterstock.com/image-photo/profile-picture-smiling-young-caucasian-600nw-1954278664.jpg",
  },
  // Additional doctors added here
  {
    name: "Dr. John Doe",
    title: "Consultant Neurologist",
    speciality: "Neurology",
    experience: "10+ Years",
    gender: "Male",
    profileLink: "#",
    imageUrl: "https://www.shutterstock.com/image-photo/profile-picture-smiling-young-caucasian-600nw-1954278664.jpg",
  },
  {
    name: "Dr. Jane Smith",
    title: "Consultant Cardiologist",
    speciality: "Cardiology",
    experience: "15+ Years",
    gender: "Female",
    profileLink: "#",
    imageUrl: "https://www.shutterstock.com/image-photo/profile-picture-smiling-young-caucasian-600nw-1954278664.jpg",
  },
  {
    name: "Dr. Amy Adams",
    title: "Orthopedic Surgeon",
    speciality: "Orthopedics",
    experience: "12+ Years",
    gender: "Female",
    profileLink: "#",
    imageUrl: "https://www.shutterstock.com/image-photo/profile-picture-smiling-young-caucasian-600nw-1954278664.jpg",
  },
  {
    name: "Dr. Mark Lee",
    title: "General Surgeon",
    speciality: "Surgery",
    experience: "20+ Years",
    gender: "Male",
    profileLink: "#",
    imageUrl: "https://www.shutterstock.com/image-photo/profile-picture-smiling-young-caucasian-600nw-1954278664.jpg",
  },
  {
    name: "Dr. Linda Ford",
    title: "Dentist",
    speciality: "Dentistry",
    experience: "18+ Years",
    gender: "Female",
    profileLink: "#",
    imageUrl: "https://www.shutterstock.com/image-photo/profile-picture-smiling-young-caucasian-600nw-1954278664.jpg",
  },
  {
    name: "Dr. Harry White",
    title: "Radiologist",
    speciality: "Radiology",
    experience: "25+ Years",
    gender: "Male",
    profileLink: "#",
    imageUrl: "https://www.shutterstock.com/image-photo/profile-picture-smiling-young-caucasian-600nw-1954278664.jpg",
  },
  {
    name: "Dr. Sarah Green",
    title: "Urologist",
    speciality: "Urology",
    experience: "22+ Years",
    gender: "Female",
    profileLink: "#",
    imageUrl: "https://www.shutterstock.com/image-photo/profile-picture-smiling-young-caucasian-600nw-1954278664.jpg",
  },
  // More additional doctors across other specialities
  {
    name: "Dr. Thomas Bright",
    title: "Consultant Dentist",
    speciality: "Dentistry",
    experience: "10+ Years",
    gender: "Male",
    profileLink: "#",
    imageUrl: "https://www.shutterstock.com/image-photo/profile-picture-smiling-young-caucasian-600nw-1954278664.jpg",
  },
  {
    name: "Dr. Oscar Lee",
    title: "Consultant Orthopedist",
    speciality: "Orthopedics",
    experience: "20+ Years",
    gender: "Male",
    profileLink: "#",
    imageUrl: "https://www.shutterstock.com/image-photo/profile-picture-smiling-young-caucasian-600nw-1954278664.jpg",
  },
];

const DoctorsPage: React.FC = () => {
  const [appointments, setAppointments] = useState<{ date: Date; doctor: string; time: string }[]>([]);
  const [selectedDate, setSelectedDate] = useState<Date | null>(null);
  const [selectedTime, setSelectedTime] = useState<string>('');
  const [selectedDoctor, setSelectedDoctor] = useState<string | null>(null);
  const [showCalendar, setShowCalendar] = useState<boolean>(false); // Show calendar only when "Book an Appointment" is clicked
  const [filters, setFilters] = useState({ speciality: '', experience: '' });

  const handleBookAppointment = (doctorName: string) => {
    if (!selectedDate || !selectedTime) {
      alert("Please select both a date and a time before booking!");
      return;
    }
    const newAppointment = {
      date: selectedDate,
      doctor: doctorName,
      time: selectedTime,
    };
    setAppointments([...appointments, newAppointment]);
    setShowCalendar(false); // Hide the calendar after booking
    setSelectedDoctor(null); // Reset the selected doctor after booking
    alert(`Appointment booked with ${doctorName} on ${selectedDate.toLocaleDateString()} at ${selectedTime}`);
  };


  const handleCancelAppointment = (doctorName: string) => {
    const updatedAppointments = appointments.filter((appointment) => appointment.doctor !== doctorName);
    setAppointments(updatedAppointments);
    setSelectedDoctor(null); // Reset the selected doctor
    alert(`Appointment with ${doctorName} has been cancelled.`);
  };
  

  const handleDateChange = (date: Date | null) => {
    if (date) {
      setSelectedDate(date);
    }
  };

  const handleTimeChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
    setSelectedTime(event.target.value);
  };

  const isDateBooked = (date: Date) => {
    return appointments.some(appointment => appointment.date.toDateString() === date.toDateString());
  };

  const timeOptions = ['9:00 AM', '10:00 AM', '11:00 AM', '7:00 PM', '8:00 PM', '9:00 PM', '10:00 PM'].map(time => (
    <option key={time} value={time}>{time}</option>
  ));

  const filteredDoctors = doctors.filter(
    (doctor) =>
      (filters.speciality === '' || doctor.speciality.toLowerCase().includes(filters.speciality.toLowerCase())) &&
      (filters.experience === '' || doctor.experience.includes(filters.experience))
  );

  return (
    <div className="flex max-w-7xl mx-auto mt-12 p-4">
      {/* Left Sidebar */}
      <div className="w-1/4 p-4 border-r">
        <h3 className="text-xl font-semibold mb-4">Filters</h3>
        {/* Speciality Filter */}
        <div className="mb-4">
          <label className="block font-medium">Speciality</label>
          <input
            type="text"
            value={filters.speciality}
            onChange={(e) => setFilters({ ...filters, speciality: e.target.value })}
            placeholder="Search by speciality"
            className="w-full p-2 border border-gray-300 rounded"
          />
        </div>
        {/* Experience Filter */}
        <div className="mb-4">
          <label className="block font-medium">Experience</label>
          <input
            type="text"
            value={filters.experience}
            onChange={(e) => setFilters({ ...filters, experience: e.target.value })}
            placeholder="Search by experience"
            className="w-full p-2 border border-gray-300 rounded"
          />
        </div>
        {/* Registered Appointments */}
        <div className="mt-8">
          <h3 className="text-xl font-semibold mb-4">Registered Appointments</h3>
          <ul>
            {appointments.map((appointment, index) => (
              <li key={index} className="mb-2">{`${appointment.doctor}: ${appointment.date.toLocaleDateString()} at ${appointment.time}`}</li>
            ))}
          </ul>
        </div>
      </div>

      {/* Doctors Available */}
      <div className="w-3/4 p-4">
        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
          {filteredDoctors.map((doctor, index) => (
            <div key={index} className="flex flex-col border rounded-lg p-4 shadow-md hover:shadow-lg transition-shadow duration-300 h-full">
              <img
                src={doctor.imageUrl}
                alt={doctor.name}
                className="w-32 h-32 object-cover rounded-full mx-auto"
              />
              <div className="flex flex-col justify-between flex-grow mt-4">
                <h2 className="font-semibold text-xl text-center">{doctor.name}</h2>
                <p className="text-sm text-gray-500 text-center">{doctor.title}</p>
                <p className="text-md text-gray-800 font-medium text-center">{doctor.speciality}</p>
                <p className="text-sm text-gray-500 text-center">Experience: {doctor.experience}</p>
                <p className="text-sm text-gray-500 text-center">Gender: {doctor.gender}</p>

                {/* Conditional Rendering: "Book Appointment" or "Cancel Appointment" */}
                <div className="mt-4 text-center">
                  {appointments.some(app => app.doctor === doctor.name) ? (
                    <button
                      onClick={() => handleCancelAppointment(doctor.name)}
                      className="bg-red-500 text-white px-4 py-2 rounded-lg hover:bg-red-700"
                    >
                      Cancel Appointment
                    </button>
                  ) : (
                    <button
                      onClick={() => {
                        setSelectedDoctor(doctor.name); // Select the doctor
                        setShowCalendar(true); // Show the calendar and time picker
                        
                      }}
                      className="bg-blue-500 text-white px-4 py-2 rounded-lg hover:bg-blue-700"
                    >
                      Book an Appointment
                    </button>
                  )}
                </div>

                {/* Show Calendar and Time Selection Only When Clicking "Book an Appointment" */}
                {showCalendar && selectedDoctor === doctor.name && (
                  <div className="mt-4 text-center">
                    <div className="mb-4">
                      <DatePicker
                        selected={selectedDate}
                        onChange={handleDateChange}
                        dateFormat="yyyy/MM/dd"
                        placeholderText="Select a date"
                        className="w-full p-2 border border-gray-300 rounded"
                        filterDate={(date) => !isDateBooked(date)} // Filter out dates that are already booked
                      />
                    </div>
                    {selectedDate && (
                      <select
                        value={selectedTime}
                        onChange={handleTimeChange}
                        className="mt-2 p-2 border border-gray-300 rounded w-full"
                      >
                        <option value="">Select a time</option>
                        {timeOptions}
                      </select>
                    )}
                    <div className="mt-4">
                      <button
                        onClick={() => handleBookAppointment(doctor.name)}
                        className="bg-blue-500 text-white px-4 py-2 rounded-lg hover:bg-blue-700 mt-2"
                      >
                        Confirm Appointment
                      </button>
                    </div>
                  </div>
                )}
              </div>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
};


export default DoctorsPage;
