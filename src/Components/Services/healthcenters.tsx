import React, { useEffect, useState } from 'react';

const HealthCenters = () => {
  const [mapUrl, setMapUrl] = useState('');
  const suppressSVGErrors = () => {
    const oldConsoleError = console.error;
    console.error = function (message, ...args) {
      if (typeof message === 'string' && message.includes('Expected length')) return;
      oldConsoleError(message, ...args);
    };
  };
  useEffect(() => {
    const fetchMap = async () => {
      try {
        const response = await fetch('/');
        if (!response.ok) {
          throw new Error('Failed to fetch the map');
        }
        setMapUrl(response.url);
      } catch (error) {
        console.error('Error fetching map:', error);
      }
    };

    fetchMap();
    suppressSVGErrors();
  }, []);

  return (
    <div>
      <h1>Nearby Health Centers Map</h1>
      {mapUrl ? (
        <iframe src={mapUrl} width="100%" height="600px" style={{ border: "none" }} title="Health Centers Map" />
      ) : (
        <p>Loading map...</p>
      )}
    </div>
  );
};

export default HealthCenters;
