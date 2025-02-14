package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Medication struct for MongoDB
type Medication struct {
	ID           string bson:"_id,omitempty" json:"id"
	Name         string bson:"name" json:"name"
	Alternatives string bson:"alternatives" json:"alternatives"
	SideEffects  string bson:"side_effects" json:"side_effects"
}

// MedicationModel struct
type MedicationModel struct {
	DB *mongo.Database
}

// AddMedication inserts a new medication
func (m *MedicationModel) AddMedication(medication *Medication) error {
	collection := m.DB.Collection("medications")
	_, err := collection.InsertOne(context.TODO(), medication)
	return err
}

// GetMedication fetches a medication by name
func (m *MedicationModel) GetMedication(name string) (*Medication, error) {
	collection := m.DB.Collection("medications")
	var medication Medication
	err := collection.FindOne(context.TODO(), bson.M{"name": name}).Decode(&medication)
	if err != nil {
		return nil, err
	}
	return &medication, nil
}