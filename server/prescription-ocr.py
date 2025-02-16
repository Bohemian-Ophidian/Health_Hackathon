from flask import Flask, request, jsonify
import pytesseract
from PIL import Image
from transformers import AutoModelForCausalLM, AutoTokenizer

# Initialize Flask
app = Flask(__name__)

# Initialize Tesseract and Hugging Face Model
pytesseract.pytesseract.tesseract_cmd = r"C:\Program Files\Tesseract-OCR\tesseract.exe"  # Windows path to Tesseract (adjust for your system)
model_name = "Qwen/Qwen2.5-1.5B-Instruct"

# Load Hugging Face Model
model = AutoModelForCausalLM.from_pretrained(
    model_name,
    torch_dtype="auto",
    device_map="auto"
)
tokenizer = AutoTokenizer.from_pretrained(model_name)

# Function to extract text from the image using Tesseract OCR
def extract_text_from_image(image_path):
    # Open the image file using PIL
    img = Image.open(image_path)
    
    # Use Tesseract to do OCR and extract text
    extracted_text = pytesseract.image_to_string(img)
    
    # Return the extracted text
    return extracted_text.strip()

# Function to query the Hugging Face model with extracted medicine name
def query_model(medicine_name):
    # Pre-prompt to guide the model's response to be in two specific sections: Alternatives and Side Effects
    prompt = f"Please provide information about all provided medicines in two sections:\n\n" \
             f"1. **Alternative Medicines**: List other medicines that can be used for the same purpose.\n" \
             f"2. **Side Effects**: List common side effects of {medicine_name}.\n\n" \
             f"Medicine Name: {medicine_name}\n"

    # Structure the conversation
    messages = [
        {"role": "system", "content": "You are Qwen, created by Alibaba Cloud. You are a helpful assistant."},
        {"role": "user", "content": prompt}
    ]
    
    # Apply chat template
    text = tokenizer.apply_chat_template(
        messages,
        tokenize=False,
        add_generation_prompt=True
    )
    
    model_inputs = tokenizer([text], return_tensors="pt").to(model.device)
    
    # Generate response from the model
    generated_ids = model.generate(
        **model_inputs,
        max_new_tokens=512
    )
    
    # Get the generated response
    generated_ids = [
        output_ids[len(input_ids):] for input_ids, output_ids in zip(model_inputs.input_ids, generated_ids)
    ]
    
    response = tokenizer.batch_decode(generated_ids, skip_special_tokens=True)[0]
    
    return response

# Route to extract the medicine name from an image (OCR)
@app.route('/extract_medicine_name', methods=['POST'])
def extract_medicine_name():
    file = request.files['image']
    # Save the file temporarily
    file_path = 'temp_image.png'
    file.save(file_path)
    
    # Extract text from the image
    medicine_name = extract_text_from_image(file_path)
    
    # Return the extracted text as JSON
    if medicine_name:
        return jsonify({"medicine_name": medicine_name}), 200
    else:
        return jsonify({"error": "No text extracted from the image"}), 400

# Route to get information about the medicine
@app.route('/get_info', methods=['POST'])
def query_medicine_info():
    # Get the medicine name from the request body
    data = request.get_json()
    medicine_name = data.get('medicine_name')
    
    if not medicine_name:
        return jsonify({"error": "No medicine name provided"}), 400
    
    # Query the model with the medicine name
    model_response = query_model(medicine_name)
    
    return jsonify({"model_response": model_response}), 200

# Main function to run the Flask app
if __name__ == '__main__':
    app.run(debug=True)
