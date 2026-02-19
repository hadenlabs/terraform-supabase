package supabase

import (
	"github.com/hadenlabs/terraform-supabase/internal/app/external/faker"
)

// Project provides a simple structure for Supabase project testing
type Project struct {
	// OrganizationID is the organization identifier for Supabase projects
	OrganizationID string

	// DatabasePassword is the database password for the project
	DatabasePassword string

	// Name is the project name
	Name string

	// Region is the AWS region for the project
	Region string

	// InstanceSize is the instance size for the project
	InstanceSize string
}

// NewProject creates a new Project instance with default values
// OrganizationID defaults to "hadenlabs", other fields use faker
func NewProject() *Project {
	fake := faker.Project()

	return &Project{
		OrganizationID:   "ysidaatusqmwbbblhrtn",
		DatabasePassword: fake.DatabasePassword(),
		Name:             fake.Name(),
		Region:           fake.Region(),
		InstanceSize:     "micro",
	}
}

// NewProjectWithFaker creates a new Project instance with all fields from faker
func NewProjectWithFaker() *Project {
	fake := faker.Project()

	return &Project{
		OrganizationID:   fake.OrganizationID(),
		DatabasePassword: fake.DatabasePassword(),
		Name:             fake.Name(),
		Region:           fake.Region(),
		InstanceSize:     fake.InstanceSize(),
	}
}

// WithOrganizationID sets a custom organization ID and returns a new Project instance
func (p *Project) WithOrganizationID(orgID string) *Project {
	// Create a new instance to maintain immutability
	return &Project{
		OrganizationID:   orgID,
		DatabasePassword: p.DatabasePassword,
		Name:             p.Name,
		Region:           p.Region,
		InstanceSize:     p.InstanceSize,
	}
}

// WithDatabasePassword sets a custom database password and returns a new Project instance
func (p *Project) WithDatabasePassword(password string) *Project {
	return &Project{
		OrganizationID:   p.OrganizationID,
		DatabasePassword: password,
		Name:             p.Name,
		Region:           p.Region,
		InstanceSize:     p.InstanceSize,
	}
}

// WithName sets a custom project name and returns a new Project instance
func (p *Project) WithName(name string) *Project {
	return &Project{
		OrganizationID:   p.OrganizationID,
		DatabasePassword: p.DatabasePassword,
		Name:             name,
		Region:           p.Region,
		InstanceSize:     p.InstanceSize,
	}
}

// WithRegion sets a custom region and returns a new Project instance
func (p *Project) WithRegion(region string) *Project {
	return &Project{
		OrganizationID:   p.OrganizationID,
		DatabasePassword: p.DatabasePassword,
		Name:             p.Name,
		Region:           region,
		InstanceSize:     p.InstanceSize,
	}
}

// WithInstanceSize sets a custom instance size and returns a new Project instance
func (p *Project) WithInstanceSize(size string) *Project {
	return &Project{
		OrganizationID:   p.OrganizationID,
		DatabasePassword: p.DatabasePassword,
		Name:             p.Name,
		Region:           p.Region,
		InstanceSize:     size,
	}
}

// ToMap converts Project to a map for use with Terraform options
func (p *Project) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"organization_id":         p.OrganizationID,
		"database_password":       p.DatabasePassword,
		"name":                    p.Name,
		"region":                  p.Region,
		"instance_size":           p.InstanceSize,
		"legacy_api_keys_enabled": false, // Default value
		"module_enabled":          true,  // Default value
	}
}

// ToMapWithCustomValues converts Project to a map with custom boolean values
func (p *Project) ToMapWithCustomValues(legacyAPIKeysEnabled, moduleEnabled bool) map[string]interface{} {
	return map[string]interface{}{
		"organization_id":         p.OrganizationID,
		"database_password":       p.DatabasePassword,
		"name":                    p.Name,
		"region":                  p.Region,
		"instance_size":           p.InstanceSize,
		"legacy_api_keys_enabled": legacyAPIKeysEnabled,
		"module_enabled":          moduleEnabled,
	}
}
