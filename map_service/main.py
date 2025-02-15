from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware
import os

app = FastAPI()

# Update the origins with the correct frontend URL
origins = [
    "http://localhost:5175",  # If your frontend might run on different ports, include them
    "http://localhost:5176"   # Add the current port the frontend is using
]

app.add_middleware(
    CORSMiddleware,
    allow_origins=origins,  # Allow all the origins in the list
    allow_credentials=True,
    allow_methods=["*"],    # Allow all methods
    allow_headers=["*"],    # Allow all headers
)

@app.get("/map")
def get_map():
    map_path = "map.html"  # Ensure this is the correct path to your map file
    if os.path.exists(map_path):
        return FileResponse(map_path)
    else:
        return {"error": "Map file not found."}

if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=8000)
