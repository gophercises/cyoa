# Exercise #3: Elige tu propia aventura

[![exercise status: released](https://img.shields.io/badge/exercise%20status-released-green.svg?style=for-the-badge)](https://gophercises.com/exercises/cyoa) [![demo: ->](https://img.shields.io/badge/demo-%E2%86%92-blue.svg?style=for-the-badge)](https://gophercises.com/demos/cyoa/)


## Exercise details

[Choose Your Own Adventure](https://en.wikipedia.org/wiki/Choose_Your_Own_Adventure) es (¿fue?) una serie de libros destinados a niños donde, a medida que lees, ocasionalmente se te darán opciones sobre cómo quieres proceder. Por ejemplo, puede leer sobre un niño que camina en una cueva cuando tropieza con un pasaje oscuro o una escalera que conduce a un nivel superior y el lector tendrá dos opciones como:

- Pase a la página 44 para subir la escalera.
- Pase a la página 87 para aventurarse por el oscuro pasaje.

El objetivo de este ejercicio es recrear esta experiencia a través de una aplicación web donde cada página será una parte de la historia, y al final de cada página se le dará al usuario una serie de opciones para elegir (o se le dirá que han llegado al final de ese arco de historia en particular).

Las historias se proporcionarán a través de un archivo JSON con el siguiente formato:

```json
{
  // Each story arc will have a unique key that represents
  // the name of that particular arc.
  "story-arc": {
    "title": "A title for that story arc. Think of it like a chapter title.",
    "story": [
      "A series of paragraphs, each represented as a string in a slice.",
      "This is a new paragraph in this particular story arc."
    ],
    // Options will be empty if it is the end of that
    // particular story arc. Otherwise it will have one or
    // more JSON objects that represent an "option" that the
    // reader has at the end of a story arc.
    "options": [
      {
        "text": "the text to render for this option. eg 'venture down the dark passage'",
        "arc": "the name of the story arc to navigate to. This will match the story-arc key at the very root of the JSON document"
      }
    ]
  },
  ...
}
```

*Vea [gopher.json](https://github.com/gophercises/cyoa/blob/master/gopher.json) para un ejemplo real de una historia JSON. Me parece que ver el archivo JSON real realmente ayuda a responder cualquier confusión o pregunta sobre el formato JSON. *

Puedes diseñar el código como quieras. Puede poner todo en un solo paquete `main`, o puede dividir la historia en su propio paquete y usarlo al crear sus controladores http.

Los únicos requisitos reales son:

1. Use el paquete `html / template` para crear sus páginas HTML. Parte del propósito de este ejercicio es practicar con este paquete.
2. Cree un `http.Handler` para manejar las solicitudes web en lugar de una función de controlador.
3. Use el paquete `encoding / json` para decodificar el archivo JSON. Puede probar paquetes de terceros posteriormente, pero le recomiendo comenzar aquí.

Algunas cosas que vale la pena señalar:

- Las historias pueden ser cíclicas si un usuario elige opciones que siguen conduciendo al mismo lugar. No es probable que esto cause problemas, pero tenlo en cuenta.
- Para simplificar, todas las historias tendrán un arco de historia llamado "introducción", que es donde comienza la historia. Es decir, cada archivo JSON tendrá una clave con el valor `intro` y aquí es donde debe comenzar su historia.
- JSON-to-Go de Matt Holt es una herramienta realmente útil cuando se trabaja con JSON en Go! Compruébalo - <https://mholt.github.io/json-to-go/>

## Bonus

Como ejercicios extra también puedes:

1. Cree una versión de línea de comandos de nuestra aplicación Choose Your Own Adventure donde las historias se imprimen en el terminal y las opciones se seleccionan escribiendo números ("Presione 1 para aventurarse ...").
2. Considere cómo alteraría su programa para apoyar las historias que comienzan desde un arco definido por la historia. Es decir, ¿qué pasa si todas las historias no comienzan en un arco llamado 'introducción'? ¿Cómo rediseñaría su programa o reestructuraría el JSON? Este ejercicio de bonificación está destinado a ser tanto un ejercicio de pensamiento como uno de codificación real.
