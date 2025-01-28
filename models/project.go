package models

type Project struct {
	
	ID           uint   `gorm:"primaryKey;autoIncrement"`
	Title       string `gorm:"not null" json:"title"`
	Description string 	`gorm:"not null" json:"description"`
	StartDate   string  `json:"start_date"`
	EndDate     string `json:"end_date"`
	StudentID   uint `gorm:"not null" json:"student_id"`
	LecturerID  uint `gorm:"not null" json:"lecturer_id"`
	Student     User `gorm:"foreignKey:StudentID"`
	Lecturer    User `gorm:"foreignKey:LecturerID"`
}

type ProjectDTO struct {
	ID          uint    `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	StartDate   string  `json:"start_date"`
	EndDate     string  `json:"end_date"`
	Student     UserDTO `json:"student"`
	Lecturer    UserDTO `json:"lecturer"`
}

func ProjectToDTO(p *Project) ProjectDTO {
	return ProjectDTO{
		ID:          p.ID,
		Title:       p.Title,
		Description: p.Description,
		StartDate:   p.StartDate,
		EndDate:     p.EndDate,
		Student:     UserToDTO(&p.Student),
		Lecturer:    UserToDTO(&p.Lecturer),
	}
}

func ProjectToDTOs(projects []Project) []ProjectDTO {
	projectDTOs := make([]ProjectDTO, len(projects))
	for i, project := range projects {
		projectDTOs[i] = ProjectToDTO(&project)
	}
	return projectDTOs
}
