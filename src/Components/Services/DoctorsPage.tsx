import React, { useState, useEffect } from 'react';
import axios from 'axios';
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

const DoctorsPage: React.FC = () => {
  const [doctors, setDoctors] = useState<Doctor[]>([]);
  const [appointments, setAppointments] = useState<{ date: Date; doctor: string; time: string }[]>([]);
  const [selectedDate, setSelectedDate] = useState<Date | null>(null);
  const [selectedTime, setSelectedTime] = useState<string>('');
  const [selectedDoctor, setSelectedDoctor] = useState<string | null>(null);
  const [showCalendar, setShowCalendar] = useState<boolean>(false);
  const [filters, setFilters] = useState({ speciality: '', experience: '' });

  useEffect(() => {
    const fetchDoctors = async () => {
      try {
        const response = await axios.get('http://localhost:3001/doctors');
        setDoctors(response.data);
      } catch (error) {
        console.error('Failed to fetch doctors', error);
      }
    };
    fetchDoctors();
  }, []);

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
    setShowCalendar(false); 
    setSelectedDoctor(null);
    alert(`Appointment booked with ${doctorName} on ${selectedDate.toLocaleDateString()} at ${selectedTime}`);
  };

  const handleCancelAppointment = (doctorName: string) => {
    const updatedAppointments = appointments.filter((appointment) => appointment.doctor !== doctorName);
    setAppointments(updatedAppointments);
    alert(`Appointment with ${doctorName} has been cancelled.`);
  };

  const handleDateChange = (date: Date | null) => {
    setSelectedDate(date);
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
      <div className="w-1/4 p-4 border-r">
        <h3 className="text-xl font-semibold mb-4">Filters</h3>
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
        <div className="mt-8">
          <h3 className="text-xl font-semibold mb-4">Registered Appointments</h3>
          <ul>
            {appointments.map((appointment, index) => (
              <li key={index} className="mb-2">{`${appointment.doctor}: ${appointment.date.toLocaleDateString()} at ${appointment.time}`}</li>
            ))}
          </ul>
        </div>
      </div>
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
                        setSelectedDoctor(doctor.name);
                        setShowCalendar(true);
                      }}
                      className="bg-blue-500 text-white px-4 py-2 rounded-lg hover:bg-blue-700"
                    >
                      Book an Appointment
                    </button>
                  )}
                </div>
                {showCalendar && selectedDoctor === doctor.name && (
                  <div className="mt-4 text-center">
                    <div className="mb-4">
                      <DatePicker
                        selected={selectedDate}
                        onChange={handleDateChange}
                        dateFormat="yyyy/MM/dd"
                        placeholderText="Select a date"
                        className="w-full p-2 border border-gray-300 rounded"
                        filterDate={(date) => !isDateBooked(date)}
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
