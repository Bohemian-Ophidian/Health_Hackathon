package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Medication represents a medication document in MongoDB
type Medication struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name         string             `bson:"name" json:"name"`
	Alternatives string             `bson:"alternatives" json:"alternatives"`
	SideEffects  string             `bson:"side_effects" json:"side_effects"`
	CreatedAt    time.Time          `bson:"created_at,omitempty" json:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at,omitempty" json:"updated_at"`
}

// MedicationModel handles operations on the medications collection
type MedicationModel struct {
	Collection *mongo.Collection
}

// NewMedicationModel initializes the medication model
func NewMedicationModel(db *mongo.Database) *MedicationModel {
	return &MedicationModel{
		Collection: db.Collection("medications"),
	}
}

// AddMedication inserts a new medication into the database
func (m *MedicationModel) AddMedication(ctx context.Context, medication *Medication) error {
	medication.ID = primitive.NewObjectID()
	medication.CreatedAt = time.Now()
	medication.UpdatedAt = time.Now()

	_, err := m.Collection.InsertOne(ctx, medication)
	return err
}

// UpdateMedication updates an existing medication
func (m *MedicationModel) UpdateMedication(ctx context.Context, id string, medication *Medication) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	update := bson.M{
		"$set": bson.M{
			"name":         medication.Name,
			"alternatives": medication.Alternatives,
			"side_effects": medication.SideEffects,
			"updated_at":   time.Now(),
		},
	}

	_, err = m.Collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	return err
}

// DeleteMedication removes a medication from the database
func (m *MedicationModel) DeleteMedication(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = m.Collection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}

// GetAllMedications retrieves all medications from the database
func (m *MedicationModel) GetAllMedications(ctx context.Context) ([]Medication, error) {
	var medications []Medication

	cursor, err := m.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var medication Medication
		if err := cursor.Decode(&medication); err != nil {
			return nil, err
		}
		medications = append(medications, medication)
	}

	return medications, nil
}
