import React, { useEffect, useState } from 'react';

const HealthCenters = () => {
  const [mapUrl, setMapUrl] = useState<string>('http://127.0.0.1:5500/map_service/map.html');

  return (
    <div>
      <iframe 
        src={mapUrl} 
        title="Health Centers Map" 
        width="100%" 
<<<<<<< HEAD
        height="850vh" 
=======
        height="800vh" 
>>>>>>> 9e2c0f7c039e70181fbcf793d4822f4554bd0c05
        style={{ border: 'none' }} 
      />
    </div>
  );
};

export default HealthCenters;
