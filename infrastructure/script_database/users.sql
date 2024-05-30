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
                               `contrasena` varchar(50) NOT NULL,
                               `sesion_activa` tinyint(1) DEFAULT 1,
                               `direccion` varchar(255) DEFAULT NULL,
                               `imagen_firma` blob NOT NULL,
                               `administrador` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `tblrolisomodulo`
--

CREATE TABLE `tblrolisomodulo` (
  `iId_RolIsoModulo` int(11) NOT NULL,
  `vCodigo` varchar(8) DEFAULT NULL,
  `vDescripcion` varchar(50) DEFAULT NULL,
  `iId_Estado` int(11) DEFAULT NULL COMMENT 'Activo, Inactivo'
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

--
-- Indices de la tabla `tblrolisomodulo`
--
ALTER TABLE `tblrolisomodulo`
  ADD PRIMARY KEY (`iId_RolIsoModulo`);

-- AUTO_INCREMENT de la tabla `tblrolisomodulo`
--
ALTER TABLE `tblrolisomodulo`
  MODIFY `iId_RolIsoModulo` int(11) NOT NULL AUTO_INCREMENT;

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `tblprivilegios`
--

CREATE TABLE `tblprivilegios` (
  `codigo_personalizado` varchar(50) NOT NULL,
  `iId_privilegios` int(8) NOT NULL,
  `nombre` varchar(50) NOT NULL,
  `descripcion` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

--
-- Indices de la tabla `tblprivilegios`
--
ALTER TABLE `tblprivilegios`
  ADD PRIMARY KEY (`iId_privilegios`);

-- AUTO_INCREMENT de la tabla `tblprivilegios`
--
ALTER TABLE `tblprivilegios`
  MODIFY `iId_privilegios` int(8) NOT NULL AUTO_INCREMENT;