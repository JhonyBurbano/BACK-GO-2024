package persistence

const (
	InsertUser = `INSERT INTO tblpersonas (
        persona_id,                 
        nombre,
        apellido,
        telefono,
        celular,
        correo,
        usuario,
        contrasena,
        sesion_activa,
        direccion,
        imagen_firma,
        administrador
    ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	SelectUser = `SELECT nombre, persona_id, correo, contrasena, celular, apellido, direccion
    FROM tblpersonas
    WHERE persona_id = ?`

	SelectLoginUser = `SELECT correo, contrasena, celular FROM tblpersonas WHERE correo = ? OR celular = ?`

	SelectUsers = `SELECT nombre, persona_id, correo, contrasena, celular, apellido, direccion
    FROM tblpersonas`
)
