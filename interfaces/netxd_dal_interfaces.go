// this interface will provide the requried methods
package interfaces

import "netxd_grpc_mongo/netxd_dal/models"



type ICustomer interface{
	CreateCustomer(user *models.Customer) (*models.Customer, error)
}