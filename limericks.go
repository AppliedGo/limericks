/*
<!--
Copyright (c) 2019 Christoph Berger. Some rights reserved.

Use of the text in this file is governed by a Creative Commons Attribution Non-Commercial
Share-Alike License that can be found in the LICENSE.txt file.

Use of the code in this file is governed by a BSD 3-clause license that can be found
in the LICENSE.txt file.

The source code contained in this file may import third-party source code
whose licenses are provided in the respective license files.
-->

<!--
NOTE: The comments in this file are NOT godoc compliant. This is not an oversight.

Comments and code in this file are used for describing and explaining a particular topic to the reader. While this file is a syntactically valid Go source file, its main purpose is to get converted into a blog article. The comments were created for learning and not for code documentation.
-->

+++
title = "Poetry time! Go proverbs as limericks"
description = "Do you need explanations for Rob Pike's Go proverbs? Try this limerick version."
author = "Christoph Berger"
email = "chris@appliedgo.net"
date = "2022-12-08"
draft = "false"
categories = ["Patterns and Paradigms"]
tags = ["proverbs", "best-practices", "poetry"]
articletypes = ["Odds And Ends"]
+++


The Go proverbs capture the essence of Go. Too concise for your taste? No worries, here is each proverb explained in a limerick.

<!--more-->

# Go Proverbs, Explained in Limericks
[*Simple, Poetic, Pithy*](https://go-proverbs.github.io/)

> With a (not so) tiny bit of help from ChatGPT.
>
> My sincere apologies to Edward Lear.

## Don't communicate by sharing memory, share memory by communicating.

Once there was a Go programmer\
Whose code was quite a bummer\
But he followed the rules\
Of the Go proverbs, and jewels\
Of wisdom were soon discovered.

## Concurrency is not parallelism.

<!--There once was a programmer named Bress,\
Who thought concurrency was best,\
But parallelism he found,\
Was a completely different ground,\
And his code finally ran with finesse.

There once was a concurrency rule\
That said it's not parallelism cool\
So don't be confused\
Or your code will be bruis-ed\
Just use it for multitasking, you fool!

There once was a language called Go,
Whose concurrency was not parallel-o,
It may seem the same,
But there's a big difference to tame,
So listen up and don't be a dillo.

There once was a programmer named P.
He thought concurrency was parallelism, oh boy.
But soon he learned that's not true,
And now his code runs through,
Thanks to the wisdom of Go.

There once was a concurrency pro
Who said, "Parallelism, you know
It's not the same thing
Concurrency's king
But parallelism's just for show"

There once was a Go programmer
Who said with a nod and a stammer
"Concurrency's not
Parallelism, you ought
To know it, it's not the same hammer."

There once was a code so fine
That concurrency was divine
But beware of the trap
Parallelism's a snap
That's not the same, oh so fine!

There once was a concurrency so fine
It was not parallelism, but divine
It allowed for great speed
In tasks that we need
And made our programs much more refined.

A concurrent program may run fast
But parallelism it does not surpass
For concurrency means
That tasks are in between
But parallelism needs more than one cast

Concurrent tasks may be run,
But parallelism is a bit more fun,
It means multiple paths,
All executing at once,
While concurrent just means they're not done.

There once was a programmer named Go
Who loved concurrent programming, don't you know
It's not quite the same
As parallel programming's game
But both seek efficiency, that's for sure, oh!

-- The next two were generated from the same prompt.

Concurrent programming involves the sharing
Of resources by processes that are daring
They execute in turn
But the order is not discerned
And no speedup is expected from pairing

Parallel programming takes a different track
With multiple threads that will never look back
They run in parallel
And the order is parallel
And speedup is the goal that they'll attack

--

There once was a programmer named Joe
Who wanted to learn concurrent and parallel Go
He learned that concurrent means things happen in turns
While parallel means they all run and burn
Now he codes with efficiency and speed, don't you know!

There once was a programmer named Moe
Who thought concurrent and parallel were the same
But when he tried to code
He found the load
Was not distributed and the program was lame

There once was a Go programmer
Whose code was concurrent but not faster
Though he thought it would run
In parallel, it was done
One CPU core at a time, not blaster.

There once was a Go programmer,
Whose concurrent code was a hammer,
But despite the best tries,
It didn't parallelize,
Because the scheduler's hands were tied.

There once was code in Go concurrent,
But not always parallel it was sent,
Due to CPU cores,
It may not run in scores,
And thus not as fast as intended.

There once was a code that was grand
It ran concurrently in demand
But alas, it was not
Run in parallel, it was caught
And thus, its speed was not as planned.

Bottom line: Apparently ChatGPT has a hard time explaining the difference between concurrency and parallelism without making concurrency look inefficient or even pointless. I picked the best of the bunch, with small edits.
-->

There once was a programmer named Joe.\
He thought concurrency was parallelism, oh no!\
But soon he learned that's not true,\
And now his code runs through,\
Thanks to the wisdom of Go.

## Channels orchestrate; mutexes serialize.

In Go, channels are quite grand,\
They help processes take a stand,\
While mutexes will make,\
Things run in a straight line,\
And keep the program in command.

## The bigger the interface, the weaker the abstraction.

<!-- There once was an interface so big\
It caused the abstraction to rig\
The bigger it grew\
The weaker it threw\
Until it collapsed with a thud.-->

There once was an interface big and wide\
It was quite the spectacle, oh how it tried\
But alas, it was weak\
For its abstraction was bleak\
Less is more, it's time to step aside

## Make the zero value useful.

There once was a zero value so bland\
It sat there with nothing in hand\
But one clever programmer\
Gave it a new demeanor\
Now it's useful and doing as planned.

## interface{} says nothing.

There once was a type named interface{}\
That didn't reveal anything concise\
It kept things quite vague\
And left programmers in a rage\
"What does this mean?" they would entice.

## Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.

<!-- Gofmt's style may not be the best\
But it's still the tool that they request\
Though it's not quite their taste\
They'll use it with haste\
For its formatting they do attest. -->

Though Gofmt's style may not please\
It's still a tool that we seize\
To format our code\
In a consistent mode\
And make our projects look sleek and concise.

## A little copying is better than a little dependency.

A little copying's not a crime\
It can save you from dependency's grime\
Just be sure to give credit\
Where credit is due\
And avoid any potential legal time.

## Syscall must always be guarded with build tags.

<!-- Syscall must always be tagged\
To avoid builds that are jagged and saggy\
Guard it with care\
And your code will be fair\
And your builds will be smooth and swanky. -->

Syscall must be guarded with care\
With build tags to protect and ensnare\
Any potential threat\
That might cause it to fret\
And keep your program running fair.

## Cgo must always be guarded with build tags.

Cgo must be guarded with tags\
To prevent any potential snags\
In the building process\
It can save you from stress\
And keep your program free of flags.

<!-- Cgo must be used with tact\
And guarded with build tags to act\
As a shield from harm\
And keep your program from alarm\
So it can function without any lack. -->

## Cgo is not Go.

<!-- Cgo is not the same as Go\
It's a different beast, you know\
With unique features\
And different creatures\
So be sure to use it with a steady flow. -->

Cgo may seem like Go at first glance\
But it's quite a different circumstance\
It allows for C code\
To be integrated, bestowed\
But it's still not the same at a glance.

## With the unsafe package there are no guarantees.

<!-- The unsafe package is quite the wild card\
It offers no guarantees, quite hard\
To use it with care\
And be aware\
Of the potential danger, be on guard. -->

The unsafe package, it's plain to see\
Comes with no guarantees, can't you agree?\
It opens up doors\
To potential pitfalls and roars\
So use it cautiously, wisely, and free.

## Clear is better than clever.

Clever code can be quite a feat\
But clear code can't be beat\
It's easy to read\
And doesn't mislead\
So make sure your code is both clear and neat.

<!-- Clever code may be a delight\
But clear code is a better sight\
It's easy to read\
And won't leave you misled\
So aim for clarity, not just insight. -->

## Reflection is never clear.

Reflection can be quite a fuss\
It's never clear, that much is a must\
It can cause confusion\
And lead to an illusion\
So use it sparingly, if at all, and adjust.

## Errors are values.

<!-- Errors are not just a message\
They're values that we must acknowledge\
They can provide insight\
Into what's not quite right\
So don't ignore them, let them collage.

Errors are values, that much is true\
They can be assigned and passed through\
They provide insight\
On what went wrong, in sight\
And allow for graceful recovery, anew. -->

Errors, they're not just a blight\
But values that must be dealt with right\
They can be stored\
And later explored\
To help your program take flight.

## Don't just check errors, handle them gracefully.

Checking errors is just the start\
You must handle them with a graceful heart\
Don't just ignore them\
Or let your program condemn\
Be mindful of how you play your part.

<!-- Errors are values, that much is true\
And they can be quite useful, that's no small feat\
They can inform us\
Of potential problems and fuss\
So handle them gracefully and sweet. -->


## Design the architecture, name the components, document the details.

Design the architecture, with care\
Name the components, with flair\
Document the details\
So others won't fail\
And your project can thrive beyond compare.


## Documentation is for users.

Documentation, it's a must\
For users, so they won't be lost\
It helps them understand\
And follow your plan\
So they can use your program without any fuss.


## Don't panic.

When panic strikes, don't despair\
Take a deep breath, and take the air\
Remember that quote\
And all will be afloat\
"Don't panic," just breathe, and repair.

## <!-- to add vertical space -->

Happy coding!

ʕ◔ϖ◔ʔ
___
*/
