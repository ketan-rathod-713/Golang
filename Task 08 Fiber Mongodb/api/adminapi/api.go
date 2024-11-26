package adminapi

import (
	"fibermongoapp/app"
	"fibermongoapp/app/adminservice"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

/* wrapper around all user related api's */
type adminApi struct {
	App      *app.App
	Service  adminservice.Service
	Validate *validator.Validate
}

func New(app *app.App) *adminApi {
	return &adminApi{App: app, Service: adminservice.New(app), Validate: validator.New(validator.WithRequiredStructEnabled())}
}

// /bacancyschool
func (u *adminApi) Routes(parentRouter fiber.Router) {

	// create teacher
	parentRouter.Post("/teacher", u.AddTeacher)
	parentRouter.Get("/subject/:subjectId/teacher/:teacherId", u.AddTeacherToSubject)

	parentRouter.Post("/subject", u.CreateSubject) // Create a new class
	//TODO parentRouter.Delete("/subject", u.CreateSubject) // Create a new class

	parentRouter.Post("/class", u.CreateClass) // without adding any subject at start
	parentRouter.Get("/class", u.GetClass)

	parentRouter.Post("/class/:classId/subject/:subjectId", u.AddSubjectToClass)        // without adding any subject at start
	parentRouter.Delete("/class/:classId/subject/:subjectId", u.DeleteSubjectFromClass) // delete subjectId From the class

	// ? Different Routes
	// Now add teacher

	/*
		1. Registration
		2. Query registration
		3. Create Class
		4. Get all class
		5. create subject
		6. get all subjects
		7. add or remove subjects from class
		8. get all students
		9. get all students enrolled for this class
	*/
}
