package models

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Prescription model for MongoDB
type Prescription struct {
	ID            string    bson:"_id,omitempty" json:"id"
	OriginalImage string    bson:"original_image" json:"original_image"
	ExtractedText string    bson:"extracted_text" json:"extracted_text"
	CreatedAt     time.Time bson:"created_at" json:"created_at"
}

// MongoDB reference
var DB *mongo.Database

// SavePrescription inserts a prescription into MongoDB
func SavePrescription(p *Prescription) error {
	p.CreatedAt = time.Now()

	collection := DB.Collection("prescriptions")
	_, err := collection.InsertOne(context.TODO(), p)
	if err != nil {
		log.Printf("❌ Failed to save prescription: %v", err)
		return err
	}
	log.Println("✅ Prescription saved successfully")
	return nil
}

// GetPrescriptions retrieves all prescriptions from MongoDB
func GetPrescriptions() ([]Prescription, error) {
	var prescriptions []Prescription

	collection := DB.Collection("prescriptions")
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var p Prescription
		if err := cursor.Decode(&p); err != nil {
			return nil, err
		}
		prescriptions = append(prescriptions, p)
	}
	return prescriptions, nil
}