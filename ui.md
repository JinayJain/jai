UI System

Manager

-   keeps track of all windows in the program
-   takes in input and either delegates it or consumes some "global command"
-   possibly maintains a flag for whether it's in some special mode like command mode
    -   nonstandard behavior would be triggered

Editor

Status Bar

-   what does the status bar show?
    -   for the focused editor:
        -   current file name (and info)
        -   cursor position
    -   command mode
        -   status bar would need to communicate with the manager
-   one status bar per manager
-   how does the status bar receive data to render?
    -   does the manager provide that data?