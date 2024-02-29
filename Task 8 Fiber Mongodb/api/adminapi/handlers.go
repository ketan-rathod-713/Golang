package adminapi

import (
	"fibermongoapp/models"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (a *adminApi) CreateSubject(ctx *fiber.Ctx) error {
	var subject *models.Subject
	err := ctx.BodyParser(&subject)

	// TODO add proper response
	if err != nil {
		log.Print("Error body parsing")
		return ctx.Status(400).JSON(models.Response{Status: 400, Message: "error", Data: nil})
	}

	err = a.Service.CreateSubject(subject)
	if err != nil {
		log.Print("Error body parsing")
		return ctx.Status(400).JSON(models.Response{Status: 400, Message: "error", Data: nil})
	}

	return ctx.Status(400).JSON(models.Response{Status: 400, Message: "error", Data: &fiber.Map{"subject_added": subject}})

}

func (a *adminApi) CreateClass(ctx *fiber.Ctx) error {
	var class *models.Class
	// Intially keep empty array

	err := ctx.BodyParser(&class)
	class.Subjects = make([]models.Subject, 0)

	// TODO add proper response
	if err != nil {
		log.Print("Error body parsing")
		return ctx.Status(400).JSON(models.Response{Status: 400, Message: "error", Data: nil})
	}

	err = a.Service.CreateClass(class)
	if err != nil {
		log.Print("Error body parsing")
		return ctx.Status(400).JSON(models.Response{Status: 400, Message: "error", Data: nil})
	}

	return ctx.Status(400).JSON(models.Response{Status: 400, Message: "error", Data: &fiber.Map{"subject_added": class}})
}

func (a *adminApi) AddTeacher(ctx *fiber.Ctx) error {
	log.Println("cal")
	var teacher *models.Teacher

	err := ctx.BodyParser(&teacher)

	// TODO add proper response
	if err != nil {
		log.Print("Error body parsing")
		return ctx.Status(400).JSON(models.Response{Status: 400, Message: "error", Data: nil})
	}

	err = a.Service.CreateTeacher(teacher)
	if err != nil {
		log.Print("Error body parsing")
		return ctx.Status(400).JSON(models.Response{Status: 400, Message: "error", Data: nil})
	}

	return ctx.Status(400).JSON(models.Response{Status: 400, Message: "error", Data: &fiber.Map{"teacher_added": teacher}})
}

// input : subject_id and params classId
func (a *adminApi) AddSubjectToClass(ctx *fiber.Ctx) error {

	ClassId := ctx.Params("classId")
	SubjectId := ctx.Params("subjectId")

	classId, _ := primitive.ObjectIDFromHex(ClassId)
	subjectId, _ := primitive.ObjectIDFromHex(SubjectId)

	err := a.Service.AddSubjectToClass(classId, subjectId)

	if err != nil {
		log.Print("Error")
		return ctx.Status(400).JSON(models.Response{Status: 400, Message: "error", Data: &fiber.Map{"error": err.Error()}})
	}

	return ctx.Status(400).JSON(models.Response{Status: 400, Message: "error", Data: &fiber.Map{"subject": ClassId, "class": classId}})
}

func (a *adminApi) AddTeacherToSubject(ctx *fiber.Ctx) error {

	TeacherId := ctx.Params("teacherId")
	SubjectId := ctx.Params("subjectId")

	log.Println(TeacherId, SubjectId)

	teacherId, _ := primitive.ObjectIDFromHex(TeacherId)
	subjectId, _ := primitive.ObjectIDFromHex(SubjectId)

	err := a.Service.AddTeacherToSubject(teacherId, subjectId)

	if err != nil {
		log.Print("Error")
		return ctx.Status(400).JSON(models.Response{Status: 400, Message: "error", Data: &fiber.Map{"error": err.Error()}})
	}

	return ctx.Status(400).JSON(models.Response{Status: 400, Message: "error", Data: &fiber.Map{"teacher": teacherId, "subject": subjectId}})
}

func (a *adminApi) DeleteSubjectFromClass(ctx *fiber.Ctx) error {

	ClassId := ctx.Params("classId")
	SubjectId := ctx.Params("subjectId")

	classId, _ := primitive.ObjectIDFromHex(ClassId)
	subjectId, _ := primitive.ObjectIDFromHex(SubjectId)

	err := a.Service.DeleteSubjectFromClass(classId, subjectId)

	if err != nil {
		log.Print("Error")
		return ctx.Status(400).JSON(models.Response{Status: 400, Message: "error", Data: &fiber.Map{"error": err.Error()}})
	}

	return ctx.Status(400).JSON(models.Response{Status: 400, Message: "error", Data: &fiber.Map{"subject": ClassId, "class": classId}})
}

func (a *adminApi) GetClass(ctx *fiber.Ctx) error {
	classes, err := a.Service.GetClass()

	if err != nil {
		log.Print("Error")
		return ctx.Status(400).JSON(models.Response{Status: 400, Message: "error", Data: &fiber.Map{"error": err.Error()}})
	}

	return ctx.Status(400).JSON(models.Response{Status: 400, Message: "error", Data: &fiber.Map{"classes": classes}})
}
