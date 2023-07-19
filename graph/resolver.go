package graph

import (
	"github.com/clerkinc/clerk-sdk-go/clerk"
	"gorm.io/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	ClerkClient *clerk.Client
	DB          *gorm.DB
}
