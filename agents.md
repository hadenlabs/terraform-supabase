# agents.md - Guía para Agentes de IA

## Información del Proyecto

**Nombre:** terraform-supabase
**Tipo:** Módulo de Terraform para gestionar proyectos de Supabase
**Versión de Terraform:** >= 1.0.0
**Proveedor:** supabase/supabase (v1.7.0)

## Estructura del Proyecto

```
terraform-supabase/
├── modules/project/          # Módulo principal
│   ├── main.tf              # Recursos principales
│   ├── variables.tf         # Variables de entrada
│   ├── outputs.tf           # Outputs del módulo
│   ├── versions.tf          # Versiones de Terraform
│   └── README.md            # Documentación
├── docs/                    # Documentación completa
├── test/                    # Tests de integración
└── Taskfile.yml            # Comandos automatizados
```

## Variables del Módulo Project

### Requeridas:
- `database_password` (string, sensitive) - Contraseña de la base de datos
- `name` (string) - Nombre del proyecto
- `organization_id` (string) - ID de la organización
- `region` (string) - Región del proyecto

### Opcionales:
- `instance_size` (string) - Tamaño de la instancia (default: null)
- `legacy_api_keys_enabled` (bool, deprecated) - Claves API legacy (default: null)

### Configuración:
- `module_enabled` (bool) - Habilitar/deshabilitar módulo (default: true)

## Comandos Principales

```bash
# Formatear y validar
task terraform:fmt
task terraform:validate

# Generar documentación
task readme

# Ejecutar tests
task test
task module_project:test

# Instalar dependencias
task setup
```

## Reglas de Desarrollo

1. **Variables sensibles**: Siempre marcar con `sensitive = true`
2. **Conditional resources**: Usar `count = var.module_enabled ? 1 : 0`
3. **Documentación**: Incluir descripciones en todas las variables y outputs
4. **Formato**: Usar `terraform fmt` antes de commit
5. **Testing**: Usar Terratest para tests de integración

## Ejemplo de Uso

```hcl
module "supabase_project" {
  source = "./modules/project"

  database_password = "secure-password"
  name              = "my-project"
  organization_id   = "org-123"
  region            = "us-east-1"

  instance_size = "small"
  module_enabled = true
}
```

## Recursos Clave

- **README principal**: `README.md`
- **Documentación completa**: `docs/`
- **Taskfile**: `Taskfile.yml` (comandos automatizados)
- **Variables actuales**: `modules/project/variables.tf`

## Notas para Agentes

- Este es un módulo de Terraform, no una aplicación web
- Usar Taskfile para comandos automatizados
- Las variables opcionales usan `null` como default para heredar valores del proveedor
- El módulo está diseñado para ser reutilizable y configurable