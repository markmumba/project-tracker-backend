package models

type User struct {
    ID           uint   `gorm:"primaryKey;autoIncrement"`
    Name         string `json:"name"`
    Email        string `json:"email" gorm:"unique;not null"`
    Password     string `json:"password"`
    RoleID       uint   `gorm:"not null" json:"role_id"`
    Role         Role   `gorm:"foreignKey:RoleID" json:"-"`
    ProfileImage string `json:"profile_image"`
}

type UserDTO struct {
	Id           uint   `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Role         string `json:"role"`
	ProfileImage string `json:"profile_image"`
}

func UserToDTO(u *User) UserDTO {
	return UserDTO{
		Id:           u.ID,
		Name:         u.Name,
		Email:        u.Email,
		Role:         u.Role.Name,
		ProfileImage: u.ProfileImage,
	}
}
func UserToDTOs(users []User) []UserDTO {
	userDTOs := make([]UserDTO, len(users))
	for i, user := range users {
		userDTOs[i] = UserToDTO(&user)
	}
	return userDTOs
}
