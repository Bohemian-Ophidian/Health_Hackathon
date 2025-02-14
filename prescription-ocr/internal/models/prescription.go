package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Prescription represents a prescription document in MongoDB
type Prescription struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	OriginalImage string             `bson:"original_image" json:"original_image"`
	ExtractedText string             `bson:"extracted_text" json:"extracted_text"`
	CreatedAt     time.Time          `bson:"created_at,omitempty" json:"created_at"`
}

// PrescriptionModel handles operations on the prescriptions collection
type PrescriptionModel struct {
	Collection *mongo.Collection
}

// NewPrescriptionModel initializes the prescription model
func NewPrescriptionModel(db *mongo.Database) *PrescriptionModel {
	return &PrescriptionModel{
		Collection: db.Collection("prescriptions"),
	}
}

// AddPrescription inserts a new prescription into the database
func (m *PrescriptionModel) AddPrescription(ctx context.Context, prescription *Prescription) error {
	prescription.ID = primitive.NewObjectID()
	prescription.CreatedAt = time.Now()

	_, err := m.Collection.InsertOne(ctx, prescription)
	return err
}

// *GetAllPrescriptions - Retrieves all prescriptions from MongoDB*
func (m *PrescriptionModel) GetAllPrescriptions(ctx context.Context) ([]Prescription, error) {
	var prescriptions []Prescription

	cursor, err := m.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var prescription Prescription
		if err := cursor.Decode(&prescription); err != nil {
			return nil, err
		}
		prescriptions = append(prescriptions, prescription)
	}

	return prescriptions, nil
}
