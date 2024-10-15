package biz

import (
	kitdi "github.com/jace996/multiapp/pkg/di"
)

// ProviderSet is biz providers.
var ProviderSet = kitdi.NewSet(NewPostSeeder)
