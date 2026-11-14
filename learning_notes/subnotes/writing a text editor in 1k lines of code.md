# Kilo editor

[Kilo editor on GitHub](https://github.com/antirez/kilo) · [Background article](https://antirez.com/news/108) · [Kilo implementation tutorial](https://viewsourcecode.org/snaptoken/kilo/)

Related links:
[Article discussing interesting projects a developer should implement](https://austinhenley.com/blog/challengingprojects.html)

### Text Editor

We use text editors every day, but do you know how they really work? Ignoring
all of the fancy features in your favorite editor, how would you implement a
textbox that supports a movable text cursor and selecting, inserting, and
deleting text? You can't use the built-in textbox component from your favorite
GUI framework!

The biggest challenge is figuring out how to store the text document in memory.
My first thought was to use an array, but that has horrible performance if the
user inserts text anywhere other than the end of the document. Luckily, there
are some nice data structures that can solve this.

Another hurdle was learning how a text cursor behaves in popular editors. For
example, if I press the up arrow key with the cursor in the middle of the
document, where will it move? To the same column? Not if that line is shorter.
Keep pressing up, and the cursor will snap back to the original column once a
line is long enough. It turns out that the cursor remembers the column and
tries to get back to it. These are details I never noticed until I tried to
implement them.

![Text cursor behavior](images/textcursor.gif)

After implementing the basic editor, try adding two more features: undo/redo and
word wrapping. Implementing undo/redo efficiently was mind-blowing to me! I
first tried keeping an array of previous states, then tried the Memento pattern,
before finally settling on the Command pattern. Word wrapping forces you to
separate the _visual_ aspects of a text line from its memory aspects.

Things to learn:

- Data structures for storing text: array, [rope](<https://en.wikipedia.org/wiki/Rope_(data_structure)>), [gap buffer](https://en.wikipedia.org/wiki/Gap_buffer), and [piece table](https://en.wikipedia.org/wiki/Piece_table).
- Behavior and implementation of the text cursor.
- Design patterns for undo/redo: [Memento](https://en.wikipedia.org/wiki/Memento_pattern) and [Command](https://en.wikipedia.org/wiki/Command_pattern).
- Abstractions that separate the visual and memory aspects of text.

Further reading:

- [Text Editor: Data Structures](https://www.averylaird.com/programming/the%20text%20editor/2017/09/30/the-piece-table/)
- [Design and Implementation of a Win32 Text Editor](http://www.catch22.net/tuts/neatpad#)
- [Data Structures and Algorithms in Java](https://amzn.to/36k1kEv)
