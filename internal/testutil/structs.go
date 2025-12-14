package testutil

// Oidc is the kind of module
type Oidc string

// Servers Types.
const (
	Company          Oidc = "evilcorp"
	AuthorizationURL Oidc = "https://keycloak.infosis.tech/realms/testing-realm/broker/google/endpoint"

	ClientID     Oidc = "6thistest15137lj54uk1e.apps.googleusercontent.com" // #nosec G101
	ClientSecret Oidc = "my-secret"
	TokenURL     Oidc = ""
	HostedDomain Oidc = "evilcorp.tech"
)
