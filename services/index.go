package services

import (
	"context"

	"fmt"

	"github.com/NAVANEESHKM/netxd_dal/interfaces"
	"github.com/NAVANEESHKM/netxd_dal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CustomerService struct {
	CustomerCollection *mongo.Collection
	ctx                context.Context
}

func InitCustomerService(collection *mongo.Collection, ctx context.Context) interfaces.ICustomer {
	return &CustomerService{collection, ctx}
}

func (p *CustomerService) CreateCustomer(user *models.Customer) (*models.Customer, error) {

	res, err := p.CustomerCollection.InsertOne(p.ctx, &user)

	if err != nil {

		return nil, err
	}

	var insertedCustomer models.Customer

	insertedID, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {

		return nil, fmt.Errorf("InsertOne did not return an ObjectID")
	}

	// Use the InsertedID to fetch the inserted document.
	filter := bson.M{"Customer_ID": insertedID}
	err = p.CustomerCollection.FindOne(p.ctx, filter).Decode(&insertedCustomer)
	if err != nil {
		// Handle the error
		return nil, err
	}

	// Return the insertedCustomer as a pointer.
	return &insertedCustomer, nil
}
