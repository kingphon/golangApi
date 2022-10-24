package database

import "go.mongodb.org/mongo-driver/mongo"

const (
	companyTypes  = "company-types"
	companies     = "companies"
	cabinets      = "cabinets"
	drawers       = "drawers"
	staffs        = "staffs"
	departments   = "departments"
	documents     = "documents"
	documentTypes = "document-types"
	permissions   = "permissions"
)

func CompanyTypeCol() *mongo.Collection {
	return db.Collection(companyTypes)
}
func CompanyCol() *mongo.Collection {
	return db.Collection(companies)
}
func CabinetCol() *mongo.Collection {
	return db.Collection(cabinets)
}
func DrawerCol() *mongo.Collection {
	return db.Collection(drawers)
}
func StaffCol() *mongo.Collection {
	return db.Collection(staffs)
}
func DocumentCol() *mongo.Collection {
	return db.Collection(documents)
}
func DocumentTypeCol() *mongo.Collection {
	return db.Collection(documentTypes)
}
func DepartmentCol() *mongo.Collection {
	return db.Collection(departments)
}
func PermissionCol() *mongo.Collection {
	return db.Collection(permissions)
}
