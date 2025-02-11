import express from "express";
import mongoose from "mongoose";
import cors from "cors";
import dotenv from "dotenv";
import bcrypt from "bcryptjs";
import { PatientModel, ChatHistoryModel } from "./models.js";

dotenv.config();
const MONGO_URI = process.env.MONGO_URI;

const app = express();
const PORT = 3001; // Hardcoded Port

// Middleware
app.use(express.json());
app.use(cors());

// Connect to MongoDB
mongoose.connect(MONGO_URI, { useNewUrlParser: true, useUnifiedTopology: true })
    .then(() => console.log("✅ Connected to MongoDB: healthDB"))
    .catch(err => {
        console.error("❌ MongoDB connection error:", err);
        process.exit(1);
    });

// Register route
app.post("/register", async (req, res) => {
    try {
        const { mobile_number, password, ...rest } = req.body;

        // Check if mobile number already exists
        const existingUser = await PatientModel.findOne({ mobile_number });
        if (existingUser) {
            return res.status(400).json({ message: "User with this mobile number already exists" });
        }

        // Create new patient
        const newUser = new PatientModel({ mobile_number, password, ...rest });
        await newUser.save();

        res.status(201).json({ message: "Registration successful", user: newUser });
    } catch (err) {
        res.status(500).json({ message: "Error registering patient", error: err.message });
    }
});

// Login route
app.post("/login", async (req, res) => {
    try {
        const { mobile_number, password } = req.body;
        const user = await PatientModel.findOne({ mobile_number });

        if (!user) {
            return res.status(400).json({ message: "Invalid Credentials" });
        }

        const isMatch = await bcrypt.compare(password, user.password);
        if (!isMatch) {
            return res.status(400).json({ message: "Password Incorrect" });
        }

        res.json({ message: "Login successful", user });
    } catch (err) {
        res.status(500).json({ message: "Error logging in", error: err.message });
    }
});

// Store chat history
app.post("/chat", async (req, res) => {
    try {
        const { patient_id, user_message, bot_response, tokens_used } = req.body;

        const chatEntry = new ChatHistoryModel({ patient_id, user_message, bot_response, tokens_used });
        await chatEntry.save();

        res.status(201).json({ message: "Chat history saved", chatEntry });
    } catch (err) {
        res.status(500).json({ message: "Error saving chat history", error: err.message });
    }
});

// Get chat history by patient ID
app.get("/chat/:patient_id", async (req, res) => {
    try {
        const { patient_id } = req.params;
        const chatHistory = await ChatHistoryModel.find({ patient_id }).sort({ createdAt: -1 });

        res.json({ chatHistory });
    } catch (err) {
        res.status(500).json({ message: "Error fetching chat history", error: err.message });
    }
});

// Start Server
app.listen(PORT, () => {
    console.log(`🚀 Server is running on port ${PORT}`);
});
