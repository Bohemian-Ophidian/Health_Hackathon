import mongoose from "mongoose";
import bcrypt from 'bcryptjs';

const DoctorSchema = new mongoose.Schema({
    name: { type: String, required: true },
    specialization: { type: String, required: true },
    years_of_experience: { type: Number, required: true },
    contact_info: { type: String, required: true }
}, { timestamps: true });

const DoctorModel = mongoose.model("doctors", DoctorSchema);



const PatientSchema = new mongoose.Schema({
    sex: { type: String, required: true },
    height: { type: Number, required: true },
    weight: { type: Number, required: true },
    mobile_number: { type: Number, required: true },
    password: { type: String, required: true },
    name: { type: String, required: true },
    age: { type: Number, required: true },
    medical_history: { type: String, default: "" }
}, { timestamps: true });

PatientSchema.pre('save', async function(next) {
    if (this.isModified('password')) {
        const salt = await bcrypt.genSalt(10);
        this.password = await bcrypt.hash(this.password, salt);
    }
    next();
});

const PatientModel = mongoose.model("patients", PatientSchema);
// Chat History Schema
const ChatHistorySchema = new mongoose.Schema({
    patient_id: { type: mongoose.Schema.Types.ObjectId, ref: "patients", required: true },
    user_message: { type: String, required: true },
    bot_response: { type: String, required: true },
    tokens_used: { type: Number, default: 0 }
}, { timestamps: true });



const ChatHistoryModel = mongoose.model("chat_history", ChatHistorySchema);

export { DoctorModel, PatientModel, ChatHistoryModel };