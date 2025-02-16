import React, { useEffect, useState } from 'react';

const HealthCenters = () => {
  const [mapUrl, setMapUrl] = useState<string>('http://127.0.0.1:5500/map_service/map.html');

const [postalCode, setPostalCode] = useState<string>('');

const handlePostalCodeChange = (event: React.ChangeEvent<HTMLInputElement>) => {
  const value = event.target.value;
  if (/^\d{0,6}$/.test(value)) {
    setPostalCode(value);
  }
};

const handleSearch = () => {
  // Logic to update mapUrl based on postalCode
  // Example: setMapUrl(`http://example.com/map?postalCode=${postalCode}`);
};

useEffect(() => {
  handleSearch();
}, [postalCode]);

return (
  <div>
    <input
      type="text"
      value={postalCode}
      onChange={handlePostalCodeChange}
      placeholder="Enter 6-digit postal code"
      maxLength={6}
    />
    <iframe 
      src={mapUrl} 
      title="Health Centers Map" 
      width="100%" 
      height="850vh" 
      style={{ border: 'none' }} 
    />
  </div>
);

  return (
    <div>
      <iframe 
        src={mapUrl} 
        title="Health Centers Map" 
        width="100%" 
        height="850vh" 
        style={{ border: 'none' }} 
      />
    </div>
  );
};

export default HealthCenters;
