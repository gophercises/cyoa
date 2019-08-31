# Exercise #3: Choose your own adventure

[![exercise status: released](https://img.shields.io/badge/exercise%20status-released-green.svg?style=for-the-badge)](https://gophercises.com/exercises/cyoa) [![demo: ->](https://img.shields.io/badge/demo-%E2%86%92-blue.svg?style=for-the-badge)](https://gophercises.com/demos/cyoa/)


## Exercise details

[Choose Your Own Adventure](https://en.wikipedia.org/wiki/Choose_Your_Own_Adventure) is (was?) a series of books intended for children where as you read you would occasionally be given options about how you want to proceed. For instance, you might read about a boy walking in a cave when he stumbles across a dark passage or a ladder leading to an upper level and the reader will be presented with two options like:

- Turn to page 44 to go up the ladder.
- Turn to page 87 to venture down the dark passage.

The goal of this exercise is to recreate this experience via a web application where each page will be a portion of the story, and at the end of every page the user will be given a series of options to choose from (or be told that they have reached the end of that particular story arc).

Stories will be provided via a JSON file with the following format:

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

*See [gopher.json](https://github.com/gophercises/cyoa/blob/master/gopher.json) for a real example of a JSON story. I find that seeing the real JSON file really helps answer any confusion or questions about the JSON format.*

You are welcome to design the code however you want. You can put everything in a single `main` package, or you can break the story into its own package and use that when creating your http handlers.

The only real requirements are:

1. Use the `html/template` package to create your HTML pages. Part of the purpose of this exercise is to get practice using this package.
2. Create an `http.Handler` to handle the web requests instead of a handler function.
3. Use the `encoding/json` package to decode the JSON file. You are welcome to try out third party packages afterwards, but I recommend starting here.

A few things worth noting:

- Stories could be cyclical if a user chooses options that keep leading to the same place. This isn't likely to cause issues, but keep it in mind.
- For simplicity, all stories will have a story arc named "intro" that is where the story starts. That is, every JSON file will have a key with the value `intro` and this is where your story should start.
- Matt Holt's JSON-to-Go is a really handy tool when working with JSON in Go! Check it out - <https://mholt.github.io/json-to-go/>

## Bonus

As a bonus exercises you can also:

1. Create a command-line version of our Choose Your Own Adventure application where stories are printed out to the terminal and options are picked via typing in numbers ("Press 1 to venture ...").
2. Consider how you would alter your program in order to support stories starting form a story-defined arc. That is, what if all stories didn't start on an arc named `intro`? How would you redesign your program or restructure the JSON? This bonus exercises is meant to be as much of a thought exercise as an actual coding one.
