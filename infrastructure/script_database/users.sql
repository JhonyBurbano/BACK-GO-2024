--
-- Estructura de tabla para la tabla `sesiones`
--

CREATE TABLE `sesiones` (
                            `sesion_id` int(8) NOT NULL,
                            `persona_id` int(8) NOT NULL,
                            `token` binary(64) NOT NULL,
                            `caducidad` text NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `tblpersonas`
--

CREATE TABLE `tblpersonas` (
                               `persona_id` int(11) NOT NULL,
                               `nombre` varchar(100) NOT NULL,
                               `apellido` varchar(50) NOT NULL,
                               `telefono` varchar(20) DEFAULT NULL,
                               `celular` varchar(20) DEFAULT NULL,
                               `correo` varchar(100) NOT NULL,
                               `usuario` varchar(50) DEFAULT NULL,
                               `contrasena` binary(64) NOT NULL,
                               `sesion_activa` tinyint(1) DEFAULT 1,
                               `direccion` varchar(255) DEFAULT NULL,
                               `imagen_firma` blob NOT NULL,
                               `administrador` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;