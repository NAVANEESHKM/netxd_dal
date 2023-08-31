// this interface will provide the requried methods
package interfaces

import "github.com/NAVANEESHKM/netxd_dal/models"



type ICustomer interface{
	CreateCustomer(user *models.Customer) (*models.Customer, error)
}