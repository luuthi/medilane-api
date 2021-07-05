package seeders

import (
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/yaml.v2"
	"gorm.io/gorm"
	"io/ioutil"
	"medilane-api/config"
	utils2 "medilane-api/core/utils"
	"medilane-api/packages/accounts/builders"
)

type UserSeeder struct {
	DB     *gorm.DB
	config *config.Config
}

type Permissions struct {
	Permissions []*Permission `json:"permissions" yaml:"permissions"`
}

type Permission struct {
	PermissionName string `json:"PermissionName" yaml:"permission_name"`
	Description    string `json:"Description" yaml:"description" `
}

type Roles struct {
	Roles []*Role `json:"roles" yaml:"roles"`
}

type Role struct {
	RoleName    string   `json:"RoleName" yaml:"role_name"`
	Description string   `json:"Description" yaml:"description"`
	Permissions []string `json:"permissions" yaml:"permissions"`
}

type Users struct {
	Users []User `json:"users" yaml:"users"`
}

type User struct {
	Email    string   `json:"Email" yaml:"email" `
	Username string   `json:"Username" yaml:"username"`
	Password string   `json:"Password" yaml:"password" `
	FullName string   `json:"Name" yaml:"full_name"`
	Status   bool     `json:"Confirmed" yaml:"status" `
	Type     string   `json:"Type" yaml:"type" `
	IsAdmin  bool     `json:"IsAdmin" yaml:"is_admin" `
	Roles    []string `json:"Roles" yaml:"roles"`
}

func NewUserSeeder(db *gorm.DB, conf *config.Config) *UserSeeder {
	return &UserSeeder{DB: db, config: conf}
}

func (userSeeder *UserSeeder) LoadInitDataPermission() (permissions *Permissions) {
	configPath := userSeeder.config.MIGRATION.InitPermissionPath
	if configPath == "" {
		configPath = "/app/permissions.yaml"
	}
	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, &permissions)
	if err != nil {
		log.Println("Error loading yaml file")
	}
	return
}

func (userSeeder *UserSeeder) LoadInitDataRole() (roles *Roles) {
	configPath := userSeeder.config.MIGRATION.InitRolePath
	if configPath == "" {
		configPath = "/app/roles.yaml"
	}
	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, &roles)
	if err != nil {
		log.Println("Error loading yaml file")
	}
	return
}

func (userSeeder *UserSeeder) LoadInitDataUser() (users *Users) {
	configPath := userSeeder.config.MIGRATION.InitUserPath
	if configPath == "" {
		configPath = "/app/users.yaml"
	}
	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, &users)
	if err != nil {
		log.Println("Error loading yaml file")
	}
	return
}

func (userSeeder *UserSeeder) SetUsers() {
	users := userSeeder.LoadInitDataUser()

	for _, u := range users.Users {
		encryptedPassword, err := bcrypt.GenerateFromPassword(
			[]byte(u.Password),
			bcrypt.DefaultCost,
		)
		if err != nil {
			continue
		}

		user := builders.NewUserBuilder().
			SetEmail(u.Email).
			SetName(u.Username).
			SetPassword(string(encryptedPassword)).
			SetFullName(u.FullName).
			SetStatus(u.Status).
			SetType(u.Type).
			SetIsAdmin(u.IsAdmin).
			SetRoles(u.Roles).
			Build()

		userSeeder.DB.Table(utils2.TblAccount).Create(&user)
	}
}

func (userSeeder *UserSeeder) SetRoles() {
	roles := userSeeder.LoadInitDataRole()

	for _, r := range roles.Roles {
		role := builders.NewRoleBuilder().
			SetName(r.RoleName).
			SetDescription(r.Description).
			SetPermissions(r.Permissions).
			Build()
		userSeeder.DB.Table(utils2.TblRole).Create(&role)
	}
}

func (userSeeder *UserSeeder) SetPermissions() {
	perms := userSeeder.LoadInitDataPermission()

	for _, p := range perms.Permissions {
		perm := builders.NewPermissionBuilder().
			SetName(p.PermissionName).
			SetDescription(p.Description).
			Build()
		userSeeder.DB.Table(utils2.TblPermission).Create(&perm)
	}
}
