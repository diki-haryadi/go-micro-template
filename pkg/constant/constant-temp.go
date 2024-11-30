package constant

type (
	PassAlgorithm string

	DB string

	IsDeleted int

	ResourcesType int

	KeyID string

	ContextKey string

	CacheKey string

	Gender int
)

var (
	Bcrypt PassAlgorithm = "bcrypt"
	MD5    PassAlgorithm = "md5"
	Argon  PassAlgorithm = "argon"
	SHA    PassAlgorithm = "sha"

	MySQL      DB = "mysql"
	PostgreSQL DB = "postgres"

	False IsDeleted = 0
	True  IsDeleted = 1

	Menu ResourcesType = 1
	API  ResourcesType = 2

	Default KeyID = "default"

	Claim ContextKey = "claim"

	UserAuth      CacheKey = "auth::user-uid:%s"
	RoleMenu      CacheKey = "resources::role-uid:%s:type:%d"
	JWKPrivateKey CacheKey = "jwk::private-key:%s"
	MenuResource  CacheKey = "resources::type:%d"

	Female Gender = 0
	Male   Gender = 1
)

func (pa PassAlgorithm) String() string {
	return string(pa)
}

func (db DB) String() string {
	return string(db)
}

func (i IsDeleted) Int() int {
	return int(i)
}

func (k KeyID) String() string {
	return string(k)
}

func (rt ResourcesType) Int() int {
	return int(rt)
}

func (rt ResourcesType) String() string {
	switch rt.Int() {
	case Menu.Int():
		return "menu"
	case API.Int():
		return "api"
	default:
		return " "
	}
}

func ResourceTypeAtoi(s string) ResourcesType {
	switch s {
	case "menu":
		return Menu
	case "api":
		return API
	default:
		return 0
	}
}

func (ct ContextKey) String() string {
	return string(ct)
}

func (g Gender) Int() int {
	return int(g)
}

func (g Gender) String() string {
	switch g.Int() {
	case Female.Int():
		return "female"
	case Male.Int():
		return "male"
	default:
		return ""
	}
}
