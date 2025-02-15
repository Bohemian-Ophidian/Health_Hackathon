from fastapi import FastAPI
from typing import List
from map_service import find_hospitals_osm
from map_service import generate_map
from pydantic import BaseModel

app = FastAPI()

# Define the response model
class Hospital(BaseModel):
    name: str
    latitude: float
    longitude: float

@app.get("/api/hospitals/{postal_code}", response_model=List[Hospital])
def get_hospitals(postal_code: str):
    """
    API endpoint to get hospitals near a postal code.
    """
    hospitals = find_hospitals_osm(postal_code)
    if isinstance(hospitals, list):
        generate_map(hospitals)
        for hospital in hospitals:
            print(f"🏥 Name: {hospital['name']}\n🌍 Location: {hospital['latitude']}, {hospital['longitude']}\n")
    if hospitals:
        return hospitals
    return []

if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=8000)
