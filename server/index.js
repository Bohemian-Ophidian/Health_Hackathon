import express from "express";
import mongoose from "mongoose";
import cors from "cors";
import dotenv from "dotenv";
import bcrypt from "bcryptjs";
import { PatientModel } from "./models.js";

dotenv.config();
const MONGO_URI = process.env.MONGO_URI;

const app = express();
const PORT = 3001; // Hardcoded Port

// Middleware
app.use(express.json());
app.use(cors()); // Enable CORS

// Register route
app.post("/register", (req, res) => {
    PatientModel.create(req.body)
        .then(data => res.json(data))
        .catch(err => res.status(500).json({ message: "Error registering patient", error: err }));
});

// Login route
app.post("/login", (req, res) => {
    const { mobile_number, password } = req.body;
    PatientModel.findOne({ mobile_number })
        .then(async (data) => {
            if (data) {
                const isMatch = await bcrypt.compare(password, data.password);
                if (isMatch) {
                    res.json("success");
                } else {
                    res.json({ message: "Password Incorrect" });
                }
            } else {
                res.json({ message: "Invalid Credentials" });
            }
        })
        .catch(err => res.status(500).json({ message: "Error logging in", error: err }));
});

// Connect to MongoDB
mongoose.connect(MONGO_URI, { useNewUrlParser: true, useUnifiedTopology: true })
    .then(() => console.log("Connected to MongoDB"))
    .catch(err => {
        console.error("MongoDB connection error:", err);
        process.exit(1); // Exit process if MongoDB connection fails
    });

// Start Server
app.listen(PORT, () => {
    console.log(`Server is running on port ${PORT}`);
});
