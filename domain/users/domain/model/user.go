package model

import "time"

type User struct {
	PersonaID     int       `json:"persona_id"`
	Nombre        string    `json:"nombre"`
	Apellido      string    `json:"apellido"`
	Telefono      string    `json:"telefono,omitempty"`
	Celular       string    `json:"celular,omitempty"`
	Correo        string    `json:"correo"`
	Usuario       string    `json:"usuario,omitempty"`
	Contrasena    string    `json:"contrasena"`
	SesionActiva  bool      `json:"sesion_activa"`
	Direccion     string    `json:"direccion,omitempty"`
	ImagendFirma  []byte    `json:"imagen_firma"`
	Administrador int       `json:"administrador,omitempty"`
	DateCreated   time.Time `json:"date_created"`
	DateModified  time.Time `json:"date_modified"`
}
