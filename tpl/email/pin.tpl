<!DOCTYPE html>
<html lang="es">
    <head>
        <meta charset="utf-8">
	    <title>Has recibido un PIN de parte de Gonitor</title>
        <style>
            body {
                font-family: verdana; 
                padding-top: 0px; 
                background-color: transparent; 
            }
            div.container { 
                padding: 20px; 
                background-color: transparent;
            }            
            .button {
                border: none;
                color: white;
                padding: 15px 32px;
                text-align: center;
                text-decoration: none;
                display: inline-block;
                font-size: 16px;
                margin: 4px 2px;
                cursor: pointer;
                background-color: #4CAF50;
            }
            .button:hover {
                background-color: #008CBA;
            }
        </style>
    </head>
    <body>
	    <div class="container">
            <h4>Has recibido un PIN de parte de GONITOR</h4>
            <p>
                El PIN que aparece en el recuadro de abajo sirve para recuperar 
                contraseña y validar correo electrónico.
            </p>
            <p>
                El PIN expira dentro de los próximos 30 minutos, y solo 
                podrá usarse una vez. Siempre puedes solicitar un nuevo PIN si 
                este ya no es válido.
            </p>
            <div class="button">{{.PIN}}</div>
        </div>
    </body>
</html>