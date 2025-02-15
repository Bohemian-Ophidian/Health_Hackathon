import { useEffect, useState } from 'react';
import axios from 'axios';
import { MapContainer, TileLayer, Marker, Popup } from 'react-leaflet';
import { LatLngExpression } from 'leaflet'; 
import 'leaflet/dist/leaflet.css';

interface Hospital {
  name: string;
  latitude: number;
  longitude: number;
}

const HealthCenters = () => {
  const [hospitals, setHospitals] = useState<Hospital[]>([]);
  const [postalCode, setPostalCode] = useState<string>('411001'); // Example postal code for Pune

  useEffect(() => {
    // Fetch the hospitals data from FastAPI
    axios
      .get(`http://localhost:8000/api/hospitals/${postalCode}`)
      .then((response) => setHospitals(response.data))
      .catch((error) => console.error('Error fetching hospitals:', error));
  }, [postalCode]);

  const defaultCenter: LatLngExpression = [20.5937, 78.9629]; // Default center (India)

  return (
    <div>
      <h1>Nearby Health Centers</h1>
      <div style={{ marginBottom: '10px' }}>
        <label htmlFor="postalCode">Enter Postal Code: </label>
        <input
          type="text"
          id="postalCode"
          value={postalCode}
          onChange={(e) => setPostalCode(e.target.value)}
        />
      </div>

      <MapContainer center={defaultCenter} zoom={5} style={{ height: '560px', width: '100%' }}>
        <TileLayer
          url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
        />
        {hospitals.map((hospital) => (
          <Marker key={hospital.name} position={[hospital.latitude, hospital.longitude]}>
            <Popup>{hospital.name}</Popup>
          </Marker>
        ))}
      </MapContainer>
    </div>
  );
};

export default HealthCenters;
