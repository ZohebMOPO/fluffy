## Inspiration
As github has the contribution graph of each user. I thought to build something like that for the local git system of each folder.
## What it does
It collects all the git contributions of a project and displays the timeline of the contributions in console.
## How we built it
I built it using go-git module which makes it easy to interact with git.
## Challenges we ran into
There is an error in windows(in my PC), it stores all the gitcontributions in a dot file namely .gitlocalstatus in the home directory of the current user. There were some administrator problems in overwriting the file. 
## Accomplishments that we're proud of
I am glad that the UI of logging the contributions is nearly same as github contribution UI.
## What we learned
I learned how we organize projects in golang, dynamic types, static types and working with go-git.
## What's next for fluffy
I am really looking forward to build an desktop application and shows the local git contributions in a separate application.
