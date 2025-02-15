import React, { useEffect, useState } from 'react';

const HealthCenters = () => {
  const [mapUrl, setMapUrl] = useState<string>('http://127.0.0.1:5500/map_service/map.html');

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
