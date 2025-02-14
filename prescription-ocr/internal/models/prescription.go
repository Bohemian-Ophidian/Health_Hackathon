package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Prescription represents a prescription document
type Prescription struct {
	ID            primitive.ObjectID bson:"_id,omitempty" json:"id"
	OriginalImage string             bson:"original_image" json:"original_image"
	ExtractedText string             bson:"extracted_text" json:"extracted_text"
	CreatedAt     time.Time          bson:"created_at,omitempty" json:"created_at"
}

// PrescriptionModel handles database operations
type PrescriptionModel struct {
	Collection *mongo.Collection
}

// NewPrescriptionModel initializes the Prescription model
func NewPrescriptionModel(db *mongo.Database) *PrescriptionModel {
	return &PrescriptionModel{Collection: db.Collection("prescriptions")}
}

// SavePrescription inserts a new prescription
func (m *PrescriptionModel) SavePrescription(ctx context.Context, p *Prescription) error {
	p.ID = primitive.NewObjectID()
	p.CreatedAt = time.Now()

	_, err := m.Collection.InsertOne(ctx, p)
	return err
}

// GetPrescriptions retrieves all prescriptions
func (m *PrescriptionModel) GetPrescriptions(ctx context.Context) ([]Prescription, error) {
	var prescriptions []Prescription

	cursor, err := m.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &prescriptions); err != nil {
		return nil, err
	}

	return prescriptions, nil
}