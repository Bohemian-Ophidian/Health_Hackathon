package models

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Medication struct {
	ID           string bson:"_id,omitempty"
	Name         string bson:"name"
	Alternatives string bson:"alternatives"
	SideEffects  string bson:"side_effects"
	CreatedAt    time.Time bson:"created_at"
	UpdatedAt    time.Time bson:"updated_at"
}

type MedicationModel struct {
	Collection *mongo.Collection
}

func NewMedicationModel(db *mongo.Database) *MedicationModel {
	return &MedicationModel{Collection: db.Collection("medications")}
}

func (m *MedicationModel) GetMedication(id string) (*Medication, error) {
	var medication Medication
	err := m.Collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&medication)
	if err != nil {
		return nil, err
	}
	return &medication, nil
}

func (m *MedicationModel) AddMedication(medication *Medication) error {
	medication.CreatedAt = time.Now()
	medication.UpdatedAt = time.Now()
	_, err := m.Collection.InsertOne(context.TODO(), medication)
	return err
}

func (m *MedicationModel) UpdateMedication(id string, medication *Medication) error {
	medication.UpdatedAt = time.Now()
	_, err := m.Collection.UpdateOne(context.TODO(), bson.M{"_id": id}, bson.M{"$set": medication})
	return err
}

func (m *MedicationModel) DeleteMedication(id string) error {
	_, err := m.Collection.DeleteOne(context.TODO(), bson.M{"_id": id})
	return err
}