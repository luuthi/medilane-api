package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"medilane-api/config"
)

func Init(cfg *config.Config) *gorm.DB {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Name)

	fmt.Println(dataSourceName)

	db, err := gorm.Open(mysql.Open(dataSourceName), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			NoLowerCase:   false,
		},
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err.Error())
	}

	if cfg.MIGRATION.Migrate {
		//_ = db.SetupJoinTable(&models.Partner{}, "Users", &models.PartnerUser{})
		//_ = db.AutoMigrate(&models.Partner{})
		//	_ = db.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{})
		//
		//	err = db.SetupJoinTable(&models.DrugStore{}, "Users", &models.DrugStoreUser{})
		//	_ = db.SetupJoinTable(&models.DrugStore{}, "Products", &models.DrugStoreProduct{})
		//	if err != nil {
		//		panic(err.Error())
		//	}
		//	_ = db.AutoMigrate(&models.DrugStore{})
		//
		//	_ = db.AutoMigrate(&models.DrugStoreRelationship{})
		//	_ = db.AutoMigrate(&models.Address{})
		//	_ = db.SetupJoinTable(&models.Area{}, "Products", &models.AreaCost{})
		//	_ = db.AutoMigrate(&models.Area{})
		//
		//	_ = db.SetupJoinTable(&models.Product{}, "Variants", &models.VariantValue{})
		//	_ = db.SetupJoinTable(&models.Product{}, "Tags", &models.ProductTag{})
		//	_ = db.SetupJoinTable(&models.Product{}, "Images", &models.ProductImage{})
		//	_ = db.SetupJoinTable(&models.ProductStore{}, "Variants", &models.VariantStoreValue{})
		//	_ = db.SetupJoinTable(&models.ProductStore{}, "Images", &models.ProductStoreImage{})
		//	_ = db.SetupJoinTable(&models.ProductStore{}, "Tags", &models.ProductStoreTag{})
		//	_ = db.AutoMigrate(&models.Image{})
		//	_ = db.AutoMigrate(&models.Product{})
		//	_ = db.AutoMigrate(&models.ProductStore{})
		//	_ = db.AutoMigrate(&models.Category{})
		//
		//	_ = db.AutoMigrate(&models.OrderDetail{})
		//	_ = db.AutoMigrate(&models.PaymentMethod{})
		//	_ = db.AutoMigrate(&models.Order{})
		//
		//	_ = db.AutoMigrate(&models.OrderStoreDetail{})
		//	_ = db.AutoMigrate(&models.OrderStore{})
		//
		//	_ = db.AutoMigrate(&models.DrugStoreConsignment{})
		//	_ = db.AutoMigrate(&models.Consignment{})
		//
		//	_ = db.AutoMigrate(&models.CartDetail{})
		//	_ = db.AutoMigrate(&models.Cart{})
		//
		//	_ = db.AutoMigrate(&models.VoucherDetail{})
		//	_ = db.AutoMigrate(&models.Voucher{})
		//	_ = db.AutoMigrate(&models.PromotionDetail{})
		//	_ = db.AutoMigrate(&models.Promotion{})
		//	_ = db.AutoMigrate(&models.Notification{})
	}

	//if cfg.MIGRATION.Migrate {
	//userSeeder := seeders.NewUserSeeder(db, cfg)
	//userSeeder.SetPermissions()
	//userSeeder.SetRoles()
	//userSeeder.SetUsers()
	//}

	return db
}
