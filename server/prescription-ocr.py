from flask import Flask, request, jsonify
import pytesseract
from PIL import Image
from transformers import AutoModelForCausalLM, AutoTokenizer
from flask_cors import CORS

app = Flask(__name__)
CORS(app)  # Allow all origins

# Initialize Tesseract and Hugging Face Model
pytesseract.pytesseract.tesseract_cmd = r"C:\Program Files\Tesseract-OCR\tesseract.exe"  # Adjust path for your system
model_name = "Qwen/Qwen2.5-1.5B-Instruct"

# Load Hugging Face Model
model = AutoModelForCausalLM.from_pretrained(
    model_name,
    torch_dtype="auto",
    device_map="auto"
)
tokenizer = AutoTokenizer.from_pretrained(model_name)

def extract_text_from_image(image_path):
    # Open the image file using PIL and extract text via Tesseract
    img = Image.open(image_path)
    extracted_text = pytesseract.image_to_string(img)
    return extracted_text.strip()

def query_model(medicine_name):
    # More explicit instructions:
    system_message = (
        "You are Qwen, created by Alibaba Cloud. You are a helpful assistant. "
        "Please strictly follow the user's instructions about how to structure your output."
    )
    #This is not working for name section. See V Aditya Teja
    user_message = (
        f"I have identified a medicine named '{medicine_name}'.\n\n"
        f"1. **Alternative Medicines**: List 2 other medicines that can be used for the same purposes as {medicine_name}.\n"
        f"2. **Side Effects**: List 2 common side effects of {medicine_name}.\n\n"
        f"Please output in Markdown, with headings 'Alternative Medicines' and 'Side Effects'."
    )

    messages = [
        {"role": "system", "content": system_message},
        {"role": "user", "content": user_message}
    ]

    text = tokenizer.apply_chat_template(messages, tokenize=False, add_generation_prompt=True)
    
    model_inputs = tokenizer([text], return_tensors="pt").to(model.device)
    generated_ids = model.generate(**model_inputs, max_new_tokens=512)
    generated_ids = [output_ids[len(input_ids):] 
                     for input_ids, output_ids in zip(model_inputs.input_ids, generated_ids)]
    response = tokenizer.batch_decode(generated_ids, skip_special_tokens=True)[0]
    
    return response

# Route to extract the medicine name from an image (OCR)
@app.route('/extract_medicine_name', methods=['POST'])
def extract_medicine_name_endpoint():
    # Look for the file under 'prescription' as used by the React code
    if 'prescription' not in request.files:
        return jsonify({"error": "No file part in the request"}), 400

    file = request.files['prescription']
    if file.filename == '':
        return jsonify({"error": "No file selected"}), 400

    # Save the file temporarily
    file_path = 'temp_image.png'
    file.save(file_path)
    
    # Extract text from the image
    medicine_name = extract_text_from_image(file_path)
    
    if medicine_name:
        return jsonify({"medicine_name": medicine_name}), 200
    else:
        return jsonify({"error": "No text extracted from the image"}), 400

# Route to get information about the medicine using the extracted name
@app.route('/get_info', methods=['POST'])
def query_medicine_info():
    data = request.get_json()
    medicine_name = data.get('medicine_name')
    
    if not medicine_name:
        return jsonify({"error": "No medicine name provided"}), 400
    
    model_response = query_model(medicine_name)
    
    # Here you could optionally parse model_response into structured fields if desired.
    # For now, we simply return the full response.
    return jsonify({"model_response": model_response}), 200

if __name__ == '__main__':
    app.run(debug=True)
