import React, { useState } from 'react';
import axios from 'axios';

// Helper function to parse the model's raw text into structured fields.
const parseMedicineInfo = (text: string) => {
  const nameMatch = text.match(/Name:\s*(.*?)(?=\n|$)/i);
  const descriptionMatch = text.match(/Description:\s*([\s\S]*?)(?=Dosage:|Side Effects:|$)/i);
  const dosageMatch = text.match(/Dosage:\s*([\s\S]*?)(?=Side Effects:|$)/i);
  const sideEffectsMatch = text.match(/Side Effects:\s*([\s\S]*?)(?=$)/i);

  return {
    name: nameMatch ? nameMatch[1].trim() : '',
    description: descriptionMatch ? descriptionMatch[1].trim() : '',
    dosage: dosageMatch ? dosageMatch[1].trim() : '',
    sideEffects: sideEffectsMatch ? sideEffectsMatch[1].trim() : '',
  };
};

const PrescriptionManager: React.FC = () => {
  const [image, setImage] = useState<File | null>(null);
  const [medicineData, setMedicineData] = useState<any>(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.target.files && e.target.files[0]) {
      setImage(e.target.files[0]);
    }
  };

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
      const validTypes = ['image/png', 'image/jpeg', 'image/jpg'];
      if (!validTypes.includes(image.type)) {
        setError('Please upload a valid image file (jpg, jpeg, png).');
        setLoading(false);
        return;
      }

      const extractResponse = await axios.post('http://127.0.0.1:5000/extract_medicine_name', formData, {
        headers: {
          'Content-Type': 'multipart/form-data',
        },
      });

      const medicineName = extractResponse.data.medicine_name;

      const infoResponse = await axios.post('http://127.0.0.1:5000/get_info', { medicine_name: medicineName });

      const parsedInfo = parseMedicineInfo(infoResponse.data.model_response);

      const { name, ...restParsedInfo } = parsedInfo;

      setMedicineData({ name: medicineName, ...restParsedInfo });
    } catch (err) {
      setError('Error uploading image or fetching data.');
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="min-h-screen flex flex-col md:flex-row items-center justify-center bg-gray-100 p-8 space-y-8 md:space-y-0">
      <div className="w-full md:w-1/3 max-w-lg bg-white p-8 rounded-lg shadow-lg md:mr-8">
        <h1 className="text-4xl font-semibold mb-6 text-center">Prescription Manager</h1>

        <form onSubmit={handleSubmit} className="space-y-6">
          <div>
            <label htmlFor="prescription" className="block text-xl font-medium text-gray-700 mb-3">
              Upload Prescription
            </label>
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

      <div className="w-full md:w-2/3 max-w-2xl bg-white p-8 rounded-lg shadow-lg">
        <h3 className="text-3xl font-semibold mb-6 text-center">
          {medicineData ? 'Medicine Information' : 'Please upload a prescription to get medicine details'}
        </h3>

        {!medicineData && (
          <p className="text-center text-lg text-gray-500">
            Upload a prescription image to get detailed information about the medicine.
          </p>
        )}

        {medicineData && (
          <table className="min-w-full bg-white">
            <thead>
              <tr>
                <th className="py-2 px-4 border-b border-gray-200 bg-gray-100 text-left text-sm leading-4 font-medium text-gray-600 uppercase tracking-wider">
                  Field
                </th>
                <th className="py-2 px-4 border-b border-gray-200 bg-gray-100 text-left text-sm leading-4 font-medium text-gray-600 uppercase tracking-wider">
                  Information
                </th>
              </tr>
            </thead>
            <tbody>
              <tr>
                <td className="py-2 px-4 border-b border-gray-200">Name</td>
                <td className="py-2 px-4 border-b border-gray-200">{medicineData.name || 'Not available'}</td>
              </tr>
              <tr>
                <td className="py-2 px-4 border-b border-gray-200">Description</td>
                <td className="py-2 px-4 border-b border-gray-200">{medicineData.description || 'Not available'}</td>
              </tr>
              <tr>
                <td className="py-2 px-4 border-b border-gray-200">Dosage</td>
                <td className="py-2 px-4 border-b border-gray-200">{medicineData.dosage || 'Not available'}</td>
              </tr>
              <tr>
                <td className="py-2 px-4 border-b border-gray-200">Side Effects</td>
                <td className="py-2 px-4 border-b border-gray-200">{medicineData.sideEffects || 'Not available'}</td>
              </tr>
            </tbody>
          </table>
        )}
      </div>
    </div>
  );
};

export default PrescriptionManager;
