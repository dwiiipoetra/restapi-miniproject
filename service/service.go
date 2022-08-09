package service

import (
	"net/http"

	"strconv"
	"time"

	"github.com/dwiiipoetra/restapi-miniproject/middleware"
	"github.com/dwiiipoetra/restapi-miniproject/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Motorcycle struct {
	// ModelName   string `json:"model_name" validate:"required"`
	// MachineType string `json:"machine_type" validate:"required"`
	// Year        int    `json:"year" validate:"required"`
	// Color       string `json:"color" validate:"required"`
	ModelName   string `json:"model_name"`
	MachineType string `json:"machine_type"`
	Year        int    `json:"year"`
	Color       string `json:"color"`
}

type Repository struct {
	DB *gorm.DB
}

const SecretKey = "celerates"

func (r *Repository) Login(context *fiber.Ctx) error {
	var data map[string]string

	if err := context.BodyParser(&data); err != nil {
		return err
	}

	var user models.Users

	r.DB.Where("email = ?", data["email"]).First(&user)
	// r.DB.Where("email = ? AND password = ?", data["email"], data["password"]).Find(&user)
	if user.ID == 0 {
		context.Status(fiber.StatusNotFound)
		return context.JSON(fiber.Map{
			"message": "invalid credentials",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		context.Status(fiber.StatusBadRequest)
		return context.JSON(fiber.Map{
			"message": "invalid password",
		})
	}

	// create token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		ExpiresAt: jwt.NewNumericDate(time.Unix(time.Now().Add(time.Hour*1).Unix(), 0)), // 1 hour
	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		context.Status(fiber.StatusInternalServerError)
		return context.JSON(fiber.Map{
			"message": "could not login",
		})
	}

	context.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "login successfully",
		"token":   token,
	})
	return nil
}

func (r *Repository) Register(context *fiber.Ctx) error {
	// user := models.Users{}

	// err := context.BodyParser(&user)
	// if err != nil {
	// 	context.Status(http.StatusUnprocessableEntity).JSON(
	// 		&fiber.Map{"message": "request failed"})
	// 	return err
	// }

	// password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	// updatedUser := models.Users{
	// 	Name:     user.Name,
	// 	Email:    user.Email,
	// 	Password: password,
	// }

	// err = r.DB.Create(&updatedUser).Error

	// if err != nil {
	// 	context.Status(http.StatusBadRequest).JSON(
	// 		&fiber.Map{"message": "could not create user"})
	// 	return err
	// }

	// context.Status(http.StatusOK).JSON(
	// 	&fiber.Map{"message": "user succesfully created"})
	// return nil

	var data map[string]string

	err := context.BodyParser(&data)
	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	user := models.Users{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}

	err = r.DB.Create(&user).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create user"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "user succesfully created",
		"data":    user})
	return nil
}

func (r *Repository) CreateMotorcycle(context *fiber.Ctx) error {
	motorcycle := Motorcycle{}

	err := context.BodyParser(&motorcycle)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}

	err = r.DB.Create(&motorcycle).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create motorcycle"})
		return err
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{
			"message": "motorcycle succesfully created",
			"data":    motorcycle,
		})
	return nil
}

func (r *Repository) UpdateMotorcycle(context *fiber.Ctx) error {
	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}

	var motorcycleModel models.Motorcycles //struct in model
	var motorcycle Motorcycle              //struct in service

	err := context.BodyParser(&motorcycle)
	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "request failed",
		})
		return err
	}

	err = r.DB.Model(motorcycleModel).Where("id = ?", id).Updates(motorcycle).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not update motorcycle",
		})
		return err
	}

	updateMotorcycle := models.Motorcycles{
		ModelName:   motorcycle.ModelName,
		MachineType: motorcycle.MachineType,
		Year:        motorcycle.Year,
		Color:       motorcycle.Color,
	}
	r.DB.Model(&motorcycleModel).Take(&updateMotorcycle)
	// r.DB.Model(&motorcycleModel).Updates(&updateMotorcycle)

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "motorcycle successfully updated",
	})
	return nil
}

func (r *Repository) DeleteMotorcycle(context *fiber.Ctx) error {
	motorcycleModel := &models.Motorcycles{}

	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}

	err := r.DB.Delete(motorcycleModel, id)

	if err.Error != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not delete motorcycle",
		})
		return err.Error
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "motorcycle successfully deleted",
	})
	return nil
}

func (r *Repository) GetMotorcycles(context *fiber.Ctx) error {
	motorcyclesModel := &[]models.Motorcycles{}

	err := r.DB.Find(motorcyclesModel).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get motorcycles"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "successfully get motorcycles",
		"data":    motorcyclesModel,
	})
	return nil
}

func (r *Repository) GetMotorcycle(context *fiber.Ctx) error {
	id := context.Params("id")
	motorcycleModel := &models.Motorcycles{}
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}

	err := r.DB.Where("id = ?", id).First(motorcycleModel).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get motorcycle"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "successfully get motorcycle",
		"data":    motorcycleModel,
	})
	return nil
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/login", r.Login)
	api.Post("/register", r.Register)
	api.Post("/create_motorcycle", middleware.Auth, r.CreateMotorcycle)
	// api.Post("/create_motorcycle", r.CreateMotorcycle)
	api.Put("/update_motorcycle/:id", r.UpdateMotorcycle)
	api.Delete("/delete_motorcycle/:id", r.DeleteMotorcycle)
	api.Get("/motorcycle/:id", r.GetMotorcycle)
	api.Get("/motorcycles", r.GetMotorcycles)
}
