# Proyecto de Go con HTTP para servir archivos estáticos

Este es un proyecto básico de Go que utiliza el paquete `net/http` para crear un servidor web que sirve archivos estáticos.

## Requisitos previos

- Go (instalado en tu sistema)
- Acceso a Internet (para descargar las dependencias)

## Configuración del proyecto

1. Clona este repositorio en tu máquina local o descarga los archivos del proyecto.

2. Asegúrate de tener una carpeta llamada `public` en el directorio raíz del proyecto. Esta carpeta contendrá tus archivos HTML, CSS, JavaScript y otros recursos estáticos.

3. Asegúrate de tener Go instalado y configurado correctamente en tu sistema.

## Ejecución del proyecto

1. Abre una terminal o línea de comandos y navega hasta el directorio raíz del proyecto.

2. Ejecuta el siguiente comando para compilar y ejecutar el proyecto:

        go run main.go
        
3. El servidor se iniciará en el puerto 8080. Abre tu navegador web y navega a `http://localhost:8080`. Verás que se muestra el archivo "index.html" desde la carpeta "public".

4. Puedes editar los archivos HTML y JavaScript en la carpeta "public" según tus necesidades. Los cambios se reflejarán automáticamente al recargar la página en el navegador.

## Personalización

Si deseas agregar más archivos estáticos o rutas adicionales, puedes modificar el archivo "main.go" según tus necesidades. Consulta la documentación de `net/http` de Go para obtener más información sobre cómo personalizar el servidor web.

¡Disfruta de tu servidor web de Go con archivos estáticos!
