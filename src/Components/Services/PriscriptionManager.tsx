import React, { useState } from 'react';
import axios from 'axios';

const PrescriptionManager: React.FC = () => {
  const [image, setImage] = useState<File | null>(null);
  const [medicineData, setMedicineData] = useState<any>(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  // Handle file input change
  const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.target.files && e.target.files[0]) {
      setImage(e.target.files[0]);
    }
  };

  // Handle form submission
  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!image) {
      setError('Please upload an image.');
      return;
    }
    setLoading(true);
    setError(null);

    const formData = new FormData();
    formData.append('prescription', image);

    try {
      // Validate file type
      const validTypes = ['image/jpeg', 'image/png', 'image/jpg'];
      if (!validTypes.includes(image.type)) {
        setError('Please upload a valid image file (jpg, jpeg, png).');
        setLoading(false);
        return;
      }

      // Send the image to the POST endpoint
      await axios.post('http://127.0.0.1:5000/extract_medicine_name', formData, {
        headers: {
          'Content-Type': 'multipart/form-data',
        },
      });

      // Fetch the data from the GET endpoint
      const response = await axios.get('http://127.0.0.1:5000/get_info');
      setMedicineData(response.data);
    } catch (err) {
      setError('Error uploading image or fetching data.');
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="min-h-screen flex flex-col md:flex-row items-center justify-center bg-gray-100 p-8 space-y-8 md:space-y-0">
      {/* Left Section */}
      <div className="w-full md:w-1/3 max-w-lg bg-white p-8 rounded-lg shadow-lg md:mr-8">
        <h1 className="text-4xl font-semibold mb-6 text-center">Prescription Manager</h1>

        <form onSubmit={handleSubmit} className="space-y-6">
          <div>
            <label htmlFor="prescription" className="block text-xl font-medium text-gray-700 mb-3">Upload Prescription</label>
            <input
              type="file"
              id="prescription"
              accept="image/*"
              onChange={handleFileChange}
              className="w-full p-4 border border-gray-300 rounded-md text-xl"
            />
          </div>

          <button
            type="submit"
            className="w-full py-4 bg-blue-500 text-white font-semibold rounded-md disabled:bg-gray-400 text-xl"
            disabled={loading}
          >
            {loading ? 'Uploading...' : 'Submit Prescription'}
          </button>
        </form>

        {error && <p className="mt-4 text-red-500 text-center">{error}</p>}
      </div>

      {/* Right Section */}
      <div className="w-full md:w-2/3 max-w-2xl bg-white p-8 rounded-lg shadow-lg">
        <h3 className="text-3xl font-semibold mb-6 text-center">
          {medicineData ? 'Medicine Information' : 'Please upload a prescription to get medicine details'}
        </h3>

        {/* Default info while no prescription is uploaded */}
        {!medicineData && (
          <p className="text-center text-lg text-gray-500">
            Upload a prescription image to get detailed information about the medicine.
          </p>
        )}

        {/* Display actual medicine information after successful upload */}
        {medicineData && (
          <div className="space-y-6 text-lg">
            <p><strong>Name:</strong> {medicineData.name}</p>
            <p><strong>Description:</strong> {medicineData.description}</p>
            <p><strong>Dosage:</strong> {medicineData.dosage}</p>
            <p><strong>Side Effects:</strong> {medicineData.sideEffects}</p>
          </div>
        )}
      </div>
    </div>
  );
};

export default PrescriptionManager;
