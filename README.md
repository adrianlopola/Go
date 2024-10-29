# Verificación de Edad para Licencia de Conducir

Este programa en Go solicita la edad del usuario y determina si cumple con la edad mínima (18 años) para obtener una licencia de conducir.

## Instalación

1. **Instala Go** desde [golang.org](https://golang.org/dl/) y sigue las instrucciones de instalación.
2. **Instala VSCode** desde [code.visualstudio.com](https://code.visualstudio.com/) y agrega la extensión de Go.

## Ejecución del Programa

Para ejecutar el programa, abre `main.go` en VSCode y presiona el botón de **Play**, o usa el siguiente comando en la terminal:

```bash
go run main.go
```
## Pruebas Unitarias
El archivo main_test.go contiene pruebas unitarias para validar el funcionamiento del programa. Se incluyen tres pruebas:

1. **Prueba de edad mayor o igual a 18**: verifica que el programa permita obtener el carnet para edades de 18 o más.
2. **Prueba de edad menor a 18**: verifica que el programa niegue el carnet para edades menores de 18.
3. **Prueba con edad negativa**: esta prueba está diseñada para fallar, mostrando cómo el programa no maneja este caso de entrada no válida.
## Ejecución de las Pruebas
Para ejecutar las pruebas unitarias, en una terminal de VSCode usa el siguiente comando:

```bash
go test -v
```
Este comando mostrará el resultado detallado de cada prueba, indicando cuáles pasaron y cuáles fallaron.