import os
import google.generativeai as genai
from fastapi import FastAPI, HTTPException
from pydantic import BaseModel
from pymongo import MongoClient
from bson import ObjectId
from dotenv import load_dotenv
from fastapi.middleware.cors import CORSMiddleware

# Load environment variables
load_dotenv()
MONGO_URI = os.getenv("MONGO_URI")
GEMINI_API_KEY = os.getenv("GEMINI_API_KEY")

# Validate API Key
if not GEMINI_API_KEY:
    raise RuntimeError("❌ Google API key is missing. Set 'GEMINI_API_KEY' in environment variables.")

# Configure Gemini API
genai.configure(api_key=GEMINI_API_KEY)

# Initialize FastAPI
app = FastAPI()

# Enable CORS for frontend communication
app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],  # Set this to your frontend domain in production
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

# Connect to MongoDB
client = MongoClient(MONGO_URI)
db = client.healthDB  # Database Name
patients_collection = db.patients  # Patients Collection
chat_history_collection = db.chat_history  # Chat History Collection

# Pydantic model for Patient
class Patient(BaseModel):
    name: str
    age: int
    gender: str
    medical_history: str

# Pydantic model for Chat Requests
class ChatRequest(BaseModel):
    patient_id: str  # MongoDB ObjectId as a string
    user_message: str

# Gemini model configuration
safety_settings = [
    {"category": "HARM_CATEGORY_HARASSMENT", "threshold": "BLOCK_NONE"},
    {"category": "HARM_CATEGORY_HATE_SPEECH", "threshold": "BLOCK_MEDIUM_AND_ABOVE"},
    {"category": "HARM_CATEGORY_SEXUALLY_EXPLICIT", "threshold": "BLOCK_MEDIUM_AND_ABOVE"},
    {"category": "HARM_CATEGORY_DANGEROUS_CONTENT", "threshold": "BLOCK_MEDIUM_AND_ABOVE"},
]

generation_config = {
    "temperature": 0.7,
    "top_p": 0.9,
    "max_output_tokens": 500,
}

model = genai.GenerativeModel(
    model_name="gemini-2.0-flash-exp",
    safety_settings=safety_settings,
    generation_config=generation_config,
    # system_instruction="",
    system_instruction="Your role is to listen to the user, understand his symptoms, tailor your response according to it, and you must be straightforward and concise about it, also provide quick and simple solutions like first aid or other remedies before asking the patient to consult a doctor.",
)

# ✅ API for Chatbot Interaction
@app.post("/chat/chatWithBot")
def chat_with_bot(chat_request: ChatRequest):
    try:
        patient_id = chat_request.patient_id
        user_message = chat_request.user_message

        # Validate patient existence
        patient = patients_collection.find_one({"_id": ObjectId(patient_id)})
        if not patient:
            raise HTTPException(status_code=404, detail="Patient not found")

        # Format patient details into context
        patient_context = (
            f"Patient Details:\n"
            f"- Name: {patient['name']}\n"
            f"- Age: {patient['age']}\n"
            f"- Sex: {patient['sex']}\n"
            f"- Medical History: {patient['medical_history']}\n\n"
            f"Now respond to the following message from the patient:"
        )

        # Combine patient context with user message
        full_prompt = f"{patient_context}\n\nUser: {user_message}"

        # Generate response from Gemini AI
        response = model.generate_content(full_prompt)
        bot_response = response.text.strip() if hasattr(response, "text") else "Sorry, I couldn't generate a response."

        # Store chat in MongoDB
        chat_entry = {
            "patient_id": ObjectId(patient_id),
            "user_message": user_message,
            "bot_response": bot_response
        }
        chat_history_collection.insert_one(chat_entry)

        return {
            "patient_id": patient_id,
            "user_message": user_message,
            "bot_response": bot_response
        }

    except Exception as e:
        raise HTTPException(status_code=500, detail=f"Chatbot error: {str(e)}")
