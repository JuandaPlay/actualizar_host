Hosts Editor (Go + Fy_ne)

Este proyecto proporciona una aplicación GUI de Windows (exe) para añadir una entrada al archivo hosts de forma segura (requiere privilegios de administrador).

Requisitos
- Go 1.20+
- Fy_ne (go get fyne.io/fyne/v2)
- Un manifest para elevar privilegios, por ejemplo: hosts-editor.exe.manifest

Cómo compilar (resumen)
- Crear el binario: go build -o hosts-editor.exe ./cmd/hosts_gui
go build -ldflags="-H windowsgui" -o hosts-editor.exe ./cmd/hosts_gui

- Incluir manifest: copiar hosts-editor.exe.manifest junto al ejecutable o usar un recurso RC para incrustarlo
- Ejecutar: ejecutar hosts-editor.exe; Windows mostrará el prompt de UAC para elevación (deberás aceptar)

Notas de seguridad
- No evites el diálogo de UAC; la elevación debe requerirse para modificar el archivo hosts.
- Esta herramienta está destinada a administradores o usuarios con permisos.
- El comportamiento por defecto evita duplicados y coloca la línea justo debajo de entradas localhost.
